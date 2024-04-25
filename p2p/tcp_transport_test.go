package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	listenAddresss := ":4000"
	opts := TCPTransportOpts{
		ListenAddr: ":4000",
	}
	tr := NewTCPTransport(opts)

	assert.Equal(t, tr.ListenAddr, listenAddresss)

	assert.Nil(t, tr.ListenAndAccept())
}
