package turl

import (
	"gopkg.in/redis.v5"
)

const (
	ID_COUNTER = "turl_id_counter"
)

type Store struct {
	client *redis.Client
}

func NewStore(addr, pwd string) *Store {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       0, // use default DB
	})
	return &Store{
		client: client,
	}
}

func (s *Store) Set(entry *Entry) (string, error) {
	data := entry.GetMapData()
	result := s.client.HMSet(entry.GetKey(), data)
	if result.Err() != nil {
		return "", result.Err()
	}
	return result.Result()
}

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

func (s *Store) GetID() (int64, error) {
	if err := s.client.Incr(ID_COUNTER).Err(); err != nil {
		return -1, err
	}
	return s.client.Get(ID_COUNTER).Int64()
}
