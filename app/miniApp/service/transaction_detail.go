package service

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/sirupsen/logrus"
	"math/big"
	"miniapp/model"
	"miniapp/utils"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func transfer(data map[string]string) {
	from := data["from"]
	password := data["password"]
	amount := data["amount"]
	toAddress := data["to"]
	bytes, err2 := utils.DecryptGCM(password, utils.PassKey)
	if err2 != nil {
		logrus.Error(err2)
	}

	//1. 本地转账地址私钥 2. 接收账户 3. 构建交易 4. 交易链签名 5. 发送事务
	addressSecret := getLocalAddressSecret(from, string(bytes))

	ct := utils.GetClient()
	to := common.HexToAddress(toAddress)
	amountNum, err2 := strconv.Atoi(amount)
	if err2 != nil {
		logrus.Error(err2)
	}

	balance := amountNum
	gasPrice, err := ct.SuggestGasPrice(context.Background())
	if err != nil {
		logrus.Error(err)
	}
	nonce, err := ct.PendingNonceAt(context.Background(), to)
	if err != nil {
		logrus.Error(err)
	}
	//创建交易
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &to,
		Value:    big.NewInt(int64(balance)),
		Gas:      uint64(21000),
		GasPrice: gasPrice,
		Data:     nil,
	})
	chainId, err := ct.ChainID(context.Background())
	if err != nil {
		logrus.Error(err)
	}
	signTx, err2 := types.SignTx(tx, types.NewEIP155Signer(chainId), addressSecret)
	if err2 != nil {
		logrus.Error(err2)
	}
	client := utils.GetClient()
	err2 = client.SendTransaction(context.Background(), signTx)
	if err2 != nil {
		logrus.Error(err2)
	}
	
}

func getLocalAddressSecret(address string, password string) *ecdsa.PrivateKey {
	var fromAccount model.AccountB
	utils.GetDb().Where("from_address = ?", address).First(&fromAccount)
	if fromAccount.Address == "" {
		logrus.Warning("current address is not exist")
		return nil
	}
	bytes, err := os.ReadFile(fromAccount.Path)
	if err != nil {
		logrus.Error(err)
	}
	decryptKey, err := keystore.DecryptKey(bytes, password)

	return decryptKey.PrivateKey
}
