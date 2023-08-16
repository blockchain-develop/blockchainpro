# optimistic rollup

## introduce
* transactions are written on ethereum，computation and storage of the contract are done off-chain
* rollup validator posts on-chain an assertion(rollup block)
* rollup systems differ is in how they ensure that the rollup blocks are correct

### rollup block
* a list of actions
* a hash of its state

## optimistic rollup
* rollup block does not contain an accompanying proof guaranteeing its validity.
* when rollup block is posted on-chain, the validator making a bond.
* in time window, anyone can post their own bond and challenge the rollup block.
* fraud proof
* dispute resolution

### multi-round interactive rollup
* challenge window
* back-and-forth interactive
* referee determines that one party make a false claim, punishes that party by taking their bond

### executing the chain
* any user can submit an rollup block about the execution of the rollup chain
* after rollup block is submitted to ethereum, a challeng period begins, any other user can challenge the correctness of the rollup block
* if a challenge has been initiated, the dispute is mediated by ethereum, it is guaranteed that an honest user will always win a challenge
* validators need place bonds

### submitting transactions
* all transactions executed on the arbitrum rollup chain are submitted to an inbox smart contract running on ethereum
* anyone monitoring the inbox can know the correct state of the arbitrum chain by simply executing the transactions from that contract

