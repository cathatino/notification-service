package manager

import "errors"

var (
	ErrUnexpectedLengthFromDb = errors.New("Unexpected Length Fetch From DB")
	ErrRecordNotFound         = errors.New("Record Not found")
)
