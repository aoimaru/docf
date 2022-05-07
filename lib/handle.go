package lib

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

func Replace(contents string) {

}

type Result struct {
	Res []string `json:res`
}

type OutPut struct {
	Results []Result `json:results`
}

func GetRun(file_path string, flag bool) OutPut {
	fp, err := os.Open(file_path)
	if err != nil {
		log.Fatal(err)
	}
	res, err := parser.Parse(fp)
	if err != nil {
		log.Fatal(err)
	}
	var output OutPut
	for _, child := range res.AST.Children {
		if child.Value == "RUN" {
			originals := strings.Split(child.Original, " ")
			contents := make([]string, 0)
			for _, original := range originals {
				// fmt.Println(original)
				if flag {
					content := EscMeta(original)
					contents = append(contents, content...)
				} else {
					content := Meta(original)
					contents = append(contents, content...)
				}
			}
			// fmt.Println("contents", contents)
			res := Result{contents}
			output.Results = append(output.Results, res)
		}
	}
	return output
}

func ListFiles(root string) ([]string, error) {
	var fps []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// ディレクトリは除く
		if !info.IsDir() {
			fps = append(fps, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return fps, nil
}
