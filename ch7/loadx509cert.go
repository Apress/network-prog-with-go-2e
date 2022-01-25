/* LoadX509Cert
 */
package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/gob"
	"fmt"
	"os"
)

func main() {
	certCerFile, err := os.Open("jan.newmarch.name.cer")
	checkError(err)
	derBytes := make([]byte, 1000) // bigger than the file
	count, err := certCerFile.Read(derBytes)
	checkError(err)
	certCerFile.Close()
	// trim the bytes to actual length in call
	cert, err := x509.ParseCertificate(derBytes[0:count])
	checkError(err)
	fmt.Printf("Name %s\n", cert.Subject.CommonName)
	fmt.Printf("Not before %s\n", cert.NotBefore.String())
	fmt.Printf("Not after %s\n", cert.NotAfter.String())

	pub, err := os.Open("public.key")
	checkError(err)
	dec := gob.NewDecoder(pub)
	publicKey := new(rsa.PrivateKey)
	err = dec.Decode(publicKey)
	checkError(err)
	pub.Close()

	if cert.PublicKey.(*rsa.PublicKey).N.Cmp(publicKey.N) == 0 && publicKey.E == cert.PublicKey.(*rsa.PublicKey).E {
		println("Same public key")
	} else {
		println("Different public key")
	}
}
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
