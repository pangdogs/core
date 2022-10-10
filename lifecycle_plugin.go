package galaxy

type _PluginInit interface {
	Init()
}

type _PluginShut interface {
	Shut()
}
