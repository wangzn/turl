package turl

//TURL describe the basic information used to operating url-hash entry.
type TURL struct {
	e *Encode
	s *Store
}

//New return a pointer to a TURL instance.
func New(salt, addr string) (*TURL, error) {
	e := NewEncode(salt)
	s, err := NewStore(addr)
	return &TURL{
		e: e,
		s: s,
	}, err
}

//GetURL returns the original url from a hashed key.
func (t *TURL) GetURL(k string) (string, error) {
	res, err := t.s.Get(k)
	if err != nil {
		return "", err
	}
	if res.IsActive() {
		return res.url, nil
	}
	return "", ErrInactiveURL
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
	key, err := t.e.Encode([]int64{id})
	if err != nil {
		return "", err
	}
	entry.key = key
	entry.url = url
	_, err = t.s.Set(entry)
	return key, err
}

//InactiveURL inactive the stored url-hash entry.
func (t *TURL) InactiveURL(key string) error {
	e, err := t.s.Get(key)
	if err != nil {
		return err
	}
	if e == nil || e.GetKey() != key {
		return ErrKeyNotFound
	}
	e.status = StInactiveURL
	_, err = t.s.Set(e)
	if err != nil {
		return err
	}
	return nil
}
