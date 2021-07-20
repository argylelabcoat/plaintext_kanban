# plaintext_kanban

This is a simple cli tool for rendering a "task board" from a directory structure of markdown files.


## Usage 

```sh
./kanban [project path] [output html path]
```

## Data Format
Each file equates to a task.  Where the file is within the heirarchy indicates its status.

Within each task file, a header is used for determining additional metadata.  Currently it looks like this:

```
---
Title: Document Design
AssignedTo: @argylelabcoat
---
```

The directory structure looks like this:

* ProjectName/
  * backlog/
    * task_a.md
    * task_b.md
  * active/
    * task_a.md
    * task_b.md
  * verify/
    * task_a.md
    * task_b.md
  * done/
    * task_a.md
    * task_b.md

## Building

Currently set to use Go 1.16

build simply with:

```sh
go build
```
