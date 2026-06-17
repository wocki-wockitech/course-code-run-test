---
id: 3b5d518f-0f1b-40c5-a63e-f11eecf94935
title: "Языки и версии"
estimated_minutes: 5
---

# Языки и версии

Playground поддерживает несколько языков. Каждый язык имеет набор версий, управляемый через каталог образов.

## Поддерживаемые языки

| Язык | Entry file | Что под капотом |
|------|-----------|-----------------|
| Go | `main.go` | `go build` + запуск бинаря |
| PHP | `main.php` | `php main.php` |
| Python (planned) | `main.py` | `python main.py` |
| JavaScript (planned) | `main.js` | `node main.js` |

## Версии

Каждый язык имеет:
- **Актуальные версии** — показываются по умолчанию в селекторе
- **Архивные версии** — доступны по запросу ("+ Старые версии")
- **Версия по умолчанию** — выбирается если автор не указал конкретную

## Управление версиями в sandbox.yaml

Поле `version` поддерживает три режима:

### Без ограничений (все версии из runme)

```yaml
language: go
# version не указан → студент видит все доступные версии
```

> [!sandbox] go-playground

### Фиксация одной версии (селектор скрыт)

```yaml
language: go
version: "1.22"
# Скаляр = жёстко зафиксировано. Студент не может переключить.
```

Полезно для уроков, привязанных к конкретному поведению версии.

> [!sandbox] go-pinned

### Выбор из подмножества (первая = дефолт)

```yaml
language: go
version: ["1.21", "1.22", "1.23"]
# Массив = студент выбирает из этих. Первая — начальная.
```

Идеально для уроков про различия версий: «в 1.23 так работает, а если переключить на 1.21 — увидите разницу».

> [!sandbox] go-versions

## Read-only файлы

### Вся песочница read-only

```yaml
language: go
files_readonly: true
```

Студент видит код и запускает, но не может редактировать. Подходит для демонстрации эталонного решения:

> [!sandbox] go-readonly

### Отдельные файлы read-only

```yaml
language: go
files:
  - path: main.go
    readonly: false    # можно редактировать
  - path: example_test.go
    readonly: true     # только для чтения (замок на табе)
```

Полезно для задач где тесты или конфиги зафиксированы, а студент пишет решение:

> [!sandbox] go-mixed

## Скрытые файлы

```yaml
files:
  - path: internal/setup.go
    hidden: true   # участвует в компиляции, но студент не видит
```

Файлы с `hidden: true` отправляются на сервер при запуске, но не отображаются в UI. Полезно для инфраструктурного кода (go.mod, helper-функции) которые загромождают интерфейс.

## Другие языки

> [!sandbox] php-playground

## Полная схема sandbox.yaml

```yaml
id: uuid                    # auto-filled (studyme-action fix-ids)
language: go                # required: go | php | python | js
version: "1.22"            # string = locked | ["1.21","1.22"] = choice | omit = all
layout: single             # single | flat | tree (auto by file count)
actions: [run]             # [run] | [test] | [run, test]
mode: run                  # default action button
files_readonly: false      # все файлы read-only
allow_add_files: true      # можно ли создавать/удалять файлы
output_position: bottom    # side | bottom
show_tests: false          # показывать тестовые файлы студенту
files:                     # explicit file list (иначе auto-discover)
  - path: main.go
    readonly: false
    hidden: false
```
