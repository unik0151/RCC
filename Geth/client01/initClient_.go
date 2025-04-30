package client01

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func InitClient() {

	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	_ = client // we'll use this in the upcoming sections
	black, err := client.HeaderByNumber(context.Background(), nil)

	fmt.Printf("bloack.Hash().Hex(): %v\n", black.Hash().Hex())
}
