package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Transaction struct {
	Type    string
	Amount  float64
	Balance float64
	Date    time.Time
}

type Wallet struct {
	Balance      float64
	Transactions []Transaction
}

func NewWallet() *Wallet {
	return &Wallet{Balance: 0, Transactions: []Transaction{}}
}

func (w *Wallet) AddMoney(amount float64) {
	w.Balance += amount
	w.Transactions = append(w.Transactions, Transaction{Type: "Add Money", Amount: amount, Balance: w.Balance, Date: time.Now()})
}

func (w *Wallet) WithdrawMoney(amount float64) {
	if amount > w.Balance {
		fmt.Println("Insufficient balance to withdraw this amount.")
		return
	}

	w.Balance -= amount
	w.Transactions = append(w.Transactions, Transaction{Type: "Withdraw Money", Amount: amount, Balance: w.Balance, Date: time.Now()})
}

func (w *Wallet) ShowBalance() {
	fmt.Printf("Current balance: %.2f MGA\n", w.Balance)
}

func (w *Wallet) DisplayHistory() {
	fmt.Println("Transaction history:")
	for i, transaction := range w.Transactions {
		formattedDate := transaction.Date.Format("02/01/2006 15:04:05")
		fmt.Printf("%d. %s: %.2f MGA (Balance: %.2f MGA) - Date: %s\n",
			i+1, transaction.Type, transaction.Amount, transaction.Balance, formattedDate)
	}
}

func DisplayMenu() {
	fmt.Println("\nMenu:")
	fmt.Println("1. Add Money")
	fmt.Println("2. Withdraw Money")
	fmt.Println("3. Show Balance")
	fmt.Println("4. Display History")
	fmt.Println("5. Exit")
}

func StartMenu(wallet *Wallet) {
	reader := bufio.NewReader(os.Stdin)

	for {
		DisplayMenu()
		fmt.Print("Select an option (1-5): ")
		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			fmt.Print("Enter the amount to add: ")
			amountStr, _ := reader.ReadString('\n')
			amount, _ := strconv.ParseFloat(strings.TrimSpace(amountStr), 64)
			wallet.AddMoney(amount)
		case "2":
			fmt.Print("Enter the amount to withdraw: ")
			amountStr, _ := reader.ReadString('\n')
			amount, _ := strconv.ParseFloat(strings.TrimSpace(amountStr), 64)
			wallet.WithdrawMoney(amount)
		case "3":
			wallet.ShowBalance()
		case "4":
			wallet.DisplayHistory()
		case "5":
			os.Exit(0)
		default:
			fmt.Println("Invalid option. Please enter a number from 1 to 5.")
		}
	}
}

func main() {
	wallet := NewWallet()
	StartMenu(wallet)
}