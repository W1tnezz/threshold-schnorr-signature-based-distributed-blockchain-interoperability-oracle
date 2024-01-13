const fs = require("fs");
const Oracle = artifacts.require("Oracle");

module.exports = async function () {

  let oracle = await Oracle.deployed();
  let fee = await oracle.totalFee();

  let message = "0x5071f5eb1f1eeea184b2a472950cfc087a211e83cf473d5fff208df28b0db0d5";
  await oracle.validateTransaction(message, {
    value: fee,
  });

};
