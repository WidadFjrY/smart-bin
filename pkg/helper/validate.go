package helper

import (
	"smart-trash-bin/pkg/exception"
	"strings"
)

func ValError(err error) {
	if err != nil {
		if strings.Contains(err.Error(), "failed on the 'email' tag") {
			panic(exception.NewBadRequestError("email format must be correct"))
		} else if strings.Contains(err.Error(), "failed on the 'min' tag") {
			panic(exception.NewBadRequestError("password must contain a minimum of 8 digits"))
		} else if strings.Contains(err.Error(), "failed on the 'required' tag") {
			panic(exception.NewBadRequestError("form must be filled out"))
		}
		panic(err)
	}
}
