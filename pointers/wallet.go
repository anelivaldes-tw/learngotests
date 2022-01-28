package pointers

import "fmt"

type Wallet struct {
	balance int
}

/*In Go, when you call a function or a method the arguments are copied.
When calling func (w Wallet) Deposit(amount int) the w is a copy of
whatever we called the method from.(Are value receiver)*/

// You can find out what the address of that bit of memory with &myVal.

func (w Wallet) Deposit(amount int) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}
func (w Wallet) Balance() int {
	fmt.Printf("address of balance in Balance is %v \n", &w.balance)
	return w.balance
}
