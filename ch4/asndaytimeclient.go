/* ASN.1 DaytimeClient
 */
package main

import (
	"bytes"
	"encoding/asn1"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: %s host:port", os.Args[0])
	}
	service := os.Args[1]
	conn, err := net.Dial("tcp", service)
	checkError(err)
	result, err := readFully(conn)
	checkError(err)
	var newtime time.Time
	_, err1 := asn1.Unmarshal(result, &newtime)
	checkError(err1)
	fmt.Println("After marshal/unmarshal: ", newtime.String())
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()
	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}

func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error: %s", err.Error())
	}
}
