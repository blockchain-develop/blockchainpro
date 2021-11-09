# bitoin return bug验证过程

我们在这里不讲解bitcoin的return bug的详情，请参考[]()。

在这篇文章我们讲解如何重现这个bug，需要
* 下载docker并在docker环境下运行bitcoin regtest网络。
* 安装Electrum的bitcoin钱包并连接regtest网络。
* 构造交易并重新bug

## bitcoin regtest网络

下载并安装docker，请参考[官方文档](https://www.docker.com)。

为了在本地docker环境下运行bitcoin regtest网络，配置Docker Compose file并保存为bitcoind-compose.yml
```
version: '3'

services:
  node:
    image: ulamlabs/bitcoind-custom-regtest:latest
  electrumx:
    image: lukechilds/electrumx:latest
    links:
      - node
    ports:
      - "51001:50001"
      - "51002:50002"
    environment:
      - NET=regtest
      - COIN=BitcoinSegwit
      - DAEMON_URL=http://test:test@node:19001
  explorer:
    image: ulamlabs/btc-rpc-explorer:latest
    links:
      - node
      - electrumx
    ports:
      - "3002:3002"
    environment:
      - BTCEXP_HOST=0.0.0.0
      - BTCEXP_BITCOIND_URI=http://test:test@node:19001
      - BTCEXP_ELECTRUMX_SERVERS=tls://electrumx:50002
      - BTCEXP_ADDRESS_API=electrumx
```

启动docker来运行bitcoin regtest网络。
```
docker-compose up -d
```

启动后，在本地可以访问bitcoin regtest的浏览器， http://localhost:3002

## 测试交易

进入docker
```
docker exec -it bitcoin-testnet_node_1 /bin/bash
```

查看balance
```
bash-5.0$ bitcoin-cli getbalance
```

这时候余额是0，因为bitcoin挖矿收入需要200个区块确认。

调整一些挖矿

```
bash-5.0$ bitcoin-cli generatetoaddress 200 `bitcoin-cli getnewaddress` 
```

再次查询余额，应该有很多了。

## 准备环境

准备一个新地址，一个legacy地址，为方便我们分析
```
bash-5.0# bitcoin-cli -regtest getnewaddress "label" "legacy"
mpJRdFjFcQM6o4JN2D1mogUt7FwpdWbDfh
```

向这个地址发送一个btc

```
bash-5.0# bitcoin-cli -regtest sendtoaddress mpJRdFjFcQM6o4JN2D1mogUt7FwpdWbDfh 1
4bd84af06f8c6749f3a5ccbb11a415bd9eaa3001a9fb66f800e2fad89def44b3
```

会有一个交易hash，我们先看看这个交易。
```
bash-5.0# bitcoin-cli -regtest getrawtransaction 4bd84af06f8c6749f3a5ccbb11a415bd9eaa3001a9fb66f800e2fad89def44b3
02000000000101cba92586976131ad746a88dfc992efd08eb078b8eb930cbbcbe11a9bda0bec930100000017160014a912cfc284c1cc04dc5f4d5d3422b59a6b182508feffffff0200e1f505000000001976a914605a8a137e7fb06ae94e8663d0fc9a83f9c3152288ace845171e0100000017a914c75e522d2e6583f48cd57e9d7463cc592feaa0088702473044022052390dbcccf8e82a27f60dc097e18d8473ae3ac87b5d23a7dfeafac802ceb3390220429e5bf4e353d9cb9ac987cfff5df784f8eca9f69072cfe62fe680d4087a453801210345fd5778c954b65a4784882e96c3fe3bf3abc5ed095ede66f04029abf7a3736e84020000
```

```
bash-5.0# bitcoin-cli -regtest decoderawtransaction 02000000000101cba92586976131ad746a88dfc992efd08eb078b8eb930cbbcbe11a9bda0bec930100000017160014a912cfc284c1cc04dc5f4d5d3422b59a6b182508feffffff0200e1f505000000001976a914605a8a137e7fb06ae94e8663d0fc9a83f9c3152288ace845171e0100000017a914c75e522d2e6583f48cd57e9d7463cc592feaa0088702473044022052390dbcccf8e82a27f60dc097e18d8473ae3ac87b5d23a7dfeafac802ceb3390220429e5bf4e353d9cb9ac987cfff5df784f8eca9f69072cfe62fe680d4087a453801210345fd5778c954b65a4784882e96c3fe3bf3abc5ed095ede66f04029abf7a3736e84020000
{
  "txid": "4bd84af06f8c6749f3a5ccbb11a415bd9eaa3001a9fb66f800e2fad89def44b3",
  "hash": "e37b9e673fe3d6e6dda6a0a70f2fbd732f61ff17df64461f33ee82bc9926bd81",
  "version": 2,
  "size": 249,
  "vsize": 168,
  "weight": 669,
  "locktime": 644,
  "vin": [
    {
      "txid": "93ec0bda9b1ae1cbbb0c93ebb878b08ed0ef92c9df886a74ad3161978625a9cb",
      "vout": 1,
      "scriptSig": {
        "asm": "0014a912cfc284c1cc04dc5f4d5d3422b59a6b182508",
        "hex": "160014a912cfc284c1cc04dc5f4d5d3422b59a6b182508"
      },
      "txinwitness": [
        "3044022052390dbcccf8e82a27f60dc097e18d8473ae3ac87b5d23a7dfeafac802ceb3390220429e5bf4e353d9cb9ac987cfff5df784f8eca9f69072cfe62fe680d4087a453801",
        "0345fd5778c954b65a4784882e96c3fe3bf3abc5ed095ede66f04029abf7a3736e"
      ],
      "sequence": 4294967294
    }
  ],
  "vout": [
    {
      "value": 1.00000000,
      "n": 0,
      "scriptPubKey": {
        "asm": "OP_DUP OP_HASH160 605a8a137e7fb06ae94e8663d0fc9a83f9c31522 OP_EQUALVERIFY OP_CHECKSIG",
        "hex": "76a914605a8a137e7fb06ae94e8663d0fc9a83f9c3152288ac",
        "reqSigs": 1,
        "type": "pubkeyhash",
        "addresses": [
          "mpJRdFjFcQM6o4JN2D1mogUt7FwpdWbDfh"
        ]
      }
    },
    {
      "value": 47.99809000,
      "n": 1,
      "scriptPubKey": {
        "asm": "OP_HASH160 c75e522d2e6583f48cd57e9d7463cc592feaa008 OP_EQUAL",
        "hex": "a914c75e522d2e6583f48cd57e9d7463cc592feaa00887",
        "reqSigs": 1,
        "type": "scripthash",
        "addresses": [
          "2NBRPSzjqfQHEtJ9AAx5vd6WEYhwQw5xX9p"
        ]
      }
    }
  ]
}
```

## 构造交易并重现bug
这是一个手动构造交易的过程，大致流程如下：
* 选择一个可以使用的output
* 使用bitcoin-cli来构造一个交易
* 在Electrum上decode交易并修改解锁脚本
* 发送交易来花费btc
在整个过程中，我们并没有去使用私钥签名。
  
### 将要花费的output
我们从上面交易4bd84af06f8c6749f3a5ccbb11a415bd9eaa3001a9fb66f800e2fad89def44b3的第0个output作为我们新交易的输入来花费这1个btc，我们再生成一个新地址来接收0.5个btc，并向当前的地址找零0.4个btc，0.1个btc作为矿工费。
我的新地址：mp6D7UjybN9ZmxQd6r91kRq9goqvneEz4o
找零地址：mpJRdFjFcQM6o4JN2D1mogUt7FwpdWbDfh

### 构造交易

先生成交易
```
bash-5.0# bitcoin-cli -regtest createrawtransaction "[{\"txid\":\"4bd84af06f8c6749f3a5ccbb11a415bd9eaa3001a9fb66f800e2fad89def44b3\",\"vout\":0}]" "[{\"mp6D7UjybN9ZmxQd6r91kRq9goqvneEz4o\":\"0.5\"},{\"mpJRdFjFcQM6o4JN2D1mogUt7FwpdWbDfh\":\"0.4999\"}]"
0200000001b344ef9dd8fae200f866fba90130aa9ebd15a411bbcca5f349678c6ff04ad84b0000000000ffffffff0280f0fa02000000001976a9145e0b1a28814bf2bfe486331566020c62e85ac18788ac70c9fa02000000001976a914605a8a137e7fb06ae94e8663d0fc9a83f9c3152288ac00000000
```

这个交易是未签名的，我们可以decode看看交易的内容
```
bash-5.0# bitcoin-cli -regtest decoderawtransaction 0200000001b344ef9dd8fae200f866fba90130aa9ebd15a411bbcca5f349678c6ff04ad84b0000000000ffffffff0280f0fa02000000001976a9145e0b1a28814bf2bfe486331566020c62e85ac18788ac70c9fa02000000001976a914605a8a137e7fb06ae94e8663d0fc9a83f9c3152288ac00000000
{
  "txid": "f3828e4f1410e83e190753ffb9fcaf6f957d7b53b2123d37c6b38463f0935af1",
  "hash": "f3828e4f1410e83e190753ffb9fcaf6f957d7b53b2123d37c6b38463f0935af1",
  "version": 2,
  "size": 119,
  "vsize": 119,
  "weight": 476,
  "locktime": 0,
  "vin": [
    {
      "txid": "4bd84af06f8c6749f3a5ccbb11a415bd9eaa3001a9fb66f800e2fad89def44b3",
      "vout": 0,
      "scriptSig": {
        "asm": "",
        "hex": ""
      },
      "sequence": 4294967295
    }
  ],
  "vout": [
    {
      "value": 0.50000000,
      "n": 0,
      "scriptPubKey": {
        "asm": "OP_DUP OP_HASH160 5e0b1a28814bf2bfe486331566020c62e85ac187 OP_EQUALVERIFY OP_CHECKSIG",
        "hex": "76a9145e0b1a28814bf2bfe486331566020c62e85ac18788ac",
        "reqSigs": 1,
        "type": "pubkeyhash",
        "addresses": [
          "mp6D7UjybN9ZmxQd6r91kRq9goqvneEz4o"
        ]
      }
    },
    {
      "value": 0.49990000,
      "n": 1,
      "scriptPubKey": {
        "asm": "OP_DUP OP_HASH160 605a8a137e7fb06ae94e8663d0fc9a83f9c31522 OP_EQUALVERIFY OP_CHECKSIG",
        "hex": "76a914605a8a137e7fb06ae94e8663d0fc9a83f9c3152288ac",
        "reqSigs": 1,
        "type": "pubkeyhash",
        "addresses": [
          "mpJRdFjFcQM6o4JN2D1mogUt7FwpdWbDfh"
        ]
      }
    }
  ]
}
```

### 尝试发送未签名交易
尝试发送一个未签名的交易：
```
bash-5.0# bitcoin-cli -regtest sendrawtransaction 0200000001b344ef9dd8fae200f866fba90130aa9ebd15a411bbcca5f349678c6ff04ad84b0000000000ffffffff0280f0fa02000000001976a9145e0b1a28814bf2bfe486331566020c62e85ac18788ac70c9fa02000000001976a914605a8a137e7fb06ae94e8663d0fc9a83f9c3152288ac00000000
error code: -26
error message:
mandatory-script-verify-flag-failed (Operation not valid with the current stack size) (code 16)
```
script verify failed, 因为没有签名，验签脚本执行失败。

### 签名再发送

我们的交易输入引用的output如下:
```
    {
      "value": 1.00000000,
      "n": 0,
      "scriptPubKey": {
        "asm": "OP_DUP OP_HASH160 605a8a137e7fb06ae94e8663d0fc9a83f9c31522 OP_EQUALVERIFY OP_CHECKSIG",
        "hex": "76a914605a8a137e7fb06ae94e8663d0fc9a83f9c3152288ac",
        "reqSigs": 1,
        "type": "pubkeyhash",
        "addresses": [
          "mpJRdFjFcQM6o4JN2D1mogUt7FwpdWbDfh"
        ]
      }
    },
```
从上面的脚本可以看出，我们需要地址mpJRdFjFcQM6o4JN2D1mogUt7FwpdWbDfh对应的私钥签名，先导出私钥:
```
bash-5.0# bitcoin-cli dumpprivkey mpJRdFjFcQM6o4JN2D1mogUt7FwpdWbDfh
cSQn2eGCUMDh1pgMK8v3SxePfdEQTZH71GjNWZt2ryyrzFbEbDTx
```

使用私钥签名交易：
```
bash-5.0# bitcoin-cli -regtest signrawtransactionwithkey 0200000001b344ef9dd8fae200f866fba90130aa9ebd15a411bbcca5f349678c6ff04ad84b0000000000ffffffff0280f0fa02000000001976a9145e0b1a28814bf2bfe486331566020c62e85ac18788ac70c9fa02000000001976a914605a8a137e7fb06ae94e8663d0fc9a83f9c3152288ac00000000 "[\"cSQn2eGCUMDh1pgMK8v3SxePfdEQTZH71GjNWZt2ryyrzFbEbDTx\"]"
{
  "hex": "0200000001b344ef9dd8fae200f866fba90130aa9ebd15a411bbcca5f349678c6ff04ad84b000000006a4730440220179ffbbf5afb409d9ec84aa9e320efcf4f52dd235694e10d87b01285e5edd97902206e6eb07a09622f60934fd4ef2fc759fdefabbb333a84476aa9f62de1366810b101210321b66191a47bc5d45a1467cc73e88b15d6b6e828e9bb6b26ef19c7d86010199dffffffff0280f0fa02000000001976a9145e0b1a28814bf2bfe486331566020c62e85ac18788ac70c9fa02000000001976a914605a8a137e7fb06ae94e8663d0fc9a83f9c3152288ac00000000",
  "complete": true
}
```

检查签名后的消息：
```
bash-5.0# bitcoin-cli -regtest decoderawtransaction 0200000001b344ef9dd8fae200f866fba90130aa9ebd15a411bbcca5f349678c6ff04ad84b000000006a4730440220179ffbbf5afb409d9ec84aa9e320efcf4f52dd235694e10d87b01285e5edd97902206e6eb07a09622f60934fd4ef2fc759fdefabbb333a84476aa9f62de1366810b101210321b66191a47bc5d45a1467cc73e88b15d6b6e828e9bb6b26ef19c7d86010199dffffffff0280f0fa02000000001976a9145e0b1a28814bf2bfe486331566020c62e85ac18788ac70c9fa02000000001976a914605a8a137e7fb06ae94e8663d0fc9a83f9c3152288ac00000000
{
  "txid": "3dd38d9b4bdbe58fe1df236bf5d230dba6111eaefea8de987f45694a05d17a47",
  "hash": "3dd38d9b4bdbe58fe1df236bf5d230dba6111eaefea8de987f45694a05d17a47",
  "version": 2,
  "size": 225,
  "vsize": 225,
  "weight": 900,
  "locktime": 0,
  "vin": [
    {
      "txid": "4bd84af06f8c6749f3a5ccbb11a415bd9eaa3001a9fb66f800e2fad89def44b3",
      "vout": 0,
      "scriptSig": {
        "asm": "30440220179ffbbf5afb409d9ec84aa9e320efcf4f52dd235694e10d87b01285e5edd97902206e6eb07a09622f60934fd4ef2fc759fdefabbb333a84476aa9f62de1366810b1[ALL] 0321b66191a47bc5d45a1467cc73e88b15d6b6e828e9bb6b26ef19c7d86010199d",
        "hex": "4730440220179ffbbf5afb409d9ec84aa9e320efcf4f52dd235694e10d87b01285e5edd97902206e6eb07a09622f60934fd4ef2fc759fdefabbb333a84476aa9f62de1366810b101210321b66191a47bc5d45a1467cc73e88b15d6b6e828e9bb6b26ef19c7d86010199d"
      },
      "sequence": 4294967295
    }
  ],
  "vout": [
    {
      "value": 0.50000000,
      "n": 0,
      "scriptPubKey": {
        "asm": "OP_DUP OP_HASH160 5e0b1a28814bf2bfe486331566020c62e85ac187 OP_EQUALVERIFY OP_CHECKSIG",
        "hex": "76a9145e0b1a28814bf2bfe486331566020c62e85ac18788ac",
        "reqSigs": 1,
        "type": "pubkeyhash",
        "addresses": [
          "mp6D7UjybN9ZmxQd6r91kRq9goqvneEz4o"
        ]
      }
    },
    {
      "value": 0.49990000,
      "n": 1,
      "scriptPubKey": {
        "asm": "OP_DUP OP_HASH160 605a8a137e7fb06ae94e8663d0fc9a83f9c31522 OP_EQUALVERIFY OP_CHECKSIG",
        "hex": "76a914605a8a137e7fb06ae94e8663d0fc9a83f9c3152288ac",
        "reqSigs": 1,
        "type": "pubkeyhash",
        "addresses": [
          "mpJRdFjFcQM6o4JN2D1mogUt7FwpdWbDfh"
        ]
      }
    }
  ]
}
```
发送交易：
```
bash-5.0# bitcoin-cli -regtest sendrawtransaction 0200000001b344ef9dd8fae200f866fba90130aa9ebd15a411bbcca5f349678c6ff04ad84b000000006a4730440220179ffbbf5afb409d9ec84aa9e320efcf4f52dd235694e10d87b01285e5edd97902206e6eb07a09622f60934fd4ef2fc759fdefabbb333a84476aa9f62de1366810b101210321b66191a47bc5d45a1467cc73e88b15d6b6e828e9bb6b26ef19c7d86010199dffffffff0280f0fa02000000001976a9145e0b1a28814bf2bfe486331566020c62e85ac18788ac70c9fa02000000001976a914605a8a137e7fb06ae94e8663d0fc9a83f9c3152288ac00000000
3dd38d9b4bdbe58fe1df236bf5d230dba6111eaefea8de987f45694a05d17a47
```

检查余额：
```
bash-5.0# bitcoin-cli -regtest getreceivedbyaddress mp6D7UjybN9ZmxQd6r91kRq9goqvneEz4o
0.50000000
```

以上是正常签名并发送交易发送btc的整个过程。
















