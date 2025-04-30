package utils

import (
	crypto2 "crypto"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sirupsen/logrus"
	"miniapp/conf"
)

/*
*
return account , secret
*/
func CreateWallet(password string) (accounts.Account, string) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		logrus.Error(err)
	}
	privateSecret := crypto.FromECDSA(privateKey)
	ks := keystore.NewKeyStore(conf.WALLET_FILE_PATH, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount(password)

	return account, hexutil.Encode(privateSecret)[2:]
}

func TestCreateWallet() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		logrus.Error(err)
	}
	privateSecret := crypto.FromECDSA(privateKey)
	fmt.Println("私钥：", hexutil.Encode(privateSecret))
	one(privateKey.PublicKey)
	two(privateKey.Public())
}

func one(pub ecdsa.PublicKey) {
	//使用自带工具
	address := crypto.PubkeyToAddress(pub)
	fmt.Println("address:", address.Hex())
}

func two(publicKey crypto2.PublicKey) {
	key := publicKey.(*ecdsa.PublicKey)
	ecdsPub := crypto.FromECDSAPub(key)
	//去除一字节 0x04 十六进制 去除0x标识符 04 = 0000 0100 八位一字节
	//这个地方keccak26哈希后，为256位bit =  32字节，取最后20个字节 2字节一
	keccak256 := crypto.Keccak256(ecdsPub[1:])[12:]
	publicHex := hexutil.Encode(keccak256)
	fmt.Println("publicHex:", publicHex)
}
