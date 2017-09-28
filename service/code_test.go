package service

import (
	"lucky-money/domain"
	"testing"
)

func TestGenerateCode(t *testing.T) {
	code := NewCode()
	codes := make(map[string]bool)
	for index := 0; index < 100000; index++ {
		c := code.GenerateTo(domain.NewEnvelope(1, 1, 1, 1))
		if _, ok := codes[c]; !ok {
			codes[c] = true
		} else {
			t.Errorf("code duplicated")
		}
	}
}
