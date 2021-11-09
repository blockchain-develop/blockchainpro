# AMM

AMM核心基础：
* 做市商函数设计
* 流动性分布的控制

## uniswap

* 恒等乘积做市商 CPMM
    * X * Y = K

* 无常损失 IL
    * 主要原因是在AMM机制下如果流动性池中的非稳定资产价格上涨或下跌，做市商会完全自动的与市场一般交易者做出相反的行为，越涨越卖，越跌越买，因此池子中资产价格上涨，其数量就会减少，资产价格下降，其数量就会增加

* 损益
    * AMM机制影响流动性做市商损益全部因素，主要有四个方面，分别是Gas费用、交易手续费、非稳定币价格变动带来的损益和无常损失

* 流动性分布与范围订单

## 问题

* 无常损失
* 价格滑点
* 抢跑  
* 资本效率低
* 单边流动性注入

## AMM探索方向

### 提升资本效率
* uniswap v3为代表的流动性主动管理
    * LP的范围订单，设置虚拟储备，公平性问题  
* Curve v2为代表的流动性被动管理
    * 恒定做市公式、内部预言机价格盯住机制  
* Balancer V2为代表的资产管理库
    * 沉淀资产投资

### 做市公式
* Bancor Network的联合曲线
* uniswap的恒等乘积
* Balancer的常数不变式和智能订单路由（SOR）
* Curve结合恒定加和和恒定乘积
* TWAMM

### 收益
* 聚合器、收益最大化、收益耕作（yield farming） 1inch
    * 如何在一个流动性池中寻找便宜的汇率
    * 如何在多个流动性源中寻找便宜的汇率
    * 如何在做市、借贷、资产管理和整个生态中套利  
* 对冲

### 解决无常损失
* Pivot算法

### 单边流动性




* [自动做市商（AMM）算法的数学原理及其未来发展（上篇）](https://www.jinse.com/blockchain/1157635.html)
* [自动做市商（AMM）算法的数学原理及其未来发展](https://www.jinse.com/blockchain/1161282.html)
* [A Mathematical View of Automated Market Maker (AMM) Algorithms and Its Future](https://medium.com/anchordao-lab/automated-market-maker-amm-algorithms-and-its-future-f2d5e6cc624a)
* [How to DeFi Beginner](https://nigdaemon.gitbook.io/how-to-defi-beginnerv2/)
* [How to DeFi Advanced](https://nigdaemon.gitbook.io/how-to-defi-advanced-zhogn-wen-b/)