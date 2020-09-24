# ERC721 Token Standard

## method

1. balanceOf
```
function balanceOf(address _owner) external view returns (uint256);
```

获取指定账户的余额

2. ownerOf
```
function ownerOf(uint256 _tokenId) external view returns (address);
```

获取指定token的发行人

3. safeTransferFrom
```
function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes data)
```

从指定账户from转账到to账户

4. safeTransferFrom
```
function safeTransferFrom(address _from, address _to, uint256 _tokenId) external payable;
```

从指定账户from转账到to账户

5. transferFrom
```
function transferFrom(address _from, address _to, uint256 _tokenId) external payable;
```

6. approve
```
function approve(address _approved, uint256 _tokenId) external payable;
```

7. setApprovalForAll
```
function setApprovalForAll(address _operator, bool _approved) external;
```

8. getApproved
```
function getApproved(uint256 _tokenId) external view returns (address);
```

9. isApprovedForAll
```
function isApprovedForAll(address _owner, address _operator) external view returns (bool);
```

10. supportsInterface
```
function supportsInterface(bytes4 interfaceID) external view returns (bool);
```

## example
```
```

## reference
[EIP-165: ERC-165 Standard Interface Detection](https://eips.ethereum.org/EIPS/eip-165)
[EIP-721: ERC-721 Non-Fungible Token Standard](https://eips.ethereum.org/EIPS/eip-721)
[ERC 721](https://docs.openzeppelin.com/contracts/3.x/api/token/erc721)
[0xcert ERC-721 Token](https://github.com/0xcert/ethereum-erc721)
[以太坊标准ERC-721]https://www.jianshu.com/p/82714a8aae40