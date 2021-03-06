package service

import "github.com/timoth-y/scrapnote-api/data.users/core/model"

type UserSerializer interface {
	Decode(input []byte) (*model.User, error)
	DecodeRange(input []byte) ([]*model.User, error)
	DecodeMap(input []byte) (map[string]interface{}, error)
	DecodeInto(input []byte, target interface{}) error
	Encode(input interface{}) ([]byte, error)
}