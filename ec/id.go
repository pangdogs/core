package ec

import (
	"github.com/segmentio/ksuid"
)

// ID 唯一ID（160位）
type ID [20]byte

func (id ID) String() string {
	return id.Encode()
}

func (id ID) Encode() string {
	return EncodeID(id)
}

var EncodeID = func(id ID) string {
	return ksuid.KSUID(id).String()
}

func (id *ID) Decode(str string) error {
	_id, err := DecodeID(str)
	if err != nil {
		return err
	}
	*id = _id
	return nil
}

var DecodeID = func(str string) (ID, error) {
	id, err := ksuid.Parse(str)
	return ID(id), err
}
