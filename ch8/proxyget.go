/* ProxyGet
 */
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: ", os.Args[0], "http://host:port/page")
	}
	rawURL := os.Args[1]
	url, err := url.Parse(rawURL)
	checkError(err)

	response, err := http.Get(url.String())

	checkError(err)
	fmt.Println("Read ok")

	if response.StatusCode != http.StatusOK {
		log.Fatalln(response.StatusCode)
	}
	fmt.Println("Response ok")

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
func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error ", err.Error())
	}
}
