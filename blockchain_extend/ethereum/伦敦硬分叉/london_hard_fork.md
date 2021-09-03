# ethereum london hard fork

## 以太坊的问题

目前以太坊的收费机制是拍卖机制，用户出价，矿工选择出价最高的交易，将起打包进区块。这种机制简单且高效，但存在问题：

1. 竞价效率低，用户发送交易时选择出价费用，而矿工选择最高出价的交易，对用户来说，出价多少合适很难预估，经常出现恶性出价竞争导致超额支付，用户交易费用应该和网络拥塞正相关，而不是纯粹的用户竞价。
2. 交易延迟，每个区块gas limit的限制导致交易可能需要等待几个区块才能被打包进区块。这需要区块大小可伸缩，在一个长周期里区块平均大小是限制的，而不是每个区块大小都是限制的。
3. 链安全性，随着区块奖励减少，未来交易费成为矿工的主要激励，这就要求有足够的交易规模才能支撑网络的安全性。


## london hard fork

* [EIP1559](https://eips.ethereum.org/EIPS/eip-1559): ETH1.0交易费用市场改革,本质上是关于以太坊网络交易定价机制的解决方案，包括交易的base fee和tips、区块大小弹性机制
* [EIP3198](https://eips.ethereum.org/EIPS/eip-3198): basefee操作码
* [EIP3529](https://eips.ethereum.org/EIPS/eip-3529): 降低gas退款
* [EIP3541](https://eips.ethereum.org/EIPS/eip-3541): 拒绝0xef开头的新合约
* [EIP3554](https://eips.ethereum.org/EIPS/eip-3554): 难度炸弹推迟到2021.12.01，区块难度将提升。区块越来越难被挖掘，从而变得无利可图，矿工将停止在eth1.0上的挖矿。

### EIP1559

#### EIP1559内容

将交易费用分为基本费用和小费，基本费用被销毁，矿工获得小费和新区块奖励。

如果一个区块的gas使用量超过该区块的gas limit，基本费用在接下来的区块增加，而当一个区块的gas使用量低于该区块的gas limit时，基本费用则下调。

bakhta观点：

减少当前以太坊的拥堵和高费用并不是EIP 1559的目的，EIP1559目的是引入"区块弹性"的概念，以太坊平台的理论最大容量将增加一倍。

交易费用是供求关系的函数。从技术上讲，平均可用区块空间不会增加，因为基本费用机制的设计倾向于达到最大区块容量的一半。因此，简短的答案是“否”，这个升级将不是以太坊解决其可扩展性问题所需的长期解决方案。

但是，从更乐观的角度来说，随着越来越多的第二层解决方案继续被采用，所有网络的费用和拥塞问题最终都将得到解决。

#### EIP1559安全

1. EIP1559将最大区块大小调大，如果矿工机器不能及时处理区块将导致一些问题。
2. 长时间周期看，平均区块大小和目前未引入EIP1559没有区别，只是区块大小偶尔会变大，需要确保节点能处理最大的区块。
3. 大部分交易只包含基础费用，打包到区块的交易顺序取决于每个节点自己的实现，建议根据接受交易的时间排序，矿工还是优先选择高tip费用的。
4. 矿工可以选择生产空块或者只包含有tip的交易，只要这类攻击持续下去，那么遵守规则的节点将获利。(It is possible that miners will mine empty blocks until such time as the base fee is very low and then proceed to mine half full blocks and revert to sorting transactions by the priority fee. While this attack is possible, it is not a particularly stable equilibrium as long as mining is decentralized. Any defector from this strategy will be more profitable than a miner participating in the attack for as long as the attack continues (even after the base fee reached 0). Since any miner can anonymously defect from a cartel, and there is no way to prove that a particular miner defected, the only feasible way to execute this attack would be to control 50% or more of hashing power. If an attacker had exactly 50% of hashing power, they would make no money from priority fee while defectors would make double the money from priority fees. For an attacker to turn a profit, they need to have some amount over 50% hashing power, which means they can instead execute double spend attacks or simply ignore any other miners which is a far more profitable strategy.)
5. 以太坊将没有固定量的增发，如果焚毁的大于挖矿奖励，eth变成通缩，如果挖矿奖励大于焚毁，eth是通胀的。

[以太坊EIP1559实施后，51%攻击更难了吗](http://www.liujia.name/?p=40462)

### EIP3529

#### EIP3529内容

[EIP-3529：减少gas返还](http://www.finacerun.com/home/news/detail/article_id/71963.html)

## Long Hard Fork所做的修改

### 区块修改

```
class Block:
	parent_hash: int = 0
	uncle_hashes: Sequence[int] = field(default_factory=list)
	author: int = 0
	state_root: int = 0
	transaction_root: int = 0
	transaction_receipt_root: int = 0
	logs_bloom: int = 0
	difficulty: int = 0
	number: int = 0
	gas_limit: int = 0 # note the gas_limit is the gas_target * ELASTICITY_MULTIPLIER
	gas_used: int = 0
	timestamp: int = 0
	extra_data: bytes = bytes()
	proof_of_work: int = 0
	nonce: int = 0
	base_fee_per_gas: int = 0
```

增加了参数base_fee_per_gas，gas_limit协议有变更。

base_fee_per_gas区块基础费用，是指当前区块所有交易的基础费用，EIP1559中交易费用包含了两部分，base fee（将被burn掉）和tip（奖励给矿工）。区块头中的base_fee_per_gas就指定了打包到该区块的所有交易的base fee。

gas_limit: 区块大小限制，EIP1559中的区块大小是弹性的，随着网络交易拥塞程度而波动。有一个gas_target概念，gas_target类似EIP1559以前的gas_limit，而此处的gas_limit是gas_target * 2，意味着下一个区块可以突然最大增大2倍。

当前区块的base_fee_per_gas计算：
```
ELASTICITY_MULTIPLIER = 2
parent_gas_target = self.parent(block).gas_limit // ELASTICITY_MULTIPLIER
parent_gas_limit = self.parent(block).gas_limit
parent_gas_used = self.parent(block).gas_used

assert block.gas_used <= block.gas_limit, 'invalid block: too much gas used'
assert block.gas_limit < parent_gas_limit + parent_gas_limit // 1024, 'invalid block: gas limit increased too much'
assert block.gas_limit > parent_gas_limit - parent_gas_limit // 1024, 'invalid block: gas limit decreased too much'
assert block.gas_limit >= 5000

parent_base_fee_per_gas = self.parent(block).base_fee_per_gas
# check if the base fee is correct
if INITIAL_FORK_BLOCK_NUMBER == block.number:
    expected_base_fee_per_gas = INITIAL_BASE_FEE
elif parent_gas_used == parent_gas_target:
    expected_base_fee_per_gas = parent_base_fee_per_gas
elif parent_gas_used > parent_gas_target:
    gas_used_delta = parent_gas_used - parent_gas_target
    base_fee_per_gas_delta = max(parent_base_fee_per_gas * gas_used_delta // parent_gas_target // BASE_FEE_MAX_CHANGE_DENOMINATOR, 1)
    expected_base_fee_per_gas = parent_base_fee_per_gas + base_fee_per_gas_delta
else:
    gas_used_delta = parent_gas_target - parent_gas_used
    base_fee_per_gas_delta = parent_base_fee_per_gas * gas_used_delta // parent_gas_target // BASE_FEE_MAX_CHANGE_DENOMINATOR
    expected_base_fee_per_gas = max(parent_base_fee_per_gas - base_fee_per_gas_delta, 0)
assert expected_base_fee_per_gas == block.base_fee_per_gas, 'invalid block: base fee not correct'
```

### 两类交易

LegacyTx

```
class TransactionLegacy:
	signer_nonce: int = 0
	gas_price: int = 0
	gas_limit: int = 0
	destination: int = 0
	amount: int = 0
	payload: bytes = bytes()
	v: int = 0
	r: int = 0
	s: int = 0
```

DynamicFeeTx

```
class Transaction1559Payload:
	chain_id: int = 0
	signer_nonce: int = 0
	max_priority_fee_per_gas: int = 0
	max_fee_per_gas: int = 0
	gas_limit: int = 0
	destination: int = 0
	amount: int = 0
	payload: bytes = bytes()
	access_list: List[Tuple[int, List[int]]] = field(default_factory=list)
	signature_y_parity: bool = False
	signature_r: int = 0
	signature_s: int = 0
```

对比交易数据模型可以看到，EIP1559没有了gas_price，增加了max_priority_fee_per_gas和max_fee_per_gas。access_list是EIP2930引入的，此处不做介绍。

在London hard fork后，Legacy交易是兼容支持的，但Legacy交易将自动被升级为DynamicFee交易，Legacy交易中的gas_price被作为DynamicFee交易中的max_priority_fee_per_gas和max_fee_per_gas对待。

交易验证：
1. 用户balance必须大于转账金额和手续费金额
2. 用户交易的max_fee_per_gas必须大于区块的base_fee_per_gas
3. 用户交易的max_fee_per_gas必须小于2^256
4. 用户交易的max_priority_fee_per_gas必须小于2^256
5. 用户交易的max_fee_per_gas必须大于max_priority_fee_per_gas

交易执行:
1. 焚毁base fee，base fee = base_fee_per_gas * gas_used
2. 计算矿工小费并支付tip，tip = min(transaction.max_priority_fee_per_gas, transaction.max_fee_per_gas - block.base_fee_per_gas) * gas_used


## London Hard Fork交易构造方式

### 两种不同的写法

1.1 使用Legacy交易，这和EIP1559以前兼容，但交易参数gas_price意义不一样，EIP1559交易的gas_price包含了base fee和tip，但仍然可以通过eth client的SuggestGasPrice获得，所以编程上仍然兼容。签名器可以用London signer，也可以是以前的非London signer。
```
	ethClient, _ := ethclient.Dial("https://ropsten.infura.io/v3/19e799349b424211b5758903de1c47ea")
	ctx := context.Background()
	privateKey,_ := crypto.HexToECDSA("994D7BC4C1DE95D4C3069F3F64A032BC55482970F40141D074141D099CC88569")
	fromAddr := crypto.PubkeyToAddress(privateKey.PublicKey)
	nonce,_ := ethClient.PendingNonceAt(ctx, fromAddr)
	contractAddr := common.HexToAddress("239100e629a9ca8e0bf45c7892b0fc72d78aa97a")
	amount := big.NewInt(0)
	gasPrice, _ := ethClient.SuggestGasPrice(ctx)
	gasLimit := uint64(8000000)
	tx := types.NewTransaction(nonce, contractAddr, amount, gasLimit, gasPrice, nil)
	/*
	tx = types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &contractAddr,
		Value:    amount,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     nil,
	})
	 */
	signer := types.MakeSigner(params.RopstenChainConfig, new(big.Int).SetInt64(12900000))
	signed_tx, err := types.SignTx(tx, signer, privateKey)
```

1.2 使用legacy交易，不支付tip，只支付base fee

```
	ethClient, _ := ethclient.Dial("https://ropsten.infura.io/v3/19e799349b424211b5758903de1c47ea")
	ctx := context.Background()
	privateKey,_ := crypto.HexToECDSA("994D7BC4C1DE95D4C3069F3F64A032BC55482970F40141D074141D099CC88569")
	fromAddr := crypto.PubkeyToAddress(privateKey.PublicKey)
	nonce,_ := ethClient.PendingNonceAt(ctx, fromAddr)
	contractAddr := common.HexToAddress("239100e629a9ca8e0bf45c7892b0fc72d78aa97a")
	amount := big.NewInt(0)
	gasPrice, _ := ethClient.SuggestGasPrice(ctx)
	gasTip, _ := ethClient.SuggestGasTipCap(ctx)
	gasLimit := uint64(8000000)
	tx := types.NewTransaction(nonce, contractAddr, amount, gasLimit, gasPrice.Sub(gasPrice, gasTip), nil)
	/*
	tx = types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &contractAddr,
		Value:    amount,
		Gas:      gasLimit,
		GasPrice: gasPrice.Sub(gasPrice, gasTip),
		Data:     nil,
	})
	*/
	signer := types.MakeSigner(params.RopstenChainConfig, new(big.Int).SetInt64(12900000))
	signed_tx, err := types.SignTx(tx, signer, privateKey)
```

2.1 使用DynamicFee交易，指定可以接受的最大base fee和最大tip

```
	ethClient, _ := ethclient.Dial("https://ropsten.infura.io/v3/19e799349b424211b5758903de1c47ea")
	ctx := context.Background()
	privateKey,_ := crypto.HexToECDSA("994D7BC4C1DE95D4C3069F3F64A032BC55482970F40141D074141D099CC88569")
	fromAddr := crypto.PubkeyToAddress(privateKey.PublicKey)
	nonce,_ := ethClient.PendingNonceAt(ctx, fromAddr)
	contractAddr := common.HexToAddress("239100e629a9ca8e0bf45c7892b0fc72d78aa97a")
	amount := big.NewInt(0)
	//gasPrice, _ := ethClient.SuggestGasPrice(ctx)
	gasTip, _ := ethClient.SuggestGasTipCap(ctx)
	gasLimit := uint64(8000000)
	tx := types.NewTx(&types.DynamicFeeTx{
		Nonce:    nonce,
		To:       &contractAddr,
		Value:    amount,
		Gas:      gasLimit,
		GasFeeCap: abi.MaxUint256,
		GasTipCap: gasTip,
		Data:     nil,
	})
	signer := types.MakeSigner(params.RopstenChainConfig, new(big.Int).SetInt64(12900000))
	signed_tx, err := types.SignTx(tx, signer, privateKey)
```

### 以上写法的区别

1.1和1.2两种写法是指定了这笔交易用户愿意支付的费用，包括base fee和tip，这和非EIP1559没有区别，如果gas price小于区块的基础费用，交易失败，如果大于区块的基础费用，则根据排序结果决定是否能打包到区块，不能则尝试下一个区块。

意义：我愿意支付这些费用，如果小于基础费用就失败，如果大于基础费用则作为tip给矿工。

2.1写法指定了可以接受的最大base fee和最大tip，只要区块的基础费用小于交易中的GasFee，则总是尝试打包进区块，收取的tip则是min(GasFeeCap - base_fee, GasTipCap)，上述GasFeeCap写的是MaxUint256意味着所有的网络基础费用用户交易都接受，而GasTipCap则是支付给矿工的tip，如果写0则不支付tip，只支付网络基础费用。

意义：任何网络波动造成的网络基础费用变动我都接受，而且我额外的支付一部分tip。

### 参考

[London JSON RPC Considerations](https://notes.ethereum.org/@timbeiko/london-json-rpc)

### 启动时间

网络|启动高度|启动时间
:--:|:--:|:--:
Ropsten|10499401|2021.06.24
Mainnet|12965000|2021.08.04


### ethereum钱包支持London Hard Fork

1 构造交易
```
rpc：
personal_signTransaction

param:
type TransactionArgs struct {
From                 *common.Address `json:"from"`
To                   *common.Address `json:"to"`
Gas                  *hexutil.Uint64 `json:"gas"`
GasPrice             *hexutil.Big    `json:"gasPrice"`
MaxFeePerGas         *hexutil.Big    `json:"maxFeePerGas"`
MaxPriorityFeePerGas *hexutil.Big    `json:"maxPriorityFeePerGas"`
Value                *hexutil.Big    `json:"value"`
Nonce                *hexutil.Uint64 `json:"nonce"`
Data  *hexutil.Bytes `json:"data"`
Input *hexutil.Bytes `json:"input"`
AccessList *types.AccessList `json:"accessList,omitempty"`
ChainID    *hexutil.Big      `json:"chainId,omitempty"`
}
```
参数可以只填写gasPrice，也可以填写maxFeePerGas和maxPriorityFeePerGas。London Hard Fork后是兼容以前的。

我建议使用目前的写法，填gasPrice，和目前的兼容。

2 gasPrice

目前gasPrice从配置、geth节点rpc的eth_gasPrice、GasStation的API三个值中取最大值。

London Hard Ford后，geth节点rpc的eth_gasPrice返回的gasPrice包含了base fee和矿工tip两部分。GasStation的API目前不确定。

在没有确定GasStation的API对London Hard Fork的支持情况之前，代码不做修改，这最坏情况下退化为从geth节点获取gasPrice。

## 动态

1. 2021.07.23  Geth升级

修复EIP1559导致的共识失败，以太坊主网上的geth用户需要在London Hard Fork激活之前升级到geth v1.10.6版本才能保持共识。

2. 2021.07.23 Vitalik回应EIP-1559推迟传言

vitalik回应EIP-1559推迟传言：大部分Geth核心开发者诶呦参与EthCC。

3. 2021.07.21 Ropsten测试网分叉

以太坊客户端OpenEthereum团队在Ropsten测试网区块高度10679538发现节点暂停出块，使用不同客户端可能导致分叉。

## 参考

* [EIP-1559创作者：EIP-1559无法解决以太坊拥堵和高费用问题](https://www.8btc.com/article/6625023)
* [一文告诉你什么是EIP-1559](https://www.btcfans.com/article/35665)
* [Vitalik：以太坊网络或面临经济安全威胁，EIP 1559可解决](https://www.8btc.com/article/625910)
* [EIP-1559究竟在吵什么](https://zhuanlan.zhihu.com/p/361104358)
* [london hard fork announcement for development](https://blog.ethereum.org/2021/06/18/london-testnets-announcement/)
* [全面梳理解密，EIP-1559为何能让市场如此兴奋？](https://www.163.com/dy/article/GEF97JCH0511QUF7.html)
* [解读以太坊提案EIP1559：降低交易费总额和交易费波动性](https://www.jinse.com/blockchain/552029.html)
* [观点：EIP-1559只是徒劳，毫无益处](https://baijiahao.baidu.com/s?id=1683869264859376411&wfr=spider&for=pc)
* [以太坊的伦敦硬分叉预计将于周四抵达，将迎来EIP-1559](https://news.huoxing24.com/20210804094939077146.html)

* [以太坊燃烧第一个24小时，中文社区在关心什么？](https://www.8btc.com/article/6670915)
* [EIP-1559上线之际，泼一点冷水](https://news.huoxing24.com/20210809171533498961.html)

