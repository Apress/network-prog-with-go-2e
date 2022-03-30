/* LookupPort
 */
package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("Usage: %s network-type service\n", os.Args[0])
	}
	networkType := os.Args[1]
	service := os.Args[2]
	port, err := net.LookupPort(networkType, service)
	if err != nil {
		log.Fatalln("Error: ", err.Error())
	}
	fmt.Println("Service port ", port)
}
