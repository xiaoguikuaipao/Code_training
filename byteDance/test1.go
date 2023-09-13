package main

import (
	"encoding/binary"
	"fmt"
)

func main() {

	buf := make([]byte, 2)
	buf[0] = 0x0f
	buf[1] = 0xff
	fmt.Println(binary.LittleEndian.Uint16(buf))

	buf2 := make([]byte, 2)
	buf2[0] = 0x0f
	buf2[1] = 0xff
	fmt.Println(binary.BigEndian.Uint16(buf2))

}
