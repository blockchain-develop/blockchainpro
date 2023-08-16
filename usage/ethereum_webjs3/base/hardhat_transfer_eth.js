const { ethers } = require('hardhat');

async function interact() {
    const [signer1, signer2] = await ethers.getSigners();
    console.log(signer1, signer2);

    let balance1, balance2;
    balance1 = await ethers.provider.getBalance(signer1.address);
    balance2 = await ethers.provider.getBalance(signer2.address);
    console.log(balance1, balance2);
    
    const transaction = {
        to: signer2.address,
        value: ethers.parseEther('1.0'),
    };

    const tx = await signer1.sendTransaction(transaction);
    console.log('transaction', tx);

    const txReceipt = await ethers.provider.getTransactionReceipt(tx.hash);
    console.log('transaction receipt', txReceipt);

    balance1 = await ethers.provider.getBalance(signer1.address);
    balance2 = await ethers.provider.getBalance(signer2.address);
    console.log(balance1, balance2);

    const fee = await ethers.provider.getFeeData();
    console.log(fee);
}

(async () => {
    await interact();
})();