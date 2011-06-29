package slices

import "reflect"

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