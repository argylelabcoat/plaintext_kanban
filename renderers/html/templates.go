package html

import (
	"text/template"
)

const taskTemplateText = `
<div class="card block">
	<div class="card-content">
		<div class="has-background-info-light">
		<p class="title is-5">{{ .Title }}</p>
		<p class="subtitle is-7">Assigned To: {{ .AssignedTo }}</p>
		</div>
		<div class="content">
		{{ .Description }}
		</div>
	</div>
</div>
`

const columnTemplateText = `
<div class="column">
	<div class="notification is-info title">
		{{ .Title }}
	</div>
	{{ .Tasks }}
</div>
`

const projectTemplateText = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>{{ .Name }}</title>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@0.9.3/css/bulma.min.css">
  </head>
  <body>
  <section class="section">
    <div class="container has-background-primary-light p-1">
        <div class="columns is-mobile is-multiline is-centered">
		{{ .Columns }}
		</div>
    </div>
  </section>
  </body>
</html>
`

var taskTemplate *template.Template

var columnTemplate *template.Template
var projectTemplate *template.Template

func init() {
	var err error
	taskTemplate, err = template.New("Task").Parse(taskTemplateText)
	if nil != err {
		panic(err)
	}
	columnTemplate, err = template.New("Column").Parse(columnTemplateText)
	if nil != err {
		panic(err)
	}
	projectTemplate, err = template.New("Project").Parse(projectTemplateText)
	if nil != err {
		panic(err)
	}
}

type renderableColumn struct {
	Title string
	Tasks string
}

type renderableProject struct {
	Name    string
	Columns string
}
