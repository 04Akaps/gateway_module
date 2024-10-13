package error

type ErrType string

const (
	CheckType_Error = ErrType("not supported get type")
)

// implement error interface for using err
func (e ErrType) Error() string {
	return string(e)
}
