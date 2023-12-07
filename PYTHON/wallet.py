from datetime import datetime
from utils import format_date

class Wallet:
    def __init__(self):
        self.balance = 0
        self.actions = []

    def add_money(self, amount):
        self.balance += amount
        self.actions.append({'type': 'Add Money', 'amount': amount, 'balance': self.balance, 'date': datetime.now()})

    def withdraw_money(self, amount):
        if amount > self.balance:
            print("Insufficient balance to withdraw this amount.")
            return

        self.balance -= amount
        self.actions.append({'type': 'Withdraw Money', 'amount': amount, 'balance': self.balance, 'date': datetime.now()})

    def show_balance(self):
        print(f"Current balance: {self.balance} MGA")

    def display_history(self):
        print("Action history:")
        for index, action in enumerate(self.actions, start=1):
            formatted_date = format_date(action['date'])
            print(f"{index}. {action['type']}: {action['amount']} MGA (Balance: {action['balance']} MGA) - Date: {formatted_date}")