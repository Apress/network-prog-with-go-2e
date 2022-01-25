/* ProtocolBuffer
 */

package main

import (
	"myapp/protos"

	"google.golang.org/protobuf/proto"

	"fmt"
)

func main() {
	name := protos.Person_Name{
		Family:   "newmarch",
		Personal: "jan",
	}
	email1 := protos.Person_Email{
		Kind:    "home",
		Address: "jan@newmarch.name",
	}
	email2 := protos.Person_Email{
		Kind:    "work",
		Address: "j.newmarch@boxhill.edu.au",
	}
	emails := []*protos.Person_Email{&email1, &email2}

	p := protos.Person{
		Name:  &name,
		Email: emails,
	}

	fmt.Println(p)

	data, _ := proto.Marshal(&p)

	newP := protos.Person{}

	proto.Unmarshal(data, &newP)

	fmt.Printf("%v\n", newP)

	if p.Name.Personal == newP.Name.Personal && p.Email[0].Address == newP.Email[0].Address {
		fmt.Println("same")
	}
}
