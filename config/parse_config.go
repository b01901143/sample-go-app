package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// configuration interface for printing
type ConfigStruct interface {
	PrintStruct()
}

// explicitly define the tags for each field
type configuration struct {
	ApiVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   metadata `yaml:"metadata"`
	Spec       spec     `yaml:"spec"`
}

type metadata struct {
	Name   string `yaml:"name"`
	Labels label  `yaml:"labels"`
}

type label struct {
	Name string `yaml:"name"`
}

type spec struct {
	Volumes    []volume    `yaml:"volumes"`
	Containers []container `yaml:"containers"`
}

type volume struct {
	Name   string `yaml:"name"`
	Secret secret `yaml:"secret"`
}
type secret struct {
	SecretName string `yaml:"secretName"`
}
type volumeMount struct {
	Name      string `yaml:"name"`
	ReadOnly  string `yaml:"readOnly"`
	MountPath string `yaml:"mountPath"`
}

type container struct {
	Name         string        `yaml:"name"`
	Image        string        `yaml:"image"`
	VolumeMounts []volumeMount `yaml:"volumeMounts"`
}

type dynamicStruct map[string]interface{}

func (conf configuration) PrintStruct() {
	d, _ := yaml.Marshal(conf)
	fmt.Println(string(d))
}

func (conf dynamicStruct) PrintStruct() {
	d, _ := yaml.Marshal(conf)
	fmt.Println(string(d))
}

// LoadConfig loads from podConfig yaml file and returns the structure
func LoadConfig(podConfig string, dynamic bool) (ConfigStruct, error) {
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

	if dynamic {
		nc := dynamicStruct{}
		err = yaml.Unmarshal(yamlFile, &nc)
		if err != nil {
			return nil, fmt.Errorf("Error reading YAML file: %s\n", err)
		}

		return &nc, nil
	} else {
		nc := configuration{}
		err = yaml.Unmarshal(yamlFile, &nc)
		if err != nil {
			return nil, fmt.Errorf("Error reading YAML file: %s\n", err)
		}

		return &nc, nil
	}

	return nil, nil
}
