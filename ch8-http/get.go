/* Get
 */
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: ", os.Args[0], "host:port")
	}
	url := os.Args[1]
	response, err := http.Get(url)
	checkError(err)
	if response.StatusCode != http.StatusOK {
		log.Fatalln(response.StatusCode)
	}
	fmt.Println("The response header is")
	b, _ := httputil.DumpResponse(response, false)
	fmt.Print(string(b))
	contentTypes := response.Header["Content-Type"]
	if !acceptableCharset(contentTypes) {
		log.Fatalln("Cannot handle", contentTypes)
	}
	fmt.Println("The response body is")
	var buf [512]byte
	reader := response.Body
	for {
		n, err := reader.Read(buf[0:])
		if err != nil {
			if err == io.EOF {
				fmt.Print(string(buf[0:n]))
				reader.Close()
				break
			}
			checkError(err)
		}
		fmt.Print(string(buf[0:n]))
	}
}
func acceptableCharset(contentTypes []string) bool {
	// each type is like [text/html; charset=utf-8]
	// we want the UTF-8 only
	for _, cType := range contentTypes {
		if strings.Index(cType, "utf-8") != -1 {
			return true
		}
	}
	return false
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
