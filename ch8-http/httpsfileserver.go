/* HTTPSFileServer
 */
package main

import (
	"net/http"
	"log"
)

func main() {
	// deliver files from the directory /tmp/www
	fileServer := http.FileServer(http.Dir("/tmp/www"))
	// register the handler and deliver requests to it
	err := http.ListenAndServeTLS(":8000", "jan.newmarch.name.pem",
		"private.pem", fileServer)
	if err != nil {
		log.Fatalln(err)
	}
}
