---
title: "Learning to Fly, While Flying"
date: 2026-07-19
slug: learning-to-fly-while-flying
summary: "The image of a pilot reading 'How to Fly' mid-flight is the most accurate metaphor for software engineering ever produced. The absurdity is the point. The absurdity is also the job."
tags: software-engineering, learning, career, debugging
---

<div style="margin: 0 0 20px;">
  <img src="https://pbs.twimg.com/media/HNhSwViakAAQlif?format=jpg&name=large" alt="A pilot flying a plane while reading a 'How to Become a Pilot' book — the software engineering experience" style="width:100%;max-width:720px;border:1px solid rgba(255,255,255,0.08);border-radius:12px;display:block;">
</div>

There is an image circulating that captures the software engineering experience more accurately than any job description ever written: a pilot, mid-flight, reading a book titled *How to Become a Pilot.*

The plane is in the air. The pilot is at the controls. The book is open. These three facts should be mutually exclusive. In software engineering, they are simultaneously true at all times.

## The absurdity is the job

Every other profession understands this image as a failure mode. A surgeon does not open a textbook during an operation. A lawyer does not flip through tort law during cross-examination. A structural engineer does not consult the bridge-building chapter while the concrete is being poured. In every other field, learning before doing is basic professional hygiene. Doing while learning is a sign that something has gone catastrophically wrong.

In software, doing while learning is Tuesday.

You are three months into a React job. The senior engineer quits. You are now the frontend lead. You have never led a frontend. You are leading the frontend. You are reading the React docs on your second monitor while reviewing pull requests on your first. The plane is at cruising altitude. The book is open. You are becoming the pilot by piloting.

> Nobody who hires a software engineer expects them to know the stack. They expect them to be able to learn the stack faster than the stack can break. The job is not knowing. The job is learning at altitude.

## Why this works (when it works)

The image is funny because it's absurd. But the reason it's a metaphor rather than an indictment is that software is uniquely learnable in production.

Pilots have simulators. Software engineers have staging environments. Pilots have checklists. Software engineers have linters and CI. Pilots have co-pilots. Software engineers have code review. These are not perfect analogues — a failed staging deploy does not kill 180 passengers — but the structure is similar. The feedback loops are tight. The cost of a mistake, in most cases, is a broken build, not a broken fuselage.

And the knowledge required to fly any particular software plane is so specific, so contingent on decisions made by people who left the company three years ago, that no amount of pre-flight training could cover it. You cannot train for a legacy monolith built by a developer who named every variable after Lord of the Rings characters. You can only encounter it, at altitude, and open the book.

> The plane you are flying was assembled in the air by seventeen previous pilots, each of whom had their own book open. The manual you're reading was written by the third pilot, updated by the seventh, and contradicted by the twelfth. The manual is wrong in three places. You will discover which three by flying.

## When the metaphor gets dark

The joke works because we recognize ourselves in it. The joke gets uncomfortable when we stop recognizing the exit.

In a healthy career, the book eventually closes. You internalize the controls. You look out the window instead of at the manual. The learning continues, but the *emergency* learning — the kind where you're reading the docs while the error is on the screen, the kind where you're Stack Overflowing in one tab and SSHing in another — should be a phase, not a permanent state.

When it becomes permanent, it stops being a metaphor for learning and becomes a metaphor for burnout. The pilot who never stops reading the manual is a pilot who never gains confidence in their ability to fly. They are permanently a beginner, in a profession where the plane keeps getting more complex and the manual keeps getting thicker and the passengers keep asking when they'll arrive. The pilot is still reading. The pilot has been reading for fifteen years. The pilot is a staff engineer.

> The difference between a junior and a senior is not that the senior knows how to fly. It's that the senior knows which pages of the manual to ignore.

## Learning to fly is flying

Here is the thing the image doesn't tell you but implies: the pilot reading the book is, against all probability, still flying the plane. The plane has not crashed. The book is being read, the controls are being operated, and the flight is proceeding. This is not a miracle. This is a skill.

The skill is not "reading manuals." The skill is not "operating aircraft." The skill is doing both simultaneously, under time pressure, without panicking. It is the ability to encounter a problem you have never seen before, in a system you partly understand, using a tool you are learning as you use it, and to produce a working outcome before the plane runs out of fuel.

That skill has no name in pilot school because pilot school doesn't teach it. Software engineering doesn't teach it either. It is acquired exclusively by doing it — by being put in the cockpit, handed the book, and told that the destination is 3,000 miles away and the fuel is good for 2,800. You figure out the rest in the air. That's the job. The book will help. The book will also be wrong. You'll know which parts when you need them.

> Every software engineer who has ever shipped anything learned to fly while flying. The ones who waited until they felt ready never left the ground.
