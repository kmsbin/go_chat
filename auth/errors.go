package auth

import (
	"errors"
)

type AuthErrors struct {
	InvalidToken   error
	MalformedToken error
	UndefinedToken error
}

func (inToken AuthErrors) New() AuthErrors {
	inToken.InvalidToken = errors.New("invalid Token")
	inToken.MalformedToken = errors.New("malformed Token")
	inToken.UndefinedToken = errors.New("token Undefined")
	return inToken
}
