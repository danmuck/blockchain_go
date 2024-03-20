package blockchain

import (
	"blockchain_go/interfaces"
	"bytes"
	"fmt"
	"os"
)


func (p *Proof_of_Work) verify() bool {
	return true
}

func NewProof(auth []byte, chain *Blockchain) *Proof_of_Work {
	proof := &Proof_of_Work{
		Author: auth,
		Blockchain: chain.Id,
	}

	return proof
}

func (bc *Blockchain) FullValidateChain() error {
	root := bc.chain[0]
	root_hash := interfaces.Hash_SHA3_256(bc.root)

	hash_previous := interfaces.Hash_SHA3_256(root)
	if !bytes.Equal(root_hash, hash_previous) {
		return fmt.Errorf("root no match")
	}

	fmt.Fprintf(os.Stderr, "root: %+v::%+v \n", root_hash, hash_previous)

	index := 0
	for idx, block := range bc.chain {
		if idx == 0 {
			continue
		}
		if idx != block.Index {
			return fmt.Errorf("index: (real: %v, block: %v)", index, block.Index)
		}
		hash_current := interfaces.Hash_SHA3_256(block)
		
		fmt.Fprintf(os.Stderr, "other: %+v::%+v \n", hash_current, hash_previous)
		if !bytes.Equal(block.Previous_Block, hash_previous) {
			return fmt.Errorf("block #%v no match", idx)
		}
		hash_previous = hash_current
	}

	return nil

}