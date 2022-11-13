package runtime

import "github.com/galaxy-kit/galaxy-go/util/container"

type _InnerGC interface {
	getInnerGC() container.GC
}
