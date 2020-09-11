# EVM Storage Lab

## Test1

合约地址: 0x452CE6c20bAb1F3766f74ad159ec85Fe1BA8A3dC

Test1.deployed().then(function(instance){return instance.test();});

交易hash: 0xcd8b3864b3910a74810fd1acc4d6f2fb0724f24569cf3661ee9d6471f949bc13

web3.eth.getStorageAt("0x452CE6c20bAb1F3766f74ad159ec85Fe1BA8A3dC", 0)

## Test2

合约地址: 0x9841007D2CFd5A913c8bc0f5D229F3dB0aFc395f

Test2.deployed().then(function(instance){return instance.test();});

交易hash: 0x7490b8353f8a16abbe5be25044caf56259f3a5b0cd04576de99c91f80e7850b5

web3.eth.getStorageAt("0x9841007D2CFd5A913c8bc0f5D229F3dB0aFc395f", 0)

## Test3

合约地址: 0x46341de214BFA32014aD07226CCe4636F6B6C7b8

Test3.deployed().then(function(instance){return instance.test();});

交易hash: 0x5f57e4e07f59a87ef7189e45d3f3b17dbb4181a27575e619a45aca5909dd2c16

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

## Test4

合约地址: 0x43d300AE7D7BAbfFd1101cd95621d3BC511a1Dab

Test4.deployed().then(function(instance){return instance.test();});

交易hash: 0x1f8b5b1d30a4206d3fef06d18b59f67f7f63d8bf932b4981396121dc9d95558b


contractAddress = "0x43d300AE7D7BAbfFd1101cd95621d3BC511a1Dab"

var mainSlot = web3.toBigNumber("0x00")
var slot_main_0 = "0x" + mainSlot.toString(16);
web3.eth.getStorageAt(contractAddress, slot_main_0, function (err, result) {
  console.log(result);
});

var slot_main_1 = "0x" + mainSlot.add(1).toString(16);
web3.eth.getStorageAt(contractAddress, slot_main_1, function (err, result) {
  console.log(result);
});

var slot_main_2 = "0x" + mainSlot.add(2).toString(16);
web3.eth.getStorageAt(contractAddress, slot_main_2, function (err, result) {
  console.log(result);
});

var slot_mapping_0 = web3.sha3(web3.padLeft("0", 64) + web3.padLeft("1", 64), { encoding: 'hex' });
web3.eth.getStorageAt(contractAddress, slot_mapping_0, function (err, result) {
  console.log(result);
});

var slot_mapping_1 = web3.sha3(web3.padLeft("1", 64) + web3.padLeft("1", 64), { encoding: 'hex' });
web3.eth.getStorageAt(contractAddress, slot_mapping_1, function (err, result) {
  console.log(result);
});

