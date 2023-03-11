/* PersonClientXML
 */
package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"os"
)

type Person struct {
	Name   string
	Emails []string
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: ", os.Args[0], "ws://host:port")
	}
	service := os.Args[1]
	conn, err := websocket.Dial(service, "", "http://localhost")
	checkError(err)
	person := Person{Name: "Jan",
		Emails: []string{"ja@newmarch.name",
			"jan.newmarch@gmail.com"},
	}
	err = XMLCodec.Send(conn, person)
	if err != nil {
		fmt.Println("Couldn't send msg " + err.Error())
	}
	os.Exit(0)
}
func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error ", err.Error())
	}
}
