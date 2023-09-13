package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

const (
	numnodes = 20
	numoftxs = 10
)

var m sync.Map
var m2 sync.Map
var Nodes = make([]*Node, 0)

//var addr []string = GetAddr()

type Node struct {
	Id                 int
	CurrentView        int
	State              string
	Blockchain         []Block
	PPMessages         chan Message
	PMessages          chan Message
	CommitMsg          chan Message
	BlocksVotes        map[string]map[string]int
	StateToBlock       map[string]string
	AlreadySeenTxs     map[string]int
	AlreadyConfirmedTx map[string]int
	TxPools            chan string
	PrimaryNode        int
	DelegatedPool      chan Message
}

type Message struct {
	B        Block
	Tx       string
	MsgType  string
	Senderid int
	View     int
}

type Block struct {
	Proposerid int
	Hash       string
	Tx         string
	Height     int
}

func NewNode(id int) *Node {
	return &Node{
		Id:                 id,
		CurrentView:        0,
		State:              "ready",
		Blockchain:         make([]Block, 0),
		PPMessages:         make(chan Message, 500),
		PMessages:          make(chan Message, 500),
		CommitMsg:          make(chan Message, 500),
		BlocksVotes:        make(map[string]map[string]int, numoftxs),
		TxPools:            make(chan string, 100),
		StateToBlock:       make(map[string]string, numoftxs),
		AlreadySeenTxs:     make(map[string]int, numoftxs),
		PrimaryNode:        1,
		AlreadyConfirmedTx: make(map[string]int, numoftxs),
		DelegatedPool:      make(chan Message, 100),
	}
}

//func GetAddr() []string {
//	Addrs := make([]string, numnodes)
//	for i := 0; i < numnodes; i++ {
//		Addrs[i] = "0.0.0.0:" + strconv.Itoa(5000+i)
//	}
//	return Addrs
//}

//func (n *Node) Listen() {
//	ln, err := net.Listen("tcp", n.Addrs[n.Id])
//	if err != nil {
//		fmt.Printf("Node %d failed to listen on port %s: %v\n", n.Id, n.Addrs[n.Id], err)
//		return
//	} else {
//		fmt.Printf("Node %d start to listen\n", n.Id)
//	}
//	defer func(ln net.Listener) {
//		err := ln.Close()
//		if err != nil {
//
//		}
//	}(ln)
//	for {
//		conn, err := ln.Accept()
//		if err != nil {
//			fmt.Printf("Node %d failed to accept connection: %v\n", n.Id, err)
//			continue
//		}
//		go n.Receive(conn)
//	}
//}

//func (n *Node) Receive(con net.Conn) {
//	defer func(con net.Conn) {
//		err := con.Close()
//		if err != nil {
//
//		}
//	}(con)
//	buf := make([]byte, 1024)
//	num, err := con.Read(buf)
//	if err != nil {
//		log.Fatal("Error reading message", err)
//	}
//	var msg Message
//	err = json.Unmarshal(buf[:num], &msg)
//	if err != nil {
//		log.Println("error decoding message", err.Error())
//		return
//	}
//	switch msg.MsgType {
//	case "preprepare":
//		n.PPMessages <- msg
//	case "prepare":
//		n.PMessages <- msg
//	case "commit":
//		n.CommitMsg <- msg
//	case "transaction":
//		if n.AlreadySeenTxs[msg.Tx] != 1 {
//			n.TxPools <- msg.Tx
//			n.AlreadySeenTxs[msg.Tx] = 1
//			n.Broadcast(msg)
//			fmt.Printf("Node%v first receive transaction:%v\n", n.Id, msg.Tx)
//		}
//	case "delegate":
//		n.DelegatedPool <- msg
//	}
//}

func (n *Node) ProcessMsg() {
	for {
		if n.Id != 0 {
			switch n.State {
			case "ready":
				n.ProcessPP()
			case "preprepare":
				n.ProcessP()
			case "prepare":
				n.ProcessCommit()
			}
		} else if n.Id == 0 {
			n.SpecialNode()
		}
	}
}

func (n *Node) ReadyForDelegate() {
	for {
		select {
		case Msg := <-n.DelegatedPool:
			if Msg.View == n.CurrentView {
				n.PrimaryNode = Msg.Senderid
				return
			} else if Msg.View > n.CurrentView {
				n.DelegatedPool <- Msg
			}
		default:
			continue
		}
	}
}

func (n *Node) Propose() {
	for {
		if n.PrimaryNode == n.Id && n.State == "ready" {
			var Tx string
			select {
			case Tx = <-n.TxPools:
				if _, ok := n.AlreadyConfirmedTx[Tx]; ok {
					continue
				}
			default:
				continue
			}
			B := Block{
				Tx:         Tx,
				Height:     len(n.Blockchain) + 1,
				Proposerid: n.Id,
			}
			block, err := json.Marshal(B)
			if err != nil {
				log.Fatal("Propose Marshal error", err)
			}
			B.Hash = hex.EncodeToString(block)
			PPMsg := Message{
				MsgType:  "preprepare",
				View:     n.CurrentView,
				B:        B,
				Senderid: n.Id,
			}
			n.Broadcast(PPMsg)
			fmt.Printf("node%v has proposed tx%v\n", n.Id, Tx)
			n.PrimaryNode = -1
			n.State = "preprepare"
		}
	}
}

