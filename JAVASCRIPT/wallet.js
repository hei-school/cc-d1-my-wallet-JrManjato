const { formatDate } = require('./utils');

class Wallet {
  constructor() {
    this.balance = 0;
    this.actions = [];
  }

  addMoney(amount) {
    this.balance += amount;
    this.actions.push({ type: 'Add Money', amount, balance: this.balance, date: new Date() });
  }

  withdrawMoney(amount) {
    if (amount > this.balance) {
      console.log("Insufficient balance to withdraw this amount.");
      return;
    }

    this.balance -= amount;
    this.actions.push({ type: 'Withdraw Money', amount, balance: this.balance, date: new Date() });
  }

  showBalance() {
    console.log(`Current balance: ${this.balance} MGA`);
  }

  displayHistory() {
    console.log("Action history:");
    this.actions.forEach((action, index) => {
      const formattedDate = formatDate(action.date);
      console.log(`${index + 1}. ${action.type}: ${action.amount} MGA (Balance: ${action.balance} MGA) - Date: ${formattedDate}`);
    });
  }
}

module.exports = Wallet;