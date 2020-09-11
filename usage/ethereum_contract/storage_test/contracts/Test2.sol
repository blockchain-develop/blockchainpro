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
