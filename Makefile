include $(GOROOT)/src/Make.inc

TARG=slices

GOFILES=\
	slices.go\
	slice.go\
	slice_value.go\
	int_slice.go\
	float32_slice.go\
	float64_slice.go

include $(GOROOT)/src/Make.pkg