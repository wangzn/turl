package turl

import (
	"errors"
	"reflect"
	"strconv"
	"time"
)

const (
	ST_VALID_URL = iota
	ST_INVALID_URL
)

type Entry struct {
	id         int64
	key        string
	url        string
	status     int
	updateTime string
}

func (e *Entry) GetKey() string {
	return e.key
}

func (e *Entry) SetStatus(status int) {
	e.status = status
}

func (e *Entry) GetMapData() map[string]string {
	m := make(map[string]string)
	m["id"] = strconv.Itoa(int(e.id))
	m["key"] = e.key
	m["url"] = e.url
	m["status"] = strconv.Itoa(e.status)
	m["updateTime"] = e.updateTime
	return m
}

func NewEntryByMap(m map[string]string) (*Entry, error) {
	err := errors.New("Invalid urlentry map data")
	entry := Entry{}
	t := reflect.TypeOf(entry)
	v := reflect.ValueOf(entry)
	for k := 0; k < t.NumField(); k++ {
		key := t.Field(k).Name
		if _, ok := m[key]; !ok {
			return nil, err
		}
	}
	entry.id, _ = strconv.ParseInt(m["id"], 10, 64)
	entry.key = m["key"]
	entry.url = m["url"]
	entry.status, _ = strconv.Atoi(m["status"])
	entry.updateTime = m["updateTime"]
	return &entry, nil
}

func NewEntry(id int64) *Entry {
	return &Entry{
		id:         id,
		status:     ST_VALID_URL,
		updateTime: time.Now().String(),
	}
}
