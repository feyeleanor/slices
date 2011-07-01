package slices

import "reflect"

const(
	IS_LESS_THAN	= iota - 1
	IS_SAME_AS
	IS_GREATER_THAN
)

type Nested interface {
	Depth() int
}

type Flattenable interface {
	Flatten()
}

type Equatable interface {
	Equal(interface{}) bool
}

type Typed interface {
	Type() reflect.Type
}