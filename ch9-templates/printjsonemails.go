/**
 * PrintJSONEmails
 */
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"text/template"
)

type Person struct {
	Name   string
	Emails []string
}

const templ = `{"Name": "{{- .Name -}}", "Emails": [
{{- range $index, $elmt := .Emails -}}
    {{- if $index -}}
        , "{{- $elmt -}}"
    {{- else -}}
         "{{- $elmt -}}"
    {{- end -}}
{{- end -}}
] }`

func main() {
	person := Person{
		Name: "jan",
		Emails: []string{"jan@newmarch.name",
			"jan.newmarch@gmail.com"},
	}
	t := template.New("Person template")
	t, err := t.Parse(templ)
	checkError(err)
	err = t.Execute(os.Stdout, person)
	checkError(err)

	// check via validity json package
	var b bytes.Buffer
	err = t.Execute(&b, person)
	checkError(err)
	if json.Valid(b.Bytes()) {
		fmt.Println("\nvalid json")
	}
}
func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error ", err.Error())
	}
}
