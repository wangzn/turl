package turl

import (
	"errors"
	"github.com/speps/go-hashids"
)

type TURL struct {
	h *hashids.HashID
	s *Store
}

func New(salt, addr, pwd string) *TURL {
	hd := hashids.NewData()
	hd.MinLength = 6
	hd.Salt = salt
	h := hashids.NewWithData(hd)
	s := NewStore(addr, pwd)
	return &TURL{
		h: h,
		s: s,
	}
}

func (t *TURL) Encode(s []int64) (string, error) {
	if t == nil {
		return "", errors.New("t is nil")
	}
	return t.h.EncodeInt64(s)
}

func (t *TURL) Decode(s string) ([]int64, error) {
	if t == nil {
		return []int64{}, errors.New("t is nil")
	}
	return t.h.DecodeInt64WithError(s)
}

func (t *TURL) GetUrl(k string) (string, error) {
	res, err := t.s.Get(k)
	if err != nil {
		return "", err
	}
	return res.url, nil
}

func (t *TURL) GetEntry(k string) (*Entry, error) {
	return t.s.Get(k)
}

func (t *TURL) Set(url string) (string, error) {
	id, err := t.s.GetID()
	if err != nil {
		return "", err
	}
	entry := NewEntry(id)
	key, err := t.Encode([]int64{id})
	if err != nil {
		return "", err
	}
	entry.key = key
	entry.url = url
	res, err := t.s.Set(entry)
	return key, err
}
