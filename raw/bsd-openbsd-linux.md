---
title: BSD is clean, OpenBSD is cleaner
date: 2026-07-12
slug: bsd-openbsd-linux
summary: "BSD is an operating system designed as a whole. Linux is a kernel assembled into an operating system by distributions. The difference in design quality follows directly from the difference in architecture. OpenBSD takes the logic to its extreme: correctness over features, every time."
tags: bsd, openbsd, linux, unix, operating-systems, software-design
---

In 1991, a Finnish student wanted a Unix-like system for his 386 PC. BSD Net/2 had been released, but he didn't know about it. He started writing his own kernel. He called it Linux. The rest is history — but the wrong history, or at least an incomplete one.

Linus Torvalds later said that if a working 386BSD had been available, he would never have created Linux. He was not making a philosophical choice between cathedral and bazaar. He was solving a practical problem with the tools he knew existed. The lawsuit that froze BSD development from 1992 to 1994 — AT&T's USL vs. BSDi, alleging that BSD contained proprietary Unix code — gave Linux a two-year window with no competition. By the time the suit settled and 4.4BSD-Lite shipped clean, Linux had momentum it would never lose.

The consequence is that the world runs on the bazaar. The cathedral — cleaner, more coherent, better designed — is a niche. The niche is worth understanding because it represents a different theory of what an operating system should be. OpenBSD is the purest expression of that theory.

## What BSD is

BSD is not a kernel with some utilities. It is an operating system. One source tree. One development group. One coherent design. The kernel, the C library, the core utilities (`ls`, `cp`, `grep`, the shell), the daemons, the manual pages — all maintained by the same people, in the same repository, under the same quality standards.

This is the model inherited from Research Unix at Bell Labs. Ken Thompson and Dennis Ritchie didn't ship a kernel and let someone else figure out the userland. They shipped a system. The system had conceptual integrity because one small group controlled all of it. BSD preserved this model. Linux abandoned it — not by choice, but by circumstance. When Linus started, the GNU project had already written the C library, the coreutils, and the shell. He only needed to write the kernel. The assembly model was the path of least resistance.

The difference is visible at every level. BSD configuration files have consistent syntax because the same people who wrote the daemons wrote the configuration parsers. BSD manual pages are complete and accurate because they are maintained as part of the source tree, not as an afterthought by a separate documentation project. BSD systems boot predictably because the init system, the device drivers, and the service scripts were designed together, not integrated after the fact by a distribution maintainer.

> "Practically the entire BSD distribution is written and packaged by the distribution maintainers. The FreeBSD people know about (and have written) everything that's part of the FreeBSD distribution."

This is not a marketing claim. It is a statement about the structure of the source repository. Every component is designed, reviewed, and tested as part of one system. When a kernel interface changes, the userland tools that depend on it are updated in the same commit. When a security vulnerability is found in a library, every program that uses that library is audited — because they are all in the same tree and maintained by the same people. The coherence is not accidental. It is the result of organizational structure. It is Conway's Law, applied to operating systems: the design mirrors the communication structure of the team that builds it. BSD's team is one team with one tree. The design has one voice.

## What Linux is

Linux is a kernel. Linus Torvalds chose to focus only on the kernel and to not ship user-level programs. Distributions — Red Hat, Debian, Ubuntu, Arch, hundreds of others — assemble an operating system from the Linux kernel, the GNU C library, GNU coreutils, a shell, a desktop environment, an init system, and hundreds of other projects. Each project has its own maintainers, its own priorities, its own release schedule, its own coding style, its own documentation format. The distribution's job is to make them work together and to resolve the conflicts when they don't.

This model has been extraordinarily successful at generating variety and breadth. There are Linux distributions for every purpose, every hardware platform, every ideological preference. The Linux kernel supports more hardware than any other operating system in history. The driver ecosystem alone is a miracle of coordination — thousands of contributors from hundreds of companies, merging code into a single tree at a rate of 10,000+ commits per release cycle.

But the model has no mechanism for internal coherence. There is no single group that understands the entire system. There is no consistent design language across components. A Linux system is a negotiated settlement between independently developed projects that happen to run on the same kernel. The negotiation is managed by distribution maintainers who did not write any of the components and whose primary job is integration, not design.

> "Linux has no similar concept [of a base system]; the kernel is maintained and distributed by one group, the usual runtime library by another, and so on. Linux distributions have the job of assembling all of the bits."

This is Brooks's committee design applied to operating systems. Each component is individually impressive. The assembly lacks conceptual integrity because no single mind controlled the interfaces. The interfaces were negotiated between projects with different goals, different timelines, and different ideas about what "good" means. The result works. It is not clean. It is not coherent. It is a patchwork quilt. Everyone who has debugged a Linux system at 3am knows the feeling of crossing a component boundary and discovering that the assumptions changed.

