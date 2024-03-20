package interfaces

import (
	"bytes"
	"crypto/rand"
	"encoding/gob"
	"log"

	"golang.org/x/crypto/sha3"
)


var KeyBytes = 32
var KeyBits = 256


var Blockchain_t = 0
var Client_t = 1
var Routing_t = 2
type Message struct {
	Sender 		string
	Recipient 	string

	Message_t	int
	Data 		[]byte
	Err 		error

	Sum 		[]byte
}

func GenerateKey() []byte {
	key := make([]byte, KeyBytes)
	_, err := rand.Read(key)
	if err != nil {
		return nil
	}
	return key
}

func Hash_SHA3_256(b any) []byte {

	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(b)
	if err != nil {
		log.Fatal("encode error: ", err)
	}

	hash := sha3.Sum256(buf.Bytes())

	return hash[:]

}