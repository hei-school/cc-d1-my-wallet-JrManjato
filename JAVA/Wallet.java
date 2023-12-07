import java.util.ArrayList;
import java.util.List;

public class Wallet {
    private double balance;
    private List<Transaction> transactions;

    public Wallet() {
        this.balance = 0;
        this.transactions = new ArrayList<>();
    }

    public void addMoney(double amount) {
        this.balance += amount;
        this.transactions.add(new Transaction("Add Money", amount, this.balance));
    }

    public void withdrawMoney(double amount) {
        if (amount > this.balance) {
            System.out.println("Insufficient balance to withdraw this amount.");
            return;
        }

        this.balance -= amount;
        this.transactions.add(new Transaction("Withdraw Money", amount, this.balance));
    }

    public void showBalance() {
        System.out.printf("Current balance: %.2f MGA%n", this.balance);
    }

    public void displayHistory() {
        System.out.println("Transaction history:");
        for (int i = 0; i < transactions.size(); i++) {
            Transaction transaction = transactions.get(i);
            String formattedDate = Transaction.formatDate(transaction.getDate());
            System.out.printf("%d. %s: %.2f MGA (Balance: %.2f MGA) - Date: %s%n",
                    i + 1, transaction.getType(), transaction.getAmount(), transaction.getBalance(), formattedDate);
        }
    }
}