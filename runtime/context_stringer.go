package runtime

import (
	"fmt"
)

// String implements fmt.Stringer
func (ctx *ContextBehavior) String() string {
	return fmt.Sprintf("{Id:%s Name:%s}", ctx.GetId(), ctx.GetName())
}
