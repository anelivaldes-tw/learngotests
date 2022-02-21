package pointers

import (
	"errors"
	"fmt"
)

//The var keyword allows us to define values global to the package.
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

/*In Go, when you call a function or a method the arguments are copied.
When calling func (w Wallet) Deposit(amount int) the w is a copy of
whatever we called the method from.(Are value receiver)*/

// You can find out what the address of that bit of memory with &myVal.

// The difference is the receiver type is *Wallet rather
// than Wallet which you can read as "a pointer to a wallet"

// Now the memory address are the same

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}
func (w *Wallet) Balance() Bitcoin {
	fmt.Printf("address of balance in Balance is %v \n", &w.balance)
	return (*w).balance // It is not required to change to (*w), but better for consistency
}

// Let's implement Stringer on Bitcoin
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
