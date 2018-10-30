package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"time"
)

// 区块
type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
	Delegate  string
}

// 生成区块
func generateBlock(oldBlock Block, BPM int, address string) (Block, error) {
	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateBlockHash(newBlock)
	newBlock.Delegate = address

	return newBlock, nil
}

// 区块链
var Blockchain []Block

// 受托人
var delegates = []string{"001", "002", "003", "004", "005"}

// 当前delegates 的索引
var indexDelegate int

// 生成hash字符串
func calculateBlockHash(block Block) string {
	data := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	return calculateHash(data)
}

// 生成区块的hash值
func calculateHash(s string) string {

	hash32 := sha256.Sum256([]byte(s))
	hash := hash32[:]
	return hex.EncodeToString(hash)
}

// 验证区块是否有效
func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}
	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}
	if calculateBlockHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

// 更换受托人的排列顺序
func randDelegates(delegates []string) []string {
	var randList []string
	randList = delegates[1:]
	randList = append(randList, delegates[0])

	fmt.Printf("%v\n", randList)
	return randList
}

func main() {
	indexDelegate = 0

	// 创世区块
	t := time.Now()
	genesiBlock := Block{}
	genesiBlock.Timestamp = t.String()
	genesiBlock.Index = 0
	genesiBlock.BPM = 0
	genesiBlock.PrevHash = ""
	genesiBlock.Hash = calculateBlockHash(genesiBlock)
	genesiBlock.Delegate = ""

	Blockchain = append(Blockchain, genesiBlock)

	indexDelegate++

	countDelegate := len(delegates)

	for indexDelegate < countDelegate {

		//3s 生成区块
		time.Sleep(time.Second * 3)
		fmt.Println(indexDelegate)

		// 创建区块
		rand.Seed(int64(time.Now().Unix()))
		bpm := rand.Intn(100)
		oldLastIndex := Blockchain[len(Blockchain)-1]
		newBlock, err := generateBlock(oldLastIndex, bpm, delegates[indexDelegate])
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Printf("Blockchain...%v\n", newBlock)

		if isBlockValid(newBlock, oldLastIndex) {
			Blockchain = append(Blockchain, newBlock)
		}

		indexDelegate = (indexDelegate + 1) % countDelegate
		if indexDelegate == 0 {
			delegates = randDelegates(delegates)
		}

	}

}
