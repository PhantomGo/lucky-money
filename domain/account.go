package domain

type Account struct {
	ID      int64
	Balance int64
}

func NewAccount(id, balance int64) *Account {
	return &Account{
		ID:      id,
		Balance: balance,
	}
}

func (account *Account) Deposit(amount int64) int64 {
	account.Balance += amount
	return account.Balance
}
