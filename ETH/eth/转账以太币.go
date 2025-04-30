package eth

import (
	"context"
	"eth/file"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func SendETH() {

	// 1. 创建客户端
	client := InitClient()
	defer client.Close()

	//2.导入keystore 本地保证私钥安全
	// ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)

	// account, err := ks.Import(file.ReadFile(), "qwerasdf", "qwerasdf")

	key, err := keystore.DecryptKey(file.ReadFile(), "qwerasdf")

	fmt.Printf("hexutil.Encode(key.PrivateKey.D.Bytes()): %v\n", hexutil.Encode(key.PrivateKey.D.Bytes()))

	address := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)

	if err != nil {
		log.Fatal(err)
	}
	// accountHex := account.Address.Hex()
	fmt.Printf("address: %v\n", address)

	//系统生成账户交易的随机数
	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		log.Fatal(err)
	}
	//设置转账金额 value
	balance := EthConvertWei(1.5)

	// eth标准gaslimit ：21000

	//这里是根据x个先前块获取平均gasPrice
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//to 发送给谁
	toAddress := common.HexToAddress("0x06d4b911c0c5623049a77966bb5b72abce77f943")

	tx := types.NewTransaction(nonce, toAddress, balance, uint64(21000), gasPrice, nil)

	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//使用EIP155规则 将私链公链的链id 组装成符合EIP155定义规则的对象，然后用这个对象对tx has,之后使用私钥进行加密
	signTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), key.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	//types.NewTransaction 生成transaction对象，
	err = client.SendTransaction(context.Background(), signTx)
	if err != nil {
		log.Fatal(err)
	}

	//获取交易hash 0x为16进制标识前缀
	fmt.Printf("signTx.Hash().Hex(): %v\n", signTx.Hash().Hex())
}

func EthConvertWei(eth float64) *big.Int {
	wei := new(big.Int)
	wei.Mul(big.NewInt(int64(eth*1e18)), big.NewInt(int64(1)))
	return wei
}

func KeyStoreConvertPrivateObj() {
	key, err := keystore.DecryptKey(file.ReadFile(), "qwerasdf")
	if err != nil {
		log.Fatal(err)
	}
	//获取私钥二进制数据
	// privateKey := key.PrivateKey.D
	privateBytes := key.PrivateKey.D.Bytes()

	fmt.Printf("hexutil.Encode(privateKey.Bytes()): %v\n", hexutil.Encode(privateBytes))

	publicKey := key.PrivateKey.PublicKey

	pubEcdsaAddress := crypto.PubkeyToAddress(publicKey)

	fmt.Printf("pubEcdsaAddress.Hex(): %v\n", pubEcdsaAddress.Hex())

}
