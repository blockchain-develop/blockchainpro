# 深入理解EVM存储

## 背景

EVM存储有三类，memory，stack和storage，storage是持久化的存储，最终会写入到区块链的世界状态。

这里主要介绍EVM的storage。

## 基本的合约存储
合约代码
```
pragma solidity 0.5.16;

contract Test1 {
    uint a;
    uint b;
    uint c;

    constructor() public {
    }

    function test() public {
        a = 100;
        b = 200;
        c = a + b;
    }
}
```
opcode
```
0:  PUSH1 0x80
1:  PUSH1 0x40
2:  MSTORE
3:  CALLVALUE
4:  DUP1
5:  ISZERO
6:  PUSH1 0x0f
7:  JUMPI
8:  PUSH1 0x00
9:  DUP1
10: REVERT
11: JUMPDEST
12: POP
13: PUSH1 0x04
14: CALLDATASIZE
15: LT
16: PUSH1 0x28
17: JUMPI
18: PUSH1 0x00
19: CALLDATALOAD
20: PUSH1 0xe0
21: SHR
22: DUP1
23: PUSH4 0xf8a8fd6d
24: EQ
25: PUSH1 0x2d
26: JUMPI
27: JUMPDEST
28: PUSH1 0x00
29: DUP1
30: REVERT
31: JUMPDEST
32: PUSH1 0x33
33: PUSH1 0x35
34: JUMP
35: JUMPDEST
36: STOP
37: JUMPDEST
38: PUSH1 0x64
39: PUSH1 0x00
40: DUP2
41: SWAP1
42: SSTORE
43: POP
44: PUSH1 0xc8
45: PUSH1 0x01
46: DUP2
47: SWAP1
48: SSTORE
49: POP
50: PUSH1 0x01
51: SLOAD
52: PUSH1 0x00
53: SLOAD
54: ADD
55: PUSH1 0x02
56: DUP2
57: SWAP1
58: SSTORE
59: POP
60: JUMP
```
solodity和opcode
solidity|opcode|说明
:---:|:--:|:---:
a = 100|38 - 43|内容
b = 200|44 - 49|内容
c = a + b|50 - 59|内容

这个合约中的a、b、c都是需要持久化的状态数据，都需要保存到storage上，从opcode可以看到，solidity使用索引的方式来访问对应的变量，在这个合约中，按照变量a、b、c的顺序依次分配了0、1、2三个索引。

我们部署合约，执行test，接下来检查storage，索引为0处值应该为0x64，索引为1处值应该为0xc8，而索引为2处值应该为0x012c。

我部署的合约地址为"0x452CE6c20bAb1F3766f74ad159ec85Fe1BA8A3dC"，取该合约的存储：
```
truffle(development)> web3.eth.getStorageAt("0x452CE6c20bAb1F3766f74ad159ec85Fe1BA8A3dC", 0);
'0x64'
truffle(development)> web3.eth.getStorageAt("0x452CE6c20bAb1F3766f74ad159ec85Fe1BA8A3dC", 1);
'0xc8'
truffle(development)> web3.eth.getStorageAt("0x452CE6c20bAb1F3766f74ad159ec85Fe1BA8A3dC", 2);
'0x012c'
```
可以看到storage已经修改并成为持久的世界状态。

solidity状态存储的基本思想是将所有的需要持久的状态变量分配索引来存储到storage中。

在EVM的storage中使用mapping(byte32 => byte32)来存储数据，这意味着索引和变量在存储时最长为32 bytes。

对于bool,int8,int256,address,fixed-size byte array(byte1...byte32)类型，按照变量的位置顺序分配对于的索引。

对于边长的动态数据，[]byte, string, mapping(uint256 => uint256)甚至mapping(uint256 => mapping(uint256 => uint256))，又是如何分配索引呢？

