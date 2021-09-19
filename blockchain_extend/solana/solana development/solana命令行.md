# solana命令行


## solana

* solana --help
* solana <command> --help
* solana --version // 本地的客户端使用的版本
* solana cluster-version // 远端节点的版本  
* solana config get/set  // 获取或者设置solana的环境配置
* solana config set --url https://api.devnet.solana.com

## solana-keygen

* solana-keygen pubkey prompt://
* solana-keygen pubkey /home/solana/my_wallet.json
* solana-keygen pubkey usb://ledger?key=0
[钱包CLI](https://docs.solana.com/wallet-guide/cli)
[钱包](https://docs.solana.com/cli/conventions#keypair-conventions)  

## 生成钱包

solana-keygen new --outfile /Users/tangaoyuan/.config/solana/my_wallet.json
[生成文件系统钱包](https://docs.solana.com/wallet-guide/file-system-wallet)

## token操作

[token](https://docs.solana.com/cli/transfer-tokens)

## 编译合约
* cargo build  // 编译本地系统的执行文件 用于测试
* cargo build-bpf  // 编译solana链上合约的执行文件

## 部署合约

* solana program deploy <program_file> 
* solana program deploy --program-id <keypair_file> <program_file>
* solana program show <account_address>

rogram Id: 3KS2k14CmtnuVv2fvYcvdrNgC94Y11WETBpMUGgXyWZL

[部署合约](https://docs.solana.com/cli/deploy-a-program)
