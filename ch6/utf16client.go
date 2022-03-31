/* UTF16 Client
 */
package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"unicode/utf16"
)

const BOM = '\ufffe'

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: ", os.Args[0], "host:port")
	}
	service := os.Args[1]
	conn, err := net.Dial("tcp", service)
	checkError(err)
	shorts := readShorts(conn)
	ints := utf16.Decode(shorts)
	str := string(ints)
	fmt.Println(str)
}
func readShorts(conn net.Conn) []uint16 {
	var buf [512]byte
	// read everything into the buffer
	n, err := conn.Read(buf[0:2]) // start with BOM
	for {
		m, err := conn.Read(buf[n:]) // read remaining byte pairs (originally unit16)
		if m == 0 || err != nil {
			break
		}
		n += m
	}
	checkError(err)
	var shorts []uint16
	shorts = make([]uint16, n/2)

	// We are checking for endianess
	// first - big endian 0xfffe
	// second - little endian 0xfeff
	// else - unknown
	// the inner loops are reading one byte per iteration
	// depending on endianess, places in the correct byte order
	// *warning* our server only supports big-endian
	if buf[0] == 0xff && buf[1] == 0xfe {
		for i := 2; i < n; i += 2 {
			shorts[i/2] = uint16(buf[i])<<8 + uint16(buf[i+1])
		}
	} else if buf[0] == 0xfe && buf[1] == 0xff {
		for i := 2; i < n; i += 2 {
			shorts[i/2] = uint16(buf[i+1])<<8 + uint16(buf[i])
		}
	} else {
		// unknown byte order
		fmt.Println("Unknown order")
	}
	return shorts
}
func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error ", err.Error())
	}
}
