package main

import (
	"os"

	"github.com/elastic/beats/metricbeat/beater"

	// Uncomment the following line to include all official metricbeat module and metricsets
	//_ "github.com/elastic/beats/metricbeat/include"

	// Make sure all your modules and metricsets are linked here
	_ "{{cookiecutter.beat_path}}/{{cookiecutter.beat}}/module/{{cookiecutter.module}}/{{cookiecutter.metricset}}"

	"github.com/elastic/beats/libbeat/beat"
)

var Name = "{{cookiecutter.beat}}"

func main() {
	if err := beat.Run(Name, "", beater.New()); err != nil {
		os.Exit(1)
	}
}