func (n *Node) ProcessPP() {
	select {
	case PPMsg := <-n.PPMessages:
		if n.State == "ready" && PPMsg.B.Proposerid == n.PrimaryNode &&
			PPMsg.View == n.CurrentView {
			n.StateToBlock[PPMsg.B.Hash] = "preprepared"
			fmt.Printf("Node %v trans to preprepare state\n", n.Id)
			n.State = "preprepare"
			PreMsg := Message{
				Senderid: n.Id,
				View:     n.CurrentView,
				B:        PPMsg.B,
				MsgType:  "prepare",
			}
			n.Broadcast(PreMsg)
		} else if PPMsg.View > n.CurrentView {
			n.PPMessages <- PPMsg
		}
	default:
		return
	}
}

func (n *Node) ProcessP() {
	select {
	case PMsg := <-n.PMessages:
		if n.State == "preprepare" && PMsg.B.Proposerid == n.PrimaryNode &&
			PMsg.View == n.CurrentView {
			if _, ok := n.BlocksVotes[PMsg.B.Hash]; ok != true {
				n.BlocksVotes[PMsg.B.Hash] = make(map[string]int)
			}
			n.BlocksVotes[PMsg.B.Hash]["prepare"] += 1
			if (n.BlocksVotes[PMsg.B.Hash]["prepare"] > 2/3*numnodes) &&
				n.StateToBlock[PMsg.B.Hash] == "preprepared" {
				n.StateToBlock[PMsg.B.Hash] = "prepared"
				fmt.Printf("Node %d trans to prepare\n", n.Id)
				n.State = "prepare"
				ComMsg := Message{
					View:     n.CurrentView,
					Senderid: n.Id,
					B:        PMsg.B,
					MsgType:  "commit",
				}
				n.Broadcast(ComMsg)
			}
		} else if PMsg.View > n.CurrentView {
			n.PMessages <- PMsg
		}
	default:
		return
	}
}

func (n *Node) ProcessCommit() {
	select {
	case ComMsg := <-n.CommitMsg:
		if n.State == "prepare" && ComMsg.B.Proposerid == n.PrimaryNode &&
			ComMsg.View == n.CurrentView {
			if _, ok := n.BlocksVotes[ComMsg.B.Hash]; ok != true {
				n.BlocksVotes[ComMsg.B.Hash] = make(map[string]int)
			}
			n.BlocksVotes[ComMsg.B.Hash]["commit"] += 1
			if (n.BlocksVotes[ComMsg.B.Hash]["commit"] >= 2/3*numnodes) &&
				n.StateToBlock[ComMsg.B.Hash] == "prepared" {
				n.StateToBlock[ComMsg.B.Hash] = "commited"
				n.Blockchain = append(n.Blockchain, ComMsg.B)
				n.CurrentView += 1
				n.AlreadyConfirmedTx[ComMsg.B.Tx] = 1
				n.ReadyForDelegate()
				n.State = "ready"
				fmt.Printf("Node %d trans to ready and changes primary to %v\n", n.Id, n.PrimaryNode)
				if num, ok := m.Load(ComMsg.B.Tx); ok {
					if num.(int) <= 1/3*numnodes {
						m.Store(ComMsg.B.Tx, num.(int)+1)
					} else {
						fmt.Printf("tx %v has been confirmed: %v\n", ComMsg.B.Tx, time.Now().UnixMilli())
						m.Store(ComMsg.B.Tx, -1000)
					}
				} else {
					m.Store(ComMsg.B.Tx, 1)
				}

			}
		} else if ComMsg.View > n.CurrentView {
			n.CommitMsg <- ComMsg
		}
	default:
		return
	}
}

//func (n *Node) Broadcast(Msg Message) {
//	for i := 0; i < numnodes-1; i++ {
//		if i != n.Id {
//			go func() {
//				conn, err := net.Dial("tcp", n.Addrs[i])
//				if err != nil {
//					log.Fatal("Error connecting:", err)
//				}
//				defer func(conn net.Conn) {
//					err := conn.Close()
//					if err != nil {
//
//					}
//				}(conn)
//				data, err := json.Marshal(Msg)
//				if err != nil {
//					log.Fatal("Error Marshal:", err)
//				}
//				_, err = conn.Write(data)
//				if err != nil {
//					log.Fatal("Error sending:", err)
//				}
//			}()
//		}
//	}
//	fmt.Printf("Node %d finish broadcasting %v\n", n.Id, Msg.MsgType)
//}

