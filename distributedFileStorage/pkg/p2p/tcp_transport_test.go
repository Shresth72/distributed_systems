package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	trConfig := TCPTransportConfig{
		ListenAddr: ":3000",
		ShakeHands: NOPHandshakeFunc,
		Decoder:    DefaultDecoder{},
	}
	tr := NewTCPTransport(trConfig)
	assert.Equal(t, ":3000", tr.ListenAddr)

	assert.NoError(t, tr.ListenAndAccept())
}
