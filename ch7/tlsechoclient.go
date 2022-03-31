package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
)

func main() {
	rootPEM, err := os.ReadFile("jan.newmarch.name.pem")
	// First, create the set of root certificates. For this example we only
	// have one. It's also possible to omit this in order to use the
	// default root set of the current operating system.
	roots := x509.NewCertPool()
	if ok := roots.AppendCertsFromPEM(rootPEM); !ok {
		panic("failed to parse root certificate")
	}

	conn, err := tls.Dial("tcp", "localhost:1200", &tls.Config{
		RootCAs: roots,
		//		InsecureSkipVerify: false,
	})
	if err != nil {
		panic("failed to connect: " + err.Error())
	}

	// Now write and read lots
	for n := 0; n < 10; n++ {
		fmt.Println("Writing...")
		conn.Write([]byte("Hello " + string(n+48)))
		var buf [512]byte
		n, _ := conn.Read(buf[0:])
		fmt.Println(string(buf[0:n]))
	}

	conn.Close()
}
