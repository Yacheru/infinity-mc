package constants

import "errors"

var (
	EmptyRequiredVar        = errors.New("empty required variable in category")
	NoNewsFoundError        = errors.New("no news found")
	ParamsDoesntExistsError = errors.New("params doesnt exists")

	ErrorPingElastic = errors.New("error pinging elasticsearch")
)
