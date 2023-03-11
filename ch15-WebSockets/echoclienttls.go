/* EchoClientTLS
 */
package main

import (
	"crypto/tls"
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: ", os.Args[0], "wss://host:port")
	}
	config, err := websocket.NewConfig(os.Args[1],
		"http://localhost")
	checkError(err)
	tlsConfig := &tls.Config{InsecureSkipVerify: false}
	config.TlsConfig = tlsConfig
	conn, err := websocket.DialConfig(config)
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
