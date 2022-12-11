package yaml

import (
	"github.com/finishy1995/codegenerator/dataloader"
	"gopkg.in/yaml.v3"
)

type Loader struct {
}

func NewYamlLoader() *Loader {
	return &Loader{}
}

func (l *Loader) LoadFromFile(filePath string) (map[string]interface{}, error) {
	// unsupported file load
	return nil, dataloader.ErrorNoNeedLog
}

func (l *Loader) LoadFromBytes(filePath string, bytes []byte) (map[string]interface{}, error) {
	// if you have many yaml files, you can use filePath as a prefix to map key, such as ${filePath}.
	result := make(map[string]interface{})
	err := yaml.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
