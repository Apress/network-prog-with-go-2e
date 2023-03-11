/* Ping
 */
package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

// change this to your own IP address or leave set to 0.0.0.0
const myIPAddress = "0.0.0.0"
const ipv4HeaderSize = 20

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: ", os.Args[0], "host")
	}
	localAddr, err := net.ResolveIPAddr("ip4", myIPAddress)
	checkError(err)

	remoteAddr, err := net.ResolveIPAddr("ip4", os.Args[1])
	checkError(err)

	conn, err := net.DialIP("ip4:icmp", localAddr, remoteAddr)
	checkError(err)

	var msg [512]byte
	msg[0] = 8  // echo
	msg[1] = 0  // code 0
	msg[2] = 0  // checksum, fix later
	msg[3] = 0  // checksum, fix later
	msg[4] = 0  // identifier[0]
	msg[5] = 13 // identifier[1] (arbitrary)
	msg[6] = 0  // sequence[0]
	msg[7] = 37 // sequence[1] (arbitrary)
	len := 8

	// now fix checksum bytes
	check := checkSum(msg[0:len])
	msg[2] = byte(check >> 8)
	msg[3] = byte(check & 255)

	// send the message
	_, err = conn.Write(msg[0:len])
	checkError(err)
	fmt.Print("Message sent:    ")
	for n := 0; n < 8; n++ {
		fmt.Print(" ", msg[n])
	}
	fmt.Println()

	// receive a reply
	size, err2 := conn.Read(msg[0:])
	checkError(err2)
	fmt.Print("Message received:")
	for n := ipv4HeaderSize; n < size; n++ {
		fmt.Print(" ", msg[n])
	}
	fmt.Println()
	os.Exit(0)
}
func checkSum(msg []byte) uint16 {
	sum := 0
	// assume even for now
	for n := 0; n < len(msg); n += 2 {
		sum += int(msg[n])*256 + int(msg[n+1])
	}
	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16)
	var answer uint16 = uint16(^sum)
	return answer
}
func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error: %s", err.Error())
	}
}
