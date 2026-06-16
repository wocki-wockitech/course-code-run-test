---
id: 802fd2f0-a526-432b-9bd3-51da2b6dae2a
title: "Challenge with Hints"
estimated_minutes: 10
---

# Challenge with Hints

Authors can add hints — students reveal them one by one when stuck.

## Hints

In `challenge.yaml`:
```yaml
hints:
  - "Use sync.Mutex to protect shared state"
  - "Lock() before writing, Unlock() after"
  - "defer mu.Unlock() — idiomatic pattern"
```

Student sees a "Hint" button → reveals one at a time. This does NOT penalize (unlike code review), but can be considered in SR repetition (recall quality is lower if hints were used).

## Try it

Implement a thread-safe counter:

> [!challenge] thread-safe-counter
