package turl

import (
	"errors"

	"github.com/speps/go-hashids"
)

//TURL describe the basic information used to operating url-hash entry.
type TURL struct {
	h *hashids.HashID
	s *Store
}

//New return a pointer to a TURL instance.
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

//Encode encodes a int64 slice into a string.
func (t *TURL) Encode(s []int64) (string, error) {
	if t == nil {
		return "", errors.New("t is nil")
	}
	return t.h.EncodeInt64(s)
}

//Decode decodes a string into a int64 slice.
func (t *TURL) Decode(s string) ([]int64, error) {
	if t == nil {
		return []int64{}, errors.New("t is nil")
	}
	return t.h.DecodeInt64WithError(s)
}

//GetURL returns the original url from a hashed key.
func (t *TURL) GetURL(k string) (string, error) {
	res, err := t.s.Get(k)
	if err != nil {
		return "", err
	}
	return res.url, nil
}

//GetEntry returns a url-hash entry from a hashed key.
func (t *TURL) GetEntry(k string) (*Entry, error) {
	return t.s.Get(k)
}

//Set sets url to TURL, and get a hashed, tiny key to represent it.
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
	_, err = t.s.Set(entry)
	return key, err
}
