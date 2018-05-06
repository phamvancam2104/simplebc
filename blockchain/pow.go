package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
	"simplebc/utils"
)

const MaxInt64 int64 = 1<<63 - 1

const difficulty = 15

type PoW struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *PoW {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-difficulty))
	pow := &PoW{b, target}
	return pow
}

func (pow *PoW) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevHash,
			pow.block.Data,
			utils.IntToHex(pow.block.Timestamp),
			utils.IntToHex(int64(difficulty)),
			utils.IntToHex(int64(nonce)),
		},
		[]byte{})
	return data
}

func (pow *PoW) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)

	for int64(nonce) < MaxInt64 {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	return nonce, hash[:]
}

func (pow *PoW) Validate() bool {
	var hashInt big.Int
	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	isValid := hashInt.Cmp(pow.target) == -1
	return isValid
}
