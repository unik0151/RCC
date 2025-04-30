package utils

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"miniapp/conf"
	"sync"
)

var client *ethclient.Client

// 调用一次
var clientOne sync.Once

func GetClient() *ethclient.Client {
	clientOne.Do(func() {
		cl, err := ethclient.Dial(conf.CLIENT_URL)
		if err != nil {
			logrus.Fatal(err)
		}
		client = cl
	})
	return client

}

func Say() {

}
