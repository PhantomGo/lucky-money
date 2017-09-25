package service

import (
	"math/rand"
	"time"

	"lucky-money/domain"
)

const letterBytes = "0123456789abcdefghijkmnpqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

type Code struct {
	envelopes map[string]*domain.Envelope
}

func NewCode() *Code {
	return &Code{
		envelopes: make(map[string]*domain.Envelope),
	}
}

func (sv *Code) Verify(code string) (envelope *domain.Envelope, ok bool) {
	envelope, ok = sv.envelopes[code]
	ok = !envelope.IsExpired()
	return
}

func (sv *Code) GenerateTo(envelope *domain.Envelope) (code string) {
	code = randString(8)
	sv.envelopes[code] = envelope
	return
}

func (sv *Code) Clean() (result []*domain.Envelope) {
	result = make([]*domain.Envelope, 0)
	for code, envlope := range sv.envelopes {
		if envlope.IsExpired() {
			delete(sv.envelopes, code)
			result = append(result, envlope)
		}
	}
	return
}

func randString(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
