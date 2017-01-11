package turl

import (
	"errors"
	"reflect"
	"strconv"
	"time"
)

//Define valid and invalid url status in url entry struct.
const (
	StValidURL = iota
	StInvalidURL
)

//Entry define the struct of a url-hash entry.
type Entry struct {
	id         int64
	key        string
	url        string
	status     int
	updateTime string
}

//GetKey return the key of url-hash entry.
func (e *Entry) GetKey() string {
	return e.key
}

//SetStatus sets the status of a url-hash entry.
func (e *Entry) SetStatus(status int) {
	e.status = status
}

//GetMapData returns a mapped data according to a url-hash entry.
func (e *Entry) GetMapData() map[string]string {
	m := make(map[string]string)
	m["id"] = strconv.Itoa(int(e.id))
	m["key"] = e.key
	m["url"] = e.url
	m["status"] = strconv.Itoa(e.status)
	m["updateTime"] = e.updateTime
	return m
}

//NewEntryByMap return a url-hash entry according to a mapped data.
//Err occurs if some filed missing.
func NewEntryByMap(m map[string]string) (*Entry, error) {
	err := errors.New("Invalid urlentry map data")
	entry := Entry{}
	t := reflect.TypeOf(entry)
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

//NewEntry return a pointer to a new url-hash entry.
func NewEntry(id int64) *Entry {
	return &Entry{
		id:         id,
		status:     StValidURL,
		updateTime: time.Now().String(),
	}
}
