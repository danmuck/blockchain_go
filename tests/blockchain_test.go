package tests

import (
	"blockchain_go/blockchain"
	"fmt"
	"os"
	"testing"
	"time"
)


type LogTimer struct {
	start 			time.Time
	duration 		time.Duration

	origin			time.Time
}

func NewTimer() *LogTimer {
	now := time.Now()
	t := &LogTimer{start: now, origin: now}
	return t
}

func (t *LogTimer) Timer(print bool) {
	t.duration = time.Since(t.start)
	t.start = time.Now()
	if print {
		fmt.Printf(" %+v sec \n", t.duration.Seconds())
	}
}


func TestInit(t *testing.T) {

	chain := blockchain.Initialize()

	fmt.Fprintf(os.Stderr, "\nchain signature: %+v \n", blockchain.Chain_sig)

	root := chain.GetRoot()


	make_blocks := 0
	max_blocks := 100
	print_time := false
	timer := NewTimer()
	for make_blocks < max_blocks {
		str := "hi"
		new_block := blockchain.NewBlock(root, blockchain.Chain_sig, []byte(str), nil)
		err := chain.AppendBlock(*new_block,
				blockchain.Chain_sig,
				*blockchain.NewProof(blockchain.Chain_sig, chain))
		if err != nil {
			fmt.Fprintf(os.Stderr, "\n\n %+v \n\n", err)

		}
		root = new_block
		make_blocks++

		err = chain.ValidateChain()
		if err != nil {
			fmt.Fprintf(os.Stderr, "\n\n %+v \n\n", err)

		}
		if make_blocks % 100 == 0 {
			print_time = true
			fmt.Printf("#%v ", make_blocks)
		} else { print_time = false }

		timer.Timer(print_time)

	}
	fmt.Printf("total time: %v sec", time.Since(timer.origin).Seconds())
	chain.ChainString()

}
