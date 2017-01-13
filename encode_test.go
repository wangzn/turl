package turl

import (
	"math/rand"
	"testing"
	"time"
)

func TestEncodeAndDecode(t *testing.T) {
	n := 100
	e := NewEncode("testsuitforturl")
	data := make(map[string][]int64)
	for i := 0; i < n; i++ {
		s := getRandomIntSlice(i)
		res, err := e.Encode(s)
		if err != nil {
			t.Error("Encode with error")
		} else {
			if exist, ok := data[res]; ok {
				if !equalSlice(exist, s) {
					t.Error("Encode collision")
				}
			} else {
				data[res] = s
			}
		}
	}

	for k, v := range data {
		decoded, err := e.Decode(k)
		if err != nil {
			t.Error("Decode with error")
		} else {
			if !equalSlice(decoded, v) {
				t.Error("Decoded slice does not equal to the original")
			}
		}
	}
}

func getRandomIntSlice(l int) []int64 {
	rand.Seed(time.Now().Unix())
	var s []int64
	for i := 0; i < 10; i++ {
		s = append(s, rand.Int63())
	}
	return s
}

func equalSlice(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if v != b[k] {
			return false
		}
	}
	return true
}
