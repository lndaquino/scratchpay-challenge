package errors

import "errors"

var ErrNotFound = errors.New("No clinic found")

var ErrInvalidFrom = errors.New("Invalid format (from). Time format HH:MM ==> Hour must be between 0 and 24 and minutes between 0 and 59")

var ErrInvalidTo = errors.New("Invalid format (to). Time format HH:MM ==> Hour must be between 0 and 24 and minutes between 0 and 59")

var ErrInvalidParameters = errors.New("Invalid parameters: [from] time must be earlier than [to] time")
