---
id: 6ac4a69f-355e-470c-830a-f94d2acddb79
title: "Hello, Playground"
estimated_minutes: 5
---

# Hello, Playground

Playground is a built-in sandbox for running code right in the browser. Code executes in an isolated container (gVisor) — safe and without installing anything on your machine.

## How it works

1. Write code in the editor
2. Click **Run**
3. See output in real-time (SSE stream)

Under the hood: your code is sent to the server → an ephemeral Kubernetes pod is created with RuntimeClass gVisor → code is compiled and executed → stdout is streamed back.

## Try it

> [!sandbox] go

Change the program — print your name. Click Run.

## Sandbox limitations

- **No network** — can't download packages or call external APIs
- **Timeout** — program is killed after 10 seconds
- **Memory** — limited (1Gi in dev, less in prod)
- **Files** — up to 20 files, 64KB each

These limitations are by design. Sandbox is for learning, not production development.

## What you see in output

Playground shows execution stages:
- **compiling** — Go compiles your code
- **running** — program is executing
- Final status: `done` (exit 0), `timeout` (limit exceeded), or compile error
