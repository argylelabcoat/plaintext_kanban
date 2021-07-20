package main

import (
	"argylelabcoat/plaintext/kanban/db"
	"argylelabcoat/plaintext/kanban/renderers/html"
	"github.com/spf13/cobra"
	"os"
)

func render(projpath, outpath string) {
	demo, err := db.LoadProjectDir("./Projects/Demo")
	if nil != err {
		panic(err)
	}

	hr := &html.HtmlRenderer{}

	var t2bytes []byte
	t2bytes, err = hr.RenderProject(demo)

	os.WriteFile("test.html", t2bytes, 0655)
}

func main() {
	var renderCmd = &cobra.Command{
		Use:   "render [path to project] [path to output html]",
		Short: "Create static HTML page representing the current state of a project",
		Long: `echo things multiple times back to the user by providing
a count and a string.`,
		Args: cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			render(args[0], args[1])
		},
	}
	var rootCmd = &cobra.Command{Use: "kanban"}
	rootCmd.AddCommand(renderCmd)
	rootCmd.Execute()
}
