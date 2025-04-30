package eth

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func SelectTx(gethHttp string, blockNumber *big.Int) {
	client, err := ethclient.Dial(gethHttp)
	if err != nil {
		log.Fatal(err)
	}
	// 通过chainId 和事务 获取对应发送方地址
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}
	// blockNumber, err := client.BlockNumber(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(blockNumber)))
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	// var txHas common.Hash
	for _, tx := range block.Transactions() {
		fmt.Printf("tx.Hash().Hex(): %v\n", tx.Hash().Hex())
		fmt.Printf("tx.Value(): %v\n", tx.Value())
		fmt.Printf("tx.Gas(): %v\n", tx.Gas())
		fmt.Printf("tx.GasPrice(): %v\n", tx.GasPrice())
		fmt.Printf("tx.Nonce(): %v\n", tx.Nonce())
		fmt.Printf("tx.Data(): %v\n", tx.Data())
		fmt.Printf("tx.To().Hex(): %v\n", tx.To().Hex())

		if sender, err := types.Sender(types.NewEIP155Signer(chainId), tx); err != nil {
			fmt.Printf("sender.Hex(): %v\n", sender.Hex())
		}
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}
		// txHas = tx.Hash()
		//每完成一笔交易 就会生成一笔收据
		fmt.Printf("receipt.Status: %v\n", receipt.Status)
		fmt.Printf("receipt.Logs: %v\n", receipt.Logs)
		break
	}

	blockHas := common.HexToHash("0x1b176bec4bc855eb37e20a80e027e924a7c25c921161377dde523ad9bce403f9")
	tx, isPendIng, err := client.TransactionByHash(context.Background(), blockHas)
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
	}
	fmt.Printf("isPendIng: %v\n", isPendIng)
	fmt.Printf("tx.Hash(): %v\n", tx.Hash())
}
