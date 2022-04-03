/* TemperatureServer
 */
package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"os/exec"
	"time"
)

var ROOT_DIR = "."

func GetTemp(ws *websocket.Conn) {
	for {
		msg, err := exec.Command(ROOT_DIR + "/sensors.sh").CombinedOutput()
		checkError(err)
		fmt.Println("Sending to client: " + string(msg[:]))
		err = websocket.Message.Send(ws, string(msg[:]))
		if err != nil {
			fmt.Println("Can't send")
			break
		}
		time.Sleep(time.Duration(2) * time.Second)
		var reply string
		err = websocket.Message.Receive(ws, &reply)
		if err != nil {
			fmt.Println("Can't receive")
			break
		}
		fmt.Println("Received back from client: " + reply)
	}
}
func main() {
	fileServer := http.FileServer(http.Dir(ROOT_DIR))
	http.Handle("/GetTemp", websocket.Handler(GetTemp))
	http.Handle("/", fileServer)
	err := http.ListenAndServe(":12345", nil)
	checkError(err)
}
func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error ", err.Error())
	}
}
