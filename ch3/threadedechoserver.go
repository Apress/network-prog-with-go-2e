/* ThreadedEchoServer
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
		// run as a goroutine
		go handleClient(conn)
	}
}
func handleClient(conn net.Conn) {
	// close connection on exit
	defer conn.Close()
	var buf [512]byte
	for {
		// read up to 512 bytes
		n, err := conn.Read(buf[0:])
		checkError(err)
		fmt.Println(string(buf[0:]))
		// write the n bytes read
		_, err = conn.Write(buf[0:n])
		checkError(err)
	}
}
func checkError(err error) {
	if err != nil {
		log.Fatalf("Fatal error: %s", err.Error())
	}
}
