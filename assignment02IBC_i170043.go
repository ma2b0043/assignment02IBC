package assignment02IBC

import (
	"fmt"
	"crypto/sha256"
)

type Block struct {
	Spender     map[string]int
	Receiver    map[string]int
	PrevPointer *Block
	PrevHash    string
	CurrentHash string
}

func CalculateBalance(userName string, chainHead *Block) int {
	var netAmount int = 0
	var counter int = 0
	if chainHead.PrevPointer != nil {
		for chainHead.PrevPointer != nil {
			for key, amount := range chainHead.Spender {
				if key == userName {
					netAmount -= amount
				}
			}
			for key, amount := range chainHead.Receiver {
				if key == userName {
					netAmount += amount
				}
			}
			chainHead = chainHead.PrevPointer
			 counter += 1 	
		}
		if userName == "Satoshi" || userName == "" {
			netAmount = netAmount + 100*counter
			return netAmount
		}
		return netAmount

	} else {
		for key, amount := range chainHead.Spender {
			if key == userName {
				netAmount -= amount
			}
		}
		for key, amount := range chainHead.Receiver {
			if key == userName {
				netAmount += amount
			}
		}
		if userName == "Satoshi" {
			netAmount = netAmount + 100
			return netAmount
		}
		return netAmount
	}
	
}

func CalculateHash(inputBlock *Block) string {
	var hashed string = ""
	for key, element := range inputBlock.Spender {
		hashed += key
		hashed += string(element)
	}
	for key, element := range inputBlock.Receiver {
		hashed += key
		hashed += string(element)
	}
	fmt.Println("Concanted String:", hashed)
	obj := sha256.New()
	obj.Write([]byte(fmt.Sprintf("%x", hashed)))

	return fmt.Sprintf("%x", obj.Sum(nil))
}

func InsertBlock(spendingUser string, receivingUser string, miner string, amount int, chainHead *Block) *Block {

	var newBlock *Block = new(Block)
	if chainHead != nil {
		if miner == "Satoshi" {
			fmt.Println("Spender's Total balance:", CalculateBalance(spendingUser, chainHead))
			fmt.Println("Transcation Amount:", amount)
			if CalculateBalance(spendingUser, chainHead) >= amount {
				newBlock.Spender = make(map[string]int)
				newBlock.Receiver = make(map[string]int)
				newBlock.Spender[spendingUser] = amount
				newBlock.Receiver[receivingUser] = amount
				newBlock.PrevPointer = chainHead
				newBlock.PrevHash = chainHead.CurrentHash
				newBlock.CurrentHash = CalculateHash(newBlock)
				fmt.Printf("New Block Added\n")
				return newBlock
			} else {
				fmt.Print(spendingUser, "doesnt have enough coins to complete this transcation\n")
				return chainHead
			}
		} else {
			fmt.Printf("Only Satoshi can be the miner, So Current Block not added. Current Blockchain Rules\n")
			return chainHead
		}
	} else {
		if miner == "Satoshi" {
			if spendingUser == "" && receivingUser == "" {
				newBlock.Spender = make(map[string]int)
				newBlock.Receiver = make(map[string]int)
				newBlock.Spender[""] = amount
				newBlock.Receiver[""] = amount
				newBlock.PrevPointer = nil
				newBlock.PrevHash = ""
				newBlock.CurrentHash = CalculateHash(newBlock)
				fmt.Println("Spender's Total balance:", CalculateBalance(spendingUser, newBlock))
				fmt.Println("Transcation Amount:", amount)
				fmt.Printf("Gensis Block Added\n")
				return newBlock
			}
			fmt.Printf("As only satoshi can start the Block chain transcation, adding spendingUser and recieveingUser is not authorized\n")
			return chainHead
		}
		fmt.Printf("Only Satoshi can be the miner, So Current Block not added. Current Blockchain Rules\n")
		return chainHead
	}

}

func ListBlocks(chainHead *Block) {
	if chainHead.PrevPointer != nil {
		for chainHead.PrevPointer != nil {
			fmt.Println("CurrentHash:",chainHead.CurrentHash)
			fmt.Println("Current Block Transcations:-")
			for key, element := range chainHead.Spender {
				fmt.Println("Spender:", key, "=>", "Receiver:", element)
			}
			for key, element := range chainHead.Receiver {
				fmt.Println("Spender:", key, "=>", "Receiver:", element)
			}
			chainHead = chainHead.PrevPointer
		}
	} else {
		fmt.Println("CurrentHash:",chainHead.CurrentHash)
		fmt.Println("Current Block Transcations:-")
		for key, element := range chainHead.Spender {
			fmt.Println("Spender:", key, "=>", "Receiver:", element)
		}
		for key, element := range chainHead.Receiver {
			fmt.Println("Spender:", key, "=>", "Receiver:", element)
		}
	}

}

func VerifyChain(chainHead *Block) {
  if chainHead.PrevPointer != nil {
		var temp *Block = new(Block)
		for chainHead.PrevPointer !=nil {
			temp = chainHead.PrevPointer
			if temp.PrevHash == chainHead.CurrentHash {
				fmt.Println("Curent block hash:", chainHead.CurrentHash)
				fmt.Println("Previous block hash:", temp.PrevHash)
				fmt.Println("Previous and Current block hash match. Block chain integreity mantained.")
			} else {
				fmt.Println("Curent block hash:", chainHead.CurrentHash)
				fmt.Println("Previous block hash:", temp.PrevHash)
				fmt.Println("Previous and Current block hash dont match. Block chain integreity compromised.")
				return 
			}
			chainHead = chainHead.PrevPointer
		}
  }
  fmt.Println("Current Blockchain has only gensis block. Gensis Block Hash:", chainHead.CurrentHash)
}

