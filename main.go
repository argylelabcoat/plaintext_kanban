package main

import (
	"argylelabcoat/plaintext/kanban/db"
	"argylelabcoat/plaintext/kanban/renderers/html"
	"os"
)

func main() {

	demo, err := db.CreateProject("./Projects", "Demo")
	if nil != err {
		panic(err)
	}

	_, err = demo.AddTask("create_models.md", "Create Models", "Create Data Models", db.Statuses[db.Active])

	if nil != err {
		panic(err)
	}

	_, err = demo.AddTask("docu_design.md", "Document Design", "Create Documentation for Design", db.Statuses[db.Backlog])

	if nil != err {
		panic(err)
	}

	hr := &html.HtmlRenderer{}

	var t2bytes []byte
	t2bytes, err = hr.RenderProject(demo)

	os.WriteFile("test.html", t2bytes, 0655)
}
