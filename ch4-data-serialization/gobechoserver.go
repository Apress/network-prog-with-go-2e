/* Gob EchoServer
 */
package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
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
	service := "0.0.0.0:1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		encoder := gob.NewEncoder(conn)
		decoder := gob.NewDecoder(conn)
		for n := 0; n < 10; n++ {
			var person Person
			decoder.Decode(&person)
			fmt.Println(person)
			encoder.Encode(person)
		}
		conn.Close() // we're finished
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error ", err.Error())
	}
}