## 固定大数据的存储
合约
```
pragma solidity 0.5.16;

contract Test2 {
    uint8[64] a;

    constructor() public {
    }

    function test() public {
        a[0] = 10;
        a[1] = 20;
        a[33] = 30;
    }
}
```
opcode
```
0:   PUSH1 0x80
1:   PUSH1 0x40
2:   MSTORE
3:   CALLVALUE
4:   DUP1
5:   ISZERO
6:   PUSH1 0x0f
7:   JUMPI
8:   PUSH1 0x00
9:   DUP1
10:  REVERT
11:  JUMPDEST
12:  POP
13:  PUSH1 0x04
14:  CALLDATASIZE
15:  LT
16:  PUSH1 0x28
17:  JUMPI
18:  PUSH1 0x00
19:  CALLDATALOAD
20:  PUSH1 0xe0
21:  SHR
22:  DUP1
23:  PUSH4 0xf8a8fd6d
24:  EQ
25:  PUSH1 0x2d
26:  JUMPI
27:  JUMPDEST
28:  PUSH1 0x00
29:  DUP1
30:  REVERT
31:  JUMPDEST
32:  PUSH1 0x33
33:  PUSH1 0x35
34:  JUMP
35:  JUMPDEST
36:  STOP
37:  JUMPDEST
38:  PUSH1 0x0a
39:  PUSH1 0x00
40:  DUP1
41:  PUSH1 0x40
42:  DUP2
43:  LT
44:  PUSH1 0x43
45:  JUMPI
46:  INVALID
47:  JUMPDEST
48:  PUSH1 0x20
49:  SWAP2
50:  DUP3
51:  DUP3
52:  DIV
53:  ADD
54:  SWAP2
55:  SWAP1
56:  MOD
57:  PUSH2 0x0100
58:  EXP
59:  DUP2
60:  SLOAD
61:  DUP2
62:  PUSH1 0xff
63:  MUL
64:  NOT
65:  AND
66:  SWAP1
67:  DUP4
68:  PUSH1 0xff
69:  AND
70:  MUL
71:  OR
72:  SWAP1
73:  SSTORE
74:  POP
75:  PUSH1 0x14
76:  PUSH1 0x00
77:  PUSH1 0x01
78:  PUSH1 0x40
79:  DUP2
80:  LT
81:  PUSH1 0x72
82:  JUMPI
83:  INVALID
84:  JUMPDEST
85:  PUSH1 0x20
86:  SWAP2
87:  DUP3
88:  DUP3
89:  DIV
90:  ADD
91:  SWAP2
92:  SWAP1
93:  MOD
94:  PUSH2 0x0100
95:  EXP
96:  DUP2
97:  SLOAD
98:  DUP2
99:  PUSH1 0xff
100: MUL
101: NOT
102: AND
103: SWAP1
104: DUP4
105: PUSH1 0xff
106: AND
107: MUL
108: OR
109: SWAP1
110: SSTORE
111: POP
112: PUSH1 0x1e
113: PUSH1 0x00
114: PUSH1 0x21
115: PUSH1 0x40
116: DUP2
117: LT
118: PUSH1 0xa1
119: JUMPI
120: INVALID
121: JUMPDEST
122: PUSH1 0x20
123: SWAP2
124: DUP3
125: DUP3
126: DIV
127: ADD
128: SWAP2
129: SWAP1
130: MOD
131: PUSH2 0x0100
132: EXP
133: DUP2
134: SLOAD
135: DUP2
136: PUSH1 0xff
137: MUL
138: NOT
139: AND
140: SWAP1
141: DUP4
142: PUSH1 0xff
143: AND
144: MUL
145: OR
146: SWAP1
147: SSTORE
148: POP
149: JUMP
```

检查storage：
```
truffle(development)> web3.eth.getStorageAt("0x9841007D2CFd5A913c8bc0f5D229F3dB0aFc395f", 0)
'0x140a'
truffle(development)> web3.eth.getStorageAt("0x9841007D2CFd5A913c8bc0f5D229F3dB0aFc395f", 1)
'0x1e00'
```

可以看到长度为64byte的数组，按照32byte切割来分配索引，前32byte分配索引0，后32byte分配索引1，对于固定长度的数据，solidity按照byte32切分来分配索引。

## 动态变长数组的存储
contract
```
pragma solidity 0.5.16;

contract Test3 {
    uint256 b;
    uint8[] a;
    uint256 c;

    constructor() public {
    }

    function test() public {
        a.length = 34;
        a[0] = 10;
        a[1] = 20;
        a[33] = 30;

        b = 100;
        c = 200;
    }
}
```
opcode
```
```
在往下看之前，可以先想一想solidity该怎么给a、b、c分配对应的索引。

我们来检查索引值：
```
truffle(development)> web3.eth.getStorageAt("0x46341de214BFA32014aD07226CCe4636F6B6C7b8", 0);
'0x64'
truffle(development)> web3.eth.getStorageAt("0x46341de214BFA32014aD07226CCe4636F6B6C7b8", 1);
'0x22'
truffle(development)> web3.eth.getStorageAt("0x46341de214BFA32014aD07226CCe4636F6B6C7b8", 2);
'0xc8'
```
索引为0的位置是b的值，2的位置为c的值，但1的位置是0x22，是a动态数据的长度数据，也就是说索引为1的位置存储了对象a的引用信息。

所以对于动态变长数据solidity分配一个索引来存储其引用信息，那么a的原始数据到底存储到了哪里？

