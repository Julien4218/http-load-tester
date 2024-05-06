package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func Init(filepath string) (*InputConfig, error) {
	config := &InputConfig{}

	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("a valid config filepath is required, please specify an input config file like `--config my-config.yaml`, detail:%s", err)
	}

	err = yaml.Unmarshal(file, config)
	if err != nil {
		return nil, fmt.Errorf("a valid YAML config filepath is required, please specify an input config file like `--config my-config.yaml`, detail:%s", err)
	}

	return config, nil
}
