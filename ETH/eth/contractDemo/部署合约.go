package eth

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"eth/contracts/store"
	"eth/eth"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
)

func DeployContract() {

	client := eth.InitClient()
	defer client.Close()
	//私钥获取后去掉0x  // b262d74dbc92d39261253438138f71d1c5135eaa36572e6d7c40b52b2ea9781b
	//0xf69984ccd90a4c6b701ef875a7bf62e5965753bb 0x1c0879ccf8c8b8d93294a223bc5be8a0535e10e73ba277ddd68a9fb024e7f58d
	privateKey, err := crypto.HexToECDSA("1c0879ccf8c8b8d93294a223bc5be8a0535e10e73ba277ddd68a9fb024e7f58d")
	if err != nil {
		log.Fatal(err)
	}
	pubKey := privateKey.Public()
	pubECDSA := pubKey.(*ecdsa.PublicKey)

	pubAddress := crypto.PubkeyToAddress(*pubECDSA)
	fmt.Printf("pubAddress.Hex(): %v\n", pubAddress.Hex())

	nonce, err := client.PendingNonceAt(context.Background(), pubAddress)
	if err != nil {
		log.Fatal(err)
	}

	// //gas , gaslimit value , gasprice
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(300000)
	auth.Value = big.NewInt(0)

	address, tx, _, err := store.DeployStore(auth, client, "v_1.0")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("address.Hex(): %v\n", address.Hex())
	fmt.Printf("tx.Hash().Hex(): %v\n", tx.Hash().Hex())

}
