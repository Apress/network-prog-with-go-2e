package main

import (
	"encoding/asn1"
	"fmt"
)

type T struct {
	S string
	I int
}

func main() {
	// using variables
	t1 := T{"ok", 1}
	mdata1, _ := asn1.Marshal(t1)
	var newT1 T
	asn1.Unmarshal(mdata1, &newT1)
	fmt.Printf("Before marshal: %v, after unmarshal: %v\n", t1, newT1)

	// using pointers
	var t2 = new(T)
	t2.S = "still ok"
	t2.I = 2
	mdata2, _ := asn1.Marshal(*t2)
	var newT2 = new(T)
	asn1.Unmarshal(mdata2, newT2)
	fmt.Printf("Before marshal: %v, after unmarshal: %v\n", t2, newT2)
}
