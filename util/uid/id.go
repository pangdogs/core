package uid

import "github.com/segmentio/ksuid"

var (
	// Nil is a nil id.
	Nil Id = ""

	// New generates a new id.
	New = func() Id {
		return Id(ksuid.New().String())
	}
)

// Id represents a global unique id.
type Id string

// IsNil checks if an Id is nil.
func (id Id) IsNil() bool {
	return id == Nil
}
