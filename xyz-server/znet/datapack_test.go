package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
	"time"
)

func TestDataPack(t *testing.T) {
	started := revForDataPack(t)
	<-started
	sendForDataPack(t)
}

func revForDataPack(t *testing.T) chan struct{} {
	started := make(chan struct{})
	socket, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		t.Error(err)
	}
	go func() {
		defer showSince()()
		started <- struct{}{}
		for {
			conn, err := socket.Accept()
			if err != nil {
				t.Error(err)
			}
			dp := &DataPack{}
			go func() {
				for {
					head := make([]byte, dp.GetHeadLen())
					_, err := io.ReadFull(conn, head)
					if err != nil {
						fmt.Println(err)
						break
					}
					msg, err := dp.UnPack(head)
					if err != nil {
						t.Error(err)
					} else {
						fmt.Printf("unpack head succ: %v\n", msg)
					}
					if msg.DataLen() > dp.GetHeadLen() {
						data := make([]byte, msg.DataLen()-dp.GetHeadLen())
						_, err := io.ReadFull(conn, data)
						if err != nil {
							t.Error(err)
						}
						msg.SetData(data)
						fmt.Printf("set data succ: %s\n", string(msg.Data()))
					}
				}
			}()
		}
	}()
	return started
}

func showSince() func() {
	fmt.Printf("begin at %s\n", time.Now().String())
	return func() {
		fmt.Printf("exit at %s\n", time.Now().String())
	}
}

func sendForDataPack(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		t.Error(err)
	}
	for i := range 10 {
		dp := &DataPack{}
		msg := &Message{
			id:   uint32(i),
			data: fmt.Appendf([]byte{}, "hello %d", i),
		}
		msg.dataLen = uint32(len(msg.data) + 8)
		data, err := dp.Pack(msg)
		if err != nil {
			t.Error(err)
		}
		_, err = conn.Write(data)
		if err != nil {
			t.Error(err)
		} else {
			fmt.Printf("send i:%d succ\n", i)
		}
		time.Sleep(1 * time.Second)
	}
}
