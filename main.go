package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
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

func InitBlock() *Block {
	firstBlock := &Block{
		Id:        0,
		Timestamp: time.Now().UnixNano(),
		PrevHash:  "",
		Hash:      "",
		Data:      "init block",
	}

	firstBlock.InitHash()

	return firstBlock
}

func CreateBlock(prev *Block, data string) *Block {
	block := &Block{
		Id:        prev.Id + 1,
		Timestamp: time.Now().UnixNano(),
		PrevHash:  prev.Hash,
		Data:      data,
	}

	block.InitHash()

	return block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Chain[len(bc.Chain)-1]
	newBlock := CreateBlock(prevBlock, data)
	bc.Chain = append(bc.Chain, newBlock)
}

func (bc *Blockchain) PrintData() {
	for _, block := range bc.Chain {
		fmt.Printf("Id: %d\n", block.Id)
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Printf("Previous Hash: %s\n", block.PrevHash)
		fmt.Printf("Current Hash: %s\n", block.Hash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("---------------------\n")
		fmt.Println()
	}
}

func main() {
	// fmt.Printf("test\n")

	blockchain := &Blockchain{
		Chain: []*Block{InitBlock()},
	}

	// test
	if len(blockchain.Chain) != 1 {
		fmt.Printf("failed to initialize blockchain")
	} else {
		fmt.Printf("blockchain initialized with firstblock\n")
		fmt.Println()
	}

	// blockchain.PrintData()

	blockchain.AddBlock("first added")
	time.Sleep(1 * time.Second)
	blockchain.AddBlock("second added")
	time.Sleep(2 * time.Second)
	blockchain.AddBlock("third added")

	blockchain.PrintData()

}
