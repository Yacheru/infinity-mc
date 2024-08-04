package constants

import "errors"

var (
	ErrLoadConfig       = errors.New("failed read in config")
	ErrParseConfig      = errors.New("failed unmarshal the config into a Struct")
	ErrEmptyVar         = errors.New("required environment variable is empty")
	ErrFailedLoadConfig = errors.New("failed load config")
)
