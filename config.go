package vivliothiki

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// ReadYamlConfig is a function to read yaml file based
func ReadYamlConfig(path, filename string, dest interface{}) error {
	src := fmt.Sprintf("%s/%s.%s.yml", path, filename, GetEnv())

	yamlFile, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, dest)
	if err != nil {
		return err
	}
	return nil
}
