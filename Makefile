include $(GOROOT)/src/Make.inc

TARG=slices

GOFILES=\
	slices.go\
	slice.go\
	slice_value.go

include $(GOROOT)/src/Make.pkg