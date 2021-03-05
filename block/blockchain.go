package block

type Blockchain struct {
	length int
	blocks []*Block
}

func InitBlockchain() *Blockchain {
	blockchain := &Blockchain{}
	initBlock := CreateNewBlock([]byte("Init Block"), []byte{})
	blockchain.blocks = append(blockchain.blocks, initBlock)
	blockchain.length++

	return blockchain 
}

func (self *Blockchain) AddNewBlock(data []byte) {
	prevBlock := self.blocks[self.length-1]
	block := CreateNewBlock(data, prevBlock.hash)
	self.blocks = append(self.blocks, block)
	self.length++
}

func (self * Blockchain) GetBlocks() []*Block {
	return self.blocks
}
