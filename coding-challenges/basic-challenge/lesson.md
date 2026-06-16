---
id: l1000004-0000-0000-0000-000000000004
title: "Первый challenge"
estimated_minutes: 10
---

# Первый challenge

Challenge — задача с автоматической проверкой. Студент пишет код, жмёт **Submit**, тесты автора запускаются скрыто, результат — pass/fail.

## Отличие от playground

| | Playground | Challenge |
|---|---|---|
| Цель | Эксперимент, "поиграть" | Решить задачу |
| Кнопка | Run | Run + Submit |
| Тесты | Нет | Авторские (hidden) |
| Вердикт | Просто вывод | Pass / Fail |
| Файлы | Все редактируемые | Есть readonly (тесты) |

## Как работает

1. Студент видит **template** — стартовый код с TODO
2. Пишет решение
3. **Run** — запускает его код (видит свой вывод)
4. **Submit** — `go test` с авторскими тестами → pass/fail
5. При fail — видит какие тесты упали (без исходника тестов)

## Попробуйте

Реализуйте функцию `Sum(a, b int) int`:

> [!challenge] implement-sum

## Структура challenge на диске

```
challenges/
└── implement-sum/
    ├── challenge.yaml      ← метаданные
    ├── solution.go         ← template (студент видит)
    ├── solution_test.go    ← тесты (hidden)
    └── solution/
        └── solution.go     ← эталонное решение (hidden)
```
