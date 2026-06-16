---
id: 8eef624a-e41c-4965-9ed5-93cc56a7b54f
title: "Бенчмарки"
estimated_minutes: 10
---

# Бенчмарки

Go имеет встроенную поддержку бенчмарков — измерение производительности кода.

## go test -bench

```go
func BenchmarkSum(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Sum(1, 2)
    }
}
```

Вывод:
```
BenchmarkSum-8    1000000000    0.29 ns/op    0 B/op    0 allocs/op
```

- **ns/op** — наносекунд на операцию
- **B/op** — байт аллоцировано на операцию
- **allocs/op** — количество аллокаций на операцию

## Как используется в challenge

```yaml
# challenge.yaml
test_mode: bench
bench_threshold:
  ns_op_max: 1000      # не медленнее 1µs/op
  allocs_max: 0        # zero allocations
```

Студент должен написать код, удовлетворяющий порогам производительности.

## Типичные задачи

- «Оптимизируй до 0 аллокаций» (используй sync.Pool, pre-allocate slices)
- «Снизь время до O(1)» (используй map вместо линейного поиска)
- «Уменьши аллокации» (strings.Builder вместо конкатенации)

## Попробуйте

Оптимизируйте функцию конкатенации строк до 0 аллокаций:

> [!challenge] zero-alloc-concat
