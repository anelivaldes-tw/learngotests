package pointers

type Wallet struct {
	balance int
}

/*In Go, when you call a function or a method the arguments are copied.
When calling func (w Wallet) Deposit(amount int) the w is a copy of
whatever we called the method from.(Are value receiver)*/

func (w Wallet) Deposit(amount int) {
	w.balance += amount
}
func (w Wallet) Balance() int {
	return w.balance
}
