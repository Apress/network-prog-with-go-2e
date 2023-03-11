/* EchoClientGorilla
 */
package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: ", os.Args[0], "ws://host:port")
	}
	service := os.Args[1]
	header := make(http.Header)
	header.Add("Origin", "http://localhost:12345")
	conn, _, err := websocket.DefaultDialer.Dial(service, header)
	checkError(err)
	for {
		_, reply, err := conn.ReadMessage()
		if err != nil {
			if err == io.EOF {
				// graceful shutdown by server
				fmt.Println(`EOF from server`)
				break
			}
			if websocket.IsCloseError(err,
				websocket.CloseAbnormalClosure) {
				fmt.Println(`Close from server`)
				break
			}
			fmt.Println("Couldn't receive msg " +
				err.Error())
			break
		}
		fmt.Println("Received from server: " +
			string(reply[:]))
		// return the msg
		err = conn.WriteMessage(websocket.TextMessage, reply)
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
