# Uniswap

## introduction
* automated token exchange
* easy-of-use, gas efficiency, censorship resistance, zero rent extraction

* order book(match between buyers and sellers)

* liquidity reserves
* constant product market maker mechanism (keep overall reserves in relative equilibrium)
* liquidity providers

## founder
* Hayden Adams

## factory
* a separate exchange contract
* a reserve of ETH and associated ERC20
* trade between ETH and ERC20

## registry
* link factory
* trade direct ERC20 to ERC20 using ETH as an intermediary

## create exchanges

## ETH - ERC20 trades

## ERC20 - ERC20 trades

## Swap and Transfer
* swap indicate seller token amount
* transfer indicate buyer token amount

## adding liquidity
* the first liquidity provider to join a pool sets the initial exchange rate by depositing what they believe to be an equivalent value of ETH and ERC20 tokens
* if the ratio is off, arbitrage traders
* if the exchange rate is bad, there is a profitable arbitrage opportunity
* depositing ETH and ERC20 using the exchange rate into the associated exchange contract

* token pool
* eth pool
* lp holders and supply of lp

amount minted of lp = total amount of lp * (eth deposited / eth pool)
token deposited = token pool * (amount minted of lp / total amount of lp)

## removing liquidity

* ETH and ERC20 are withdrawn at the current exchange rate (reserve ratio)

eth withdraw = (amount of lp) / (total amount of lp) * (eth pool)
token withdraw = (amount of lp) / (total amount of lp) * (token pool)

## front run
* set minimum/maximum





































