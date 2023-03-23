package model

type Chain struct {
	blocks []*Block
}

func (c *Chain) AddBlock(data string) {
	c.blocks = append(c.blocks, NewBlock(data, c.blocks[len(c.blocks)-1].Hash))
}

func NewBlockChain() *Chain {
	return &Chain{
		blocks: []*Block{NewGenesisBlock()},
	}
}