## The lawsuit that chose the bazaar

The USL vs. BSDi lawsuit (1992-1994) is the pivot on which operating system history turned. AT&T's Unix Systems Laboratories sued Berkeley Software Design, Inc., alleging that the BSD Net/2 release still contained proprietary AT&T code. The suit froze BSD development for nearly two years. Developers couldn't contribute. Users couldn't trust the codebase. Companies couldn't build products on it. The uncertainty was total.

During those two years, Linux — which had no AT&T code, was written from scratch, and faced no legal threat — absorbed the energy that would have gone to BSD. Developers who would have contributed to the BSD kernel contributed to Linux instead. Companies that would have built BSD distributions built Linux distributions instead. The network effects tipped. By 1994, when the lawsuit settled and 4.4BSD-Lite shipped with all AT&T code removed, Linux had won the mindshare. It has never lost it.

Linus himself acknowledged the contingency:

> "If 386BSD had been available when I started on Linux, Linux would probably never have happened."

The world's dominant server operating system was not chosen on technical merit. It was chosen because a lawsuit froze the technically superior alternative during the critical window when network effects were forming. This is not a criticism of Linux. It is a fact about history. The cathedral didn't lose because it was worse. It lost because it was sued. The bazaar didn't win because it was better designed. It won because it was available. Availability beats design quality in the short run. The short run became the long run.

## OpenBSD: the extreme of the philosophy

If BSD represents the cathedral model, OpenBSD represents the cathedral with the strictest building code. Theo de Raadt forked OpenBSD from NetBSD in 1995. The project's goals are explicit and uncompromising: correctness over features, security through design, and code quality as the primary metric.

> "We are non-stop trying to find ways across our entire source tree that small little programmer errors result in problems. At some point, we have to start asking ourselves whether features are the thing, or whether quality is the issue. I really think we have to focus on the quality before the features." — Theo de Raadt

This is not a slogan. It is enforced by process. The six-month release cycle — May and November, every year, twenty-five consecutive on-time releases with no critical bugs — structures all development.

The cycle has phases. Four months of development: features are written, code lands, the tree is open. One month of API lockdown: interfaces freeze, testing intensifies, bugs are fixed. Final weeks of code freeze: only the simplest edits — documentation, minor fixes — are accepted. The tree must build and boot on every supported architecture before release. If a feature isn't ready, it waits. The next release is never more than six months away.

> "Pretty soon, you get a clue, and start working with the release. Also, when you have a guarantee that you WILL have a six months release schedule, trying to wedge that last improvement in loses some of its attraction. You know that you will be able to do it for the next release, which is only six months away..." — Marc Espie, OpenBSD developer

The six-month cadence is a forcing function. Features that aren't ready are deferred. Features that are deferred are refined. Features that are refined land clean. The discipline produces a system where every component has been through the freeze cycle multiple times, where every interface has been tested on every architecture, where the manual pages match the code because both were updated before the freeze. The result is not just a secure operating system. It is a well-engineered one.

## Security as design, not feature

OpenBSD's approach to security is the clearest expression of its design philosophy. Most operating systems treat security as a feature to be added: firewalls, anti-virus, intrusion detection, patches for vulnerabilities as they are discovered. OpenBSD treats security as a property of the design itself. Eliminate entire classes of vulnerability. Make the correct thing the only possible thing. If the API can be misused, change the API.

The innovations list is long because the approach is systematic:

- **W^X:** Memory is either writable or executable. Never both. This eliminates the mechanism that most exploits rely on. Implemented in 2003. Now standard everywhere. OpenBSD did it first.

- **pledge(2):** A process declares at startup what system calls it will use. After `pledge()`, any other system call kills the process. A file server that pledges only `stdio` and `sendmsg` cannot open new files, cannot fork, cannot exec. If an attacker compromises it, they gain a process that can do almost nothing. This is not a mitigation. It is a design constraint enforced by the kernel.

- **unveil(2):** A process declares which parts of the filesystem it can access. After `unveil()`, the rest of the filesystem does not exist as far as that process is concerned. A web browser that unveils only `~/Downloads` cannot read `~/.ssh`. If the browser is compromised, the attacker can't either.

- **Secure malloc:** The memory allocator randomizes allocations, guards pages, and detects use-after-free. These are not optional hardening flags. They are the default allocator.

- **pf(4):** The packet filter, written by OpenBSD, with a clean syntax, default-deny semantics, and integration with the rest of the system's security model. Not bolted on. Part of the base system.

The project's security record is stated without marketing: "Only two remote holes in the default install, in a heck of a long time." The statement is verifiable. The record is public. The claim is modest. The achievement is extraordinary.

> "The problem with security is that people learn what they're supposed to by example, learn they're supposed to use APIs in a certain way, and they're just wrong." — Theo de Raadt

