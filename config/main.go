package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func Init(filepath string) (*InputConfig, error) {

	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("a valid config filepath is required, please specify an input config file like `--config my-config.yaml`, detail:%s", err)
	}
	return InitWithContent(file)
}

func InitWithContent(bytes []byte) (*InputConfig, error) {
	config := &InputConfig{}

	err1 := yaml.Unmarshal(bytes, config)
	if err1 != nil {
		err2 := json.Unmarshal(bytes, config)
		if err2 != nil {
			return nil, fmt.Errorf("a valid YAML or JSON config is required, the content provided wasn't valid, detail:%s and detail:%s", err1, err2)
		}
	}

	return config, nil
}
