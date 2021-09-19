# hello world on solana

本文演示了solana区块链上合约安装部署和交互的全过程。主要包括：
* 一个hello world链上合约
* 一个交互过程

## 依赖环境

本文中的整个流程需要依赖下面的环境：
* install node
* install npm
* install the latest Rust
* install Solana v1.7.11

## 配置本地客户端环境

1. 配置节点的url为测试网
```
solana config set --url https://api.devnet.solana.com
```

2. 生成新的key
```
solana-keygen new --outfile /Users/tangaoyuan/.config/solana/my_wallet.json 
```

3. 配置钱包
```
solana config set --wallet /Users/tangaoyuan/.config/solana/my_wallet.json 
```

## 编译合约

进入到项目目录 src/program-rust
```
cargo build-bpf
```

## 部署合约

## 运行合约

## web3.js
```
npm install
npm run build:program-rust
solana program depoly dist/program/helloworld.so
npm run start
```

## 参考
[example-helloworld](https://github.com/solana-labs/example-helloworld/blob/master/README.md)
