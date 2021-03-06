package core

import (
	"crypto/rsa"
)

type AuthService interface {
	PublicKey() *rsa.PublicKey
	Authenticate() (string, error)
}
