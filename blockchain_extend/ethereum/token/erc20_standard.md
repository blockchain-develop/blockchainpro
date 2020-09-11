# ERC20 Token Standard

## method

1. name
```
function name() public view returns (string)
```

返回token的名称。

2. symbol
```
function symbol() public view returns (string)
```

返回token的单位。

3. decimals
```
function decimals() public view returns (uint8)
```

返回token的精度。

4. totalSupply
```
function totalSupply() public view returns (uint256)
```

返回token的总供应量。

5. balanceOf
```
function balanceOf(address _owner) public view returns (uint256 balance)
```

获取一个账户的余额。

6. transfer
```
function transfer(address _to, uint256 _value) public returns (bool success)
```

从这个函数的调用者账户转账到指定的to账户。

7. transferFrom
```
function transferFrom(address _from, address _to, uint256 _value) public returns (bool success)
```

从指定的from账户转账到指定的to账户，这要求from账户有approve资产给这个函数的调用者。

8. approve
```
function approve(address _spender, uint256 _value) public returns (bool success)
```

函数调用者账户授权指定的spender账户可以转账自己的资产。

9. allowance
```
function allowance(address _owner, address _spender) public view returns (uint256 remaining)
```

获取指定的owner账户授权给账户spender转账资产余额。

## event

1. Transfer
```
event Transfer(address indexed _from, address indexed _to, uint256 _value)
```

2. Approval
```
event Approval(address indexed _owner, address indexed _spender, uint256 _value)
```

## example

sonsensys的ERC20 Token

```
/*
Implements EIP20 token standard: https://github.com/ethereum/EIPs/blob/master/EIPS/eip-20.md
.*/


pragma solidity ^0.4.21;

import "./EIP20Interface.sol";


contract EIP20 is EIP20Interface {

    uint256 constant private MAX_UINT256 = 2**256 - 1;
    mapping (address => uint256) public balances;
    mapping (address => mapping (address => uint256)) public allowed;
    /*
    NOTE:
    The following variables are OPTIONAL vanities. One does not have to include them.
    They allow one to customise the token contract & in no way influences the core functionality.
    Some wallets/interfaces might not even bother to look at this information.
    */
    string public name;                   //fancy name: eg Simon Bucks
    uint8 public decimals;                //How many decimals to show.
    string public symbol;                 //An identifier: eg SBX
    uint256 totalSupply;

    function EIP20(
        uint256 _initialAmount,
        string _tokenName,
        uint8 _decimalUnits,
        string _tokenSymbol
    ) public {
        balances[msg.sender] = _initialAmount;               // Give the creator all initial tokens
        totalSupply = _initialAmount;                        // Update total supply
        name = _tokenName;                                   // Set the name for display purposes
        decimals = _decimalUnits;                            // Amount of decimals for display purposes
        symbol = _tokenSymbol;                               // Set the symbol for display purposes
    }

    function transfer(address _to, uint256 _value) public returns (bool success) {
        require(balances[msg.sender] >= _value);
        balances[msg.sender] -= _value;
        balances[_to] += _value;
        emit Transfer(msg.sender, _to, _value); //solhint-disable-line indent, no-unused-vars
        return true;
    }

    function transferFrom(address _from, address _to, uint256 _value) public returns (bool success) {
        uint256 allowance = allowed[_from][msg.sender];
        require(balances[_from] >= _value && allowance >= _value);
        balances[_to] += _value;
        balances[_from] -= _value;
        if (allowance < MAX_UINT256) {
            allowed[_from][msg.sender] -= _value;
        }
        emit Transfer(_from, _to, _value); //solhint-disable-line indent, no-unused-vars
        return true;
    }

    function balanceOf(address _owner) public view returns (uint256 balance) {
        return balances[_owner];
    }

    function approve(address _spender, uint256 _value) public returns (bool success) {
        allowed[msg.sender][_spender] = _value;
        emit Approval(msg.sender, _spender, _value); //solhint-disable-line indent, no-unused-vars
        return true;
    }

    function allowance(address _owner, address _spender) public view returns (uint256 remaining) {
        return allowed[_owner][_spender];
    }
}
```

## reference
[EIP-20: ERC-20 Token Standard](https://eips.ethereum.org/EIPS/eip-20)
[ConsenSys implementation of ERC20](https://github.com/ConsenSys/Tokens/blob/fdf687c69d998266a95f15216b1955a4965a0a6d/contracts/eip20/EIP20.sol)
[以太坊代币标准ERC20与ERC223的区别](http://www.jouypub.com/2018/caf472b0ea3025ea6ed370a23b23eea3/)
