
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

//IntToBytes 256 => [0 0 0 0 0 0 1 0]
func IntToBytes(n int) []byte {
	data := int64(n)
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes()
}

//BytesToInt [0 0 0 0 0 0 1 0] => 256
func BytesToInt(bys []byte) int {
	bytebuff := bytes.NewBuffer(bys)
	var data int64
	binary.Read(bytebuff, binary.BigEndian, &data)
	return int(data)
}

func main() {
	fmt.Println(IntToBytes(256))
	fmt.Println(BytesToInt([]byte{0, 0, 0, 0, 0, 0, 1, 0}))
}
