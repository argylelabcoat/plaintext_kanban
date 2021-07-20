package db

import (
	"os"
	"path"
)

type Project struct {
	Name  string
	Board map[string][]Task
	Path  string
}
type StatusType int

const (
	Backlog StatusType = iota
	Active
	Verify
	Done
)

var Statuses = map[StatusType]string{
	Backlog: "backlog",
	Active:  "active",
	Verify:  "verify",
	Done:    "done",
}
var DirPerm os.FileMode = 0775

func (project *Project) AddTask(filename, title, description, status string) (task *Task, err error) {
	task = &Task{
		Title:       title,
		Description: description,
		Status:      status,
		FileName:    filename,
		ProjectPath: project.Path,
	}
	project.Board[status] = append(project.Board[status], *task)
	err = task.Save()
	return
}

func CreateProject(root, name string) (project *Project, err error) {
	projectRoot := path.Join(root, name)
	project = &Project{
		Name:  name,
		Path:  projectRoot,
		Board: map[string][]Task{},
	}
	for _, col := range Statuses {
		project.Board[col] = []Task{}
		colDir := path.Join(project.Path, col)
		err = os.MkdirAll(colDir, DirPerm)
		if err != nil {
			return
		}
	}
	return
}

func LoadProjectDir(projpath string) (project *Project, err error) {
	_, name := path.Split(projpath)
	project = &Project{
		Name:  name,
		Path:  projpath,
		Board: map[string][]Task{},
	}
	for _, col := range Statuses {
		colDir := path.Join(project.Path, col)
		project.Board[col], err = TasksFromDir(colDir)
		if nil != err {
			return
		}
	}
	return
}
