/*
ASN1 example.
*/
package main

import (
	"encoding/asn1"
	"fmt"
	"time"
)

func thirteen() {
	val := 13
	mdata, _ := asn1.Marshal(val)
	var n int
	_, _ = asn1.Unmarshal(mdata, &n)
	fmt.Printf("Before marshal: %v, After unmarshal: %v\n", val, n)
}

func ascii() {
	s := "hello"
	mdata, _ := asn1.Marshal(s)
	var newstr string
	_, _ = asn1.Unmarshal(mdata, &newstr)
	fmt.Printf("Before marshal: %v, After unmarshal: %v\n", s, newstr)
}

func myTime() {
	t := time.Now()
	mdata, _ := asn1.Marshal(t)
	var newtime = new(time.Time)
	_, _ = asn1.Unmarshal(mdata, newtime)
	fmt.Printf("Before marshal: %v, After unmarshal: %v\n", t, newtime)
}

func main() {
	thirteen()
	ascii()
	myTime()
}
