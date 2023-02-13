//go:generate stringer -type Scope
package ec

// Scope 可访问作用域
type Scope int32

const (
	Scope_Local  Scope = iota // 本地可以访问
	Scope_Global              // 全局可以访问
)
