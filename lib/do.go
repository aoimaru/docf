package lib

import (
	"os"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

func Do(file_path string) (string, error) {
	fp, err := os.Open("./Dockerfile")
	if err != nil {
		return "", err
	}
	res, err := parser.Parse(fp)
	if err != nil {
		return "", err
	}
	ast := res.AST.Dump()
	return ast, nil
}
