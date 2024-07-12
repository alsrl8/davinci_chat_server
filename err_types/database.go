package err_types

import "errors"

var ErrUserExists = errors.New("user already exists")

var ErrEmailPasswordNotMatch = errors.New("email password not match")
