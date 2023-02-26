package ec

import "encoding/base64"

// ID 唯一ID（160位）
type ID [20]byte

func (id ID) String() string {
	return base64.RawURLEncoding.EncodeToString(id[0:])
}
