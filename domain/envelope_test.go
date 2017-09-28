package domain

import "testing"

func TestEnvelope(t *testing.T) {
	e := NewEnvelope(1, 1, 3, 3)
	_, err := e.Open(1, 1)
	_, err = e.Open(1, 1)
	if err == nil {
		t.Errorf("account illegal")
	}
	print(err.Error())
	e.Open(2, 1)
	e.Open(3, 1)
	if !e.IsExpired() {
		t.Errorf("status error")
	}
}
