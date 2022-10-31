package timer

import "github.com/pangdogs/galaxy/define"

var Plugin = define.DefinePlugin[Timer, any]().RuntimePlugin(newTimer)
