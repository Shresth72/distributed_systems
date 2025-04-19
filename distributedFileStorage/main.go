package main

import (
	"fmt"
	"log"

	"github.com/Shresth72/dfs/pkg/p2p"
)

func onPeer(peer p2p.Peer) error {
	fmt.Printf("Doing some logic with the peer logic, outside of the TCP transport\n")
	return nil
}

func main() {
	trConfig := p2p.TCPTransportConfig{
		ListenAddr: ":3000",
		ShakeHands: p2p.NOPHandshakeFunc,
		Decoder:    p2p.DefaultDecoder{},
		OnPeer:     onPeer,
	}
	tr := p2p.NewTCPTransport(trConfig)

	go func() {
		for {
			msg := <-tr.Consume()
			fmt.Printf("%+v from %s\n", msg.Payload, msg.From)
		}
	}()

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
