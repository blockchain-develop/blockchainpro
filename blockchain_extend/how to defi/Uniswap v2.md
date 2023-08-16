# Uniswap

## Ecosystem participants
* liquidity providers
* traders

## smart contract
* singleton factory
    * generic bytecode responsible for powering pairs
    * turn on the protocol charge
* many pairs
    * serving as automated market makers and keeping track of pool token balance
    * expose data which can be used to build decentralized price oracles

## core
* swap
    * trader
* pool
    * liquidity provider
* flash swaps
    * capital free arbitrage
    * instant leverage
* price oracles
    * TWAPs (time weighted average prices)
![](./uniswap%20v2%20price%20oracle.png) 
    * TWAP = (price cumulative 2 - price cumulative 1) / (timestamp 2 - timestamp 2)
![](./uniswap%20price%20oracle%20v2%20twap.png)  
      
## fees
* liquidity provider fees (0.3% fee for swapping tokens)
* protocol fee(0.05% fee for swapping tokens)













   




 