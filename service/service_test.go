package service

import (
	"fmt"
	"testing"
)

func TestService(t *testing.T) {
	srv := NewService()
	srv.Account(1)
	srv.Account(2)
	srv.Account(3)
	if code, err := srv.Fill(1, 100, 3); err == nil {
		if r, err1 := srv.Open(1, code); err1 == nil {
			fmt.Println(r.Amount)
		} else {
			print(err.Error())
		}
		if r1, err2 := srv.Open(2, code); err2 == nil {
			fmt.Println(r1.Amount)
		} else {
			print(err.Error())
		}
		if r2, err3 := srv.Open(3, code); err3 == nil {
			fmt.Println(r2.Amount)
		} else {
			print(err.Error())
		}
	} else {
		print(err.Error())
	}
}
