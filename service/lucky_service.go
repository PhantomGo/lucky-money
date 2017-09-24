package service

import "math/rand"

type LuckyService struct {
}

func NewLuckyService() *LuckyService {
	return &LuckyService{}
}

func (sv *LuckyService) Draw(account *Account, el *Envelope) (result *OpendEnvelope, err error) {
	amount := money(int64(el.AvailableNumber), el.AvailableAmount)
	if result, err = el.Open(account.ID, amount); err != nil {
		return
	}
	account.Deposit(amount)
	return
}

func money(n, a int64) int64 {
	if n == 1 {
		return a
	}
	var money int64
	if a > n {
		money = rand.Int63n(a - n)
	} else {
		money = 1
	}
	if money <= 1 {
		money = 1
	}

	return money
}
