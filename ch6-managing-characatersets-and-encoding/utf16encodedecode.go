package main

import (
	"unicode/utf16"
	"fmt"
)

func main() {
	str := "百度一下, 你就知道"
	fmt.Println("Before encoding:", str)

	runes := utf16.Encode([]rune(str))
	ints := utf16.Decode(runes)

	str = string(ints)
	fmt.Println("After encoding:", str)
}
