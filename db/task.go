package db

import (
	"fmt"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/text"
	"os"
	"path"
	"strings"
)

type Task struct {
	Title       string
	AssignedTo  string
	Description string
	Status      string
	FileName    string
	ProjectPath string
}

var markdown goldmark.Markdown
var FilePerm os.FileMode = 0644

func init() {
	markdown = goldmark.New(
		goldmark.WithExtensions(
			meta.Meta,
		),
	)
}

func (task *Task) Save() (err error) {
	metalinefmt := "%v: %v\n"
	metadelim := "---\n"
	builder := &strings.Builder{}
	builder.WriteString(metadelim)
	builder.WriteString(fmt.Sprintf(metalinefmt, "Title", task.Title))
	builder.WriteString(fmt.Sprintf(metalinefmt, "AssignedTo", task.AssignedTo))
	builder.WriteString(metadelim)

	builder.WriteString("\n")
	builder.WriteString(task.Description)

	filepath := task.GetPath()
	err = os.WriteFile(filepath, []byte(builder.String()), FilePerm)

	return
}

func (task *Task) ChangeStatus(newstatus string) error {
	oldpath := task.GetPath()
	task.Status = newstatus
	newpath := task.GetPath()
	return os.Rename(oldpath, newpath)
}

func (task *Task) GetPath() (filepath string) {
	filepath = path.Join(path.Join(task.ProjectPath, task.Status), task.FileName)
	return
}

func LoadTaskFromFile(filepath string) (task *Task, err error) {
	statpath, filename := path.Split(filepath)
	projpath, status := path.Split(statpath)

	source, err := os.ReadFile(filepath)
	if nil != err {
		return
	}
	doc := markdown.Parser().Parse(text.NewReader(source))
	title, ok := doc.OwnerDocument().Meta()["Title"]
	if !ok {
		err = NewErrorTaskMissingMeta(projpath, "Title")
	}

	assignedto, ok := doc.OwnerDocument().Meta()["AssignedTo"]
	if !ok {
		assignedto = ""
	}

	parts := strings.Split(string(source), "---\n")
	body := parts[len(parts)-1]

	task = &Task{
		Title:       title.(string),
		AssignedTo:  assignedto.(string),
		Description: body,
		FileName:    filename,
		Status:      status,
		ProjectPath: projpath,
	}

	return
}

func TasksFromDir(dirpath string) (tasks []Task, err error) {
	var entries []os.DirEntry
	tasks = []Task{}
	entries, err = os.ReadDir(dirpath)
	for _, entry := range entries {
		var fi os.FileInfo
		var task *Task
		fi, err = entry.Info()
		if nil != err {
			return
		}
		if !fi.IsDir() && fi.Size() > 0 {
			task, err = LoadTaskFromFile(path.Join(dirpath, fi.Name()))
			if nil != err {
				return
			}
			tasks = append(tasks, *task)
		}
	}
	return
}
