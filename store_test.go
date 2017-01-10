package turl

import (
	"strconv"
	"testing"
	"time"
)

func TestStoreUtils(t *testing.T) {
	data := make(map[string]string)
	s := NewStore("127.0.0.1:6379", "")
	id, err := s.GetID()
	if err != nil {
		t.Error("Fail to get id", err.Error())
	}
	data["id"] = strconv.Itoa(int(id))
	data["key"] = "thisiskey"
	data["url"] = "thisisurl"
	data["status"] = "0"
	data["updateTime"] = time.Now().String()

	e, err := NewEntryByMap(data)
	if err != nil {
		t.Error("Fail to init entry", err.Error())
	}

	_, err = s.Set(e)
	if err != nil {
		t.Error("Fail to set", err.Error())
	}

}
