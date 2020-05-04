package blockchain

//Block struct
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

//CreateBlock creates a block
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Nonce = nonce
	block.Hash = hash[:]
	return block
}

//Genesis initializes a block chain
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}
