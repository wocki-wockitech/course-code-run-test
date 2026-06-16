---
id: d1354dd6-d8cd-4da9-b0d3-4c370afa4a96
title: "First Challenge"
estimated_minutes: 10
---

# First Challenge

A challenge is a task with automatic grading. Student writes code, clicks **Submit**, author's tests run behind the scenes, result is pass/fail.

## Difference from playground

| | Playground | Challenge |
|---|---|---|
| Goal | Experiment, "play around" | Solve a task |
| Button | Run | Run + Submit |
| Tests | None | Author's (hidden) |
| Verdict | Just output | Pass / Fail |
| Files | All editable | Some readonly (tests) |

## How it works

1. Student sees **template** — starter code with TODO
2. Writes their solution
3. **Run** — runs their code (sees own output)
4. **Submit** — `go test` with author's tests → pass/fail
5. On fail — sees which tests failed (without test source code)

## Try it

Implement the function `Sum(a, b int) int`:

> [!challenge] implement-sum

## Challenge structure on disk

```
challenges/
└── implement-sum/
    ├── challenge.yaml      ← metadata
    ├── solution.go         ← template (student sees)
    ├── solution_test.go    ← tests (hidden)
    └── solution/
        └── solution.go     ← reference solution (hidden)
```
