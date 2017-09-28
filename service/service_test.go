package service

import "testing"

func TestService(t *testing.T) {
	srv := NewService()
	srv.Account(1)
	srv.Account(2)
	srv.Account(3)

	var (
		code string
		err  error
	)

	if code, err = srv.Fill(1, 100, 3); err == nil {
		if r, err1 := srv.Open(1, code); err1 == nil {
			print(r.Amount)
			b, _ := srv.Balance(1)
			if b != r.Amount {
				t.Errorf("banlance error %d : %d", b, r.Amount)
			}
		} else {
			t.Errorf(err.Error())
		}
		if r, err1 := srv.Open(2, code); err1 == nil {
			print(r.Amount)
			b, _ := srv.Balance(2)
			if b != r.Amount {
				t.Errorf("banlance error %d : %d", b, r.Amount)
			}
		} else {
			t.Errorf(err.Error())
		}
		if r, err1 := srv.Open(3, code); err1 == nil {
			print(r.Amount)
			b, _ := srv.Balance(3)
			if b != r.Amount {
				t.Errorf("banlance error %d : %d", b, r.Amount)
			}
		} else {
			t.Errorf(err.Error())
		}
	} else {
		t.Errorf(err.Error())
	}

	hs, _ := srv.Histories(1)
	for _, h := range hs {
		print(" Account ", h.AccountID, " Balance ", h.Amount)
	}

	srv.ClearExpired()
	if r, err1 := srv.Open(1, code); err1 != nil {
		print(r)
	} else {
		t.Errorf("envelope is now expired")
	}
}
