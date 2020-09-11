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
