package main

import (
	"log"

	"github.com/Jitulteron7/p2p"
)

func main() {
	tr := p2p.NewTCPTrasport(":3000")

	if err := tr.ListenAndAccept(); err !=nil{
		log.Fatal(err)
	}
	select {}
	// fmt.Println("Hello")
}