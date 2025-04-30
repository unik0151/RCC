# Ganache 用于本地测试的rpc框架
Ganache(正式名称为 testrpc)是一个用 Node.js 编写的以太坊实现，用于在本地开发去中心化应用程序时进行测试。现在我们将带着您完成安装并连接到它。

# geth搭建私有链 版本geth 1.10.26-stable
password: qwerasdf
account:  0xf69984CCD90A4c6B701Ef875A7bF62e5965753bb

<!-- 创建创世块 -->
geth -datadir test_ethereum/data init test_ethereum/genesis.json

<!-- 创建账户 -->
geth --datadir test_ethereum/data account new


-datadir 选项，以便相关区块链数据都存入这个路径，方便统一管理。

genesis.json模板
{
  "config": {
    "chainId": 1,
    "homesteadBlock": 0,
    "eip150Block": 0,
    "eip155Block": 0,
    "eip158Block": 0,
    "byzantiumBlock": 0,
    "constantinopleBlock": 0,
    "petersburgBlock": 0,
    "istanbulBlock": 0,
    "berlinBlock": 0,
    "londonBlock": 0
  },
  "alloc": {
    "0x0000000000000000000000000000000000000001": {
      "balance": "111111111"
    },
    "0x0000000000000000000000000000000000000002": {
      "balance": "222222222"
    }
  },
  "coinbase": "0x0000000000000000000000000000000000000000",
  "difficulty": "0x20000",
  "extraData": "",
  "GasLimit": "0x2fefd8",
  "nonce": "0x0000000000000042",
  "mixhash": "0x0000000000000000000000000000000000000000000000000000000000000000",
  "parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
  "timestamp": "0x00"
}

geth -datadir test_ethereum/data -networkid 1 -port "30303" -http -http.api "admin,debug,eth,miner,net,personal,txpool,web3" -http.port "8999" -http.corsdomain "*" -nodiscover -mine -miner.threads=1 -miner.etherbase 0xf69984CCD90A4c6B701Ef875A7bF62e5965753bb --allow-insecure-unlock --rpc.txfeecap 0 --rpc.gascap=0

<!-- --rpc.txfeecap 0 --rpc.gascap 0 允许任何金额进行gas -->



<!-- 连接本地geth启动的私有链节点 -->
geth attach http://127.0.0.1:8999

```bash
# 创建账户
personal.Account("密码")
# 0xf69984CCD90A4c6B701Ef875A7bF62e5965753bb qwerasdf 搭建geth设置的本人账户，geth启动后会自动挖矿，将挖矿的奖励和gas费存入这个
// 0x06d4b911c0c5623049a77966bb5b72abce77f943 123456
# 查看所有账户
eth.accounts

0xd0260b303492cf5517c1a37e43a62d8cf12d18ba6daf605cf626d2c68f56d0ac
# 解锁账户 300s
personal.unlockAccount("0xf69984CCD90A4c6B701Ef875A7bF62e5965753bb" , "qwerasdf",300)

# 查看余额 单位: wei 
eth.getBalance("账户") 

# 发起转账 ， 须先解锁账户
personal.unlockAccount("0xf69984CCD90A4c6B701Ef875A7bF62e5965753bb","qwerasdf",300)
eth.sendTransaction({from: "0xf69984CCD90A4c6B701Ef875A7bF62e5965753bb",to: "0x06d4b911c0c5623049a77966bb5b72abce77f943",value: web3.toWei(2,"ether")})
# 交易txHas
# 0x1b176bec4bc855eb37e20a80e027e924a7c25c921161377dde523ad9bce403f9

# 查询交易所在区块
eth.getTransaction("0x1b176bec4bc855eb37e20a80e027e924a7c25c921161377dde523ad9bce403f9").then(function(tx){console.log(tx.blockNumber)})

web3.fromWei(eth.getBalance(eth.accounts[0]))

# 开始挖矿
miner.start()
miner.stop()
```

