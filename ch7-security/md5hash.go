/* MD5Hash
 */
package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	hash := md5.New()
	bytes := []byte("hello\n")
	hash.Write(bytes) // add data to running hash
	hashValue := hash.Sum(nil) // retrieve the hashed data
	hashSize := hash.Size() // how many bytes Sum returns (could just len(hashValue)
	
	// for every 4 bytes of hashValue
	// we stuff into an byte of val by shifting
	// val[first_byte] = hashValue[n] after shifting 24
	// val[second_byte = hashValue[n+1] after shifting 16
	// val[third_byte] = hashValue[n+2] after shifting 8
	// val[fourth_byte] = hashValue[n+3]
	// in the end, we have unint32 value that we print
	for n := 0; n < hashSize; n += 4 {
		var val uint32
		val = uint32(hashValue[n])<<24 +
			uint32(hashValue[n+1])<<16 +
			uint32(hashValue[n+2])<<8 +
			uint32(hashValue[n+3])
		fmt.Printf("%x ", val)
	}
	fmt.Println()
}
