/* ThreadedIPEchoServer
 */
package main

import (
	"log"
	"net"
)

func main() {
	service := ":1200"
	listener, err := net.Listen("tcp", service)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}
func handleClient(conn net.Conn) {
	defer conn.Close()
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		checkError(err)

		_, err = conn.Write(buf[0:n])
		checkError(err)
	}
}
func checkError(err error) {
	if err != nil {
		log.Fatalf("Fatal error: %s", err.Error())
	}
}
