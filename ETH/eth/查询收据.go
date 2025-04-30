package eth

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func SelectReciped(gethHttp string, blockNumber *big.Int) {
	client, err := ethclient.Dial(gethHttp)
	if err != nil {
		log.Fatal(err)
	}

	block, err := client.BlockByNumber(context.Background(), blockNumber)
	receiptByHash, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithHash(block.Hash(), false))
	if err != nil {
		log.Fatal(err)
	}
	receiptByNum, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blockNumber.Int64())))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("receiptByNum[0]: %v\n", receiptByNum[0])

	for _, receipt := range receiptByHash {
		fmt.Printf("receipt.Status: %v\n", receipt.Status)
		fmt.Printf("receipt.Logs: %v\n", receipt.Logs)
		fmt.Printf("receipt.TxHash.Hex(): %v\n", receipt.TxHash.Hex())
		fmt.Printf("receipt.TransactionIndex: %v\n", receipt.TransactionIndex)
		fmt.Printf("receipt.BlockHash.Hex(): %v\n", receipt.BlockHash.Hex())
	}
	// 根据单独的txHash 获取单独的信息
	// client.TransactionReceipt()
}
