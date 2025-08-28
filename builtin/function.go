package builtin

import (
	"reflect"
)

// ArgType defines the type of function argument
type ArgType byte

const (
	ExprArg ArgType = 1 << iota
	PredicateArg
)

const OptionalArg ArgType = 1 << 7

type Function struct {
	Name      string
	Fast      func(arg any) any
	Func      func(args ...any) (any, error)
	Safe      func(args ...any) (any, uint, error)
	Types     []reflect.Type
	Validate  func(args []reflect.Type) (reflect.Type, error)
	Deref     func(i int, arg reflect.Type) bool
	Predicate bool
	FuncArgs  []ArgType // 新增：定义参数类型，包括谓语参数
}

func (f *Function) Type() reflect.Type {
	if len(f.Types) > 0 {
		return f.Types[0]
	}
	return reflect.TypeOf(f.Func)
}
