package interfaces


type Block interface {
	GetIndex() int
}


type Blockchain interface {
	
	GetRoot() Block
	GetTop() Block

	FindBlock(block_hash []byte) *Block
	
	AppendBlock(block Block, signature []byte, proof Proof_of_Work) error
	
	ValidateChain() error
	
}


type Transaction interface {

}


type Proof_of_Work interface {

}