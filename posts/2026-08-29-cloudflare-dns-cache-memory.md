---
title: "953 Bytes to 420: How Cloudflare Freed 100TB of RAM by Redesigning Its DNS Cache"
date: 2026-08-29
slug: cloudflare-dns-cache-memory
summary: "Five Rust-level changes to Big Pineapple, the platform behind 1.1.1.1, cut each DNS cache entry from 953 to 420 bytes — freeing roughly 100TB across the fleet (the DDR5 of 130 Gen 13 servers) while making the cache faster: insert throughput up 43%, lookup latency down 19%. At 250 billion entries, every byte is a budget line."
tags: [cloudflare, dns, rust, memory, cache, systems-design, performance, 1.1.1.1, jemalloc, big-pineapple]
---

## The arithmetic of 250 billion entries

DNS caching is a memory business, and the volumes are so large that the smallest inefficiency becomes a line item. [Big Pineapple](https://blog.cloudflare.com/big-pineapple-intro/), the Rust platform behind [1.1.1.1](https://blog.cloudflare.com/the-as112-project/) — plus Gateway DNS, DNS Firewall, and AS112 — stores **over 250 billion cache entries at any given time**. Cloudflare's own math: wasting a single byte per entry costs more than 250 gigabytes across the fleet. Waste one plausible structural inefficiency per entry, and you have a data center.

Which is exactly what five successive changes just removed. In [a deep dive by systems engineer Sebastiaan Neuteboom](https://blog.cloudflare.com/dns-cache-memory-optimization-1111/) ([Tom's Hardware coverage](https://www.tomshardware.com/tech-industry/big-tech/cloudflare-frees-100tb-of-ram-by-shrinking-dns-cache-entries)), each cached entry shrank from **953 bytes to 420 bytes — a 56% reduction** — with no physical RAM added to a single server. Across the fleet that is roughly **100 terabytes of freed memory, equivalent to the DDR5 in 130 of Cloudflare's 768GB Gen 13 servers**. The cache also got *faster*: insert throughput up 43%, lookup latency down 19%. No reconfiguration of memory modules, no added capacity — a redesign of how each entry is laid out in memory, done one Rust type at a time.

The timing makes the win sharper than the raw numbers. Server-grade DDR5 prices are [on track to double year over year](https://www.tomshardware.com/tech-industry/big-tech/cloudflare-frees-100tb-of-ram-by-shrinking-dns-cache-entries) — the 100TB Cloudflare reclaimed is the memory it did *not* have to buy at the top of a price spike.

## What a cache entry actually is

Every cache item is a key-value pair. The key identifies what was queried:

```rust
pub struct CacheKey {
    qname: Name,          // the domain being resolved
    qtype: Rtype,         // A, AAAA, TXT, ...
    authenticated: bool,  // DNSSEC-validated
    tag: Vec<u8>,         // per-client-subnet tag
}
```

The value holds the DNS response — answer, authority, and additional record sections — plus metadata: creation time, TTL, a hit counter:

```rust
pub struct CacheEntry {
    timestamp: UnixTimeStamp,
    inception: Instant,
    ttl: Ttl,
    hits: u32,
    answers: Vec<Record>,
    authority: Vec<Record>,
    additional: Vec<Record>,
    errors: Vec<ExtendedError>,
    // ...
}
```

Nothing about either struct is wrong. Both are idiomatic Rust, and that is precisely the problem: idiomatic Rust is written for a working set of *one* entry. Big Pineapple holds a quarter of a trillion of them, so every field is multiplied by 250 billion. When EDNS Client Subnet (ECS) is in use, authoritative servers return different answers per client network and the cache holds multiple entries per query — multiplying both entry count and per-entry size, which is why the optimizations hit ECS-heavy data centers hardest.

## Change 1 — Fire the capacity field

`Vec<T>` stores three words: a pointer to heap data, a length, and a **capacity**. Capacity exists so a growing vector can append without reallocating. But once a DNS response is stored in the cache, it is never modified again. The capacity field is dead weight — 8 bytes per vector — and the *slack* capacity reserves is dead weight too: a `Vec` with capacity for eight records but only five stored leaves three slots of heap memory permanently unused.

Swapping `Vec<T>` for `Box<[T]>` and `String` for `Box<str>` removes both. A boxed slice can't grow, so it has no capacity field and reserves nothing beyond its exact contents. Each cache entry stores eight `Vec`/`String` fields; replacing them saves 8 bytes per field, 64 bytes per entry, plus the over-allocated heap. At 250 billion entries that single mechanical change is worth **over 15 terabytes**.

## Change 2 — Fewer lists, fewer pointers

The answer, authority, and additional sections were three separate boxed slices — six words of pointers and lengths per entry. DNS record counts per section fit in a `u16`, so Cloudflare stored one contiguous list plus **2-byte offsets** to each section's start. Two 8-byte pointer/length pairs become two 2-byte offsets: 28 bytes saved per entry.

This is where the post's most subtle point appears: memory savings don't always map to the bytes you deleted. Rust pads structs to alignment and rounds sizes up. Several booleans were packed into a single bitflag — and shrinking the field count eliminated surrounding padding, so the struct shrank by *more* than the booleans were worth. When your unit of account is 250 billion copies, padding is not decoration; it is a tax with a compounding multiplier.

## Change 3 — Drop the owner

Every DNS record carries an *owner* — the domain the record belongs to. And in the overwhelmingly common case, the owner is identical to the domain that was queried. A query for `example.com A` returns two records, both owned by `example.com`. Only in chains — `example.com → CNAME cdn.example.com → A records` — do owners diverge.

On the wire, DNS handles repeated owners with RFC 1035 name compression: a 2-byte pointer back to a previously-encoded name, so `www.example.com` encodes `www` plus a pointer to where `example.com` already appears. Cloudflare deliberately does *not* follow compression pointers in cache memory — doing so on the hot path is too expensive — but storing full owners for every record is waste. So the owner became optional:

```rust
pub struct Record {
    owner: Option<Box<Name>>, // None: owner == queried domain, inferred at read time
    class: Class,
    ttl: Ttl,
    rtype: Rtype,
    data: RecordData,
}
```

When `owner` is `None`, response construction restores the queried domain from the cache key — already in hand on every lookup — and avoids a heap allocation entirely. When it differs (the CNAME case), `Some` keeps a pointer to the full name. In practice the majority of cached records need no allocation for the owner at all.

## Change 4 — Size the enum, box the outliers

Rust enums are sum types: the enum is always the size of its *largest* variant, plus a tag. For record data, the natural design stores each DNS record type as a variant — and the natural design is catastrophically wasteful:

```rust
pub enum RecordData {
    A(Ipv4Addr),      // 4 bytes
    Aaaa(Ipv6Addr),   // 16 bytes
    Txt(Txt),
    Naptr(Naptr),     // 136 bytes — sets the enum size
    Svcb(Svcb),
    // ...
}
```

NAPTR carries three variable-length text fields, a domain name, and two integers: 136 bytes of payload, 144 bytes once tag and padding are counted. Every `A` record — 4 bytes of actual data — occupies that same 144-byte slot. `A` and `AAAA` are over **80% of traffic**, so the typical cached record wastes 120+ bytes on padding. This is the single biggest line item in the whole exercise, and the fix is a classic space/time trade:

```rust
pub enum RecordData {
    // Small and common variants stored inline
    A(Ipv4Addr),
    Aaaa(Ipv6Addr),
    // Large variants moved to the heap
    Txt(Box<Txt>),
    Naptr(Box<Naptr>),
    Svcb(Box<Svcb>),
}
```

The enum shrinks to 24 bytes; `A` and `AAAA` are inline, saving 120 bytes per record. Smaller types like TXT still occupy the 24-byte enum but their heap allocation is sized to their actual data instead of padded to 144. NAPTR itself *pays* — it now costs a pointer plus allocation overhead — but NAPTR records are rare, so the trade is cheap.

Boxing has two costs worth naming, because they come back in Change 5. First, **allocator overhead**: each boxed variant is a separate allocation, and allocators round up to size classes. Big Pineapple uses [jemalloc](https://jemalloc.net/), whose fixed-size bins fit a 32-byte TXT request perfectly but round a 40-byte MX request up to 48, wasting 8 bytes. Second, **locality**: boxed data scatters across the heap, so reading a record means following a pointer that may land on a cold cache line. Neither is fatal alone; together they are the reason the next change exists.

## Change 5 — Store the wire format, not the parse tree

The tempting endpoint — store the entire DNS response as one wire-format message and patch per-client fields like the message ID on lookup — has real problems. DNSSEC records are only included when the client sets the DO flag, so Cloudflare would have to cache two variants of every response or filter an already-built message, and parsing a full message on every lookup costs time the parsed-enum layout avoids.

The middle ground Cloudflare took: **store just the record data as raw bytes** — a single `Box<[u8]>` where each record is a 2-byte length prefix followed by its encoded bytes — while the rest of the entry stays structured. This eliminates the per-variant enum overhead *and* the boxed heap allocations from Change 4 in one move. All record data sits in one contiguous allocation, so locality improves instead of degrading. And response construction gets cheaper: most record types — A, AAAA, TXT, and all DNSSEC types — are **copied byte-for-byte into the outgoing message**, skipping the field-by-field reserialization the parsed representation required. Only records containing domain names (CNAME, NS, MX, SOA) still need parsing, for name compression. Since directly-copyable records are the vast majority of traffic, this change alone cut lookup latency 5% in benchmarks and raised insert throughput 13%.

The implementation detail that makes it affordable is a reusable **scratchspace buffer** that persists across insertions: records are serialized into it, then one `Box<[u8]>` is allocated and filled with a `memcpy`. That replaces N per-record allocations with a single one and avoids the allocator's inability to reclaim the unused tail when shrinking a `Vec<u8>`. The cost of the design is that records can no longer be randomly indexed — iteration through the buffer is sequential — which complicates round-robin rotation of A/AAAA records. With one to four records per entry, the price is negligible.

## The scoreboard

The benchmark (56% A, 25% AAAA, 19% TXT, one to four records per entry) tells the whole story:

| Metric | Before | After | Change |
|---|---|---|---|
| Per-entry net footprint | 953 bytes | 420 bytes | −56% |
| Per-entry allocations | 1.1 KB | 461 bytes | −58% |
| Cache insert throughput | 625,000 entries/s | 893,000 entries/s | +43% |
| Cache lookup latency | 828 ns | 670 ns | −19% |

Production numbers land close, with the rollout stepping across releases from May 18 to July 6, 2026: p99 resident memory fell from 9.3 GB to 5.3 GB per instance (−43%), p90 from 6.5 GB to 3.8 GB (−42%), and the aggregate working set settled roughly 100 terabytes lower. The planned reinvestment is the quiet best part: the freed memory goes back into **cache capacity**, raising hit rates and cutting upstream query volume — the same memory pays a second dividend.

## What this means

Four lessons, in increasing order of importance.

**Memory is a capex line now.** When DDR5 prices double year over year, the decision "buy more RAM vs. waste less of it" tilts hard. Cloudflare's answer was not a hardware purchase order but five releases of type-level discipline, delivered on a platform already holding a quarter trillion entries. The cost of the project was engineering; the avoided cost was a server generation's worth of DDR5.

**Type design is memory design.** Every one of the five changes is a Rust semantics fact applied at scale: `Vec` carries capacity for growth you never use; enums are sized by their largest variant; structs are padded to alignment; ownership is a pointer you can sometimes drop. These are the features that make Rust safe and ergonomic — and each one is overhead you stop paying the moment data becomes immutable at rest. Cache entries are never mutated after insert, which turns out to be a license to strip every field that existed only to support mutation.

**Layout beats cleverness.** Nothing here is algorithmic. No new data structure, no eviction breakthrough, no distributed trick. The win is the unglamorous work of removing overhead: capacity fields, redundant lists, duplicated owners, enum padding, scattered allocations. The compounding insight is the multiplier — 250 billion entries turns "a few bytes" into "a data center."

**Every win has a tax, and the taxes interact.** Boxing bought space and cost locality; the wire-format layout bought back the locality and cost random access; dropping owners bought space and cost self-containedness (the cache key must be in hand on every lookup). The discipline that makes it all work is measurement — a custom allocator wrapper tracking the number and size of allocations per entry, alongside insert/lookup benchmarks — followed by staged rollout. The same economic lens this blog applies to tokens applies to bytes: [every unit has a price tag](https://blog.hackspree.com/#every-token-has-a-price-tag), and at fleet scale the unit count makes the price.

The larger lesson is the one Cloudflare's engineers would probably insist on: the 100TB was never going to come from a single heroic optimization. It came from five boring ones, each individually worth terabytes because each one multiplied a small saving by 250 billion. Optimize the layout, not the algorithm — and when you find yourself saving a byte per entry, remember what the next byte is worth.
