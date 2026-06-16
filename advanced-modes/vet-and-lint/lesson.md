---
id: 57a9fc52-fa43-4a69-b0e4-b2db05d3c2d5
title: "Vet и Lint"
estimated_minutes: 8
---

# Vet и Lint

Два инструмента статического анализа, которые ловят ошибки ДО запуска программы.

## go vet

Встроенный в Go — ищет распространённые ошибки:
- Неправильный формат в `fmt.Printf`
- Невозможные условия (unreachable code)
- Shadow variables
- Копирование sync.Mutex

```go
// go vet найдёт ошибку:
fmt.Printf("%d", "hello")  // %d ожидает int, получил string
```

## golangci-lint

Мета-линтер — запускает десятки проверок одновременно:
- `errcheck` — необработанные ошибки
- `ineffassign` — присвоение в переменную которая не читается
- `staticcheck` — advanced static analysis
- `gocritic` — стилистические замечания

## Как используется в challenge

```yaml
# challenge.yaml
test_mode: lint    # вместо обычного test
```

Вердикт: exit 0 = код чистый (pass), non-0 = есть замечания (fail).

## Попробуйте

Исправьте все замечания линтера:

> [!challenge] fix-lint-issues
