package main

import (
	"blockhain-go/domain"
	"encoding/json"
	"fmt"
)

func main() {
	bc := domain.NewBlockchain()

	// bc.GiveMandate("KPU", "shellrean", 1)
	bc.GiveMandate("KPU", "namira", 1)
	bc.CreateBlock(bc.LastestBlock().Hash())

	data, _ := json.Marshal(bc)
	fmt.Println(string(data))
}
