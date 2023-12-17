package apis

// Error message
const (
	ErrNotFound = Error("not found")
)

type Error string

func (e Error) Error() string {
	return string(e)
}