The API is the problem. Fixing the API fixes the vulnerability class. Fixing individual instances of the vulnerability fixes only the instances. The OpenBSD approach is to find the API that produces the vulnerability, change the API so it cannot be misused, and then fix every caller to use the new API correctly. This is information hiding applied to security: hide the dangerous operation behind an interface that makes the dangerous operation impossible. The caller cannot do the wrong thing because the wrong thing is not exposed.

## What Brooks and Parnas would see

Brooks would recognize BSD as a system with conceptual integrity. One small group controls the entire design. Every interface is reviewed by the same people who wrote the implementations that use it. The system speaks with one voice. The manual pages match the code. The configuration syntax is consistent. This is what conceptual integrity looks like at the scale of an operating system. It is achieved by the one-mind rule — one development group with authority over the entire tree.

He would recognize Linux as a system assembled from components designed by different groups with different visions. The kernel has one design philosophy. Systemd has another. The GNU tools have a third. GNOME has a fourth. Each component is internally coherent. The assembly is not. This is committee design at the scale of an operating system. It works. It is not clean. The seams between components are where the complexity lives. The seams are also where the bugs live.

Parnas would recognize OpenBSD's security architecture as information hiding applied to attack surfaces. The kernel hides the hardware. Libc hides the kernel. The daemons hide the services. Each layer exposes the minimum interface. The pledge and unveil mechanisms enforce that a compromised process cannot exceed its declared interface. The attacker gains access to a process. The process has no access to anything else. The information hiding is not advisory. It is enforced by the kernel. The volatile decision — what this process can do — is hidden behind a stable interface that the process itself cannot change.

Parnas would also recognize the release cycle as a form of modular development. The base system is a module with a stable interface (the release). The ports tree is a separate module with a different interface (the package). The two modules evolve at different speeds, by different processes, with different quality standards. The release boundary is the contract. What ships in the release must work together, on every architecture, with no known bugs. What is in ports is best-effort. The separation is clean. The hiding is real. The architecture absorbs the difference in quality requirements because the volatile part (ports) is isolated from the stable part (base).

## The dirty and the clean

Linux is dirty in the way that a city is dirty. It works. It is full of life. It contains multitudes. The streets don't follow a plan because the city grew organically, each neighborhood added by different builders at different times with different ideas about what a street should be. The result functions. It is not elegant. Nobody would design it from scratch this way. But it is what exists, and it runs most of the internet, and the sheer variety and energy and pace of development are extraordinary.

BSD is clean in the way that a well-designed building is clean. The structure is visible. The materials are consistent. The wiring is labeled. The documentation matches the implementation. You can understand the whole thing by studying any part because the same design language is used throughout. It does fewer things than Linux. The things it does, it does correctly. The release ships on time, with no critical bugs. The manual pages are accurate. The security model is coherent. The system makes sense as a system.

OpenBSD is cleaner still — the building with the strictest code, inspected continuously, where every door has a lock and every lock has a key and every key opens exactly one door. The tradeoff is clear: less hardware support, fewer features, a ports tree that lags behind Linux package availability. The benefit is also clear: a system where correctness is not aspirational but enforced by process, where security is not a feature but a property of the design, where the release ships on schedule and you can trust it.

The world runs on Linux because the lawsuit froze BSD at the wrong moment and network effects did the rest. The cathedral lost. The bazaar won. But the cathedral is still standing. For work where correctness matters more than breadth — firewalls, routers, secure servers, any system facing the internet — the cathedral is the right choice. OpenBSD is the cathedral at its most uncompromising. The design principles are visible in every layer. The discipline is enforced by process. The result is an operating system that does less than Linux and does it correctly. That is not a limitation. That is the point.

---

**References:**
- OpenBSD Project Goals, [openbsd.org/goals.html](https://www.openbsd.org/goals.html)
- Theo de Raadt, "OpenBSD: Maintaining the quality mindset," ZDNet interview, 2004.
- Warner Losh, TUHS mailing list, January 2020: "AT&T lawsuit was after Linus started."
- Linus Torvalds, quoted in multiple sources: "If 386BSD had been available... Linux would probably never have happened."
- Related posts: [Brooks on Software Design: conceptual integrity](https://blog.hackspree.com/#brooks-design-conceptual-integrity), [Parnas's Information Hiding](https://blog.hackspree.com/#parnas-information-hiding), [Git is a Unix tool](https://blog.hackspree.com/#git-unix-philosophy)


Engineering is the discipline of building things that work within constraints. Every topic on this blog — operating systems, AI models, trading infrastructure, research labs, innovation economics — is examined through the lens of systems design. The lens is engineering. The method is: understand the constraints, design within them, verify the design works, iterate. The domain provides the specifics. The method is universal.


> Clean code is not an aesthetic preference. It is a property of systems designed as a whole. BSD is clean because one group designed the whole. Linux is not because a thousand groups designed the parts.
