---
id: 8fe56b96-02c0-441f-aaf6-392d5c7e6d76
title: "Race Detector"
estimated_minutes: 10
---

# Race Detector

Go has a built-in data race detector. This is a unique feature — no other mainstream language has an equivalent out of the box.

## What is a data race

A data race occurs when two goroutines access the same variable concurrently and at least one of them is writing. The result is unpredictable.

```go
// DATA RACE:
var counter int

go func() { counter++ }()
go func() { counter++ }()
// counter could be 1 or 2 — undefined behavior
```

## How -race works

`go test -race` instruments the code — adds checks on every memory access. If a race is detected — the test fails with a detailed stack trace of both accesses.

**Cost:** memory ×5-10, speed ×2-10. That's why the race detector is used during testing, not in production.

## How it's used in challenges

```yaml
# challenge.yaml
test_mode: test_race    # go test -race
limits:
  memory_mb: 2048       # race detector is memory-hungry
```

Verdict: if the race detector finds a race — fail (even if tests "pass").

## Try it

Write a thread-safe cache without data races:

> [!challenge] safe-cache
