package html

import (
	"argylelabcoat/plaintext/kanban/db"
	"bytes"
	"github.com/yuin/goldmark"
)

type HtmlRenderer struct{}

func (hr *HtmlRenderer) RenderTask(task *db.Task) (data []byte, err error) {
	var buf bytes.Buffer
	var descbuf bytes.Buffer
	if err = goldmark.Convert([]byte(task.Description), &descbuf); err != nil {
		return
	}
	task.Description = descbuf.String()
	err = taskTemplate.Execute(&buf, task)
	if nil != err {
		return
	}
	data = buf.Bytes()
	return
}

func (hr *HtmlRenderer) RenderProject(project *db.Project) (data []byte, err error) {
	var buf bytes.Buffer

	var tasksBuf bytes.Buffer
	for status := db.Backlog; status <= db.Done; status++ {
		columnName := db.Statuses[status]
		tasks := project.Board[columnName]
		var colbuf bytes.Buffer
		for _, task := range tasks {
			var taskB []byte
			taskB, err = hr.RenderTask(&task)
			colbuf.Write(taskB)
		}
		var coltxt string = colbuf.String()
		column := &renderableColumn{
			Title: columnName,
			Tasks: coltxt,
		}
		err = columnTemplate.Execute(&tasksBuf, column)
		if nil != err {
			return
		}
	}

	rProject := &renderableProject{
		Name:    project.Name,
		Columns: tasksBuf.String(),
	}

	err = projectTemplate.Execute(&buf, rProject)
	if nil != err {
		return
	}
	data = buf.Bytes()
	return
}
