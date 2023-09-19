# Ambire Wallet技术解读

[Ambire Wallet Whitepaper](https://ambire.notion.site/ambire/Ambire-Wallet-Whitepaper-d502e54caf584fe7a67f9b0a018cd10f)
[Ambire Wallet Srouce](https://github.com/ambireTech/wallet)

## intro
* web3 wallet
* self-custody
* secure
* smart contract account & account abstraction & ERC-4337
* account can be authenticated via hardware wallets / in-browser keys & recovery

## ethereum wallet tough problem

### transaction fee
* transaction fee(gas price & gas limit & eip-1559), exposing the underlying complexities to users
* must be ETH

### nonce
* transactions are sequential
* stuck transactions

### 交易可视化和执行安全
* ETH transfer可以可视化，但合约执行完全没法可视化
* 无法可视化会导致用户执行错误的操作

### ERC20 approvals UX & security
* ERC20 approval是user授权另一个地址可以transfer user的token。比如授权给DEX，DEX和ERC20都是合约。用户发送token到DEX做swap，如果用户call erc20合约发送token到dex，则token丢失，dex合约不会被执行，用户需要call dex合约，dex合约call erc20合约将token transfer给自己。分别为push token和pull token。合约pull token，则需要用户事先额外执行approve transaction
* leaky abstraction: ERC20 approval的底层设计泄漏到UX layer
  * execute a separate transaction before the actual action
  * infinite approval
  * approving a smart contract maybe not safe. contract can pull your funds without you action & contract maybe upgradable

### key management
* private key / seed phrase are notoriously hard to keep safe
* dapp经常需要使用hot wallet
* 不能transfer所有权，例如将aave或者uniswap的持仓切换到更安全的wallet，更换wallet是一个巨大的stuck

## smart contract solve most UX issues
* batched transactions: one transaction can execute multiple user operations
* erc20 minimal approvals, batch together the approval with the specific user action, approval seamless to the user
* fee payments in stable coins or any erc20
* upgradable security. hot wallet & email/password account, transfer control to new wallet
* account recovery & social recovery
* transaction batching allows complex interactions
* no nonce management & replacing & canceling transactions
* simulating transaction & 可视化交易
* wallet stability，用户只需要对用户操作签名，无关链升级，如EIP 1559升级就修改了transaction接口

## smart contract wallet， why no adoption up until now
* technical limitations that have been solved recently (EIP 1271)

## ambire wallet features
* hardware wallet support
* connect any dapp through WalletConnect
* automatic transaction fee management
* paying transaction fees in stable coins
* dashboard of all assets
* transaction preview
* build-in swaps and cross-chain transfers
* multiple network
* sign up with email/password
* deposit FIAT
* gas tank
* transaction batching
* front-running/sandwiching protection via flashbots
* multiples signers can be used to control the account

## email/password accounts
* authentication is non-custodial
* on-chain 2/2 multisig, encrypted with the user's password & ambire backend via a HSM
* access the funds need two keys
* recovery procedure can be started with on key, with timelocked
