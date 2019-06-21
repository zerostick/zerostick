// +build ignore

package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/shurcooL/vfsgen"
)

func main() {
	var cwd, _ = os.Getwd()
	templates := http.Dir(filepath.Join(cwd, "zerostick_web/templates"))
	if err := vfsgen.Generate(templates, vfsgen.Options{
		Filename:     "build/templates_vfsdata.go",
		PackageName:  "templates",
		BuildTags:    "deploy_build",
		VariableName: "Templates",
	}); err != nil {
		log.Fatalln(err)
	}
}
