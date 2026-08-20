# Task Tracker CLI

A small command-line task manager written in Go. It stores tasks in a local JSON file and supports creating, updating, deleting, listing, and changing task status.

This project is an implementation of the [roadmap.sh Task Tracker project](https://roadmap.sh/projects/task-tracker).

## Features

- Add tasks with an automatically assigned ID
- Update a task's description
- Delete tasks
- Mark tasks as in progress or done
- List all tasks or filter them by status
- Persist tasks in `task.json` in the current working directory
- Create the storage file automatically when it does not exist
- Use only the Go standard library

## Requirements

- Go 1.26.5 or later

## Installation

Clone the repository, enter the project directory, and build the executable:

```bash
git clone https://github.com/ubpat16/task-tracker.git
cd task-tracker
go build -o task-cli .
```

You can then run the application with `./task-cli`. Alternatively, replace `./task-cli` in the examples below with `go run .` to run it without building first.

## Usage

The CLI accepts positional arguments. Put descriptions containing spaces inside quotes.

### Add a task

```bash
./task-cli add "Buy groceries"
```

Example output:

```text
Task Created:  1
```

New tasks start with the `todo` status.

### Update a task

```bash
./task-cli update 1 "Buy groceries and cook dinner"
```

### Delete a task

```bash
./task-cli delete 1
```

### Change a task's status

Mark a task as in progress:

```bash
./task-cli mark-in-progress 1
```

Mark a task as done:

```bash
./task-cli mark-done 1
```

### List tasks

List every task:

```bash
./task-cli list
```

Filter tasks by status:

```bash
./task-cli list todo
./task-cli list in-progress
./task-cli list done
```

Example listing:

```text
ID   TASK              Status
--   ----              ----
1    Buy groceries     todo
2    Write the README  in progress
```

## Command Reference

| Command                                | Description                           |
| -------------------------------------- | ------------------------------------- |
| `task-cli add "<description>"`         | Add a new task                        |
| `task-cli update <id> "<description>"` | Update a task's description           |
| `task-cli delete <id>`                 | Delete a task                         |
| `task-cli mark-in-progress <id>`       | Mark a task as in progress            |
| `task-cli mark-done <id>`              | Mark a task as done                   |
| `task-cli list`                        | List all tasks                        |
| `task-cli list todo`                   | List tasks that have not been started |
| `task-cli list in-progress`            | List tasks currently in progress      |
| `task-cli list done`                   | List completed tasks                  |

## Data Storage

Tasks are saved to `task.json` in the directory from which the CLI is run. The file is created automatically and contains a JSON array similar to this:

```json
[
  {
    "id": 1,
    "description": "Buy groceries",
    "status": "todo",
    "created_at": "2026-08-20T12:00:00+01:00",
    "updated_at": "0001-01-01T00:00:00Z"
  }
]
```

The supported stored status values are `todo`, `in progress`, and `done`. Updating a description or status sets the task's `updated_at` timestamp.

## Project Structure

```text
.
├── cmd/                 # CLI command handlers
├── internal/fs/         # JSON-backed task storage
├── main.go              # Argument routing and application entry point
└── go.mod               # Go module definition
```

## Development

Format, vet, and build the project with:

```bash
gofmt -w main.go cmd/*.go internal/fs/*.go
go vet ./...
go build ./...
```
