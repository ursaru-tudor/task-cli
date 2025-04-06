# Task Tracker CLI

task-cli is a CLI application for tracking tasks and managing a to-do list written in Go. It stores all data in a local JSON file.

Designed for: https://roadmap.sh/projects/task-tracker

## Installation

### Compilation

To build task-cli you must first ensure you have git and go installed. Then run

```
git clone https://github.com/ursaru-tudor/task-cli.git

cd task-cli

go build
```

## Usage

### Add a task

```
> task-cli add "Task name"
```

**You can add mutliple tasks with a single call to task-cli by listing mutiple task names**

```
> task-cli add "Task 1" "Task 2" ...
```

### Update a task's name

```
> task-cli update <task-id> "New task name"
```

### Delete a task

```
> task-cli delete <task-id>
```


**You can delete mutliple tasks with a single call to task-cli by listing multiple task-ids**

```
> task-cli delete <task-id 1> <task-id 2> <task-id 3>
```

### Marking a task "in-progress" or "done"

```
> task-cli mark-in-progress <task-id>
```

```
> task-cli mark-done <task-id>
```

**You can mark mutliple tasks with a single call to task-cli by listing multiple task-ids, similar to delete**

### Listing tasks

```
> task-cli list
```

Will list all tasks.

You can provide one or multiple of 'todo', 'in-progress', or 'done' to only see tasks with at specific stages. For example:

```
> task-cli list todo in-progress
```

Will show all tasks that are not 'done'.

### See information on a single task

This will show information about one or more tasks by providing their task-ids

```
> task-cli info <task-id>
```

Help can be accessed at any time by running

```
> task-cli help
```