package turl

import (
	"errors"
	"net/url"

	"gopkg.in/redis.v5"
)

//Define the key used in redis to store id counter.
const (
	IDCOUNTER = "turl_id_counter"
)

//Store describes a struct to manage client information.
type Store struct {
	client *redis.Client
}

//NewStore return a Store instance according to the addr or other params.
func NewStore(addr string) (*Store, error) {

	u, err := url.Parse(addr)
	if err != nil {
		return nil, errors.New("Invalid store addr")
	}
	if u.Scheme != "redis" {
		return nil, errors.New("Unsupported store type")
	}
	pwd := ""
	if u.User != nil {
		pwd, _ = u.User.Password()
	}
	//use arbitary redis.
	client := redis.NewClient(&redis.Options{
		Addr:     u.Host,
		Password: pwd,
		DB:       0, // use default DB
	})
	return &Store{
		client: client,
	}, nil
}

//Set sets the url-hash entry into the store.
func (s *Store) Set(entry *Entry) (string, error) {
	data := entry.GetMapData()
	result := s.client.HMSet(entry.GetKey(), data)
	if result.Err() != nil {
		return "", result.Err()
	}
	return result.Result()
}

//Get read information from store, and save it into url-hash entry.
func (s *Store) Get(k string) (*Entry, error) {
	result := s.client.HGetAll(k)
	if result.Err() != nil {
		return nil, result.Err()
	}
	m, err := result.Result()
	if err != nil {
		return nil, err
	}
	return NewEntryByMap(m)
}

//GetID returns a id used to hash url.
func (s *Store) GetID() (int64, error) {
	if err := s.client.Incr(IDCOUNTER).Err(); err != nil {
		return -1, err
	}
	return s.client.Get(IDCOUNTER).Int64()
}
