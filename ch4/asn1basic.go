/* ASN.1 Basic
 */

package main

import (
	"encoding/asn1"
	"fmt"
	"time"
)

func main() {
	// time pointer to time value
	t := time.Now()
	fmt.Println("Before marshalling: ", t.String())
	mdata, _ := asn1.Marshal(t)
	var newtime = new(time.Time)
	_, _ = asn1.Unmarshal(mdata, newtime)
	fmt.Println("After marshal/unmarshal: ", newtime.String())

	// vulgar fraction, string to string
	s := "hello \u00bc"
	fmt.Println("Before marshalling: ", s)
	mdata2, _ := asn1.Marshal(s)
	var newstr string
	_, _ = asn1.Unmarshal(mdata2, &newstr)
	fmt.Println("After marshal/unmarshal: ", newstr)
}
