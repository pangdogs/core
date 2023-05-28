package ec

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"kit.golaxy.org/golaxy/internal"
)

func (entity *EntityBehavior) String() string {
	return EntitySerializer.String(entity.opts.CompositeFace.Iface)
}

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

func (entity *EntityBehavior) Value() (driver.Value, error) {
	return EntitySerializer.Value(entity.opts.CompositeFace.Iface)
}

func (entity *EntityBehavior) Scan(src interface{}) error {
	return EntitySerializer.Scan(entity.opts.CompositeFace.Iface, src)
}

var EntitySerializer internal.Serializer[Entity] = DefaultEntitySerializer{}

type DefaultEntitySerializer struct{}

func (DefaultEntitySerializer) String(entity Entity) string {
	var parentInfo string
	if parent, ok := entity.GetParent(); ok {
		parentInfo = parent.GetId().String()
	} else {
		parentInfo = "nil"
	}

	return fmt.Sprintf("{Id:%s SerialNo:%d Prototype:%s Parent:%s State:%s}",
		entity.GetId(), entity.GetSerialNo(), entity.GetPrototype(), parentInfo, entity.GetState())
}

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

func (DefaultEntitySerializer) Value(entity Entity) (driver.Value, error) {
	return nil, errors.New("not support")
}

func (DefaultEntitySerializer) Scan(entity Entity, src interface{}) error {
	return errors.New("not support")
}
