package learn_bitcoin

import (
	"crypto/sha256"
	"fmt"
	"time"
)

// Block represents the blockchain block
type Block struct {
	timestamp    time.Time
	transactions []string
	prevHash     []byte
	Hash         []byte
}

func NewBlock(transactions []string, prevHash []byte) *Block {
	currentTime := time.Now()
	return &Block{
		timestamp:    currentTime,
		transactions: transactions,
		prevHash:     prevHash,
		Hash:         NewHash(currentTime, transactions, prevHash),
	}
}

func NewHash(currentTime time.Time, transactions []string, prevHash []byte) []byte {
	input := append(prevHash, currentTime.String()...)
	for _, transaction := range transactions {
		input = append(input, string(transaction)...)
	}
	hash := sha256.Sum256(input)
	return hash[:]
}

func PrintBlockInformation(block *Block) {
	fmt.Printf("\ttime: %s\n", block.timestamp.String())
	fmt.Printf("\tprevHash: %x\n", block.prevHash)
	fmt.Printf("\tHash: %x\n", block.Hash)
	printTransactions(block)
}

func printTransactions(block *Block) {
	fmt.Println("\tTransactions: ")
	for i, transaction := range block.transactions {
		fmt.Printf("\t\t%v: %q\n", i, transaction)
	}
}
