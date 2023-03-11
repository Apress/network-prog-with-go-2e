/* Marshal
 */
package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    Name     `xml:"name"`
	Email   []Email  `xml:"email"`
}
type Name struct {
	Family   string `xml:"family"`
	Personal string `xml:"personal"`
}
type Email struct {
	Kind    string "attr"
	Address string "chardata"
}

func main() {
	person := Person{
		Name: Name{Family: "Newmarch", Personal: "Jan"},
		Email: []Email{Email{Kind: "home", Address: "jan"},
			Email{Kind: "work", Address: "jan"}}}
	buff, _ := xml.Marshal(person)
	fmt.Println(string(buff))
}
