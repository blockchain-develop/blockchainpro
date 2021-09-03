pragma solidity 0.5.16;

contract Test1 {
    uint a;
    constructor() public {
    }
    function test(address toAddress, uint256 amount) public {
        address(uint160(toAddress)).transfer(amount);
    }
}
