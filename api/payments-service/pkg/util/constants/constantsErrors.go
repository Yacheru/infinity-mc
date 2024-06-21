package constants

import "errors"

var (
	ErrLoadConfig  = errors.New("failed loading config")
	ErrParseConfig = errors.New("failed to parse env to config struct")
	ErrEmptyVar    = errors.New("required variable environment is empty")
)
