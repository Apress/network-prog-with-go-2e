/* XMLCodec
 */
package main

import (
	"encoding/xml"
	"golang.org/x/net/websocket"
)

func xmlMarshal(v interface{}) (msg []byte, payloadType byte, err error) {
	msg, err = xml.Marshal(v)
	return msg, websocket.TextFrame, nil
}
func xmlUnmarshal(msg []byte, payloadType byte, v interface{}) (err error) {
	err = xml.Unmarshal(msg, v)
	return err
}

var XMLCodec = websocket.Codec{xmlMarshal, xmlUnmarshal}
