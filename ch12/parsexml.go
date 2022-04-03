/* Parse XML
 */
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: ", os.Args[0], "file")
	}
	file := os.Args[1]
	bytes, err := ioutil.ReadFile(file)
	checkError(err)
	r := strings.NewReader(string(bytes))
	parser := xml.NewDecoder(r)
	depth := 0
	for {
		token, err := parser.Token()
		if err != nil {
			break
		}
		switch elmt := token.(type) {
		case xml.StartElement:
			name := elmt.Name.Local
			printElmt(name+":start", depth)
			depth++
		case xml.EndElement:
			depth--
			name := elmt.Name.Local
			printElmt(name+":end", depth)
		case xml.CharData:
			printElmt(string([]byte(elmt)), depth)
		case xml.Comment:
			printElmt("Comment", depth)
		case xml.ProcInst:
			printElmt("ProcInst", depth)
		case xml.Directive:
			printElmt("Directive", depth)
		default:
			fmt.Println("Unknown")
		}
	}
}
func printElmt(s string, depth int) {
	slimS := strings.TrimSpace(s)
	if len(slimS) == 0 {
		return
	}
	for n := 0; n < depth; n++ {
		fmt.Print("  ")
	}
	fmt.Println(slimS)
}
func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error ", err.Error())
	}
}
