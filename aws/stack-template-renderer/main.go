package main

import (
	"os"
	"text/template"

	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/halt"
)

func main() {
	cloudformationTemplate := template.New("root.gotemplate")

	cloudformationTemplate.Delims("{{{", "}}}")

	_, err :=
		cloudformationTemplate.ParseGlob("aws/stack-templates/*.gotemplate")

	if err != nil {
		halt.Panic(flaw.From(err))
	}

	err = cloudformationTemplate.Execute(os.Stdout, nil)

	if err != nil {
		halt.Panic(flaw.From(err))
	}
}
