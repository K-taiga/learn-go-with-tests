package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

// typeの中にtypeを記載できる
type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

// グローバル変数
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// レシーバーのメモリ上のポインターを指定する じゃないとCOPYを取得する
func (w *Wallet) Deposit(amount Bitcoin) {
	// fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

func (w Wallet) Balance() Bitcoin {
	// (*w).balanceのように取得する際はポインタからの参照が逆方向に必要だがそれは自動でできるのでこれでok
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount

	return nil
}
