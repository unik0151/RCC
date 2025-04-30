package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
	"math/big"
	"os"
	"os/signal"
	"p-cy/api/sol"
	"p-cy/log"
	"strings"
	"syscall"
)

var EthClient *ethclient.Client

//func init() {
//	dial, err := ethclient.Dial(os.Getenv("SEPOLIA"))
//	if err != nil {
//		//log.Logger.Error("Init eth client error" + err.Error())
//	}
//	EthClient = dial
//}

func Transfer() {
	client := EthClient
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Logger.Error("Suggest gas price error" + err.Error())
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Logger.Error("ChainID error" + err.Error())
	}
	toAddress := common.HexToAddress(os.Getenv("0x6666B328DA0Fc2CE25f794c77Ddd58385330eC9d"))
	//1. 组装交易数据 2. 交易数据根据规则hash签名 ， 前缀加到链上，标识是哪个链的交易
	tx := types.NewTx(&types.LegacyTx{
		To:       &toAddress,
		Gas:      21000,
		GasPrice: gasPrice,
		Value:    ethConvWei(0.001),
		Data:     []byte(""),
	})
	ecdsa, err := crypto.HexToECDSA(os.Getenv("SECRET"))
	if err != nil {
		log.Logger.Error("Init ecdsa error" + err.Error())
	}
	transaction, err := types.SignTx(tx, types.NewLondonSigner(chainID), ecdsa)
	if err != nil {
		log.Logger.Error("SignTx error" + err.Error())
	}
	log.Logger.Info("transaction hash " + transaction.Hash().Hex())

}

func ContractTransfer() {
	//client := EthClient
	//
	//privateKey, err := crypto.HexToECDSA("b262d74dbc92d39261253438138f71d1c5135eaa36572e6d7c40b52b2ea9781b")
	//if err != nil {
	//	return
	//}
	//fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	//
	////合约地址
	//toContractAddress := common.HexToAddress("")
	//
	//chainID, err := client.ChainID(context.Background())
	//if err != nil {
	//	log.Logger.Error("ChainID error" + err.Error())
	//}
	//
	//transactor := bind.NewKeyedTransactor(privateKey, chainID)
	//transactor.Value = big.NewInt(0.0001 * 1e18)
	////deploySol, transaction, _, err := sol.DeploySol(transactor, client)
	//tx, err := bind.NewBoundContract(
	//	toContractAddress,
	//	abi.ABI{}, // 无 ABI（直接转账）
	//	client,
	//	client,
	//	nil,
	//).Transact(transactor, "") // 空 calldata 触发 receive()
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Printf("Transaction Hash: 0x%x\n", tx.Hash())

}

func Deploy() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/zIEmQKN6l1Yc42TZvo8C9QbzzPmdhMFx")
	if err != nil {
		fmt.Println("client error" + err.Error())
	}
	defer client.Close()

	privateKeyHex := "b262d74dbc92d39261253438138f71d1c5135eaa36572e6d7c40b52b2ea9781b"
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		//log.Logger.Error("Init ecdsa error" + err.Error())
	}

	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	chainID, err := client.ChainID(context.Background())
	if err != nil {

	}
	auth := bind.NewKeyedTransactor(privateKey, chainID)
	nonceAt, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {

	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	fmt.Printf("gasPrice is %v\n", gasPrice)
	//gas, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
	//	From: fromAddress,
	//})
	if err != nil {
	}
	auth.Nonce = big.NewInt(int64(nonceAt))
	auth.Value = big.NewInt(0)
	auth.GasLimit = 210000
	auth.GasPrice = gasPrice
	auth.From = fromAddress

	_, _, _, err = sol.DeploySol(auth, client)
	if err != nil {
		panic(err)
	}
}

func QueryCurrentTxs() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/zIEmQKN6l1Yc42TZvo8C9QbzzPmdhMFx")
	if err != nil {
		fmt.Println("client error" + err.Error())
	}
	defer client.Close()
	address := common.HexToAddress("0x82bBDc134F0ED6917Beee35a46Bb505Bf54c022E")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{address},
	}
	logs, _ := client.FilterLogs(context.Background(), query)
	for _, vLog := range logs {
		fmt.Printf("交易哈希：%s\n区块高度：%d\n",
			vLog.TxHash.Hex(), vLog.BlockNumber)
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		fmt.Println("ChainID error" + err.Error())
	}
	// 遍历区块交易
	blockNumber, _ := client.BlockNumber(context.Background())
	for i := blockNumber - 10; i <= blockNumber; i++ { // 最近10个区块
		block, _ := client.BlockByNumber(context.Background(), new(big.Int).SetUint64(i))
		for _, tx := range block.Transactions() {
			sender, _ := types.Sender(types.NewLondonSigner(chainID), tx)
			if sender == address && *tx.To() == address {
				fmt.Printf("交易哈希：%s\n金额：%s\nGas费用：%d\n",
					tx.Hash().Hex(), tx.Value().String(), tx.Gas())
			}
		}
	}
}

