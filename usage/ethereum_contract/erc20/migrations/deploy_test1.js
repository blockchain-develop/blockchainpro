const test1 = artifacts.require("test1");

module.exports = function (deployer) {
    deployer.deploy(test1);
};