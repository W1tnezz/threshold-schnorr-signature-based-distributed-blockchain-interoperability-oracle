const BN256G1 = artifacts.require("BN256G1");
const Registry = artifacts.require("Registry");
const Sakai = artifacts.require("Sakai");
const IBSAS = artifacts.require("IBSAS");
module.exports = function (deployer) {
    deployer.deploy(Registry).then(
      function(){
        return deployer.deploy(BN256G1);
      }
    ).then(
      function(){
        return deployer.link(BN256G1, Sakai);
      }
    ).then(
      function(){
        return deployer.deploy(Sakai, Registry.address);
      }
    ).then(
      function(){
        return deployer.link(BN256G1, IBSAS);
      }
    ).then(
      function(){
        return deployer.deploy(IBSAS, Registry.address);
      }
    )
    
    // deployer.deploy(BN256G1);
    // deployer.deploy(Sakai, Registry.address);
    // deployer.link(BN256G1, IBSAS);
    // deployer.deploy(IBSAS, Registry.address);
};
