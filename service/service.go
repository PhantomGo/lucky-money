package service

import "errors"

type Service struct {
	accounts  map[int64]*Account
	histories map[int64][]*OpendEnvelope
	maxID     int64

	code  *CodeService
	lucky *LuckyService
}

func NewService() *Service {
	sv := &Service{
		code:      NewCodeService(),
		lucky:     NewLuckyService(),
		accounts:  make(map[int64]*Account),
		histories: make(map[int64][]*OpendEnvelope),
		maxID:     0,
	}
	return sv
}

func (sv *Service) Balance(id int64) (result int64, err error) {
	if _, ok := sv.accounts[id]; !ok {
		return 0, errors.New("account does not exist")
	}
	result = sv.accounts[id].Balance
	return
}

func (sv *Service) Histories(id int64) (result []*OpendEnvelope, err error) {
	var ok bool
	if result, ok = sv.histories[id]; !ok {
		result = make([]*OpendEnvelope, 0)
	}
	return
}

func (sv *Service) Open(id int64, code string) (result *OpendEnvelope, err error) {
	var (
		ok       bool
		envelope *Envelope
	)
	if _, ok = sv.accounts[id]; !ok {
		err = errors.New("account does not exist")
		return
	}
	if envelope, ok = sv.code.Verify(code); !ok {
		err = errors.New("code does not exist")
		return
	}
	result, err = sv.lucky.Draw(sv.accounts[id], envelope)
	if _, ok := sv.histories[id]; !ok {
		sv.histories[id] = make([]*OpendEnvelope, 0)
	}
	sv.histories[id] = append(sv.histories[id], result)
	return
}

func (sv *Service) Fill(id, amount int64, number int) (result string, err error) {
	evID := sv.nextEnvelopeID()
	envelope := NewEnvelop(evID, id, amount, number)
	result = sv.code.GenerateTo(envelope)
	return
}

func (sv *Service) ClearExpired() {
	expiredEnvelopes := sv.code.Clean()
	for _, envelope := range expiredEnvelopes {
		if envelope.AvailableAmount > 0 {
			sv.accounts[envelope.CreatorID].Deposit(envelope.AvailableAmount)
			envelope.AvailableAmount = 0
		}
	}
}

func (sv *Service) nextEnvelopeID() int64 {
	sv.maxID++
	return sv.maxID
}
