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

type TCPTrasport struct {
	listenAddress string
	listener net.Listener

	mu sync.RWMutex  // good practice to have mutex above the part which u want to project 
	peers map[net.Addr]Peer
}

func NewTCPTrasport(listenAddr string) *TCPTrasport{
	return &TCPTrasport{
		listenAddress: listenAddr,
	}
}



func (t *TCPTrasport) ListenAndAccept() error {
	var err error 
	
	t.listener, err = net.Listen("tcp",t.listenAddress)
	if err !=nil {
		return err  
	}

	go t.startAcceptLoop()


	return nil 
}


func (t *TCPTrasport) startAcceptLoop(){
	for {
		conn, err := t.listener.Accept()

		if err !=nil {
			fmt.Printf("")
		}

		go t.handleConn(conn)
	}
}	



func (t *TCPTrasport) handleConn(conn net.Conn){

}