const usdtABI = require('./../abis/usdt.json');
const { ethers, BigNumber } = require('hardhat');

async function interact() {
    await hre.network.provider.request({
        method: "hardhat_reset",
        params: [{
            forking: {
                jsonRpcUrl:"https://rpc.ankr.com/eth",
                blockNumber: 17807829
            }
        }]
    })

    // user账户
    const [user1, user2] = await ethers.getSigners();
    console.log(user1, user2);

    // 添加usdt的owner账户
    const usdtOwnerAddress = "0xC6CDE7C39eB2f0F0095F41570af89eFC2C1Ea828";
    await hre.network.provider.request({
        method: "hardhat_impersonateAccount",
        params: [usdtOwnerAddress],
    });
    const usdtOwner = await ethers.getSigner(usdtOwnerAddress);

    // 给usdt owner账户准备ETH
    const transaction = {
        to: usdtOwner.address,
        value: ethers.parseEther('1.0'),
    };
    const tx = await user1.sendTransaction(transaction);
    console.log('sent eth transaction', tx);

    //
    const usdtAddress = "0xdAC17F958D2ee523a2206206994597C13D831ec7";
    //const usdtABI1 = new ethers.Interface(usdtABI);
    const usdt = new ethers.Contract(usdtAddress, usdtABI, usdtOwner);
    let owner = await usdt.owner();
    let symbol = await usdt.symbol();
    let name = await usdt.name();
    let totalSupply = await usdt.totalSupply();
    let decimal = await usdt.decimals();
    console.log(`name: ${name}, symbol: ${symbol}, totalSupply: ${totalSupply}, decimal: ${decimal}, owner: ${owner}`);

    //
    let balance = await usdt.balanceOf(usdtOwner.address);
    console.log(`usdt balance of address ${usdtOwner.address} is ${balance}`);

    const amount = 1000000000;
    const issuetx = await usdt.issue(amount);
    console.log(`issue tx hash: ${issuetx.hash}`);

    balance = await usdt.balanceOf(usdtOwner.address);
    console.log(`usdt balance of address ${usdtOwner.address} is ${balance}`);

    balance = await usdt.balanceOf(user1.address);
    console.log(`usdt balance of address ${user1.address} is ${balance}`);

    const transferTx = await usdt.transfer(user1.address, amount);
    console.log(`transfer tx hash: ${transferTx.hash}`);

    balance = await usdt.balanceOf(usdtOwner.address);
    console.log(`usdt balance of address ${usdtOwner.address} is ${balance}`);

    balance = await usdt.balanceOf(user1.address);
    console.log(`usdt balance of address ${user1.address} is ${balance}`);

    console.log(`${amount} usdt is sent from ${usdtOwner.address} to ${user1.address}.`);
}

(async () => {
    await interact();
})();