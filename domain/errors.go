package domain

import "errors"

var ErrNotFound = errors.New("record not found")
var ErrEmailInUse = errors.New("email already in use")
var ErrUsernameInUse = errors.New("username already taken")
