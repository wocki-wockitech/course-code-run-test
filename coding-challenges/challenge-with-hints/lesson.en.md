---
id: l1000005-0000-0000-0000-000000000005
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
