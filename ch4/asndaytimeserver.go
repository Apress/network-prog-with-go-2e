/* ASN1 DaytimeServer
 */
package main

import (
	"encoding/asn1"
	"log"
	"net"
	"time"
)

func main() {
	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		daytime := time.Now()
		// ignore returned errors
		mdata, _ := asn1.Marshal(daytime)
		conn.Write(mdata)
		conn.Close() // we're finished
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error: %s", err.Error())
	}
}
