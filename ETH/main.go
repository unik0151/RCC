package main

import (
	eth "eth/eth/contractDemo"
	"math/big"
)

var (
	gethHttp             = "http://127.0.0.1:8999"
	blockNumber *big.Int = big.NewInt(2632)
)

func main() {

	// eth.CreateAccount()

	// eth.SelectTx(gethHttp, blockNumber)
	// eth.CreateWallet()
	// eth.ListFilesByPath()
	// eth.CreateWalletByKeyStore()
	// eth.ReadKeyStore()

	// eth.CheckAddressIsVaild()
	// eth.SendETH()
	// eth.KeyStoreConvertPrivateObj()

	// eth.SelectTxReceipt()
	//合约 0xf69984CCD90A4c6B701Ef875A7bF62e5965753bb
	// eth.DeployContract()
	// eth.LoadSmartContract()
	eth.SelectSmartContract()
}
