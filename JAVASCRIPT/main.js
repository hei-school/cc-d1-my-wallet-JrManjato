const readline = require('readline');
const Wallet = require('./wallet');
const { startMenu } = require('./menu');

const wallet = new Wallet();
const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout
});

startMenu(wallet, rl);