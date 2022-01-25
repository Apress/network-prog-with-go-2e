package main

import (
	"encoding/asn1"
	"fmt"
)

type T1 struct {
	F1 rune
	F2 int
}

type T2 struct {
	F3 rune
}

type T3 struct {
	F4 byte
}

func main() {
	// this first example works
	t1 := T1{'ロ', 1}
	mdata1, _ := asn1.Marshal(t1)

	t2 := new(T2)
	_, err := asn1.Unmarshal(mdata1, t2)
	fmt.Printf("Before marshal: %v, after unmarshal: %v\n", t1, t2)
	if err != nil {
		fmt.Println(err)
	}

	// syntax error (fails to fill all fields)
	y := T2{'ロ'}
	mdata2, _ := asn1.Marshal(y)
	z := new(T1)
	_, err = asn1.Unmarshal(mdata2, z)
	fmt.Printf("Before marshal: %v, after unmarshal: %v\n", y, z)
	if err != nil {
		fmt.Println(err)
	}

	// structural error (incorrect Go type byte != rune)
	t3 := new(T3)
	_, err = asn1.Unmarshal(mdata1, t3)
	fmt.Printf("Before marshal: %v, after unmarshal: %v\n", t1, t3)
	if err != nil {
		fmt.Println(err)
	}
}
