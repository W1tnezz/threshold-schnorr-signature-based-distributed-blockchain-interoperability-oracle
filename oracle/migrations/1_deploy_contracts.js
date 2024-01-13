const BN256G1 = artifacts.require("BN256G1");
const Registry = artifacts.require("Registry");
const DKG = artifacts.require("DKG");
const Oracle = artifacts.require("Oracle");
module.exports = function (deployer) {
    deployer.deploy(Registry).then(
      function(){
        return deployer.deploy(BN256G1);
      }
    ).then(
      function(){
        return deployer.link(BN256G1, DKG);
      }
    ).then(
      function(){
        return deployer.deploy(DKG, Registry.address);
      }
    ).then(
      function(){
        return deployer.deploy(Oracle, Registry.address, DKG.address);
      }
    )
    
    // deployer.deploy(BN256G1);
    // deployer.deploy(Sakai, Registry.address);
    // deployer.link(BN256G1, IBSAS);
    // deployer.deploy(IBSAS, Registry.address);
};
