package turl

import (
	"errors"

	"github.com/speps/go-hashids"
)

//Encode describe encode and decode methods of hashing url to key
type Encode struct {
	h *hashids.HashID
}

//NewEncode return a new instance of encoder for url hashing.
func NewEncode(salt string) *Encode {
	hd := hashids.NewData()
	hd.MinLength = 6
	hd.Salt = salt
	h := hashids.NewWithData(hd)
	return &Encode{
		h: h,
	}
}

//Encode encodes a int64 slice into a string.
func (e *Encode) Encode(s []int64) (string, error) {
	if e == nil {
		return "", errors.New("t is nil")
	}
	return e.h.EncodeInt64(s)
}

//Decode decodes a string into a int64 slice.
func (e *Encode) Decode(s string) ([]int64, error) {
	if e == nil {
		return []int64{}, errors.New("t is nil")
	}
	return e.h.DecodeInt64WithError(s)
}
