package assignment02IBC

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

type Block struct {
	PreviousBlock *Block
	HashValue     string
	Transaction   string
}

func hashBlock(block string) string {

	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte("secret"))

	// Write Data to it
	h.Write([]byte(block))

	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(h.Sum(nil))

	return sha

}

func InsertBlock(transaction string, chainHead *Block) *Block {

	var newBlock *Block = new(Block)

	if chainHead == nil {

		newBlock.transaction = transaction
		newBlock.previousBlock = chainHead
		newBlock.hashValue = ""
		println("Genesis Block Added")
	} else {

		newBlock.transaction = transaction
		newBlock.previousBlock = chainHead
		newBlock.hashValue = hashBlock(chainHead.transaction + chainHead.hashValue)
		println("New Block Added")
	}

	return newBlock
}

func VerifyChain(chainHead *Block) string {

	if chainHead.previousBlock == nil { //genesis Node

		return hashBlock(chainHead.transaction)

	} else {
		blockHash := VerifyChain(chainHead.previousBlock)

		if blockHash == chainHead.hashValue {
			println("Hash Matches")

		} else {
			println("Hash Does Not Match")

		}

		return hashBlock(chainHead.transaction + blockHash)

	}
}

func ChangeBlock(oldTrans string, newTrans string, chainHead *Block) {

	if chainHead.previousBlock == nil { //Genesis Block

		if chainHead.transaction == oldTrans { //If required block is found

			chainHead.transaction = newTrans
			println("Block Changed")
		}

	} else { //recursivly iterate to the required block

		ChangeBlock(oldTrans, newTrans, chainHead.previousBlock)

		if chainHead.transaction == oldTrans { //If required block is found

			chainHead.transaction = newTrans
			println("Block Changed")
		}
	}
}

func ListBlocks(chainHead *Block) {

	if chainHead.previousBlock == nil { //genesis Node

		println("Transaction: " + chainHead.transaction)
		println("Genesis Block.\n")

	} else {
		println("Transaction " + chainHead.transaction)
		println("Hash Value: " + chainHead.hashValue)

		ListBlocks(chainHead.previousBlock)
	}
}
