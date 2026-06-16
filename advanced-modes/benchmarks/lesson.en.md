---
id: 8eef624a-e41c-4965-9ed5-93cc56a7b54f
title: "Benchmarks"
estimated_minutes: 10
---

# Benchmarks

Go has built-in support for benchmarks — measuring code performance.

## go test -bench

```go
func BenchmarkSum(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Sum(1, 2)
    }
}
```

Output:
```
BenchmarkSum-8    1000000000    0.29 ns/op    0 B/op    0 allocs/op
```

- **ns/op** — nanoseconds per operation
- **B/op** — bytes allocated per operation
- **allocs/op** — number of allocations per operation

## How it's used in challenges

```yaml
# challenge.yaml
test_mode: bench
bench_threshold:
  ns_op_max: 1000      # no slower than 1µs/op
  allocs_max: 0        # zero allocations
```

Student must write code that meets performance thresholds.

## Typical tasks

- "Optimize to 0 allocations" (use sync.Pool, pre-allocate slices)
- "Reduce time to O(1)" (use map instead of linear search)
- "Reduce allocations" (strings.Builder instead of concatenation)

## Try it

Optimize string concatenation to zero allocations:

> [!challenge] zero-alloc-concat
