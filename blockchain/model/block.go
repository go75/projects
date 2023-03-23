package model

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	// 当前区块的创建时间
	Timestamp int64
	// 当时区块实际存放的数据
	Data []byte
	// 上一个区块的hash
	PreHash [32]byte
	// 当前区块的hash
	Hash [32]byte
}

func (b *Block) SetHash() {
	b.Hash = sha256.Sum256(bytes.Join([][]byte{b.PreHash[:], b.Data, []byte(strconv.FormatInt(b.Timestamp, 10))}, []byte{}))
}

func NewGenesisBlock() *Block {
	return NewBlock("", [32]byte{})
}

func NewBlock(data string, preHash [32]byte) *Block {
	b := &Block{time.Now().Unix(), []byte(data), preHash, [32]byte{}}
	b.SetHash()
	return b
}