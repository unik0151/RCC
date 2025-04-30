package eth

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func CheckAddressIsVaild() {
	var myAddress string = "0xa0a6c68e4b7a6873898cdf4ba1d2145291cd276b8fa4b6c4572a4326814b30fe"
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	fmt.Printf("re.MatchString(\"0xf69984CCD90A4c6B701Ef875A7bF62e5965753bb\"): %v\n", re.MatchString(myAddress))
	fmt.Printf("re.MatchString(\"0xf69984CCD90A4c6B701Ef875A7bF62e5965753bba\"): %v\n", re.MatchString("0xf69984CCD90A4c6B701Ef875A7bF62e5965753bba"))

	//0xf69984CCD90A4c6B701Ef875A7bF62e5965753bb
	client, err := ethclient.Dial("http://127.0.0.1:8999")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(myAddress)

	bytecode, err := client.CodeAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}
	//外部账户 这个字段位空
	isContract := len(bytecode) > 0
	fmt.Printf("isContract: %v\n", isContract)

}
