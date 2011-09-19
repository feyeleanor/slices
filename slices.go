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

type Insertable interface {
	Len() int
	Insert(int, interface{})
}

type Container interface {
	Len() int
	At(int) interface{}
	Set(int, interface{})
}

func Prepend(i Insertable, value interface{}) {
	i.Insert(0, value)
}

func Append(i Insertable, value interface{}) {
	i.Insert(i.Len(), value)
}