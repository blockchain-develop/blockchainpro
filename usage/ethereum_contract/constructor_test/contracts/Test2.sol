pragma solidity 0.5.16;

contract Test2 {
    uint a;
    uint b;

    constructor(uint _a) public {
        a = _a;
        b = 20;
    }

    function test() public {
        b = 30;
    }
}
