package blockchain

import (
	"blockchain_go/interfaces"

	"time"
)

func (b *Block) GetIndex() int {

	return b.Index

}

func NewBlock(prev interfaces.Block, author []byte, data []byte, txns []Transaction) *Block {
	key := interfaces.GenerateKey()

	prev_hashed := interfaces.Hash_SHA3_256(prev)

	b := &Block{
		Index: prev.GetIndex()+1,
		Timestamp: time.Now().UnixNano(),
		Signature: author,
		Chain_Data: data,
		Transactions: txns,
		Nonce: key,
		Previous_Block: prev_hashed,
	}
	return b

}