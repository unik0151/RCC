package client01

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// 以太坊实现 cmd ganache-cli 有测试账号之类的，连接本地之后直接用测试账号就行了
func AccountBalance() {
	client, err := ethclient.Dial("http://localhost:8999")
	if err != nil {
		log.Fatal(err)
	}

	account := common.HexToAddress("0xEc05d894f3492B269CFDF9cC9c908f0c999E7faA")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("balacne : ", balance) // 25893180161173005034

	// 获取ETH 因为获取的是wei单位 eth = wei / 10 18次方
	fBalance := new(big.Float)
	fBalance.SetString(balance.String())
	fValue := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))
	fmt.Printf("fValue: %v\n", fValue)

	// nil获取最新的区块，最新区块头信息
	blackLast, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("blackLast.Difficulty().String(): %v\n", blackLast.Difficulty().String())

	headerV, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("headerV.Number: %v\n", headerV.Number)
	fmt.Printf("headerV.ParentHash: %v\n", headerV.ParentHash)
	// pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	// fmt.Println(pendingBalance) // 25729324269165216042
}
