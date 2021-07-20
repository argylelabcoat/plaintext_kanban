package renderers

import "argylelabcoat/plaintext/kanban/db"

type Renderer interface {
	RenderProject(project *db.Project) ([]byte, error)
	RenderTask(task *db.Task) ([]byte, error)
}
