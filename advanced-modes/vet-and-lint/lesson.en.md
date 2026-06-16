---
id: l1000007-0000-0000-0000-000000000007
title: "Vet and Lint"
estimated_minutes: 8
---

# Vet and Lint

Two static analysis tools that catch bugs BEFORE running the program.

## go vet

Built into Go — finds common mistakes:
- Wrong format verbs in `fmt.Printf`
- Impossible conditions (unreachable code)
- Shadow variables
- Copying sync.Mutex

```go
// go vet will find this:
fmt.Printf("%d", "hello")  // %d expects int, got string
```

## golangci-lint

Meta-linter — runs dozens of checks simultaneously:
- `errcheck` — unhandled errors
- `ineffassign` — assignment to unused variable
- `staticcheck` — advanced static analysis
- `gocritic` — style suggestions

## How it's used in challenges

```yaml
# challenge.yaml
test_mode: lint    # instead of regular test
```

Verdict: exit 0 = code is clean (pass), non-0 = issues found (fail).

## Try it

Fix all linter warnings:

> [!challenge] fix-lint-issues
