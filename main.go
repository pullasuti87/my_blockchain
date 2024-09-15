package main

import (
	"crypto/sha256"
	"fmt"
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
