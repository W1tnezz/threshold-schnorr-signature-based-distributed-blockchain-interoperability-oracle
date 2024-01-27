const fs = require("fs");
const Oracle = artifacts.require("Oracle");

module.exports = async function () {

  let oracle = await Oracle.deployed();
  let fee = await oracle.totalFee();

  let message = "0x73c31a7abba3368ef85c0c519f658edea5d8494c4e3fc6c61d023c0671d75cc7";
  await oracle.validateTransaction(message, {
    value: fee,
  });

};
