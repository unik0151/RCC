package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
)

func main() {
	// 1. 生成私钥
	privateKey := createPrivateKeyBtc()
	pubKey := privateKey.PubKey()
	// 未压缩的公钥二进制
	notCompressPubk := pubKey.SerializeUncompressed()

	//获取非压缩公钥
	// 对未压缩的公钥二进制进行十六进制转码，为人类可看的 此时04为前缀 此为公钥
	fmt.Printf("hex.EncodeToString(notCompressPubk): %v\n", hex.EncodeToString(notCompressPubk))
	// 3. 生成账户地址对象
	account, err := btcutil.NewAddressPubKey(notCompressPubk, &chaincfg.Params{})
	if err != nil {
		panic(err)
	}
	//对address进行双重加密后base58进行编码，获取账户地址 ， hash160.has 直接对数据进行sha258编码，ripemd160编码 ， 获取160后的20个字节，base58编码
	fmt.Printf("account: %v\n", account.EncodeAddress())
}

func createPrivateKeyBtc() *btcec.PrivateKey {
	privKey, err := btcec.NewPrivateKey()
	if err != nil {
		log.Fatal("Create PrivK Error ! ", err.Error())
		panic(err)
	}
	//32字节私钥
	fmt.Printf("privKey: %v\n", privKey)
	return privKey
}
