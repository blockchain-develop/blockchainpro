const Test1 = artifacts.require("Test1");

module.exports = function (deployer) {
  deployer.deploy(Test1, 10);
};
