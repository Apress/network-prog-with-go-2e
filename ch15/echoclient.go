/* EchoClient
 */
package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"os"
	"log"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: ", os.Args[0], "ws://host:port")
	}
	service := os.Args[1]
	conn, err := websocket.Dial(service, "",
		"http://localhost:12345")
	checkError(err)
	var msg string
	for {
		err := websocket.Message.Receive(conn, &msg)
		if err != nil {
			if err == io.EOF {
				// graceful shutdown by server
				break
			}
			fmt.Println("Couldn't receive msg " +
				err.Error())
			break
		}
		fmt.Println("Received from server: " + msg)
		// return the msg
		err = websocket.Message.Send(conn, msg)
		if err != nil {
			fmt.Println("Couldn't return msg")
			break
		}
	}
}
func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error ", err.Error())
	}
}
