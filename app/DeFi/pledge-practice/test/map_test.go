package test

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"p-cy/util"
	"testing"
	"time"
)

func TestMapGet(t *testing.T) {
	var res = map[int]map[string]string{
		0: {
			"a": "b",
		},
	}

	s, e := res[0]
	if e != true {

		log.Fatal(e)
	}
	fmt.Println(s["a"])
}

func TestPointer(t *testing.T) {
	var users = make([]int, 0)
	addInt(&users)
	fmt.Println(users)
	fmt.Println(users)
	fmt.Println(&users)
	fmt.Println(*&users)
}

func addInt(users *[]int) {

	*users = append(*users, 1)
}

func TestAes(t *testing.T) {

	//hash := sha256.New()
	//hash.Write([]byte("333333"))
	//hashSum := hash.Sum(nil)
	//fmt.Println(hex.EncodeToString(hashSum))
}

func TestCode(t *testing.T) {
	num := util.GenerateCode(6)
	fmt.Println(num)
}
func TestTime(t *testing.T) {
	now := time.Now()
	add := now.Add(5 * time.Minute)
	fmt.Println(now, " ", add)
}

/**
转账要素
1. 连接节点
2. 转账gas固定21000  value转账额度 to
basefee 从区块头获取
tp手续费
gasfee = base + tp

3. 构造tx types.newTx
4. types.signtx  EIP1559协议 或者Long加密
5. sendTransation


*/

func TestSendEth(t *testing.T) {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/2a65c5cabb974f0798f29f03827483a4")
	defer client.Close()
	if err != nil {
		log.Fatal(err)
	}
	var address = "0x82bBDc134F0ED6917Beee35a46Bb505Bf54c022E"
	var toAddress = "0x6666B328DA0Fc2CE25f794c77Ddd58385330eC9d"
	//查询地址对应余额
	//ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	balanceAt, err := client.BalanceAt(context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		log.Fatal(err)
	}
	eth := weiToEth(balanceAt)
	log.Println("eth", eth)

	chainID, _ := client.ChainID(context.Background())
	nonceAt, _ := client.PendingNonceAt(context.Background(), common.HexToAddress(address))
	gasTipCap, _ := client.SuggestGasTipCap(context.Background())
	header, _ := client.HeaderByNumber(context.Background(), nil)
	baseFee := header.BaseFee
	//gasPrice, _ := client.SuggestGasPrice(context.Background())
	toAddressCom := common.HexToAddress(toAddress)
	valueEth := EthConvertWei(0.0001)
	//gas, _ := client.EstimateGas(ctx, ethereum.CallMsg{})
	tx := types.DynamicFeeTx{
		ChainID: chainID,
		Nonce:   nonceAt,
		//小费
		GasTipCap: gasTipCap,
		//此次总花费限制 , gasPrice
		GasFeeCap: new(big.Int).Add(baseFee, gasTipCap),
		//eth转账固定gas
		Gas:   21000,
		To:    &toAddressCom,
		Value: valueEth,
	}
	transaction := types.NewTx(&tx)
	//transaction.Hash()
	//tx := types.NewTransaction(nonceAt, toAddressCom, valueEth, uint64(21000), gasPrice, nil)
	privateKeyMy, _ := crypto.HexToECDSA("b262d74dbc92d39261253438138f71d1c5135eaa36572e6d7c40b52b2ea9781b")
	newTx, _ := types.SignTx(transaction, types.NewLondonSigner(chainID), privateKeyMy)

	err = client.SendTransaction(context.Background(), newTx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("tx sent")
	//client.
}
func EthConvertWei(eth float64) *big.Int {
	wei := new(big.Int)
	wei.Mul(big.NewInt(int64(eth*1e18)), big.NewInt(int64(1)))
	return wei
}

func weiToEth(int2 *big.Int) *big.Float {
	setInt := new(big.Float).SetInt(int2)
	//转换
	setInt = setInt.Quo(setInt, big.NewFloat(1e18))
	return setInt
}
