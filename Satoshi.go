package assignment02IBC

import (
	"encoding/gob"
	"log"
	"math/rand"
	"net"
)

func SendBlockchain(port net.Conn, chainHead *Block) {
	gobEncoder := gob.NewEncoder(port)
	err := gobEncoder.Encode(&chainHead)
	if err != nil {
		log.Println(err)
	}
}

func SendPeers(port net.Conn, peers []string) {
	gobEncoder := gob.NewEncoder(port)
	err := gobEncoder.Encode(peers)
	if err != nil {
		log.Println(err)
	}
}

func SendTrans(trans string, miner string) {
	conn, err := net.Dial("tcp", miner)
	if err != nil {
		log.Println(err)
	}
	println("\nMiner: ", miner)
	trans1 := "Transaction"

	gobEncoder := gob.NewEncoder(conn)
	err = gobEncoder.Encode(&trans1)

	gobEncoder = gob.NewEncoder(conn)
	err = gobEncoder.Encode(&trans)

	println("\nTransaction Sent")

	if err != nil {
		log.Println(err)
	}

	conn.Close()
}

func SelectMiner(nodes int, peersPort []string, port string) string {
	miner := 0
	for {
		miner = rand.Intn(nodes)
		if port != peersPort[miner] {
			break
		}
	}
	return peersPort[miner]
}
