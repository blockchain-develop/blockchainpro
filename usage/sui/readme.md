# sui testnet零基础发token

* 准备环境
* issue token
* transfer token

## 准备环境

### 安装sui cli
* [guide](https://docs.sui.io/build/install)
* 选择testnet
* 测试确保sui已经安装完成，which sui

### 配置sui cli
* sui client 查看可用的命令
* sui client envs 查看配置的所有可用的环境，active的是当前使用的
* sui client new-env --alias <ALIAS> --rpc <RPC-SERVER-URL> 创建一个新的环境， 例如，增加测试环境，sui client new-env --alias testnet --rpc https://fullnode.testnet.sui.io:443
* sui client switch --env <ALIAS> 切换环境

### 准备账户
* 创建新地址  sui client new-address ed25519
* 查看所有地址 sui client addresses
* 查看当前使用的地址 sui client active-address
* 切换当前使用的地址 sui client switch --address 0xa3c00467938b392a12355397bdd3d319cea5c9b8f4fc9c51b46b8e15a807f030
* 获取token [guide](https://docs.sui.io/build/faucet)

## 发行token

### 配置环境
* 获取[sui source](https://github.com/MystenLabs/sui) 
* 使用testnet branch

### token合约
* [token合约source](./erc20)
* Move.toml中的dependencies.sui修改为你本地获取的sui source目录
* 合约发行USDC并给当前地址mint 1b

### 部署合约/发行token
* sui client publish /Users/tangaoyuan/Documents/gopath/src/github.com/blockchainpro/usage/sui/sui_ts/basic/erc20 --gas-budget 10000000
* 确保你当前使用地址有足够的SUI来支付gas
* publish后，从浏览器检查当前使用地址是否完成交易，是否有新发行的USDC以及金额是否正确 [我的测试地址](https://suiexplorer.com/address/0x59ca15713248686598a7f114eda77b55b690e7bfb5fb20c46c551499d7608712?network=testnet)

## transfer token

* sui client objects 查看该地址拥有的objects，sui、erc20都是objects，一个地址可以有多个sui objects
* sui client object <OBJECT_ID> --json 查看某个object的详细信息
* sui client transfer --to <TO> --object-id <OBJECT_ID> --gas-budget <GAS_BUDGET> 将某个object transfer给to
* sui client merge-coin --primary-coin <PRIMARY_COIN> --coin-to-merge <COIN_TO_MERGE> --gas-budget <GAS_BUDGET> 将多个coin objects合并为一个coin object
* sui client split-coin [OPTIONS] --coin-id <COIN_ID> --gas-budget <GAS_BUDGET> (--amounts <AMOUNTS>... | --count <COUNT>) 将一个coin object分成多个coin objects

### transfer 0.001 sui from 0x59ca15713248686598a7f114eda77b55b690e7bfb5fb20c46c551499d7608712 to 0x115d569aa1e7bd2527f48efa1eecf03df3bc3105bedcd729250b6384eb3b9d97
method one:
* sui client objects from地址上的sui object为0xec0f158a7786e25b87543228847c598473ddb8bc9545ea6bec2360086d42db70
* sui client split-coin --coin-id 0xec0f158a7786e25b87543228847c598473ddb8bc9545ea6bec2360086d42db70 --amounts 10000000 --gas-budget 1000000 分离出来一个新的sui object 0xaee1f7299ed601fe056a8f1491add33d11c6c31285debab9eeb656ad8f5619d9，balance为0.001 SUI。可以在浏览器地址页面观察变化。
* sui client transfer --to 0x115d569aa1e7bd2527f48efa1eecf03df3bc3105bedcd729250b6384eb3b9d97 --object-id 0xaee1f7299ed601fe056a8f1491add33d11c6c31285debab9eeb656ad8f5619d9 --gas-budget 10000000 将balance为0.001 SUI的object transfer给目标地址
* [my example](https://suiexplorer.com/txblock/B256mfRGq5uwYMnHiEhaBsrUr5ajEaBcASBrauiizfht?network=testnet)

method two:
* sui client transfer-sui --to 0x115d569aa1e7bd2527f48efa1eecf03df3bc3105bedcd729250b6384eb3b9d97 --sui-coin-object-id 0xec0f158a7786e25b87543228847c598473ddb8bc9545ea6bec2360086d42db70 --amount 1000000 --gas-budget 10000000
* [my example](https://suiexplorer.com/txblock/6DnXujoKYL3yiPvnc1PMerc9YV4DhyQFzzR5spswtMoK?network=testnet)

method three:
* sui client pay-sui --input-coins 0xec0f158a7786e25b87543228847c598473ddb8bc9545ea6bec2360086d42db70 --recipients 0x115d569aa1e7bd2527f48efa1eecf03df3bc3105bedcd729250b6384eb3b9d97 0x799ca22edfa622e550da77f145ef811c7e7ddb6e9ca5fbba285fc5ec0effe4b9 --amounts 1000000 1000000 --gas-budget 10000000
* [my example](https://suiexplorer.com/txblock/GnugFhU9LZUgpp8nGEtLK9NRyQzy3PS2QKaSdgDAd7x5?network=testnet)

### transfer erc20 from 0x59ca15713248686598a7f114eda77b55b690e7bfb5fb20c46c551499d7608712 to 0x115d569aa1e7bd2527f48efa1eecf03df3bc3105bedcd729250b6384eb3b9d97
method one:

method two:
* sui client pay --input-coins 0xed2b739ea880ad500723a9bd72bde3c1be120d3c52b93b13f6dd9efd3ade0526 --recipients 0x115d569aa1e7bd2527f48efa1eecf03df3bc3105bedcd729250b6384eb3b9d97 0x799ca22edfa622e550da77f145ef811c7e7ddb6e9ca5fbba285fc5ec0effe4b9 --amounts 1000000000 1000000000 --gas 0xec0f158a7786e25b87543228847c598473ddb8bc9545ea6bec2360086d42db70 --gas-budget 10000000
* [my example](https://suiexplorer.com/txblock/7LAuujjz1is3EvWierY96fKC28JVit7LUTgW7925uCCW?network=testnet)

