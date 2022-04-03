/* Read HTML
 */
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: ", os.Args[0], "file")
	}
	file := os.Args[1]
	bytes, err := os.ReadFile(file)
	checkError(err)
	r := strings.NewReader(string(bytes))
	z := html.NewTokenizer(r)
	depth := 0
	for {
		tt := z.Next()
		for n := 0; n < depth; n++ {
			fmt.Print(" ")
		}
		switch tt {
		case html.ErrorToken:
			if z.Err() == io.EOF {
				fmt.Println("EOF")
			} else {
				fmt.Println("Error ", z.Err().Error())
			}
			os.Exit(0)
		case html.TextToken:
			fmt.Println("Text: \"" + z.Token().String() + "\"")
		case html.StartTagToken, html.EndTagToken:
			fmt.Println("Tag: \"" + z.Token().String() + "\"")
			if tt == html.StartTagToken {
				depth++
			} else {
				depth--
			}
		}
	}
}
func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error ", err.Error())
	}
}
