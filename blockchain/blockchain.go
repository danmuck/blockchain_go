package blockchain

import (
	"blockchain_go/interfaces"
	"bytes"
	"fmt"
	"os"

	"time"
)


func (bc *Blockchain) GetRoot() *Block {

	return &bc.root

}


func (bc *Blockchain) GetTop() *Block {

	return &bc.top

}


func (bc *Blockchain) FindBlock(block_hash []byte) *Block {
	for _, block := range bc.chain {
		hash_current := interfaces.Hash_SHA3_256(block)
		
		if bytes.Equal(hash_current, block_hash) {
			return &block
		}
	}
	
	return nil

}


func (bc *Blockchain) ValidateChain() error {
	root := bc.chain[0]
	root_hash := interfaces.Hash_SHA3_256(bc.root)

	hash_previous := interfaces.Hash_SHA3_256(root)
	if !bytes.Equal(root_hash, hash_previous) {
		return fmt.Errorf("root no match")
	}

	for idx, block := range bc.chain {
		if idx == 0 {
			continue
		}
		hash_current := interfaces.Hash_SHA3_256(block)
		
		if !bytes.Equal(block.Previous_Block, hash_previous) {
			return fmt.Errorf("block #%v no match", idx)
		}
		hash_previous = hash_current
	}

	return nil

}


func (bc *Blockchain) AppendBlock(block Block, signature []byte, proof Proof_of_Work) error {

	if bytes.Equal(block.Signature, signature) &&
		bytes.Equal(proof.Author, signature) &&
		bytes.Equal(block.Signature, proof.Author) &&
		bytes.Equal(interfaces.Hash_SHA3_256(bc.top), block.Previous_Block) {

			bc.chain = append(bc.chain, block)
			bc.top = block
	}

	return nil

}


var Chain_sig []byte
func Initialize() *Blockchain {
	key := interfaces.GenerateKey()
	if key == nil {
		return nil
	}
	time := time.Now().UnixNano()

	welcome := fmt.Sprintf(`Chain signed: %v at: %v`, key, time)
	g := Block{
		Index: 0,
		Timestamp: time,
		Signature: key,

		Nonce: interfaces.GenerateKey(),
		Previous_Block: key,

		Chain_Data: []byte(welcome),
		Transactions: nil,
	}

	b := &Blockchain{
		Id: key,
		timestamp: time,
		root: g,
		top: g,
		chain: make([]Block, 0),
		Chain_Data: make([]byte, 0),
		Txn_Blocks: make([]*Block, 0),
	}
	b.chain = append(b.chain, g)
	Chain_sig = interfaces.Hash_SHA3_256(b)

	return b

}


func (bc *Blockchain) ChainString() error {
	for idx, block := range bc.chain {
		fmt.Fprintf(os.Stderr ,`
Block: [%v|%v]
	time:    %+v
	nonce:   %+v
	signed:  %+v
	data:    %+v
	txns:    %+v
	prev:	 %+v
`,
		idx,
		block.Index, 
		time.Unix(0, block.Timestamp), 
		block.Nonce, 
		block.Signature, 
		block.Chain_Data, 
		block.Transactions,
		block.Previous_Block)

	}

	return nil
}