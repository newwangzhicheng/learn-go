package pointer

import (
	"errors"
	"fmt"
)

type Stringer interface {
	string() string
}

type Bitcoin int

func (b Bitcoin) String() string {
	//格式化输出
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

//var 可以创建一个包的全局变量
var InsufficientFundsError = errors.New("wanted an error but didn't got one")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return InsufficientFundsError
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	//return (*w).balance
	//也是可以的，这里只是语法糖，可以省略
	return w.balance
}
