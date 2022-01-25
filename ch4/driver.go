package main

/* driver.go
 */

import (
	"encoding/asn1"
	"fmt"
	"badtype"
)

func main() {
	// using variables
	t1 := p.T{F:1}
	mdata1, err := asn1.Marshal(t1)
	fmt.Println(err)
	var newT1 p.T
	_, err1 := asn1.Unmarshal(mdata1, &newT1)
	fmt.Println(err1)
}

