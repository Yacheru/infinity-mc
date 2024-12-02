package entities

import "time"

type EmailVerification struct {
	Email      string    `bson:"email"`
	Code       int       `bson:"code"`
	Expiration time.Time `bson:"expiration"`
}
