const ERC20Template = artifacts.require("ERC20Template");

module.exports = function (deployer) {
  deployer.deploy(ERC20Template);
};
