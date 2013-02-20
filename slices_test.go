package slices

import (
	"sort"
	"testing"
)

func TestPrepend(t *testing.T) {
	ConfirmPrepend := func(c C128Slice, v interface{}, r Equatable) {
		if Prepend(&c, v); !r.Equal(c) {
			t.Fatalf("Prepend(%v) should be %v but is %v", v, r, c)
		}
	}

	ConfirmPrepend(C128Slice{}, complex128(0), C128Slice{0})
	ConfirmPrepend(C128Slice{0}, complex128(1), C128Slice{1, 0})

	ConfirmPrepend(C128Slice{}, C128Slice{0}, C128Slice{0})
	ConfirmPrepend(C128Slice{}, C128Slice{0, 1}, C128Slice{0, 1})
	ConfirmPrepend(C128Slice{0, 1, 2}, C128Slice{3, 4}, C128Slice{3, 4, 0, 1, 2})
}

func TestAppend(t *testing.T) {
	ConfirmAppend := func(c C128Slice, v interface{}, r Equatable) {
		if Append(&c, v); !r.Equal(c) {
			t.Fatalf("Append(%v) should be %v but is %v", v, r, c)
		}
	}

	ConfirmAppend(C128Slice{}, complex128(0), C128Slice{0})
	ConfirmAppend(C128Slice{}, C128Slice{0}, C128Slice{0})
	ConfirmAppend(C128Slice{}, C128Slice{0, 1}, C128Slice{0, 1})
	ConfirmAppend(C128Slice{0, 1, 2}, C128Slice{3, 4}, C128Slice{0, 1, 2, 3, 4})

}

func TestShuffle(t *testing.T) {
	ConfirmShuffle := func(s, r interface{}) {
		if Shuffle(s.(Deck)); Equal(s, r) {
			t.Fatalf("Shuffle(%v) should change order of elements", s)
		}
		if sort.Sort(s.(sort.Interface)); !Equal(s, r) {
			t.Fatalf("Shuffle() when sorted should be %v but is %v", r, s)
		}
	}

	ConfirmShuffle(C64Slice{0, 1, 2, 3, 4, 5}, C64Slice{0, 1, 2, 3, 4, 5})
	ConfirmShuffle(C128Slice{0, 1, 2, 3, 4, 5}, C128Slice{0, 1, 2, 3, 4, 5})
	ConfirmShuffle(F64Slice{0, 1, 2, 3, 4, 5}, F64Slice{0, 1, 2, 3, 4, 5})
	ConfirmShuffle(F64Slice{0, 1, 2, 3, 4, 5}, F64Slice{0, 1, 2, 3, 4, 5})
	ConfirmShuffle(ISlice{0, 1, 2, 3, 4, 5}, ISlice{0, 1, 2, 3, 4, 5})
	ConfirmShuffle(I8Slice{0, 1, 2, 3, 4, 5}, I8Slice{0, 1, 2, 3, 4, 5})
	ConfirmShuffle(I16Slice{0, 1, 2, 3, 4, 5}, I16Slice{0, 1, 2, 3, 4, 5})
	ConfirmShuffle(I32Slice{0, 1, 2, 3, 4, 5}, I32Slice{0, 1, 2, 3, 4, 5})
	ConfirmShuffle(I64Slice{0, 1, 2, 3, 4, 5}, I64Slice{0, 1, 2, 3, 4, 5})

	ConfirmShuffle(SSlice{"A", "B", "A", "B", "A"}, SSlice{"A", "A", "A", "B", "B"})
	ConfirmShuffle(USlice{0, 1, 2, 3, 4, 5}, USlice{0, 1, 2, 3, 4, 5})
	ConfirmShuffle(U8Slice{0, 1, 2, 3, 4, 5}, U8Slice{0, 1, 2, 3, 4, 5})
	ConfirmShuffle(U16Slice{0, 1, 2, 3, 4, 5}, U16Slice{0, 1, 2, 3, 4, 5})
	ConfirmShuffle(U32Slice{0, 1, 2, 3, 4, 5}, U32Slice{0, 1, 2, 3, 4, 5})
	ConfirmShuffle(U64Slice{0, 1, 2, 3, 4, 5}, U64Slice{0, 1, 2, 3, 4, 5})
	ConfirmShuffle(ASlice{0, 1, 2, 3, 4, 5}, ASlice{0, 1, 2, 3, 4, 5})
}

func TestShuffleWithoutSorting(t *testing.T) {
	ConfirmShuffle := func(s Deck, r interface{}) {
		if Shuffle(s); Equal(s, r) {
			t.Fatalf("Shuffle(%v) should change order of elements", s)
		}
	}
	t.Log("Implement Sort for RSlice")
	ConfirmShuffle(RList(0, 1, 2, 3, 4, 5), RList(0, 1, 2, 3, 4, 5))

	t.Log("Implement Sort for Slice")
	ConfirmShuffle(Slice{0, 1, 2, 3, 4, 5}, Slice{0, 1, 2, 3, 4, 5})

	t.Log("Implement Sort for VSlice")
	ConfirmShuffle(VList(0, 1, 2, 3, 4, 5), VList(0, 1, 2, 3, 4, 5))
}