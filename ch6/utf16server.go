/* UTF16 Server
 */
package main

import (
	"log"
	"net"
	"unicode/utf16"
)

// warning, our server currently only supports big endian
const BOM = '\ufeff'

func main() {
	service := "0.0.0.0:1210"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		// eg. Ŵ  is 0x0174, Ã is 0x00c3
		str := "Ŵj'ai arrÃªtÃ©"
		shorts := utf16.Encode([]rune(str))
		writeShorts(conn, shorts)
		conn.Close()
	}
}
func writeShorts(conn net.Conn, shorts []uint16) {
	var bytes [2]byte
	// send the BOM as first two bytes
	bytes[0] = BOM >> 8             // taking ff from BOM
	bytes[1] = BOM & 255            // taking fe from BOM
	_, err := conn.Write(bytes[0:]) // send BOM
	checkError(err)
	for _, v := range shorts {
		// breakup the unit16 into two bytes, then send
		bytes[0] = byte(v >> 8)
		bytes[1] = byte(v & 255)
		_, err = conn.Write(bytes[0:])
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
