/* ResolveIP
 */
package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: %s hostname\n", os.Args[0])
	}
	name := os.Args[1]
	addr, err := net.ResolveIPAddr("ip", name)
	if err != nil {
		log.Fatalln("Resolution error", err.Error())
	}
	fmt.Println("Resolved address is ", addr.String())
}
