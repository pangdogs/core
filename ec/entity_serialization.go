package ec

import (
	"errors"
	"kit.golaxy.org/golaxy/internal"
)

var EntitySerializer internal.Serializer[Entity] = DefaultEntitySerializer{}

func (entity *EntityBehavior) MarshalText() ([]byte, error) {
	return EntitySerializer.MarshalText(entity.opts.CompositeFace.Iface)
}

func (entity *EntityBehavior) UnmarshalText(b []byte) error {
	return EntitySerializer.UnmarshalText(entity.opts.CompositeFace.Iface, b)
}

func (entity *EntityBehavior) MarshalBinary() ([]byte, error) {
	return EntitySerializer.MarshalBinary(entity.opts.CompositeFace.Iface)
}

func (entity *EntityBehavior) UnmarshalBinary(b []byte) error {
	return EntitySerializer.UnmarshalBinary(entity.opts.CompositeFace.Iface, b)
}

type DefaultEntitySerializer struct{}

func (DefaultEntitySerializer) MarshalText(entity Entity) ([]byte, error) {
	return nil, errors.New("not support")
}

func (DefaultEntitySerializer) UnmarshalText(entity Entity, b []byte) error {
	return errors.New("not support")
}

func (DefaultEntitySerializer) MarshalBinary(entity Entity) ([]byte, error) {
	return nil, errors.New("not support")
}

func (DefaultEntitySerializer) UnmarshalBinary(entity Entity, b []byte) error {
	return errors.New("not support")
}
