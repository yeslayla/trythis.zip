package api

type MethodNotAllowedError struct {
	Err error
}

func (e *MethodNotAllowedError) Error() string {
	return e.Err.Error()
}

type PageNotFoundError struct {
	Err error
}

func (e *PageNotFoundError) Error() string {
	return e.Err.Error()
}
