
package main

import (
	"flag"
	"fmt"
	"errors"
	"gopkg.in/yaml.v2"

	"b01901143.git/sample-go-app/config"
)

type options struct {
	configPath    string
}

func (o *options) Validate() error {
	if o.configPath == "" {
		return errors.New("required flag --config-path was unset")
	}

	return nil
}

func gatherOptions() options {
	o := options{}
	flag.StringVar(&o.configPath, "config-path", "", "Path to config.yaml.")
	flag.Parse()
	return o
}

func main() {
	o := gatherOptions()

	if err := o.Validate(); err != nil {
		fmt.Println("Invalid options: %v", err)
	}

	conf, err := config.LoadConfig(o.configPath)
	if err != nil {
		fmt.Println("Error loading config.")
	}
	// fmt.Println("main: ", *conf)

	d, err := yaml.Marshal(conf)
	fmt.Println(string(d))

}
