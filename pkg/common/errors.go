package common

var (
	ErrNotFound = Error{404, "Not found"}
)

// Error is a type of error used for meta.
type Error struct {
	code int
	msg  string
}

// Error returns the message in MetaError.
func (e Error) Error() string {
	return e.msg
}
