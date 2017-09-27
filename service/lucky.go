package service

import (
	"lucky-money/domain"
	"math/rand"
)

type Lucky struct {
}

func NewLucky() *Lucky {
	return &Lucky{}
}

func (sv *Lucky) Draw(account *domain.Account, el *domain.Envelope) (result *domain.OpendEnvelope, err error) {
	amount := money(int64(el.AvailableNumber), el.AvailableAmount)
	if result, err = el.Open(account.ID, amount); err != nil {
		return
	}
	account.Deposit(amount)
	return
}

var rn = rand.New(src)

func money(n, a int64) int64 {
	if n == 1 {
		return a
	}

	var money int64
	if a > n {
		money = rn.Int63n(a - n)
	} else {
		money = 1
	}
	if money <= 1 {
		money = 1
	}

	return money
}
