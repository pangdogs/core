package meta

import "git.golaxy.org/core/utils/generic"

type Meta = generic.SliceMap[string, any]

type _MetaCtor struct {
	meta Meta
}

func (c _MetaCtor) Add(k string, v any) _MetaCtor {
	c.meta.Add(k, v)
	return c
}

func (c _MetaCtor) Delete(k string) _MetaCtor {
	c.meta.Delete(k)
	return c
}

func (c _MetaCtor) Combine(m map[string]any) _MetaCtor {
	for k, v := range m {
		c.meta.TryAdd(k, v)
	}
	return c
}

func (c _MetaCtor) Override(m map[string]any) _MetaCtor {
	for k, v := range m {
		c.meta.Add(k, v)
	}
	return c
}

func (c _MetaCtor) Clean() _MetaCtor {
	c.meta = c.meta[:0]
	return c
}

func (c _MetaCtor) Get() Meta {
	return c.meta
}

func Make() _MetaCtor {
	return _MetaCtor{}
}
