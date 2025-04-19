package p2p

import (
	"fmt"
	"io"
	"net"
	"sync"
)

// TCPPeer represents the remote node over a TCP established connection
type TCPPeer struct {
	// conn is the underlying connection of the peer
	conn net.Conn

	// if we dial and retrieve a conn => outbound == true
	// if we accept and retrieve a conn > outbound == false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransportConfig struct {
	ListenAddr string
	ShakeHands HandshakeFunc
	Decoder    Decoder
}

type TCPTransport struct {
	TCPTransportConfig
	listener net.Listener

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(config TCPTransportConfig) *TCPTransport {
	return &TCPTransport{
		TCPTransportConfig: config,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.ListenAddr)
	if err != nil {
		return err
	}

	go t.startAcceptLoop()
	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %v\n", err)
		}

		go t.handleConn(conn)
	}
}

type Temp struct{}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)
	fmt.Printf("New incoming connection: %+v\n", peer)

	if err := t.ShakeHands(peer); err != nil {
		fmt.Printf("TCP handshake error: %s\n", err)
		conn.Close()
		return
	}

	msg := &Message{}
	for {
		if err := t.Decoder.Decode(conn, msg); err != nil {
			if err == io.EOF {
				fmt.Printf("Connection closed by: %v\n", conn.RemoteAddr())
				return
			}
			fmt.Printf("TCP error: %s\n", err)
			continue
		}

		msg.From = conn.RemoteAddr()
		fmt.Printf("%+v from %s\n", msg.Payload, msg.From)
	}

}
