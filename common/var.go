package common

import "errors"

var (
	ErrRecordNotFound      = errors.New("record not found")
	ErrNoRecordsAreChanged = errors.New("no records are changed")
)
