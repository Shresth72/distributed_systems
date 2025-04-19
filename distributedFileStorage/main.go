package main

import (
	"log"

	"github.com/Shresth72/dfs/pkg/p2p"
)

func main() {
	trConfig := p2p.TCPTransportConfig{
		ListenAddr: ":3000",
		ShakeHands: p2p.NOPHandshakeFunc,
		Decoder:    p2p.DefaultDecoder{},
	}
	tr := p2p.NewTCPTransport(trConfig)

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
