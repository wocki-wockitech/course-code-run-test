---
id: 505e3396-7091-4925-8d63-00b70be8920a
title: "SQL Challenge"
estimated_minutes: 8
---

# SQL Challenge

A separate challenge type — student writes an SQL query, the platform compares the result with a reference.

## How it works

1. Author provides `setup.sql` — schema + seed data
2. Student writes a query in `solution.sql`
3. Platform executes both (reference + student) on an ephemeral DB
4. Compares result sets

## Comparison modes

| Mode | What it checks |
|------|---------------|
| `exact_rows` | Exact match (order matters) |
| `set_match` | Multiset (order doesn't matter) |
| `column_subset` | Student returned at least the required columns |

## Try it

Write a query returning users older than 18:

> [!challenge] users-over-18
