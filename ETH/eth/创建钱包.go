package eth

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

// y
func CreateWallet() {
	//1. 使用secp256k1 crypty创建私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}

	//0x04是十六进制，byte是二进制， 十六进制转二进制就为一字节s
	publicKey := privateKey.Public()
	publicKeyToECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic(ok)
	}

	//第一种
	publicAddress := crypto.PubkeyToAddress(*publicKeyToECDSA)
	fmt.Printf("publicAddress.Hex(): %v\n", publicAddress.Hex())
	fmt.Printf("len(publicAddress.Hex()): %v\n", len(publicAddress.Hex()))
	//第二种 ecdsa公钥编码 没有压缩未压缩之分 但是会出现0x04前缀要去掉[1:]
	publicBytes := crypto.FromECDSAPub(publicKeyToECDSA)

	// 去掉0x04前缀 这一步直接转换账户地址
	fmt.Printf("hexutil.Encode(publicBytes): %v\n", hexutil.Encode(publicBytes)[4:])

	//第三种
	shaHash := sha3.NewLegacyKeccak256()
	shaHash.Write(publicBytes[1:])
	//对ecdsa编码，需要先把标识0x04去除 然后 shakeccak256编码 ,获取256位 8bit = 1 byte  取最后20hash字节 ，160位,所以是12:
	fmt.Printf("hexutil.Encode(shaHash.Sum(nil)[12:]): %v\n", hexutil.Encode(shaHash.Sum(nil)[12:]))

}
