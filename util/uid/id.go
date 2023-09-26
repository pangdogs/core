package uid

import (
	"database/sql/driver"
	"github.com/segmentio/ksuid"
)

var (
	// Id_String converts an Id to its string representation.
	Id_String = func(id Id) string {
		return ksuid.KSUID(id).String()
	}

	// Id_Bytes converts an Id to its byte slice representation.
	Id_Bytes = func(id Id) []byte {
		return ksuid.KSUID(id).Bytes()
	}

	// Id_Set sets the value of an Id from a string.
	Id_Set = func(id *Id, s string) error {
		return ((*ksuid.KSUID)(id)).Set(s)
	}

	// Id_MarshalText marshals an Id to a textual representation.
	Id_MarshalText = func(id Id) ([]byte, error) {
		return ksuid.KSUID(id).MarshalText()
	}

	// Id_MarshalBinary marshals an Id to a binary representation.
	Id_MarshalBinary = func(id Id) ([]byte, error) {
		return ksuid.KSUID(id).MarshalBinary()
	}

	// Id_UnmarshalText unmarshals a textual representation into an Id.
	Id_UnmarshalText = func(id *Id, b []byte) error {
		return ((*ksuid.KSUID)(id)).UnmarshalText(b)
	}

	// Id_UnmarshalBinary unmarshals a binary representation into an Id.
	Id_UnmarshalBinary = func(id *Id, b []byte) error {
		return ((*ksuid.KSUID)(id)).UnmarshalBinary(b)
	}

	// Id_Value returns the driver Value for an Id.
	Id_Value = func(id Id) (driver.Value, error) {
		return ksuid.KSUID(id).Value()
	}

	// Id_Scan scans a value from a database driver source into an Id.
	Id_Scan = func(id *Id, src any) error {
		return ((*ksuid.KSUID)(id)).Scan(src)
	}

	// Nil is a nil id.
	Nil Id = Id(ksuid.Nil)

	// New generates a new id.
	New = func() Id {
		return Id(ksuid.New())
	}

	// UnmarshalText unmarshals a textual representation into an Id.
	UnmarshalText = func(v []byte) (id Id, err error) {
		err = id.UnmarshalText(v)
		return
	}

	// UnmarshalBinary unmarshals a binary representation into an Id.
	UnmarshalBinary = func(v []byte) (id Id, err error) {
		err = id.UnmarshalBinary(v)
		return
	}
)

// Id represents a global unique id (160bit).
type Id [20]byte

// String implements fmt.Stringer, flag.Value.
func (id Id) String() string {
	return Id_String(id)
}

// Bytes returns the byte slice representation of an Id.
func (id Id) Bytes() []byte {
	return Id_Bytes(id)
}

// IsNil checks if an Id is nil.
func (id Id) IsNil() bool {
	return id == Nil
}

// Get implements flag.Getter.
func (id Id) Get() any {
	return id
}

// Set implements flag.Value.
func (id *Id) Set(s string) error {
	return Id_Set(id, s)
}

// MarshalText implements encoding.TextMarshaler.
func (id Id) MarshalText() ([]byte, error) {
	return Id_MarshalText(id)
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (id Id) MarshalBinary() ([]byte, error) {
	return Id_MarshalBinary(id)
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (id *Id) UnmarshalText(b []byte) error {
	return Id_UnmarshalText(id, b)
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (id *Id) UnmarshalBinary(b []byte) error {
	return Id_UnmarshalBinary(id, b)
}

// Value implements driver.Valuer.
func (id Id) Value() (driver.Value, error) {
	return Id_Value(id)
}

// Scan implements sql.Scanner.
func (id *Id) Scan(src any) error {
	return Id_Scan(id, src)
}
