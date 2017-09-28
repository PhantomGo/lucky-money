package service

import (
	"lucky-money/domain"
	"testing"
)

func TestLuckyDraw(t *testing.T) {
	lucky := NewLucky()
	//account init
	a1 := domain.NewAccount(1, 0)
	a2 := domain.NewAccount(2, 0)
	a3 := domain.NewAccount(3, 0)
	a4 := domain.NewAccount(3, 0)
	//envelope init
	e1 := domain.NewEnvelope(1, 1, 100, 3)
	e2 := domain.NewEnvelope(1, 1, 100, 3)

	var amount1, amount2, amount3 int64
	t.Run("total", func(t *testing.T) {
		var total int64
		if r, err := lucky.Draw(a1, e1); err != nil {
			t.Errorf(err.Error())
		} else {
			amount1 = r.Amount
			total += r.Amount
		}
		if r, err := lucky.Draw(a2, e1); err != nil {
			t.Errorf(err.Error())
		} else {
			amount2 = r.Amount
			total += r.Amount
		}
		if r, err := lucky.Draw(a3, e1); err != nil {
			t.Errorf(err.Error())
		} else {
			amount3 = r.Amount
			total += r.Amount
		}
		if total != 100 {
			t.Errorf("actual total is %d", total)
		}
	})

	t.Run("rand", func(t *testing.T) {
		var am1, am2, am3 int64
		if r, err := lucky.Draw(a1, e2); err != nil {
			t.Errorf(err.Error())
		} else {
			am1 = r.Amount
		}
		if r, err := lucky.Draw(a2, e2); err != nil {
			t.Errorf(err.Error())
		} else {
			am2 = r.Amount
		}
		if r, err := lucky.Draw(a3, e2); err != nil {
			t.Errorf(err.Error())
		} else {
			am3 = r.Amount
		}
		if am1 == amount1 && am2 == amount2 && am3 == amount3 {
			t.Errorf("rand fail")
		}
	})

	t.Run("over", func(t *testing.T) {
		if _, err := lucky.Draw(a4, e1); err == nil {
			t.Errorf("envelope is not over")
		} else {
			if !e1.IsExpired() {
				t.Errorf("envelope is not over")
			}
		}
	})
}
