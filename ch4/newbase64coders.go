package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
)

var (
	enc   *base64.Encoding = base64.StdEncoding.WithPadding('|')
	input []byte           = []byte("∞a∞\x02ab")
	w     bytes.Buffer
)

func printer(action string, data []byte) {
	fmt.Printf("%11s: %s %v\n", action, string(data), data)
}

func restoreViaDecoder() {
	var buf *bytes.Buffer = bytes.NewBuffer(w.Bytes())
	var ior io.Reader = base64.NewDecoder(enc, buf)
	l := len(input)

	if l > 3 && l%3 != 0 {
		l = l + 2
	}

	restored := make([]byte, l)
	ior.Read(restored)
	printer("viaDecoder", restored)
}

func restoreViaEncoding() {
	var dst []byte = make([]byte, len(input))
	enc.Decode(dst, w.Bytes())
	printer("viaEncoding", dst)
}

func main() {
	printer("input", input)

	var wc io.WriteCloser = base64.NewEncoder(enc, &w)

	wc.Write(input)
	wc.Close()

	printer("encoded", w.Bytes())

	restoreViaDecoder()
	restoreViaEncoding()
}
