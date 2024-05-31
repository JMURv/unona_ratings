package repository

import "errors"

var ErrUserUUIDIsRequired = errors.New("userUUID is required")
var ErrRatingIsRequired = errors.New("rating is required")
var ErrTextIsRequired = errors.New("text is required")
