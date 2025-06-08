package main

import (
	"fmt"
	"net"
	"zinx/znet"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	for i := range 15 {
		msg := znet.NewMessage(0, fmt.Appendf([]byte{}, "hello %d", i))
		dp := znet.NewDataPack()
		bs, err := dp.Pack(msg)
		if err != nil {
			panic(err)
		}
		_, err = conn.Write(bs)
		if err != nil {
			panic(err)
		}
	}
}
