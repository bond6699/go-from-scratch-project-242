package errors

import errs "errors"

var (
	ErrPathNotFound = errs.New("path not found")
	ErrPathArgsNotProvided = errs.New("path argument not provided, only one path argument is allowed")
	ErrInvalidPath = errs.New("invalid path")

	ErrAccessDenied = errs.New("access denied")
	ErrSymslinkRead = errs.New("cannot read symlink")
	ErrReadPath = errs.New("cannot read path")
)