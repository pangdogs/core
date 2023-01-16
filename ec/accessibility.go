package ec

// Accessibility 可访问性
type Accessibility int32

const (
	Local  Accessibility = iota // 本地可以访问
	Global                      // 全局可以访问
)
