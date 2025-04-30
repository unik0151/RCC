package eth

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
)

func SelectTxReceipt() {

	client := InitClient()

	Hash := common.HexToHash("0x6dd2b2faabb84382a2b6add93bb0d9903a2058dddde67da1e9ba7192b74f8bcf")

	tx, err := client.TransactionReceipt(context.Background(), Hash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx.Status: %v\n", tx.Status)
	fmt.Printf("tx.Logs: %v\n", tx.Logs)
	fmt.Printf("tx.ContractAddress: %v\n", tx.ContractAddress)

}
