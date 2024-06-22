package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T){
	listenAddr := ":4000"
	tr := NewTCPTrasport(listenAddr)
	assert.Equal(t, tr.listenAddress, listenAddr)
}