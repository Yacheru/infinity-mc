package constants

import "errors"

var (
	ErrLoadConfig  = errors.New("failed loading config")
	ErrParseConfig = errors.New("failed to parse env to config struct")
	ErrEmptyVar    = errors.New("required variable environment is empty")

	ErrService   = errors.New("service is invalid")
	ErrReqParams = errors.New("nickname, price and duration are required params")
)
