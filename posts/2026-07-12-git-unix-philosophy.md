---
title: Git is a Unix tool
date: 2026-07-12
slug: git-unix-philosophy
summary: "Git was designed by someone who hated CVS, understood filesystems, and believed in conceptual integrity. Its best features — the object model, plumbing and porcelain, branches as 40-byte pointers — are the ones Brooks and Parnas would have designed."
tags: git, unix, linus-torvalds, conceptual-integrity, information-hiding
---

In April 2005, the Linux kernel project lost access to BitKeeper, its proprietary version control system. Andrew Tridgell had reverse-engineered the BitKeeper protocol. The license was revoked. Linus Torvalds halted all kernel development and wrote a replacement. He called it Git.

> "Writing code is easy. Getting a good design is what matters. So there was a fair amount of background to those few days that is pretty important, and that part doesn't show up in the history." — Linus Torvalds

The "few days" is the myth. The background is the reality. Linus had been thinking about version control design for months. He had strong opinions about what was wrong with every existing system — CVS, SVN, BitKeeper, the commercial alternatives. When the crisis hit, the design was ready. The implementation took two weeks.

> "I based pretty much all of the git design on three basic goals: performance, distribution, and integrity checking. Everything else pretty much flows from those three things."

Performance. Distribution. Integrity. Three goals. Everything else emerges. This is conceptual integrity in Brooks's sense: a small set of orthogonal primitives from which the entire system follows. The design feels like one mind made it — because one mind did.

## The Unix philosophy, applied to version control

Linus describes Git's relationship to Unix explicitly:

> "I kind of compare it to Unix. Unix has like a core philosophy of everything is a process, everything is a file, you pipe things between things. There's the simple concepts that underlie the philosophy, but then all the details are very complicated. I think Git has some of the same kind of — there's a fundamental core simplicity to the design and then there's the complexity of implementation."

Unix: everything is a file. Git: everything is an object identified by its content hash.

The mapping is direct. In Unix, the file is the universal abstraction — storage, devices, sockets, pipes all present the same interface. In Git, the content-addressed object is the universal abstraction — files (blobs), directories (trees), history (commits), releases (tags) all use the same storage mechanism. You insert content. Git returns a hash. You ask for the hash. Git returns the content. The key-value store is the kernel. The version control system is a user interface built on top.

This is the core design insight. Git is not a version control system with a storage layer. It is a content-addressable filesystem with a version control interface. The Pro Git book states it directly: "Git is fundamentally a content-addressable filesystem with a VCS user interface written on top of it." The filesystem is the mechanism. The VCS is the policy. Mechanism, not policy — the oldest Unix principle, applied to the problem Linus understood best.

## The object model: four primitives

Git's conceptual integrity rests on four object types. No more. Each is content-addressed. Each is immutable once written. Each composes with the others.

**Blob.** File contents. No filename. No metadata. No permissions. Just the bytes, compressed, identified by their SHA-1 hash. Two files with the same content produce the same blob. Deduplication is free. Identity is content. Content is identity. This is the simplest possible file abstraction. It is also the most powerful — because it makes no assumptions about what the bytes mean.

**Tree.** A directory listing. Names mapped to hashes. A tree points to blobs (files) and other trees (subdirectories). It records the structure without recording anything about the content. The tree is the namespace. The blob is the data. Separate concerns. Separate objects.

**Commit.** A pointer to a tree (the snapshot), plus metadata: author, committer, timestamp, message, and zero or more parent commits. The parent pointers form a directed acyclic graph. History is the DAG. The DAG is traversed. Branches are just refs pointing to commits. Merges are commits with multiple parents. The entire history model falls out of the commit object's parent pointers.

**Tag.** A named reference to any object, typically a commit, with an optional message and GPG signature. Used for releases. Lightweight. Immutable once created. The tag on the object. Not the object itself.

Four types. That is the entire storage model. Every Git repository is a key-value store containing blobs, trees, commits, and tags. Every Git operation — commit, merge, rebase, cherry-pick, bisect — is a manipulation of these four types, referenced by hash, organized into a DAG. The complexity is in the operations. The substrate is minimal.

