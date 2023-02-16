package main

import "fmt"

type wallet struct{
	balance int
}


func newWallet() *wallet{
	return &wallet{
		balance: 0,
	}
}

func (w *wallet) creditBalance(amount int){
	w.balance += amount
	fmt.Println("Wallet balance added successfully")
}

func (w *wallet) debitBalance(amount int) error{
	if w.balance < amount{
		return fmt.Errorf("Insufficient balance")
	}
	fmt.Println("Wallet balance successfully debited")
	w.balance -= amount
	return nil
}