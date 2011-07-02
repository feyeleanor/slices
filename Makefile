include $(GOROOT)/src/Make.inc

TARG=github.com/feyeleanor/slices

GOFILES=\
	slices.go\
	slice.go\
	slice_value.go\
	string.go\
	uintptr.go\
	uint.go\
	uint8.go\
	uint16.go\
	uint32.go\
	uint64.go\
	int.go\
	int8.go\
	int16.go\
	int32.go\
	int64.go\
	float32.go\
	float64.go\
	complex64.go\
	complex128.go

include $(GOROOT)/src/Make.pkg