package models

import "errors"

type ErrRespons struct {
	Code    uint16 `json:"code"`
	Message string `json:"massage"`
}

var (
	ErrNoSuchPerson = errors.New("no such person")
)
