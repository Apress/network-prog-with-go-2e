/* PrintEmails
 */
package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

type Person struct {
	Name   string
	Emails []string
}

const templ = `The name is {{.Name}}.
{{range .Emails}}
An email is "{{. | emailExpand}}"
{{end}}`

func main() {
	person := Person{
		Name: "jan",
		Emails: []string{"jan@newmarch.name",
			"jan.newmarch@gmail.com"},
	}
	t, err := template.New("Person template").Funcs(
		template.FuncMap{
			"emailExpand": func(emailAddress string) string {
				return strings.Replace(emailAddress, "@", " at ", -1)
			},
		},
	).Parse(templ)

	err = t.Execute(os.Stdout, person)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error ", err.Error())
	}
}
