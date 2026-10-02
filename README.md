# TDay

A personal task-managing command-line tool for keep tracking of what I need done
for the day.

## Features

1. `tday init` initialize TDay's configuration.
2. `tday new` create a new task.
3. `tday list` list all tasks.

## Configuration

TDay has a configuration directory (`/.tday`.) within the home directory.
For example, environment variables that TDay relies on are stored in the
configuration directory. This specifically allows TDay to be used globally within
the terminal.

- TDay's Configuration: [/docs/tday-config.md](./docs/tday-config.md.md)

## Infrastructure

| Layer   | Tool                  |
| ------- | --------------------- |
| Main    | Go, Cobra             |
| Tooling | golangci-lint, Go     |
| Data    | PostgreSQL (Supabase) |

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
tday --help
```

Init TDay:

```bash
tday init
```

Create a task:

```bash
tday new
```

List all tasks:

```bash
tday ls
```

## Docs

- [/docs/architecture.md](./docs/architecture.md)
- [/docs/planned.md](./docs/planned.md)

## Images

- [/.github/images/](/.github/images)

## License

MIT
