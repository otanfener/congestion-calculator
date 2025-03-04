package domain

type Error string

func (e Error) Error() string {
	return string(e)
}

var (
	ErrNotFound   = Error("ErrNotFound")
	ErrBadRequest = Error("ErrBadRequest")

	ErrInternal = Error("ErrInternal")
)
