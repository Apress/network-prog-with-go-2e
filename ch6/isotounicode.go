package main

import (
	"fmt"
	"unicode/utf8"
)

func str2int(str string) []int {
	r := []rune(str)
	b := make([]int, utf8.RuneCountInString(str))
	for i, v := range r {
		b[i] = int(v)
	}
	return b
}

//unicode to 8859-4
var unicodeToISOMap = map[int]uint8{
	// example match ascii 0x0021: 0x21, // !
	0x012e: 0xc7, // Į
	0x010c: 0xc8, // Č
	0x0112: 0xaa, // Ē
	0x0118: 0xca, // Ę
	// example match 0x00c9: 0xc9, // É
	// plus more
}

/* Turn a UTF-8 string into an ISO 8859 encoded byte array
 */
func unicodeStrToISO(str string) []byte {
	// get the unicode code points
	codePoints := str2int(str) //[]int(str)
	// create a byte array of the same length
	bytes := make([]byte, len(codePoints))
	for n, v := range codePoints {
		// see if the point is in the exception map
		iso, ok := unicodeToISOMap[v]
		if !ok {
			// just use the value
			iso = uint8(v)
		}
		bytes[n] = iso
	}
	return bytes
}

// inverse of unicodeToISOMap
var isoToUnicodeMap = map[uint8]int{
	0xc7: 0x012e,
	0xc8: 0x010c,
	0xaa: 0x0112,
	0xca: 0x0118,
	// and more
}

func isoBytesToUnicode(bytes []byte) string {
	codePoints := make([]int, len(bytes))
	for n, v := range bytes {
		unicode, ok := isoToUnicodeMap[v]
		if !ok {
			unicode = int(v)
		}
		codePoints[n] = unicode
	}
	return fmt.Sprintf("%q == %U", codePoints, codePoints)
}

func main() {
	x := "ĮĘa!"
	fmt.Printf("UTF-8: %s\n", x)

	fmt.Println("unicode to 8859-4")
	b := unicodeStrToISO(x)
	fmt.Printf("8859-4(hex): %x\n\n", b)

	fmt.Println("8859-4 to Unicode")
	fmt.Printf("Unicode: %v\n", isoBytesToUnicode(b))
}