[How is an array of structs accessed in getStorageAt and where is it stored?](https://ethereum.stackexchange.com/questions/41157/how-is-an-array-of-structs-accessed-in-getstorageat-and-where-is-it-stored)

可以看到，对象引用的索引的keccak256对应的索引是动态变长数组的数据位置:
```
contractAddress = "0x46341de214BFA32014aD07226CCe4636F6B6C7b8"
var m = 0;
var n = 1;
var p = 2;
var startSlot = web3.toBigNumber(web3.sha3(web3.padLeft("1", 64), { encoding: 'hex' }));
var slot_m = "0x" + startSlot.add(m).toString(16);
web3.eth.getStorageAt(contractAddress, slot_m, function (err, result) {
  console.log(result);
});

var slot_n = "0x" + startSlot.add(n).toString(16);
web3.eth.getStorageAt(contractAddress, slot_n, function (err, result) {
  console.log(result);
});

var slot_p = "0x" + startSlot.add(p).toString(16);
web3.eth.getStorageAt(contractAddress, slot_p, function (err, result) {
  console.log(result);
});
```
```
> contractAddress = "0x46341de214BFA32014aD07226CCe4636F6B6C7b8"
"0x46341de214BFA32014aD07226CCe4636F6B6C7b8"
> var m = 0;
undefined
> var n = 1;
undefined
> var p = 2;
undefined
> var startSlot = web3.toBigNumber(web3.sha3(web3.padLeft("1", 64), { encoding: 'hex' }));
undefined
> var slot_m = "0x" + startSlot.add(m).toString(16);
undefined
> web3.eth.getStorageAt(contractAddress, slot_m, function (err, result) {
......   console.log(result);
...... });
0x140a
undefined
>
> var slot_n = "0x" + startSlot.add(n).toString(16);
undefined
> web3.eth.getStorageAt(contractAddress, slot_n, function (err, result) {
......   console.log(result);
...... });
0x1e00
undefined
>
> var slot_p = "0x" + startSlot.add(p).toString(16);
undefined
> web3.eth.getStorageAt(contractAddress, slot_p, function (err, result) {
......   console.log(result);
...... });
0x0
undefined
>
```
分别取了数据索引0、1、2位置的数据，可以看到数据也是按照byte32切分放到不同的索引位置。

## mapping的存储
contract
```
pragma solidity 0.5.16;

contract Test4 {
    uint256 b;
    mapping(uint256 => uint256) a;
    uint256 c;

    constructor() public {
    }

    function test() public {
        a[0] = 10;
        a[1] = 20;

        b = 100;
        c = 200;
    }
}
```

我们直接验证数据的存储位置吧:
```
> contractAddress = "0x43d300AE7D7BAbfFd1101cd95621d3BC511a1Dab"
"0x43d300AE7D7BAbfFd1101cd95621d3BC511a1Dab"
>
> var mainSlot = web3.toBigNumber("0x00")
undefined
> var slot_main_0 = "0x" + mainSlot.toString(16);
undefined
> web3.eth.getStorageAt(contractAddress, slot_main_0, function (err, result) {
......   console.log(result);
...... });
0x64
undefined
>
> var slot_main_1 = "0x" + mainSlot.add(1).toString(16);
undefined
> web3.eth.getStorageAt(contractAddress, slot_main_1, function (err, result) {
......   console.log(result);
...... });
0x0
undefined
>
> var slot_main_2 = "0x" + mainSlot.add(2).toString(16);
undefined
> web3.eth.getStorageAt(contractAddress, slot_main_2, function (err, result) {
......   console.log(result);
...... });
0xc8
undefined
>
> var slot_mapping_0 = web3.sha3(web3.padLeft("0", 64) + web3.padLeft("1", 64), { encoding: 'hex' });
undefined
> web3.eth.getStorageAt(contractAddress, slot_mapping_0, function (err, result) {
......   console.log(result);
...... });
0x0a
undefined
>
> var slot_mapping_1 = web3.sha3(web3.padLeft("1", 64) + web3.padLeft("1", 64), { encoding: 'hex' });
undefined
> web3.eth.getStorageAt(contractAddress, slot_mapping_1, function (err, result) {
......   console.log(result);
...... });
0x14
undefined
>
```

可以看到，solidity照样按照变量顺序，依次给b、a、c分配索引0、1、2，对于b和c，索引对应的值位置直接为b和c的值，而a是引用对象，索引对应的值位置为引用信息，而mapping的引用信息为空。

mapping的值并不像变长数组那样根据对象引用的索引的keccak256来计算值索引，而是将mapping的key索引连接引用索引后计算keccak256得出该mapping的key的storage索引。