## Plumbing and porcelain: information hiding applied

Git was architected from the start as two layers:

**Plumbing.** Low-level commands that manipulate the object store: `git hash-object`, `git cat-file`, `git write-tree`, `git commit-tree`, `git update-ref`. These are stable, scriptable, composable. They do one thing. They can be piped. They are the Unix toolset applied to a content-addressable database.

**Porcelain.** High-level commands that provide user workflows: `git commit`, `git merge`, `git log`, `git rebase`. Built on plumbing. Cosmetic. Replaceable. The first version of Git had no porcelain — you committed by running `git commit-tree` and manually writing the resulting hash to `.git/HEAD`. Linus built the mechanism first. The user interface came later, contributed by others, layered on top.

This is Parnas's information hiding in system architecture. The plumbing hides the object store's implementation — the compression algorithm, the pack format, the delta encoding — behind stable interfaces. The porcelain hides the workflow complexity behind user-friendly commands. Each layer's internals can change without affecting the other. Git's pack format has been revised multiple times. The SHA-1 to SHA-256 migration is underway. The porcelain has grown from a handful of commands to hundreds. The plumbing interfaces are the same.

> "I approached it more like I would a distributed journaling filesystem, not really a traditional SCM." — Linus Torvalds

He designed a filesystem. The filesystem's interface is stable — put content, get hash; give hash, get content. The VCS is a consumer of the filesystem. The separation is clean. The hiding is real. Twenty years later, the original plumbing still works. The porcelain has changed beyond recognition. The architecture absorbed the change because the volatile decisions were hidden behind the right interfaces.

## The ultra-minimalist features

Git's best features are the ones it doesn't have. Or the ones it implements so simply they seem trivial. Brooks would recognize them as examples of conceptual integrity achieved through refusal to add complexity.

**Branches are 40-byte pointers.** A branch is a file in `.git/refs/heads/` containing a single SHA-1 hash. Creating a branch is writing 40 bytes. Switching branches is changing which ref `HEAD` points to. Merging is creating a commit with two parents. Branching in CVS meant copying the entire repository. In Git, it means writing 40 bytes. The cost is so low that branching stops being a decision. It becomes a reflex. This changes how people work — not because Git preached branching, but because the cost made it invisible.

**No rename tracking.** Traditional SCMs track file renames by recording metadata: "file A was renamed to file B." Git doesn't. Renames are inferred from content similarity at query time. If you rename a file and modify 90% of it, Git sees a delete and a create. If you rename a file and modify 10%, Git sees a rename with modifications. The algorithm runs when you ask `git log --follow`. The decision is deferred to the reader, not encoded by the writer. This is information hiding applied to history: the rename is an interpretation, not a fact. Different tools can interpret differently. The history is not polluted with metadata that later proves wrong.

**Everything is local.** `git init` creates a `.git` directory. That directory contains the entire repository — objects, refs, config, hooks, the index. Clone it, copy it, back it up. Every clone is a full backup. Every developer has the complete history. No network required for commit, branch, merge, log, bisect, or blame. The server is just another clone. The distributed nature is not a feature bolted onto a centralized model. It is the model. The centralized workflow is a convention built on top of a distributed substrate.

**Integrity is automatic.** Every object is named by its content hash. Any corruption anywhere is immediately detectable. History is tamper-evident — changing any past commit changes all subsequent hashes. Trust is cryptographic, not social. You don't need to trust the server. You don't need to trust your colleagues. You can verify. The SHA-1 was never about security. It was about detecting accidental corruption. The fact that it also detects malicious tampering is a side effect.

> "People kind of think that using the SHA-1 hashes was a huge mistake. But to me, SHA-1 hashes were never about the security. It was about finding corruption." — Linus Torvalds

**The design heuristic: WWCVSND.** What Would CVS Not Do? Linus hated CVS. He saw SVN as "lipstick on a pig." His design method was systematic inversion: distributed instead of centralized, whole-tree snapshots instead of per-file history, content-addressable instead of incrementally versioned, local instead of networked, lightweight branches instead of heavy copies. Every design decision was the opposite of what CVS had done. The result was not a better CVS. It was a different category of thing.

