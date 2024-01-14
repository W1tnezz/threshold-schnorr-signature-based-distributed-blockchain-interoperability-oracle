const fs = require("fs");
const Oracle = artifacts.require("Oracle");

module.exports = async function () {

  let oracle = await Oracle.deployed();
  let fee = await oracle.totalFee();

  let message = "0xd693c74460d7f75b302c50eb6bc9c8463476af92419d8e7c9b14994c8993a91b";
  await oracle.validateTransaction(message, {
    value: fee,
  });

};
