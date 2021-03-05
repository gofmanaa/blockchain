package block

import (
	"bytes"
	"crypto/sha1"
	"fmt"
)

type Block struct {
	hash     []byte
	data     []byte
	prevHash []byte
}

func(b *Block) hashBlock() {
	payload := bytes.Join([][]byte{b.data, b.prevHash}, []byte("0000")) 
	hash := sha1.Sum(payload)
	b.hash = hash[:]
}

func CreateNewBlock(data []byte, prevHash []byte) *Block {
	block := &Block{
		hash: []byte{},
		data: data,
		prevHash: prevHash,	
	}

	block.hashBlock()

	return block
}

func (b *Block) GetPrevHash() []byte {
	return b.prevHash
}

func (b *Block) GetData() []byte {
	return b.data
}

func (b *Block) GetHash() []byte {
	return b.hash
}

func (b *Block) String() string {
	return fmt.Sprintf("hash:%x, prevHash:%x, data:%s", b.hash, b.prevHash,b.data)
}


