/* LoadJSON
 */
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
	Name  Name
	Email []Email
}

type Name struct {
	Family   string
	Personal string
}

type Email struct {
	Kind    string
	Address string
}

func main() {
	var person Person
	loadJSON("person.json", &person)
	fmt.Printf("%v\n", person)
}

func loadJSON(fileName string, key interface{}) {
	data, err := os.ReadFile(fileName)
	checkError(err)
	err = json.Unmarshal(data, key)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error ", err.Error())
	}
}
