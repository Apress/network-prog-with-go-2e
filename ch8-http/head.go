/* Head
 */
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: ", os.Args[0], "host:port")
	}
	url := os.Args[1]
	response, err := http.Head(url)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(response.Status)
	for k, v := range response.Header {
		fmt.Println(k+":", v)
	}
}
