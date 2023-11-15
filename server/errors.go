package server

type ErrCacheNotFound struct {
	Message string
}

func NewErrCacheNotFound(message string) error {
	return &ErrCacheNotFound{Message: message}
}

func (e *ErrCacheNotFound) Error() string {
	return e.Message
}

type ErrCacheUnknown struct {
	Message string
}

func NewErrCacheUnknown(message string) error {
	return &ErrCacheUnknown{Message: message}
}

func (e *ErrCacheUnknown) Error() string {
	return e.Message
}
