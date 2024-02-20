package configfile

type Marshaler interface {
	// Backend configuration marshal
	Marshal() ([]byte, error)

	// Check if the struct value is nil
	IsNil() bool
}
