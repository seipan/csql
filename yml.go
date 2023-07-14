package csql

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Yml struct {
	dns string `yaml:"dns"`
}

type YmlFile struct {
	path string
}

func (y *YmlFile) ReadYmlFile() (any, error) {
	yml := Yml{}
	b, err := os.ReadFile(y.path)
	if err != nil {
		return nil, err
	}

	yaml.Unmarshal(b, &yml)
	return yml, nil
}
