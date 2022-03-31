/* GenX509Cert
 */
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/gob"
	"encoding/pem"
	"log"
	"math/big"
	"os"
	"time"
)

func main() {
	random := rand.Reader
	var key rsa.PrivateKey
	loadKey("private.key", &key)
	now := time.Now()
	then := now.Add(60 * 60 * 24 * 365 * 1000 * 1000 * 1000)
	// one year
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName:   "jan.newmarch.name",
			Organization: []string{"Jan Newmarch"},
		},
		NotBefore:    now,
		NotAfter:     then,
		SubjectKeyId: []byte{1, 2, 3, 4},
		KeyUsage: x509.KeyUsageCertSign |
			x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		BasicConstraintsValid: true,
		IsCA:                  true,
		DNSNames: []string{"jan.newmarch.name",
			"localhost"},
	}
	derBytes, err := x509.CreateCertificate(random, &template,
		&template, &key.PublicKey, &key)
	checkError(err)
	certCerFile, err := os.Create("jan.newmarch.name.cer")
	checkError(err)
	certCerFile.Write(derBytes)
	certCerFile.Close()
	certPEMFile, err := os.Create("jan.newmarch.name.pem")
	checkError(err)
	pem.Encode(certPEMFile, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certPEMFile.Close()
	keyPEMFile, err := os.Create("private.pem")
	checkError(err)
	pem.Encode(keyPEMFile, &pem.Block{Type: "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(&key)})
	keyPEMFile.Close()
}
func loadKey(fileName string, key interface{}) {
	inFile, err := os.Open(fileName)
	checkError(err)
	decoder := gob.NewDecoder(inFile)
	err = decoder.Decode(key)
	checkError(err)
	inFile.Close()
}
func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error ", err.Error())
	}
}