func (n *Node) Broadcast(Msg Message) {
	<-time.After(100 * time.Millisecond)
	switch Msg.MsgType {
	case "transaction":
		n.Sendtx(Msg)
	case "preprepare":
		n.SendPreprepare(Msg)
	case "prepare":
		n.SendPrepare(Msg)
	case "commit":
		n.SendCommit(Msg)
	case "delegate":
		n.SendDelegate(Msg)
	default:
	}
}

func (n *Node) Sendtx(Msg Message) {
	for _, Ni := range Nodes {
		if _, ok := Ni.AlreadySeenTxs[Msg.Tx]; ok != true {
			Ni.TxPools <- Msg.Tx
			Ni.AlreadySeenTxs[Msg.Tx] = 1
			fmt.Printf("Node %v first receives tx %v\n", Ni.Id, Msg.Tx)
		}
	}
}

func (n *Node) SendPreprepare(Msg Message) {
	for _, Ni := range Nodes {
		if Ni.Id != n.Id {
			Ni.PPMessages <- Msg
		}
	}
}
func (n *Node) SendPrepare(Msg Message) {
	for _, Ni := range Nodes {
		if Ni.Id != n.Id {
			Ni.PMessages <- Msg
		}
	}
}

func (n *Node) SendCommit(Msg Message) {
	for _, Ni := range Nodes {
		if Ni.Id != n.Id {
			Ni.CommitMsg <- Msg
		}
	}
}

func (n *Node) SendDelegate(Msg Message) {
	for _, Ni := range Nodes {
		if Ni.Id != n.Id {
			Ni.DelegatedPool <- Msg
		}
	}
}

func (n *Node) SpecialNode() {
	select {
	case Msg := <-n.CommitMsg:
		if Msg.View == n.CurrentView {
			if _, ok := n.BlocksVotes[Msg.B.Hash]; ok != true {
				n.BlocksVotes[Msg.B.Hash] = make(map[string]int)
			}
			n.BlocksVotes[Msg.B.Hash]["commit"] += 1
			if n.BlocksVotes[Msg.B.Hash]["commit"] > 2/3*numnodes {
				source := rand.NewSource(time.Now().UnixNano())
				random := rand.New(source)
				rannum := random.Intn(numnodes-1+1) + 1
				n.CurrentView += 1
				Msg := Message{
					MsgType:  "delegate",
					Senderid: rannum,
					View:     Msg.View + 1,
				}
				n.Broadcast(Msg)
				fmt.Printf("SP delegate Node%v in view %v\n", rannum, n.CurrentView)
			}
		} else if Msg.View > n.CurrentView {
			n.CommitMsg <- Msg
		}
	default:
		return
	}
}

func main() {
	for i := 0; i < numnodes; i++ {
		Nodes = append(Nodes, NewNode(i))
	}
	for i := 0; i < numnodes; i++ {
		//go Nodes[i].Listen()
		go Nodes[i].ProcessMsg()
		go Nodes[i].Propose()
	}
	time.Sleep(3 * time.Second)
	//for i := 1; i <= numoftxs; i++ {
	//	var Tx = strconv.Itoa(i)
	//	TxMsg := Message{
	//		Tx:      Tx,
	//		MsgType: "transaction",
	//	}
	//	conn, err := net.Dial("tcp", "0.0.0.0:"+strconv.Itoa(5000+i))
	//	if err != nil {
	//		log.Fatal("Send tx error", err)
	//	}
	//	data, err := json.Marshal(TxMsg)
	//	if err != nil {
	//		log.Fatal("Tx marshal err", err)
	//	}
	//	_, err = conn.Write(data)
	//	<-time.After(10 * time.Millisecond)
	//
	//}
	for i := 1; i <= numoftxs; i++ {
		tx := strconv.Itoa(i)
		TxMsg := Message{MsgType: "transaction", Tx: tx}
		<-time.After(200 * time.Millisecond)
		Nodes[i].Broadcast(TxMsg)
		fmt.Printf("tx %v has been sent,time:%v\n", tx, time.Now().UnixMilli())
	}

	//Mmonitor := func(key interface{}, value interface{}, quit chan struct{}) {
	//	for {
	//		select {
	//		case <-quit:
	//			return
	//		default:
	//			_, already := m2.Load(key)
	//			if val, ok := value.(int); ok && val > 1/3*numnodes && already == false {
	//				fmt.Printf(
	//					"Transaction:%s has been confirmed,time:%v\n",
	//					key,
	//					time.Now().UnixMilli(),
	//				)
	//				m2.Store(key, 1)
	//				close(quit)
	//			}
	//		}
	//	}
	//}
	//quitofkey := make(map[interface{}]chan struct{})
	//for {
	//	m.Range(func(key, value any) bool {
	//		quitofkey[key] = make(chan struct{})
	//		go Mmonitor(key, value, quitofkey[key])
	//		return true
	//	})
	//}
}
