package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"b01901143.git/sample-go-app/config"
)

type options struct {
	configPath string
	dynamic    bool
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
	flag.BoolVar(&o.dynamic, "dynamic", false, "Load config.yaml to dynamic structure.")
	flag.Parse()
	return o
}

func printHelp() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	o := gatherOptions()

	if err := o.Validate(); err != nil {
		fmt.Println("Invalid options: ", err)
		printHelp()
		os.Exit(1)
	}

	conf, err := config.LoadConfig(o.configPath, o.dynamic)
	if err != nil {
		fmt.Println("Error loading config: ", err)
		printHelp()
		os.Exit(1)
	}

	conf.PrintStruct()

}
