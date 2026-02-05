package services

import "errors"

var (
	ErrMacAddressRequired  = errors.New("mac_address is required")
	ErrFeederAlreadyExists = errors.New("feeder already exists")
	ErrEmailAlreadyExists  = errors.New("email already exists")
	ErrInvalidCredentials  = errors.New("invalid credentials")
)
