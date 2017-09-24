package service

type Account struct {
	ID      int64
	Balance int64
}

func (account *Account) Deposit(amount int64) int64 {
	account.Balance += amount
	return account.Balance
}
