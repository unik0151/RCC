package eth

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

const (
	wallets string = "./wallets"
)

func CreateWalletByKeyStore() {

	//
	ks := keystore.NewKeyStore(wallets, keystore.StandardScryptN, keystore.StandardScryptP)
	password := "123456"
	account, err := ks.NewAccount(password)
	if err != nil {
		panic(err)
	}
	//直接生成账户
	fmt.Printf("account.Address.Hex(): %v\n", account.Address.Hex())

}

func ReadKeyStore() {
	fileFull, err := ListFilesByPath()
	if err != nil {
		panic(err)
	}
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	fileBytes, errFile := os.ReadFile(fileFull)
	if errFile != nil {
		panic(errFile)
	}
	password := "123456"
	account, err := ks.Import(fileBytes, password, password)
	if err != nil {
		panic(err)
	}
	fmt.Printf("account.Address.Hex(): %v\n", account.Address.Hex())
	if err := os.Remove(fileFull); err != nil {
		log.Fatal(err)
	}

}

func ListFilesByPath() (string, error) {
	_, stOk := os.Stat(wallets)
	if stOk != nil {
		os.Mkdir(wallets, fs.ModePerm)
	}
	dirEntry, err := os.ReadDir(wallets)
	if err != nil {
		panic(err)
	}

	for _, v := range dirEntry {
		if v.Type().IsDir() {
			log.Fatal("当前是文件夹！")
			continue
		}

		fmt.Printf("v.Name(): %v\n", v.Name())

		return wallets + "/" + v.Name(), nil
	}
	return "", errors.New("sss")
}