## What Brooks and Parnas would say

Brooks would recognize Git's conceptual integrity immediately. Four object types. One identification scheme. One DAG model. One storage layer. A system that feels like one mind designed it — because one mind did. The complexity is in the operations on the model, not in the model itself. The model is stable. The operations evolve. This is the architecture of systems that age well.

Brooks argued that conceptual integrity requires one mind. Git had one mind for its critical design phase. Linus designed the core. Junio Hamano maintained it. The porcelain grew from contributions. The plumbing stayed stable. The one-mind rule held where it mattered — at the object model, the storage layer, the fundamental abstractions. The community built on top. The foundation didn't shift.

Parnas would recognize the information hiding. The plumbing hides the object store. The porcelain hides the plumbing. The object model hides the storage format. The content addressing hides the transport mechanism. Each volatile decision is encapsulated behind a stable interface. The pack format changed. The network protocol changed. The hashing algorithm is migrating. The interface — insert content, get hash; give hash, get content — has not changed in twenty years. This is Parnas's criterion applied to system infrastructure. Hide the decisions likely to change. Expose stable interfaces. Let the rest of the system evolve behind them.

Parnas would also recognize the design method as a form of information hiding applied to project structure. The plumbing is the stable core, maintained by a small group who understand it deeply. The porcelain is the volatile periphery, contributed by the community, evolving rapidly. The interface between them — the plumbing commands — is the contract. As long as the contract holds, the two layers can evolve independently, at different speeds, by different people, with different governance. This is modularity in organizational form.

## The thing Git got right that most software gets wrong

Git's defining achievement is not any individual feature. It is that the core design has not needed to change. The object model is twenty years old. The DAG model is twenty years old. The plumbing is twenty years old. The system has scaled from the Linux kernel — the largest software project in history — to individual developers' dotfiles, without changing the fundamental abstractions. The same `git commit` works for a monorepo with millions of files and a single-file hobby project. The primitives composed.

Most software systems fail this test. The abstractions that worked at version 1.0 cannot handle version 10.0. The design decays. The model fractures. Git's model didn't fracture because the model was minimal. Four object types. One DAG. One content-addressing scheme. There was nothing to fracture. The minimalism was not an aesthetic choice. It was an engineering strategy. Less to design. Less to implement. Less to break. Less to regret.

Linus understood this, whether or not he would use Brooks's language. He built the simplest thing that could work for his problem — Linux kernel development — and nothing more. He refused to design for use cases he didn't have. He refused to add features he didn't need. He refused to generalize beyond the concrete problem in front of him.

> "I'll do something that works for me, and I won't care about anybody else. And really that showed in the first few months and years — people were complaining that it was kind of hard to use, not intuitive enough. And then something happened, like there was a switch that was thrown."

The switch was GitHub. The community built the porcelain. The community built the hosting. The community built the tutorials, the GUIs, the integrations. Linus built the content-addressable filesystem with the DAG on top. The world built everything else. This is the pencil argument applied to version control: no single person knows how to make the whole Git ecosystem. Millions contributed their tiny know-how. The core, the part one mind designed, remained stable. The periphery, the part millions contributed to, evolved rapidly. The interface between them — the plumbing — held. That is the architecture of a system that will outlive its creator.

---

**References:**
- Linus Torvalds, "Git turns 20: A Q&A with Linus Torvalds," GitHub Blog, April 2025.
- Linus Torvalds, Google Tech Talk on Git, 2007.
- Scott Chacon and Ben Straub, *Pro Git*, Chapter 10: Git Internals, Apress.
- Related posts: [Brooks on Software Design: conceptual integrity](https://blog.hackspree.com/#brooks-design-conceptual-integrity), [Parnas's Information Hiding](https://blog.hackspree.com/#parnas-information-hiding), [I, Pencil](https://blog.hackspree.com/#i-pencil)
