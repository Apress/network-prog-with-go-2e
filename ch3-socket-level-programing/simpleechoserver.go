/* SimpleEchoServer
 */
package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	service := ":1201"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		handleClient(conn)
		conn.Close() // we're finished
	}
}
func handleClient(conn net.Conn) {
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		checkError(err)
		fmt.Println(string(buf[0:]))
		_, err = conn.Write(buf[0:n])
		checkError(err)
	}
}
func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error: %s", err.Error())
	}
}
