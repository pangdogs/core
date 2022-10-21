package ec

import (
	"github.com/pangdogs/galaxy/util"
)

type ContextHolder interface {
	getContext() util.IfaceCache
}
