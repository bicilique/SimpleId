var SayHello = artifacts.require("SayHello"); // Remove the file path

module.exports = function(deployer) {
   deployer.deploy(SayHello);
};

