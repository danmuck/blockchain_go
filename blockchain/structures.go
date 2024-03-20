package blockchain


type Block struct {

	Index 			int
	Timestamp 		int64
	Signature 		[]byte

	Nonce 			[]byte
	Previous_Block	[]byte

	Chain_Data 		[]byte
	Transactions	[]Transaction

}


type Transaction struct {

	id 				[]byte
	root			[]byte

	Data			[]byte

}


type Blockchain struct {

	Id 				[]byte
	timestamp 		int64
	root			Block
	top			 	Block

	chain 			[]Block
	Chain_Data 		[]byte
	Txn_Blocks	 	[]*Block

}


type Proof_of_Work struct {

	Author 			[]byte
	Blockchain		[]byte
	Reward 			uint64

	start	 		int64
	finish			int64
	final_block 	*Block

}

