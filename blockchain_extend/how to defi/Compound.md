# Compound

## 参考

[五分钟看懂借贷机制，什么是DeFi银行Compound](https://www.wwsww.cn/btbzixun/2068.html)

[去中心化算法银行Compound技术解析之概述篇](https://zhuanlan.zhihu.com/p/114319666)

[一文读懂Compound的清算机制](https://mzh1.com/762.html)

## introduction

### ecosystem

* vibrant
* investor
* speculator (arbitrage)
* trader

* the time value of assets

* non-zero-sum wealth

* two major flaws exist()

### centralized exchange

* trade assets on margin
* borrowing markets built into the exchange

* trust-based system (not hacked, abscond with your asset, incorrectly close out your position)
* limited to certain customer groups
* limited to a small number of assets (the most mainstream)
* balances and positions are virtual (cannot move a position on chain, cannot use borrowed token in a smart contract or ICO, inaccessible to dapps)

### peer to peer protocol

* collateralized loan & un-collateralized loan
* significant cost & friction
* lenders are required to post, manager, supervise loan offers and active loan
* loan fulfillment is often slow and asynchronous

### compound

* decentralized system for the frictionless borrowing without the flaws of existing approaches
* enabling proper money markets to function
* a positive-yield approach to storing assets

## Protocol

* establish money markets
* algorithmically derived interest rates （based on the supply and demand for the asset)
* supplier / borrower
* a floating interest rate
* negotiate term (such as maturity, interest rate, collateral)
  
* transparent
* publicly-inspectable ledger
* record of all transactions and historical interest rates.

### supplying assets

* supply an asset
* withdraw an asset
* cToken
* underlying asset

* long-term investment

### borrowing assets

* maturity dates
* market forces

* collateral factor

* borrowing capacity = sum of the value of an accounts underlying token balance * the collateral factors

* user are able to borrow up to , but not exceeding, their borrowing capacity
* an account can token no action(borrow, transfer cToken collateral, redeem cToken collateral) that would raise the total value of borrowed asset above their borrowing capacity.

* outstanding

* liquidation & liquidation discount

* arbitrageurs step in to reduce the borrower's exposure and eliminate the protocol's risk

* close factor

* portfolio

### interest rate model

* interest rates should increase as a function of demand
* utilization ratio U = Borrows / (Cash + Borrows)
* Borrowing Interest Rate = 2.5% + U * 20%
* Supplier Interest Rate = Borrowing Interest Rate * U

## architecture

### cToken

* mint - transfer an underlying asset into the market, updates msg.sender's cToken balance
* redeem - transfer an underlying asset out of the market, updates msg.sender's cToken balance
* borrow - check msg.sender collateral value, and if sufficient, transfer the underlying asset out of the market to msg.sender, updates msg.sender's borrow balance
* repayBorrow - transfer the underlying asset into the market, updates the borrower's borrow balance
* liquidate - transfer the underlying asset into the market, updates the borrower's borrow balance, then transfers cToken collateral from the borrower to msg.sender

* cToken and the underlying asset exchange rate = (underlying balance + total borrow balance - reserves) / cToken supply
* exchange rate increase over time

### interest rate mechanic

* interest rate changes result from a user mint, redeem, borrow, repay and liquidate
* the history of each interest rate is captured by an Interest Rate Index

### market dynamic

* Interest Rate Index(n) = Index(n - 1) * (1 + r * t)
* total borrow balance(n) = total borrow balance(n - 1) * (1 + r * t)
* reserve = reserve(n - 1) + total borrow balance(n - 1) * (r * t * reserve factor)

### borrow dynamic

* borrower's balance (including accrued interest) = (the current interest rate index) / (the interest rate index in the user's last check pointed)

### borrowing

* checks the user's account value
* given sufficient collateral
* update the user's borrow balance
* transfer the underlying tokens to the user's ethereum address
* update the money market's floating interest rate

### liquidation

* the value of collateral falling
* borrowed assets increasing in value

* user's borrowing balance exceeds their total collateral value

* exchanges the invoking user's asset for the borrower's collateral, at a slightly better than market price

### price feeds

* price oracle

### comptroller

* 




















































