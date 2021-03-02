// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 261.
//!+

// Package bank provides a concurrency-safe bank with one account.
package bank

type withdrawal struct {
	amount              int
	isBalanceSufficient chan bool
}

var deposits = make(chan int)           // send amount to deposit
var balances = make(chan int)           // receive balance
var withdrawals = make(chan withdrawal) // request withdrawals

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

// Withdraw withdraws amount from balance. It returns false if it failed, due to
// insufficient balance. Otherwise it returns true.
func Withdraw(amount int) bool {
	isBalanceSufficient := make(chan bool)
	withdrawals <- withdrawal{amount, isBalanceSufficient}
	return <-isBalanceSufficient
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case w := <-withdrawals:
			if w.amount > balance {
				w.isBalanceSufficient <- false
				continue
			}
			balance -= w.amount
			w.isBalanceSufficient <- true
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

//!-
