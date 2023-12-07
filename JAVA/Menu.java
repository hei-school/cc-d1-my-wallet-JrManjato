import java.util.Scanner;

public class Menu {
    public static void displayMenu() {
        System.out.println("\nMenu:");
        System.out.println("1. Add Money");
        System.out.println("2. Withdraw Money");
        System.out.println("3. Show Balance");
        System.out.println("4. Display History");
        System.out.println("5. Exit");
    }

    public static void startMenu(Wallet wallet) {
        Scanner scanner = new Scanner(System.in);

        while (true) {
            displayMenu();
            System.out.print("Select an option (1-5): ");
            String option = scanner.nextLine();

            switch (option) {
                case "1":
                    System.out.print("Enter the amount to add: ");
                    double amountToAdd = Double.parseDouble(scanner.nextLine());
                    wallet.addMoney(amountToAdd);
                    break;
                case "2":
                    System.out.print("Enter the amount to withdraw: ");
                    double amountToWithdraw = Double.parseDouble(scanner.nextLine());
                    wallet.withdrawMoney(amountToWithdraw);
                    break;
                case "3":
                    wallet.showBalance();
                    break;
                case "4":
                    wallet.displayHistory();
                    break;
                case "5":
                    scanner.close();
                    System.exit(0);
                default:
                    System.out.println("Invalid option. Please enter a number from 1 to 5.");
            }
        }
    }
}
