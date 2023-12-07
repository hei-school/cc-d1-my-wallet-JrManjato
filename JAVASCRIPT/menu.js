function displayMenu() {
  console.log("\nMenu:");
  console.log("1. Add Money");
  console.log("2. Withdraw Money");
  console.log("3. Show Balance");
  console.log("4. Display History");
  console.log("5. Exit");
}

function startMenu(wallet, rl) {
  displayMenu();
  rl.question("Select an option (1-5): ", function (option) {
    switch (option) {
      case '1':
        rl.question("Enter the amount to add: ", function (amount) {
          wallet.addMoney(parseFloat(amount));
          startMenu(wallet, rl);
        });
        break;

      case '2':
        rl.question("Enter the amount to withdraw: ", function (amount) {
          wallet.withdrawMoney(parseFloat(amount));
          startMenu(wallet, rl);
        });
        break;

      case '3':
        wallet.showBalance();
        startMenu(wallet, rl);
        break;

      case '4':
        wallet.displayHistory();
        startMenu(wallet, rl);
        break;

      case '5':
        rl.close();
        break;

      default:
        console.log("Invalid option. Please enter a number from 1 to 5.");
        startMenu(wallet, rl);
        break;
    }
  });
}

module.exports = { displayMenu, startMenu };