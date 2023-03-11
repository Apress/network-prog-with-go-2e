/* JSON EchoClient
 */
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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
	defer conn.Close()
	for n := 0; n < 10; n++ {
		data, _ := json.Marshal(person)
		conn.Write(data)

		var newPerson Person
		buf, _ := readFully(conn)
		err = json.Unmarshal(buf, &newPerson)
		fmt.Println(newPerson)
	}
}

func readFully(conn net.Conn) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])

		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		if n < 512 {
			break
		}
	}
	return result.Bytes(), nil
}

func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error ", err.Error())
	}
}
