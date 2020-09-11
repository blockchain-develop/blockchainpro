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
