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
