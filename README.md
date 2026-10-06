# TDay

TDay, pronounced as _"**Today**"_, is a simple and striaghtforward command-line
tool for task management.

_TDay was built for my personal use for keep tracking of what I need done for
the day._

## Features

1. Initializes the global configuration directory `~/.tday` with a prepared .env
   file
2. Interactive task creation
3. Interactive task updates
4. Task completion
5. Task retrieval (all tasks)
6. Task deletion
7. Structured output

## Configuration

TDay's configuration directory, `~/.tday`, holds environment variables that TDay
relies on are stored in the configuration directory. This specifically allows
enables global usage.

- `~/.tday` directory layout document (coming soon): [/docs/tday.md](./docs/tday.md)

## Infrastructure

| Layer   | Tool                  |
| ------- | --------------------- |
| Main    | Go, Cobra             |
| Tooling | golangci-lint, Go     |
| UI      | PostgreSQL (Supabase) |

## Requirements

- Go 1.27.1 or later

## Quick Start

Clone the repo & build the binary:

```bash
git clone https://github.com/yuriongit/tday.git
cd tday
go build
go install
```

Start:

```bash
tday help
```

Init TDay and create a task:

```bash
tday init
tday new
```

## Images

To see a preview of TDay, view the [/docs/preview.md](./docs/preview.md) document.

_Images directory - [/.github/images](./github/images)_

## Docs

- [/docs/architecture.md](./docs/architecture.md)
- [/docs/planned.md](./docs/planned.md)

## License

MIT
