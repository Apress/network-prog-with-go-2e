/* File Server
 */
package main

import (
	"log"
	"net/http"
)

func main() {
	// deliver files from the directory /tmp/www
	fileServer := http.FileServer(http.Dir("/tmp/www"))

	// register the handler and deliver requests to it
	err := http.ListenAndServe(":8000", fileServer)
	if err != nil {
		log.Fatalln(err)
	}
}
