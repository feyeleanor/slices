package slices

import "testing"

func TestPrepend(t *testing.T) {
	ConfirmPrepend := func(c Insertable, v interface{}, r Equatable) {
		if Prepend(c, v); !r.Equal(c) {
			t.Fatalf("Prepend(%v) should be %v but is %v", v, r, c)
		}
	}

	ConfirmPrepend(C128List(), complex128(0), C128List(0))
	ConfirmPrepend(C128List(0), complex128(1), C128List(1, 0))

	ConfirmPrepend(C128List(), C128List(0), C128List(0))
	ConfirmPrepend(C128List(), C128List(0, 1), C128List(0, 1))
	ConfirmPrepend(C128List(0, 1, 2), C128List(3, 4), C128List(3, 4, 0, 1, 2))
}

func TestAppend(t *testing.T) {
	ConfirmAppend := func(c Insertable, v interface{}, r Equatable) {
		if Append(c, v); !r.Equal(c) {
			t.Fatalf("Append(%v) should be %v but is %v", v, r, c)
		}
	}

	ConfirmAppend(C128List(), complex128(0), C128List(0))
	ConfirmAppend(C128List(), C128List(0), C128List(0))
	ConfirmAppend(C128List(), C128List(0, 1), C128List(0, 1))
	ConfirmAppend(C128List(0, 1, 2), C128List(3, 4), C128List(0, 1, 2, 3, 4))

}