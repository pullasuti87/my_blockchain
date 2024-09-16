package main

import (
	"crypto/sha256"
	"fmt"
	// "time"
	"encoding/hex"
	"strconv"
)

/*
this is one block
very simplified version
https://www.investopedia.com/terms/b/block-bitcoin-block.asp
*/
type Block struct {
	Id        int
	Timestamp int64
	PrevHash  string
	Hash      string
	Data      string
}

type Blockchain struct {
	Chain []*Block
}

// pointer receiver, reference to block
func (b *Block) InitHash() {
	// strconv --> integer to str (integer to ascii)
	data := strconv.Itoa(b.Id) + strconv.FormatInt(b.Timestamp, 10) + b.PrevHash + b.Data
	hash := sha256.Sum256([]byte(data))
	b.Hash = hex.EncodeToString(hash[:])
}

func FirstBlock() *Block {
	firstBlock := &Block{
		Id:        0,
		Timestamp: 1234567890,
		PrevHash:  "",
		Hash:      "",
		Data:      "",
	}

	return firstBlock
}

func CreateBlock() {}

func (bc *Blockchain) AddBlock() {}

func main() {
	fmt.Printf("test\n")

	//	blockchain := &Blockchain{
	//		Chain: []*Block{FirstBlock()},
	//	}
	//
	//    blockchain.AddBlock()
	//    blockchain.AddBlock()

}
