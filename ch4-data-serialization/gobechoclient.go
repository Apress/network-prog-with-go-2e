/* Gob EchoClient
 */
package main

import (
	"encoding/gob"
	"log"
	"net"
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
	if len(os.Args) != 2 {
		log.Fatalln("Usage: ", os.Args[0], "host:port")
	}
	person := Person{
		Name: Name{Family: "Newmarch", Personal: "Jan"},
		Email: []Email{
			Email{Kind: "home", Address: "jan@newmarch.name"},
			Email{Kind: "work", Address: "j.newmarch@boxhill.edu.au"},
		},
	}
	service := os.Args[1]
	conn, err := net.Dial("tcp", service)
	checkError(err)
	encoder := gob.NewEncoder(conn)
	decoder := gob.NewDecoder(conn)
	for n := 0; n < 10; n++ {
		encoder.Encode(person)
		var newPerson Person
		decoder.Decode(&newPerson)

	}
}

func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error ", err.Error())
	}
}
