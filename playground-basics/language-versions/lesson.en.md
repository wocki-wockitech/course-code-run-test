---
id: 3b5d518f-0f1b-40c5-a63e-f11eecf94935
title: "Languages and Versions"
estimated_minutes: 5
---

# Languages and Versions

Playground supports multiple languages. Each language has a set of versions managed through the image catalog.

## Supported languages

| Language | Entry file | Under the hood |
|----------|-----------|----------------|
| Go | `main.go` | `go build` + run binary |
| PHP | `main.php` | `php main.php` |
| Python (planned) | `main.py` | `python main.py` |
| JavaScript (planned) | `main.js` | `node main.js` |

## Versions

Each language has:
- **Current versions** — shown by default in the selector
- **Archived versions** — available on demand ("+ Older versions")
- **Default version** — selected if the author doesn't specify one

## Fixing version in a course

Authors can fix the version for a sandbox in a lesson:

```markdown
> [!sandbox] go 1.23
```

Student cannot switch — learning on a specific version.

## Try different languages

> [!sandbox] go-playground

> [!sandbox] php-playground
