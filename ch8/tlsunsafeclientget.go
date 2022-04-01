/* TLSUnsafeClientGet
 */
package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: ", os.Args[0], "https://host:port/page")
	}
	url, err := url.Parse(os.Args[1])
	checkError(err)
	if url.Scheme != "https" {
		log.Fatalln("Not https scheme ", url.Scheme)
	}

	transport := &http.Transport{}
	transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: false}
	client := &http.Client{Transport: transport}

	request, err := http.NewRequest("GET", url.String(), nil)
	// only accept UTF-8
	checkError(err)

	response, err := client.Do(request)
	checkError(err)

	if response.StatusCode != http.StatusOK {
		log.Fatalln(response.Status)
	}
	fmt.Println("get a response")

	chSet := getCharset(response)
	fmt.Printf("got charset %s\n", chSet)
	if chSet != "UTF-8" {
		log.Fatalln("Cannot handle", chSet)
	}

	var buf [512]byte
	reader := response.Body
	fmt.Println("got body")
	for {
		n, err := reader.Read(buf[0:])
		checkError(err)
		fmt.Print(string(buf[0:n]))
	}
}
func getCharset(response *http.Response) string {
	contentType := response.Header.Get("Content-Type")
	if contentType == "" {
		// guess
		return "UTF-8"
	}
	idx := strings.Index(contentType, "charset:")
	if idx == -1 {
		// guess
		return "UTF-8"
	}
	return strings.Trim(contentType[idx:], " ")
}
func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error ", err.Error())
	}
}
