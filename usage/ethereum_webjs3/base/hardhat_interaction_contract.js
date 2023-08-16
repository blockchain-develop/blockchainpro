const { ethers, BigNumber } = require('hardhat');

async function interact() {
    const [signer1, signer2] = await ethers.getSigners();
    console.log(signer1, signer2);

    const myTokenFactory = await ethers.getContractFactory("MyToken");
    const myToken = await myTokenFactory.deploy();

    const deployTxHash = myToken.deploymentTransaction().hash;
    console.log(`my token deploy tx hash: ${deployTxHash}`);

    const deployTxReceipt = await ethers.provider.getTransactionReceipt(deployTxHash);
    console.log("my token deploy tx receipt", deployTxReceipt);

    myTokenAddress = await myToken.getAddress();
    console.log("my token address: ", myTokenAddress);

    let balance = await myToken.balanceOf(signer1.address);
    console.log(`my token balance of address ${signer1.address} is ${ethers.formatEther(balance)}`);

    const mintTx = await myToken.mint(signer1.address, ethers.parseEther("10.0"));
    console.log(`mint tx hash: ${mintTx.hash}`);

    balance = await myToken.balanceOf(signer1.address);
    console.log(`my token balance of address ${signer1.address} is ${ethers.formatEther(balance)}`);

    //
    const amount = "1.0"
    const transferTx = await myToken.transfer(signer2.address, ethers.parseEther(amount));
    await transferTx.wait();
    console.log(`${amount} my token is sent from sender ${signer1.address} to ${signer2.address}. tx hash: ${transferTx.hash}`);

    balance = await myToken.balanceOf(signer1.address);
    console.log(`my token balance of address ${signer1.address} is ${ethers.formatEther(balance)}`);

    balance = await myToken.balanceOf(signer2.address);
    console.log(`my token balance of address ${signer2.address} is ${ethers.formatEther(balance)}`);
}

(async () => {
    await interact();
})();