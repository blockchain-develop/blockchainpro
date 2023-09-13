# vote的实现

## onchain voting & offchain voting

## 应用场景
* consensus(Pos, DPos)
* dao

## curve voting escrow
* non-standard ERC20
* Aragon DAO
* cannot be transferred
* time-weight vote

### 用户balance计算
* 用户在指定时刻t的vote balance B = T(l)/T(max)*Amount. 其中T(l)是指staking剩余的时间，T(max)是最大的staking时间，Amount是用户staking的token amount.
* 在staking期间，用户可以增加staking amount，增加staking time，staking结束后，才可以withdraw.
* 用户在指定时刻t1的vote balance B1 = T1(l)/T(max)*Amount, 用户在指定时刻t2的vote balance B2 = T2(l)/T(max)*Amount. 由t1的B1来计算t2的B2, B2 = B1 - (T2 - T1)/T(max)*Amount. T2必须小于这个staking的最小结束时间。
* 用户修改staking(create, increase amount, increase time)后，用户的vote balance变更，计算出时刻t的B，记录为一个Point.

### totalSupply计算
* totalSupply就可以这么计算，S2 = S1 - (T2 - T1)/T(max)*TotalAmount. T2必须小于所有staking的最小结束时间内.
* 用户修改staking(create, increase amount, increase time)后，用户vote balance变更，这个时候根据vote balance变更差值可以计算出时刻t的S，记录为一个Point.
* 用户修改staking(create, increase amount, increase time)后，记录该staking的scope，从而可以知道该staking会在什么时候end. 用于totalSupply的计算.