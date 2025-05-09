package p2p

import "net"

// RPC holds any arbitary data sent over each transport
// between two nodes in the network
type RPC struct {
	From    net.Addr
	Payload []byte
}
