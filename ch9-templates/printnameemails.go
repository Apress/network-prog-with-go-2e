/**
 * PrintNameEmails
 */
package main

import (
	"log"
	"os"
	"text/template"
)

type Person struct {
	Name   string
	Emails []string
}

const templ = `{{$name := .Name}}
{{ $numEmails := .Emails | len -}}
{{range $idx, $email := .Emails -}}
Name is {{$name}}, email {{$email}} is {{ $idx | increment }} of {{ $numEmails }}
{{end}}
`

func main() {
	person := Person{
		Name: "jan",
		Emails: []string{"jan@newmarch.name",
			"jan.newmarch@gmail.com"},
	}
	t, err := template.New("Person template").Funcs(
		template.FuncMap{
			"increment": func(val int) int {
				return val + 1
			},
		},
	).Parse(templ)
	checkError(err)
	err = t.Execute(os.Stdout, person)
	checkError(err)
}
func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error ", err.Error())
	}
}
