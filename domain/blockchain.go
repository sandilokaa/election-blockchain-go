package domain

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Blockchain struct {
	Pool  []*Mandate `json:"pool"`
	Chain []*Block   `json:"chain"`
}

func NewBlockchain() *Blockchain {
	bc := LoadDatase()
	if len(bc.Chain) == 0 {
		bc.CreateBlock(fmt.Sprintf("%x", [32]byte{}))
	}
	return bc
}

func LoadDatase() *Blockchain {
	f, err := os.OpenFile("database/blockchain.db", os.O_APPEND|os.O_RDWR, 0600)
	if err != nil {
		os.Exit(1)
	}

	scanner := bufio.NewScanner(f)
	blockchain := Blockchain{}

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			os.Exit(1)
		}

		var blockSerialized BlockSerialized
		err = json.Unmarshal(scanner.Bytes(), &blockSerialized)
		if err != nil {
			os.Exit(1)
		}

		blockchain.Chain = append(blockchain.Chain, blockSerialized.Value)
	}

	return &blockchain
}

func (bc *Blockchain) CreateBlock(prevHash string) *Block {
	b := NewBlock(prevHash, bc.Pool)
	bc.Chain = append(bc.Chain, b)
	bc.Pool = []*Mandate{}
	return b
}

func (bc *Blockchain) LastestBlock() *Block {
	return bc.Chain[len(bc.Chain)-1]
}

func (bc *Blockchain) GiveMandate(from, to string, value int8) {
	m := NewMandate(from, to, value)
	bc.Pool = append(bc.Pool, m)
}
