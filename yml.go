package csql

import (
	"fmt"

	"github.com/goccy/go-yaml/ast"

	goyaml "github.com/goccy/go-yaml"
)

type YmlFile struct {
	path string
	file *ast.File
}

func (y *YmlFile) ReadYmlFile() (any, error) {
	if y.path == "" {
		return nil, fmt.Errorf("path is empty")
	}
	yamlPath, err := goyaml.PathString(y.path)
	if err != nil {
		return nil, fmt.Errorf("failed to parse path %s: %w", y.path, err)
	}

	node, err := yamlPath.FilterFile(y.file)
	if err != nil {
		return nil, err
	}

	var value interface{}
	if err := goyaml.Unmarshal([]byte(node.String()), &value); err != nil {
		return nil, err
	}

	return value, nil
}
