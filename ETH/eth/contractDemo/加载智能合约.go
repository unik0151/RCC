package eth

import (
	"eth/contracts/store"
	"eth/eth"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
)

func LoadSmartContract() {
	// 0x257AdAf0c24b5339701E30bCd968C95B7BC66Cdb
	client := eth.InitClient()
	defer client.Close()
	address := common.HexToAddress("0x3A9467d18AA11F30119dd18079ecA8Ccd3EB72d6")

	instant, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("contract is loaded")
	_ = instant
}
