---
id: c5eb29fb-9d64-4f7f-bc8b-3bac0e653c15
title: "Multi-file Projects"
estimated_minutes: 7
---

# Multi-file Projects

Playground supports multiple files — like a real Go project with packages.

## Single file vs tree

| Layout | When to use |
|--------|------------|
| `single` | One main.go, quick example |
| `flat` | 2-3 files, switch via tabs |
| `tree` | Package structure, folders |

## Example: package + main

> [!sandbox] multi-pkg

Try creating a file `utils/math.go` with a function `Add(a, b int) int` and call it from `main.go`.

## How playground processes files

1. All files are sent to the server as `[]FileEntry`
2. Server places them at their paths in the working directory `/work/`
3. If no `go.mod` exists — it's created automatically (`module playground`)
4. `go build ./...` compiles the entire tree

## Read-only files

Authors can mark a file as `readonly: true` — student sees it but cannot edit. Useful for:
- Tests (in challenges)
- Ready-made infrastructure (main with server)
- Examples that shouldn't be modified
