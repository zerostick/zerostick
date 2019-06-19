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
	assets := http.Dir(filepath.Join(cwd, "zerostick_web"))
	if err := vfsgen.Generate(assets, vfsgen.Options{
		Filename:     "build/templates_vfsdata.go",
		PackageName:  "assets",
		BuildTags:    "deploy_build",
		VariableName: "Assets",
	}); err != nil {
		log.Fatalln(err)
	}
}
