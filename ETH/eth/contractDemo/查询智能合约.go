package eth

import (
	"eth/contracts/store"
	"eth/eth"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
)

func SelectSmartContract() {
	client := eth.InitClient()
	defer client.Close()
	//部署合约地址 0xc9A4816A0fe623cb8aCf1e04905A6786Df76765E
	address := common.HexToAddress("0x41d734F101607EB92778ab889957aabfeF8fEe5c")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(instance.Version(nil))
}
