package dadjoke

import "errors"

var ErrSavingJokes = errors.New("Error occured saving dad jokes")
var ErrFetchingJokes = errors.New("Error occured retrieving jokes")
var ErrNoRecords = errors.New("No records found")
