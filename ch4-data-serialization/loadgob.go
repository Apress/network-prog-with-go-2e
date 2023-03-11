/* LoadGob
 */
package main

import (
	"encoding/gob"
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
	loadGob("person.gob", &person)
	fmt.Println(person)
}

func loadGob(fileName string, key interface{}) {
	inFile, err := os.Open(fileName)
	checkError(err)
	decoder := gob.NewDecoder(inFile)
	err = decoder.Decode(key)
	checkError(err)
	inFile.Close()
}

func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error ", err.Error())
	}
}
