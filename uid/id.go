package uid

import (
	"database/sql/driver"
	"github.com/segmentio/ksuid"
)

var Id_String = func(id Id) string {
	return ksuid.KSUID(id).String()
}

var Id_Bytes = func(id Id) []byte {
	return ksuid.KSUID(id).Bytes()
}

var Id_Set = func(id *Id, s string) error {
	return ((*ksuid.KSUID)(id)).Set(s)
}

var Id_MarshalText = func(id Id) ([]byte, error) {
	return ksuid.KSUID(id).MarshalText()
}

var Id_MarshalBinary = func(id Id) ([]byte, error) {
	return ksuid.KSUID(id).MarshalBinary()
}

var Id_UnmarshalText = func(id *Id, b []byte) error {
	return ((*ksuid.KSUID)(id)).UnmarshalText(b)
}

var Id_UnmarshalBinary = func(id *Id, b []byte) error {
	return ((*ksuid.KSUID)(id)).UnmarshalBinary(b)
}

var Id_Value = func(id Id) (driver.Value, error) {
	return ksuid.KSUID(id).Value()
}

var Id_Scan = func(id *Id, src any) error {
	return ((*ksuid.KSUID)(id)).Scan(src)
}

var New = func() Id {
	return Id(ksuid.New())
}

// Nil nil id
var Nil Id

// Id 唯一Id（160位）
type Id [20]byte

func (id Id) String() string {
	return Id_String(id)
}

func (id Id) Bytes() []byte {
	return Id_Bytes(id)
}

func (id Id) IsNil() bool {
	return id == Nil
}

func (id Id) Get() any {
	return id
}

func (id *Id) Set(s string) error {
	return Id_Set(id, s)
}

func (id Id) MarshalText() ([]byte, error) {
	return Id_MarshalText(id)
}

func (id Id) MarshalBinary() ([]byte, error) {
	return Id_MarshalBinary(id)
}

func (id *Id) UnmarshalText(b []byte) error {
	return Id_UnmarshalText(id, b)
}

func (id *Id) UnmarshalBinary(b []byte) error {
	return Id_UnmarshalBinary(id, b)
}

func (id Id) Value() (driver.Value, error) {
	return Id_Value(id)
}

func (id *Id) Scan(src any) error {
	return Id_Scan(id, src)
}
