package model

import (
	"github.com/go-playground/validator/v10"
)

var UpdateValidate = validator.New()

func ValidateUpdateWatchList(wl *UpdateWatchList) error {
	return UpdateValidate.Struct(wl)
}

type UpdateWatchList struct {
	Id            string   `json:"id" validate:"required,min=3,max=20,alphanum"`
	WatchListName string   `json:"watchListName" validate:"required,min=1,max=30"`
	Scripts       []string `json:"scripts"`
}
