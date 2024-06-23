package p2p

import (
	"fmt"
	"net"
	"sync"
)

/**
Note:
	1. public code should be at top and private code should be at bottom
	2. important functions should always be at the top visa versa
*/

// TCPPeer represents the remote node over a TCP established connection.
type TCPPeer struct{
	// conn is the underlying connection of the peer 
	conn 		net.Conn

	// if we dail and retrieve a conn => outbound == true 
	// if we accept and retrieve a conn => outbound == false 
	outbound 	bool 
}


func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer{
	return &TCPPeer{
		conn: conn, 
		outbound:  outbound,
	}
}


type TCPTransport struct {
	listenAddress string
	listener net.Listener

	mu sync.RWMutex  // good practice to have mutex above the part which u want to project 
	peers map[net.Addr]Peer
}

func NewTCPTransport(listenAddr string) *TCPTransport{
	return &TCPTransport{
		listenAddress: listenAddr,
	}
}



func (t *TCPTransport) ListenAndAccept() error {
	var err error 
	
	t.listener, err = net.Listen("tcp",t.listenAddress)
	if err !=nil {
		return err  
	}

	go t.startAcceptLoop()


	return nil 
}


func (t *TCPTransport) startAcceptLoop(){
	for {
		conn, err := t.listener.Accept()

		if err !=nil {
			fmt.Printf("TCP accept error: %s\n", err)
		}

		go t.handleConn(conn)
	}
}	



func (t *TCPTransport) handleConn(conn net.Conn){
	peer := NewTCPPeer(conn, true)

	fmt.Printf("new incoming connection %+v\n", peer)
}

