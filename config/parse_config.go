
package config

import (
	"os"
	"fmt"

	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type volumeMount struct {
	Name string 
	ReadOnly string `yaml:"readOnly"`
	MountPath string `yaml:"mountPath"`
}

type container struct {
	Name string
	Image string
	VolumeMounts []volumeMount `yaml:"volumeMounts"`
}


// LoadConfig loads from podConfig yaml file and returns the structure
// func LoadConfig(podConfig string) (*map[string]interface{}, error) {
func LoadConfig(podConfig string) (*container, error) {
	stat, err := os.Stat(podConfig)
	if err != nil {
		return nil, err
	}

	if stat.IsDir() {
		return nil, fmt.Errorf("podConfig cannot be a dir - %s", podConfig)
	}

	yamlFile, err := ioutil.ReadFile(podConfig)
    if err != nil {
        return nil, fmt.Errorf("Error reading YAML file: %s\n", err)
    }

	nc := container{}

	err = yaml.Unmarshal(yamlFile, &nc)
	if err != nil {
        return nil, fmt.Errorf("Error reading YAML file: %s\n", err)
    }
	
	return &nc, nil
}
