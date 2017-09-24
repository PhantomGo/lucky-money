package service

import (
	"errors"
	"time"
)

type Envelope struct {
	ID              int64
	CreatorID       int64
	TotalAmount     int64
	AvailableAmount int64
	TotalNumber     int
	AvailableNumber uint
	Ended           bool
	Created         time.Time
	OpendList       map[int64]*OpendEnvelope
}

type OpendEnvelope struct {
	EnvelopeID int64 `json:"envelop_id"`
	AccountID  int64
	Amount     int64 `json:"amount"`
}

func NewEnvelop(id, aid, amount int64, number int) *Envelope {
	return &Envelope{
		ID:              id,
		CreatorID:       aid,
		TotalAmount:     amount,
		AvailableAmount: amount,
		TotalNumber:     number,
		AvailableNumber: uint(number),
		Ended:           false,
		Created:         time.Now(),
		OpendList:       make(map[int64]*OpendEnvelope),
	}
}

func (el *Envelope) Open(aid, amount int64) (result *OpendEnvelope, err error) {
	if amount > el.AvailableAmount {
		err = errors.New("system error")
		return
	}
	if el.IsExpired() {
		err = errors.New("game is over")
		return
	}
	if _, exist := el.OpendList[aid]; exist {
		err = errors.New("envelope has been opened")
		return
	}

	el.AvailableNumber--
	el.AvailableAmount -= amount
	result = &OpendEnvelope{
		EnvelopeID: el.ID,
		AccountID:  aid,
		Amount:     amount,
	}
	el.OpendList[aid] = result
	return
}

func (el *Envelope) IsExpired() bool {
	if el.isOver() || el.Created.AddDate(0, 0, 1).Before(time.Now()) {
		el.Ended = true
	}
	return el.Ended
}

func (el *Envelope) isOver() bool {
	return el.AvailableNumber == 0
}
