package blockchain

import "blockchain_go/interfaces"

func (t *Transaction) newTxn() *Transaction {

	t = &Transaction{
		id: interfaces.GenerateKey(),
		root: interfaces.GenerateKey(),
		Data: interfaces.GenerateKey(),
	}

	return t
}