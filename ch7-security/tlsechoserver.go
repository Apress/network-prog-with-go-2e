/* TLSEchoServer
 */
package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
)

func main() {
	cert, err := tls.LoadX509KeyPair("jan.newmarch.name.pem",
		"private.pem")
	checkError(err)
	config := tls.Config{Certificates: []tls.Certificate{cert}}
	service := "0.0.0.0:1200"
	listener, err := tls.Listen("tcp", service, &config)
	checkError(err)
	fmt.Println("Listening")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		fmt.Println("Accepted")
		go handleClient(conn)
	}
}
func handleClient(conn net.Conn) {
	defer conn.Close()
	var buf [512]byte
	for {
		fmt.Println("Trying to read")
		n, err := conn.Read(buf[0:])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(buf[0:]))
		_, err = conn.Write(buf[0:n])
		if err != nil {
			return
		}
	}
}
func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error ", err.Error())
	}
}
