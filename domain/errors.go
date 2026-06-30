package domain

import "errors"

var ErrNotFound = errors.New("record not found")
var ErrEmailInUse = errors.New("email already in use")
var ErrUsernameInUse = errors.New("username already taken")
var ErrInvalidCredentials = errors.New("invalid email or password")
var ErrInvalidToken = errors.New("invalid or expired token")
var ErrSchoolAlreadyExists = errors.New("school already exists for this admin")
var ErrPlateNumberInUse = errors.New("plate number already in use")
var ErrDriverAlreadyConnected = errors.New("driver already connected to this school")
