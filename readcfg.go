package cfgreader

import (
	"gopkg.in/yaml.v3"
	"os"
)

func ReadCfg(filename string, cfg any) error {
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(yamlFile, cfg)
}
