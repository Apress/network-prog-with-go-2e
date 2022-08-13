/* LookupHost
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
		log.Fatalf("Usage: %s hostname\n", os.Args[0])
	}
	name := os.Args[1]
	cname, _ := net.LookupCNAME(name)
	fmt.Println(cname)
	addrs, err := net.LookupHost(cname)
	if err != nil {
		log.Fatalln("Error: ", err.Error())
	}
	for _, addr := range addrs {
		fmt.Println(addr)
	}
}
