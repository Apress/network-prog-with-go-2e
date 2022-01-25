/* textproto
 */

package main

import (
	"fmt"
	"net/textproto"
	"os"
)

func checkerror(err error) {
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
}

func main() {
	conn, e := textproto.Dial("unix", "/tmp/fakewebserver")
	checkerror(e)
	defer conn.Close()
	fmt.Println("Sending request to retrieve /mypage")
	id, e := conn.Cmd("GET /mypage")
	checkerror(e)
	conn.StartResponse(id)
	defer conn.EndResponse(id)
	// fake sending back a 200 via nc or your own server
	code, stringResult, error := conn.ReadCodeLine(200)
	checkerror(e)
	fmt.Println(code, "\n", stringResult, "\n", error)

}
