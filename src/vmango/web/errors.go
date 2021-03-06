package web

type HandlerError struct {
	Message string
}

func (e *HandlerError) Error() string {
	return e.Message
}

type ErrNotFound struct {
	*HandlerError
}

func NotFound(msg string) *ErrNotFound {
	return &ErrNotFound{&HandlerError{msg}}
}

type ErrForbidden struct {
	*HandlerError
}

func Forbidden(msg string) *ErrForbidden {
	return &ErrForbidden{&HandlerError{msg}}
}

type ErrBadRequest struct {
	*HandlerError
}

func BadRequest(msg string) *ErrBadRequest {
	return &ErrBadRequest{&HandlerError{msg}}
}

type ErrNotImplemented struct {
	*HandlerError
}

func NotImplemented() *ErrNotImplemented {
	return &ErrNotImplemented{&HandlerError{"Not implemented"}}
}
