
package config

import (
	"os"
	"fmt"

	"io/ioutil"
	"gopkg.in/yaml.v2"
)

// configuration interface for printing
type ConfigStruct interface {
	PrintStruct()
}

// explicitly define the tags for each field
type configuration struct {
	ApiVersion string 				`yaml:"apiVersion"`
	Kind string 					`yaml:"kind"`
	Metadata struct {
		Name string 				`yaml:"name"`
		Labels struct {
			Name string 			`yaml:"name"`
		}							`yaml:"labels"`

	}								`yaml:"metadata"`
	Spec struct {
		Volumes []struct {
			Name string 			`yaml:"name"`
			Secret struct {
				SecretName string 	`yaml:"secretName"`
			}						`yaml:"secret"`
		}							`yaml:"volumes"`
		Containers []container 		`yaml:"containers"`
	}								`yaml:"spec"`
}

type volumeMount struct {
	Name string 		`yaml:"name"`
	ReadOnly string 	`yaml:"readOnly"`
	MountPath string 	`yaml:"mountPath"`
}

type container struct {
	Name string 				`yaml:"name"`
	Image string 				`yaml:"image"`
	VolumeMounts []volumeMount 	`yaml:"volumeMounts"`
}

type dynamicStruct map[string]interface{}

func (conf configuration) PrintStruct() {
	d, _ := yaml.Marshal(conf)
	fmt.Println(string(d))
}

func (conf dynamicStruct) PrintStruct() {
	fmt.Println(conf)
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
    } else{
		nc := configuration{}
		err = yaml.Unmarshal(yamlFile, &nc)
		if err != nil {
	        return nil, fmt.Errorf("Error reading YAML file: %s\n", err)
	    }
		
		return &nc, nil
	}

	return nil, nil
}
