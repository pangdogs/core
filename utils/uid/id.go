package uid

import "github.com/segmentio/ksuid"

var (
	// Nil is a nil id.
	Nil Id = ""

	// New generates a new id.
	New = func() Id {
		return Id(ksuid.New().String())
	}

	// From generate id from string.
	From = func(str string) Id {
		return Id(str)
	}
)

// Id represents a global unique id.
type Id string

// IsNil checks if an Id is nil.
func (id Id) IsNil() bool {
	return id == Nil
}

// String implements fmt.Stringer
func (id Id) String() string {
	return string(id)
}
