package turl

import (
	"strconv"
	"testing"
	"time"
)

var (
	KEY       = "thisiskey"
	URL       = "thisisurl"
	REDISADDR = "redis://@127.0.0.1:6379"
)

func TestStoreSet(t *testing.T) {
	data := make(map[string]string)
	s, err := NewStore(REDISADDR)
	if err != nil {
		t.Errorf("Fail to init store client, %s", err.Error())
	}
	id, err := s.GetID()
	if err != nil {
		t.Error("Fail to get id", err.Error())
	}
	data["id"] = strconv.Itoa(int(id))
	data["key"] = KEY
	data["url"] = URL
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

func TestStoreGet(t *testing.T) {
	key := "thisiskey"
	s, err := NewStore(REDISADDR)
	if err != nil {
		t.Errorf("Fail to init store client, %s", err.Error())
	}
	entry, err := s.Get(key)
	if err != nil {
		t.Error("Fail to get key")
	}
	if entry.key != KEY {
		t.Error("Get an invalid key")
	}
	if entry.url != URL {
		t.Error("Get an invalid url")
	}
}

func TestStoreGetID(t *testing.T) {
	s, err := NewStore(REDISADDR)
	if err != nil {
		t.Errorf("Fail to init store client, %s", err.Error())
	}
	id, err := s.GetID()
	if err != nil {
		t.Error("Fail to get id")
		if id != -1 {
			t.Error("Invalid id if error occurs")
		}
	}
	if id < 0 {
		t.Error("Got invalid id")
	}
}
