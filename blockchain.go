package main

import (
	"crypto/sha256"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

type Block struct {
	index        int64
	hash         string
	previousHash string
	timestamp    int64
	data         string
}

var (
	genesisBlock = Block{
		0,
		"816534932c2b7154836da6afc367695e6337db8a921823784c14378abed4f7d7",
		"",
		1465154705,
		"my genesis block!!",
	}
	blockchain = []Block{genesisBlock}
)

func GetBlockchain() []Block { return blockchain }

func GetLatestBlock() Block { return blockchain[len(blockchain)-1] }

func isValidNewBlock(newBlock Block, previousBlock Block) bool {
	if previousBlock.index+1 != newBlock.index {
		fmt.Println("invalid index")
		return false
	} else if previousBlock.hash != newBlock.previousHash {
		fmt.Println("invalid previoushash")
		return false
	} else if calculateHashForBlock(newBlock) != newBlock.hash {
		fmt.Println("invalid hash: " + calculateHashForBlock(newBlock) + " " + newBlock.hash)
		return false
	}
	return true
}

func addBlock(newBlock Block) {
	if isValidNewBlock(newBlock, GetLatestBlock()) {
		blockchain = append(blockchain, newBlock)
	}
}

func isValidChain(blockChain []Block) bool {
	if !reflect.DeepEqual(blockChain[0], genesisBlock) {
		return false
	}
	for i := 1; i < len(blockChain); i++ {
		if !isValidNewBlock(blockChain[i], blockChain[i-1]) {
			return false
		}
	}
	return true
}

func AddToBlockChain(newBlock Block) bool {
	if isValidNewBlock(newBlock, GetLatestBlock()) {
		blockchain = append(blockchain, newBlock)
		return true
	}
	return false
}

func ReplaceChain(newBlockChain []Block) {
	if isValidChain(newBlockChain) && len(newBlockChain) > len(GetBlockchain()) {
		fmt.Println("Received blockchain is valid. Replacing current blockchain with received blockchain")
		blockchain = newBlockChain
		// TODO broadcastLatest
	} else {
		fmt.Println("Received blockchain invalid")
	}
}

func GenerateNextBlock(blockData string) Block {
	previousBlock := GetLatestBlock()
	nextIndex := previousBlock.index + 1
	nextTimestamp := time.Now().Unix()
	nextHash := calculateHash(nextIndex, previousBlock.hash, nextTimestamp, blockData)
	newBlock := Block{nextIndex, nextHash, previousBlock.hash, nextTimestamp, blockData}
	addBlock(newBlock)
	// TODO broadcastLatest
	return newBlock
}

func calculateHash(index int64, previousHash string, nextTimestamp int64, blockData string) string {
	hash := sha256.New()
	hash.Write([]byte(strconv.FormatInt(index, 10)))
	hash.Write([]byte(previousHash))
	hash.Write([]byte(strconv.FormatInt(nextTimestamp, 10)))
	hash.Write([]byte(blockData))
	return string(hash.Sum(nil))
}

func calculateHashForBlock(block Block) string {
	return calculateHash(block.index, block.previousHash, block.timestamp, block.data)
}
