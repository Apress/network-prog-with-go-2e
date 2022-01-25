package main

import "fmt"

func main() {
	str := "百度一下, 你就知道"
	fmt.Println("String length: ", len([]rune(str)))
	fmt.Println("Byte length: ", len(str))
}
