package eth

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	gethHttp = "http://127.0.0.1:8999"
	// gethHttp = "https://eth-sepolia.g.alchemy.com/v2/zIEmQKN6l1Yc42TZvo8C9QbzzPmdhMFx"
)

func InitClient() *ethclient.Client {

	client, err := ethclient.Dial(gethHttp)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
