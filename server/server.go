package server

import (
	// "blockchain/api"
)

type Message struct {
	Sender 		string
	Recipient 	string
	data 		[]byte
	sum 		[]byte
}


var RootServer = 0
var FullNode = 1
var HalfNode = 2
type Node struct {
	Node_t		int
	Id 			[]byte
	routing 	[]byte
	chain_id 	[]byte
	data 		[]byte
}
