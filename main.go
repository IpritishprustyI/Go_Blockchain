package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type BlockChain struct { //this represents the block chain
	blocks []*Block
}

func (chain *BlockChain) AddBlock(data string) { //this will add blocks to the block chain
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{}) // this the first block function with data as Genesis and an empty previous hash

}
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}} //this will create the first blockchain by returning the block with genesis file

}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{}) //this takes a 2d slice and passes Data and previous hash
	//and joins then into a slice of byte and its stored inside the info var
	hash := sha256.Sum256(info) //the above process is then used to create hash
	b.Hash = hash[:]            //the created hash is pushed inside the block Hash
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash} //Here block constructor is used and block is created using the
	//hash value(the 1st arg in &Block),then the string data is converted into byte(2nd arg),then previous hash is taken

	block.DeriveHash() //Then the deriveHash function is called on our block
	return block
}
func main() {
	chain := InitBlockChain()
	chain.AddBlock("First block after Genesis")
	chain.AddBlock("Second block after Genesis")
	chain.AddBlock("third block after Genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

	}

}