//	func main() {
//		client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/zIEmQKN6l1Yc42TZvo8C9QbzzPmdhMFx")
//		if err != nil {
//			fmt.Println("client error" + err.Error())
//		}
//		defer client.Close()
//		var logs = make(chan types.Log, 2)
//		contractAddress := common.HexToAddress("0x9CC1565aaebf8fB270AA3237bA7A325Ca08173DB")
//		query := ethereum.FilterQuery{
//			Addresses: []common.Address{contractAddress},
//		}
//		filterLogs, err := client.SubscribeFilterLogs(context.Background(), query, logs)
//		if err != nil {
//			panic(err)
//		}
//		for {
//			select {
//			case msg := <-logs:
//				fmt.Println(msg)
//			case err = <-filterLogs.Err():
//				fmt.Println(err)
//
//			}
//		}
//	}
func Subscribe() {
	// 连接到以太坊客户端
	client, err := ethclient.Dial("wss://eth-sepolia.g.alchemy.com/v2/zIEmQKN6l1Yc42TZvo8C9QbzzPmdhMFx")
	if err != nil {
		fmt.Printf("Failed to connect to Ethereum client: %v", err)
	}
	defer client.Close() // 确保在main函数结束时关闭客户端连接

	// 创建一个用于接收日志的通道
	logs := make(chan types.Log, 2)

	// 指定要监听的事件日志的合同地址
	contractAddress := common.HexToAddress("0x9CC1565aaebf8fB270AA3237bA7A325Ca08173DB")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	// 订阅事件日志
	ctx, cancel := context.WithCancel(context.Background())
	filterLogs, err := client.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		fmt.Printf("Failed to subscribe to logs: %v", err)
	}
	defer cancel() // 确保在main函数结束时取消订阅

	// 设置一个通道来监听操作系统信号
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	// 启动一个goroutine来处理日志
	go func() {
		for {
			select {
			case log := <-logs:
				fmt.Println("Received log:", log)
			case err := <-filterLogs.Err():
				fmt.Printf("Subscription error: %v", err)
			case <-interrupt:
				fmt.Println("Interrupt signal received, stopping log processing...")
				return
			}
		}
	}()

	// 阻塞直到收到中断信号
	<-interrupt

	// 可选：等待一段时间以确保所有日志都已处理（根据实际需要调整）

	fmt.Println("Exiting program...")
}

func main() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/zIEmQKN6l1Yc42TZvo8C9QbzzPmdhMFx")
	if err != nil {
		fmt.Println("client error" + err.Error())
	}
	defer client.Close()
	contractAddress := common.HexToAddress("0x9CC1565aaebf8fB270AA3237bA7A325Ca08173DB")
	//header, err := client.HeaderByNumber(context.Background(), nil)
	query := ethereum.FilterQuery{
		//FromBlock: header.Number.Sub(header.Number, big.NewInt(10)),
		//ToBlock:   header.Number,
		Addresses: []common.Address{
			contractAddress,
		},
	}
	if err != nil {
		panic(err)
	}
	filterLog, err := client.FilterLogs(context.Background(), query)
	abi, err := abi.JSON(strings.NewReader(sol.SolMetaData.ABI))
	if err != nil {
		panic(err)
	}
	//1. keccak256 2. bytesToHash
	//common.BytesToHash()
	etherReceiveSigture := crypto.Keccak256Hash([]byte("EtherReceived(address,uint256)"))
	etherFallbackSigture := crypto.Keccak256Hash([]byte("FallbackCalled(address,uint256,bytes)"))
	fmt.Printf("EtherReceive signature %v\n", etherReceiveSigture.Hex())
	fmt.Printf("EtherFallback signature %v\n", etherFallbackSigture.Hex())
	for _, vlog := range filterLog {
		fmt.Println(vlog.BlockNumber)
		fmt.Println(vlog.TxHash.Hex())
		//fmt.Printf("vlog.Topics[0].Hex() %v\n", vlog.Topics[0].Hex())
		switch vlog.Topics[0].Hex() {
		case etherReceiveSigture.Hex():
			event := struct {
				Sender common.Address
				Amount *big.Int
			}{}
			//data := make([]byte, 5024)
			err = abi.UnpackIntoInterface(&event, "EtherReceived", vlog.Data)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(event.Amount)
		case etherFallbackSigture.Hex():
			var etherFallback = &struct {
				sender common.Address
				amount *big.Int
			}{}
			err = abi.UnpackIntoInterface(etherFallback, "FallbackCalled", vlog.Data)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(etherFallback.sender, etherFallback.amount)
		}
	
	}
}
func ethConvWei(amount float64) *big.Int {
	return big.NewInt(int64(amount) * 1e18)
}
