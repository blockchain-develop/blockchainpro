# Build Private Ethereum Network

* 安装geth
* 准备genesis
* 配置bootnode

## 安装geth
[installation guide](https://geth.ethereum.org/docs/install-and-build/installing-geth)

## 准备genesis
如果是ethash共识算法，则配置如下
```
{
  "config": {
    "chainId": <arbitrary positive integer>,
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
  "alloc": {},
  "coinbase": "0x0000000000000000000000000000000000000000",
  "difficulty": "0x20000",
  "extraData": "",
  "gasLimit": "0x2fefd8",
  "nonce": "0x0000000000000042",
  "mixhash": "0x0000000000000000000000000000000000000000000000000000000000000000",
  "parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
  "timestamp": "0x00"
}
```

如果需要配置初始账户资产
```
"alloc": {
  "0x0000000000000000000000000000000000000001": {
    "balance": "111111111"
  },
  "0x0000000000000000000000000000000000000002": {
    "balance": "222222222"
  }
}
```

## 配置bootnode
由于ethereum网络是P2P网络，需要有节点发现，private net，所以配置一个bootnode，用于节点发现。
```
$ bootnode --genkey=boot.key
$ bootnode --nodekey=boot.key
```
在bootnode启动后，有一个enode的url，这个url在后面启动geth节点时作为参数。

## 启动节点（非miner节点）
准备目录， 进入目录。

### 配置geth
```
geth init --datadir data <genesis.json>
```

### 启动geth
```
geth --datadir=path/to/custom/data/folder --bootnodes=<bootnode-enode-url-from-above> --networkid 10000 --port 30304 --rpc --rpcport "8085"
```
```
geth --datadir=./data --bootnodes=enode://bc3b0b020b680b2cd91826b5eaa413030af329809d5db583190b2f24ccee47e240e058823dda0007de69cc04b130deade8e19576f4f2c0471b02d7946702df31@127.0.0.1:0?discport=30301 --port 30304 --rpc --rpcport "8085"
```
## 启动miner节点
准备目录，进入目录。

### 配置geth

```
geth init --datadir data <genesis.json>
```

### 启动geth
```
geth --datadir=path/to/custom/data/folder --bootnodes=<bootnode-enode-url-from-above> --port 30305 --mine --miner.threads=1 --miner.etherbase=0x0000000000000000000000000000000000000000
```

```
geth --datadir=./data --bootnodes=enode://bc3b0b020b680b2cd91826b5eaa413030af329809d5db583190b2f24ccee47e240e058823dda0007de69cc04b130deade8e19576f4f2c0471b02d7946702df31@127.0.0.1:0?discport=30301 --port 30305 --mine --miner.threads=1 --miner.etherbase=0xd8d50Be55FE241B3c026361a793aA950BceAE845
```


```
./geth --ropsten --datadir=./data --syncmode=fast --http --http.addr 0.0.0.0 --http.port 8085 --port 30304 --http.port 30060 --http.corsdomain * --http.api eth,net,web3,personal,txpool
```




