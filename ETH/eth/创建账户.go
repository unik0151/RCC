package eth

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
)

func CreateAccount() {
	client := InitClient()
	id, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("id: %v\n", id)
	number, err := client.BlockNumber(context.Background())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("number: %v\n", number)
	defer client.Close()
	balance, err := client.BalanceAt(context.Background(), common.HexToAddress("0x82bBDc134F0ED6917Beee35a46Bb505Bf54c022E"), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("balance: %v\n", balance)
}
