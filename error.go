package turl

import (
	"errors"
)

//Define all error type.
var (
	ErrInactiveURL            = errors.New("ErrInactiveURL")
	ErrKeyNotFound            = errors.New("ErrKeyNotFound")
	ErrInvalidStoreAddr       = errors.New("ErrInvalidStoreAddr")
	ErrStoreNotConnected      = errors.New("ErrStoreNotConnected")
	ErrInvalidEntryField      = errors.New("ErrInvalidEntryField")
	ErrUnsupportedStoreType   = errors.New("ErrUnsupportedStoreType")
	ErrInvalidInstancePointer = errors.New("ErrInvalidInstancePointer")
)
