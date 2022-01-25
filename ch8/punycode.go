package main

import (
	"fmt"
	"golang.org/x/net/idna"

	"net/url"
)

func main() {
	s := "https://日本語.jp:8443"
	r1, _ := idna.ToASCII(s)
	r2, _ := idna.ToUnicode(r1)
	fmt.Println(r1)
	fmt.Println(r2)

	fmt.Println(url.QueryEscape(s))
}
