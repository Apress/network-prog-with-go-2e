/* Mask
 */
package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		log.Fatalf("Usage: %s dotted-ip-addr ones bits\n", os.Args[0])
	}
	dotAddr := os.Args[1]
	ones, _ := strconv.Atoi(os.Args[2])
	bits, _ := strconv.Atoi(os.Args[3])
	addr := net.ParseIP(dotAddr)
	if addr == nil {
		log.Fatalln("nil Invalid address")
	}
	mask := net.CIDRMask(ones, bits)
	computedOnes, computedBits := mask.Size()
	network := addr.Mask(mask)
	fmt.Println("Address is ", addr.String(),
		"\nMask length is ", computedBits,
		"\nLeading ones count is ", computedOnes,
		"\nMask is (hex) ", mask.String(),
		"\nNetwork is ", network.String())
}
