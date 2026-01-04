package services

import "errors"

var (
	ErrMacAddressRequired = errors.New("mac_address is required")
	ErrFeederAlreadyExists = errors.New("feeder already exists") 

)
