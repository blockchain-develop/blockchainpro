const {Web3} = require('web3');
const nodeUrl = 'https://mainnet.infura.io/v3/dbdf5388d0934810962da46d5d7a8f23'
const httpProvider = new Web3.providers.HttpProvider(nodeUrl);
const web3 = new Web3(httpProvider);

async function interact() {
    const accounts = await web3.eth.getAccounts();
    console.log(accounts);

    let balance1, balance2;
    balance1 = await web3.eth.getBalance(accounts[0]);
    balance2 = await web3.eth.getBalance(accounts[1]);
    console.log(balance1, balance2);

    const transaction = {
        from: accounts[0],
        to: accounts[1],
        value: web3.utils.toWei('1', 'ether'),
    };

    const transactionHash = await web3.eth.sendTransaction(transaction);
    console.log('transaction', transactionHash);

    balance1 = await web3.eth.getBalance(accounts[0]);
    balance2 = await web3.eth.getBalance(accounts[1]);
    console.log(balance1, balance2);

    const gasPrice = await web3.eth.getGasPrice();
    console.log(gasPrice);
}

(async () => {
    await interact();
})();