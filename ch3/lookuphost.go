/* LookupHost
 */
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	cname, _ := net.LookupCNAME(name)
	fmt.Println(cname)
	addrs, err := net.LookupHost(cname)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(2)
	}
	for _, s := range addrs {
		fmt.Println(s)
	}
	os.Exit(0)
}
