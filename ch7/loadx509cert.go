/* LoadX509Cert
 */
package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

func main() {
	// load certificate so we can access embedded public key
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

	// load non-emdedded public key
	// should be the same as above embedded key
	pub, err := os.Open("public.key")
	checkError(err)
	dec := gob.NewDecoder(pub)
	publicKey := new(rsa.PrivateKey)
	err = dec.Decode(publicKey)
	checkError(err)
	pub.Close()

	// genx509cert.go created a public key and certificate
	// certificates also embed the public key
	// we are comparing the public key and the embedded public key fields
	// see go doc crypto/rsa.PublicKey for more
	if cert.PublicKey.(*rsa.PublicKey).N.Cmp(publicKey.N) == 0 {
		if publicKey.E == cert.PublicKey.(*rsa.PublicKey).E {
			fmt.Println("Same public key")
			return
		}
	}
	fmt.Println("Different public key")
}
func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error ", err.Error())
	}
}
