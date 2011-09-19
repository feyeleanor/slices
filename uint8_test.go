package slices

import "testing"

func TestU8SliceString(t *testing.T) {
	ConfirmString := func(s *U8Slice, r string) {
		if x := s.String(); x != r {
			t.Fatalf("%v erroneously serialised as '%v'", r, x)
		}
	}

	ConfirmString(U8List(), "()")
	ConfirmString(U8List(0), "(0)")
	ConfirmString(U8List(0, 1), "(0 1)")
}

func TestU8SliceLen(t *testing.T) {
	ConfirmLength := func(s *U8Slice, i int) {
		if x := s.Len(); x != i {
			t.Fatalf("%v.Len() should be %v but is %v", s, i, x)
		}
	}
	
	ConfirmLength(U8List(0), 1)
	ConfirmLength(U8List(0, 1), 2)
}

func TestU8SliceSwap(t *testing.T) {
	ConfirmSwap := func(s *U8Slice, i, j int, r *U8Slice) {
		if s.Swap(i, j); !r.Equal(s) {
			t.Fatalf("Swap(%v, %v) should be %v but is %v", i, j, r, s)
		}
	}
	ConfirmSwap(U8List(0, 1, 2), 0, 1, U8List(1, 0, 2))
	ConfirmSwap(U8List(0, 1, 2), 0, 2, U8List(2, 1, 0))
}

func TestU8SliceSort(t *testing.T) {
	ConfirmSort := func(s, r *U8Slice) {
		if s.Sort(); !r.Equal(s) {
			t.Fatalf("Sort() should be %v but is %v", r, s)
		}
	}

	ConfirmSort(U8List(3, 2, 1, 4, 5, 0), U8List(0, 1, 2, 3, 4, 5))
}

func TestU8SliceCompare(t *testing.T) {
	ConfirmCompare := func(s *U8Slice, i, j, r int) {
		if x := s.Compare(i, j); x != r {
			t.Fatalf("Compare(%v, %v) should be %v but is %v", i, j, r, x)
		}
	}

	ConfirmCompare(U8List(0, 1), 0, 0, IS_SAME_AS)
	ConfirmCompare(U8List(0, 1), 0, 1, IS_LESS_THAN)
	ConfirmCompare(U8List(0, 1), 1, 0, IS_GREATER_THAN)
}

func TestU8SliceZeroCompare(t *testing.T) {
	ConfirmCompare := func(s *U8Slice, i, r int) {
		if x := s.ZeroCompare(i); x != r {
			t.Fatalf("ZeroCompare(%v) should be %v but is %v", i, r, x)
		}
	}

	ConfirmCompare(U8List(0, 1, 2), 0, IS_SAME_AS)
	ConfirmCompare(U8List(0, 1, 2), 1, IS_LESS_THAN)
	ConfirmCompare(U8List(0, 1, 2), 2, IS_LESS_THAN)
}

func TestU8SliceCut(t *testing.T) {
	ConfirmCut := func(s *U8Slice, start, end int, r *U8Slice) {
		if s.Cut(start, end); !r.Equal(s) {
			t.Fatalf("Cut(%v, %v) should be %v but is %v", start, end, r, s)
		}
	}

	ConfirmCut(U8List(0, 1, 2, 3, 4, 5), 0, 1, U8List(1, 2, 3, 4, 5))
	ConfirmCut(U8List(0, 1, 2, 3, 4, 5), 1, 2, U8List(0, 2, 3, 4, 5))
	ConfirmCut(U8List(0, 1, 2, 3, 4, 5), 2, 3, U8List(0, 1, 3, 4, 5))
	ConfirmCut(U8List(0, 1, 2, 3, 4, 5), 3, 4, U8List(0, 1, 2, 4, 5))
	ConfirmCut(U8List(0, 1, 2, 3, 4, 5), 4, 5, U8List(0, 1, 2, 3, 5))
	ConfirmCut(U8List(0, 1, 2, 3, 4, 5), 5, 6, U8List(0, 1, 2, 3, 4))

	ConfirmCut(U8List(0, 1, 2, 3, 4, 5), -1, 1, U8List(1, 2, 3, 4, 5))
	ConfirmCut(U8List(0, 1, 2, 3, 4, 5), 0, 2, U8List(2, 3, 4, 5))
	ConfirmCut(U8List(0, 1, 2, 3, 4, 5), 1, 3, U8List(0, 3, 4, 5))
	ConfirmCut(U8List(0, 1, 2, 3, 4, 5), 2, 4, U8List(0, 1, 4, 5))
	ConfirmCut(U8List(0, 1, 2, 3, 4, 5), 3, 5, U8List(0, 1, 2, 5))
	ConfirmCut(U8List(0, 1, 2, 3, 4, 5), 4, 6, U8List(0, 1, 2, 3))
	ConfirmCut(U8List(0, 1, 2, 3, 4, 5), 5, 7, U8List(0, 1, 2, 3, 4))
}

func TestU8SliceTrim(t *testing.T) {
	ConfirmTrim := func(s *U8Slice, start, end int, r *U8Slice) {
		if s.Trim(start, end); !r.Equal(s) {
			t.Fatalf("Trim(%v, %v) should be %v but is %v", start, end, r, s)
		}
	}

	ConfirmTrim(U8List(0, 1, 2, 3, 4, 5), 0, 1, U8List(0))
	ConfirmTrim(U8List(0, 1, 2, 3, 4, 5), 1, 2, U8List(1))
	ConfirmTrim(U8List(0, 1, 2, 3, 4, 5), 2, 3, U8List(2))
	ConfirmTrim(U8List(0, 1, 2, 3, 4, 5), 3, 4, U8List(3))
	ConfirmTrim(U8List(0, 1, 2, 3, 4, 5), 4, 5, U8List(4))
	ConfirmTrim(U8List(0, 1, 2, 3, 4, 5), 5, 6, U8List(5))

	ConfirmTrim(U8List(0, 1, 2, 3, 4, 5), -1, 1, U8List(0))
	ConfirmTrim(U8List(0, 1, 2, 3, 4, 5), 0, 2, U8List(0, 1))
	ConfirmTrim(U8List(0, 1, 2, 3, 4, 5), 1, 3, U8List(1, 2))
	ConfirmTrim(U8List(0, 1, 2, 3, 4, 5), 2, 4, U8List(2, 3))
	ConfirmTrim(U8List(0, 1, 2, 3, 4, 5), 3, 5, U8List(3, 4))
	ConfirmTrim(U8List(0, 1, 2, 3, 4, 5), 4, 6, U8List(4, 5))
	ConfirmTrim(U8List(0, 1, 2, 3, 4, 5), 5, 7, U8List(5))
}

func TestU8SliceDelete(t *testing.T) {
	ConfirmDelete := func(s *U8Slice, index int, r *U8Slice) {
		if s.Delete(index); !r.Equal(s) {
			t.Fatalf("Delete(%v) should be %v but is %v", index, r, s)
		}
	}

	ConfirmDelete(U8List(0, 1, 2, 3, 4, 5), -1, U8List(0, 1, 2, 3, 4, 5))
	ConfirmDelete(U8List(0, 1, 2, 3, 4, 5), 0, U8List(1, 2, 3, 4, 5))
	ConfirmDelete(U8List(0, 1, 2, 3, 4, 5), 1, U8List(0, 2, 3, 4, 5))
	ConfirmDelete(U8List(0, 1, 2, 3, 4, 5), 2, U8List(0, 1, 3, 4, 5))
	ConfirmDelete(U8List(0, 1, 2, 3, 4, 5), 3, U8List(0, 1, 2, 4, 5))
	ConfirmDelete(U8List(0, 1, 2, 3, 4, 5), 4, U8List(0, 1, 2, 3, 5))
	ConfirmDelete(U8List(0, 1, 2, 3, 4, 5), 5, U8List(0, 1, 2, 3, 4))
	ConfirmDelete(U8List(0, 1, 2, 3, 4, 5), 6, U8List(0, 1, 2, 3, 4, 5))
}

func TestU8SliceDeleteIf(t *testing.T) {
	ConfirmDeleteIf := func(s *U8Slice, f interface{}, r *U8Slice) {
		if s.DeleteIf(f); !r.Equal(s) {
			t.Fatalf("DeleteIf(%v) should be %v but is %v", f, r, s)
		}
	}

	ConfirmDeleteIf(U8List(0, 1, 0, 3, 0, 5), uint8(0), U8List(1, 3, 5))
	ConfirmDeleteIf(U8List(0, 1, 0, 3, 0, 5), uint8(1), U8List(0, 0, 3, 0, 5))
	ConfirmDeleteIf(U8List(0, 1, 0, 3, 0, 5), uint8(6), U8List(0, 1, 0, 3, 0, 5))

	ConfirmDeleteIf(U8List(0, 1, 0, 3, 0, 5), func(x interface{}) bool { return x == uint8(0) }, U8List(1, 3, 5))
	ConfirmDeleteIf(U8List(0, 1, 0, 3, 0, 5), func(x interface{}) bool { return x == uint8(1) }, U8List(0, 0, 3, 0, 5))
	ConfirmDeleteIf(U8List(0, 1, 0, 3, 0, 5), func(x interface{}) bool { return x == uint8(6) }, U8List(0, 1, 0, 3, 0, 5))

	ConfirmDeleteIf(U8List(0, 1, 0, 3, 0, 5), func(x uint8) bool { return x == uint8(0) }, U8List(1, 3, 5))
	ConfirmDeleteIf(U8List(0, 1, 0, 3, 0, 5), func(x uint8) bool { return x == uint8(1) }, U8List(0, 0, 3, 0, 5))
	ConfirmDeleteIf(U8List(0, 1, 0, 3, 0, 5), func(x uint8) bool { return x == uint8(6) }, U8List(0, 1, 0, 3, 0, 5))
}

func TestU8SliceEach(t *testing.T) {
	var count	uint8
	U8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9).Each(func(i interface{}) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})

	U8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9).Each(func(index int, i interface{}) {
		if i != uint8(index) {
			t.Fatalf("element %v erroneously reported as %v", index, i)
		}
	})

	U8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9).Each(func(key, i interface{}) {
		if i != uint8(key.(int)) {
			t.Fatalf("element %v erroneously reported as %v", key, i)
		}
	})

	count = 0
	U8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9).Each(func(i uint8) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})

	U8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9).Each(func(index int, i uint8) {
		if i != uint8(index) {
			t.Fatalf("element %v erroneously reported as %v", index, i)
		}
	})

	U8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9).Each(func(key interface{}, i uint8) {
		if i != uint8(key.(int)) {
			t.Fatalf("element %v erroneously reported as %v", key, i)
		}
	})
}

func TestU8SliceBlockCopy(t *testing.T) {
	ConfirmBlockCopy := func(s *U8Slice, destination, source, count int, r *U8Slice) {
		s.BlockCopy(destination, source, count)
		if !r.Equal(s) {
			t.Fatalf("BlockCopy(%v, %v, %v) should be %v but is %v", destination, source, count, r, s)
		}
	}

	ConfirmBlockCopy(U8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, 0, 4, U8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockCopy(U8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 9, 9, 4, U8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockCopy(U8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 2, 4, U8List(0, 1, 2, 3, 4, 2, 3, 4, 5, 9))
	ConfirmBlockCopy(U8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 2, 5, 4, U8List(0, 1, 5, 6, 7, 8, 6, 7, 8, 9))
}

func TestU8SliceBlockClear(t *testing.T) {
	ConfirmBlockClear := func(s *U8Slice, start, count int, r *U8Slice) {
		s.BlockClear(start, count)
		if !r.Equal(s) {
			t.Fatalf("BlockClear(%v, %v) should be %v but is %v", start, count, r, s)
		}
	}

	ConfirmBlockClear(U8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, 4, U8List(0, 0, 0, 0, 4, 5, 6, 7, 8, 9))
	ConfirmBlockClear(U8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, 4, U8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockClear(U8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 4, U8List(0, 1, 2, 3, 4, 0, 0, 0, 0, 9))
}

func TestU8SliceOverwrite(t *testing.T) {
	ConfirmOverwrite := func(s *U8Slice, offset int, v, r *U8Slice) {
		s.Overwrite(offset, *v)
		if !r.Equal(s) {
			t.Fatalf("Overwrite(%v, %v) should be %v but is %v", offset, v, r, s)
		}
	}

	ConfirmOverwrite(U8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, U8List(10, 9, 8, 7), U8List(10, 9, 8, 7, 4, 5, 6, 7, 8, 9))
	ConfirmOverwrite(U8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, U8List(10, 9, 8, 7), U8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmOverwrite(U8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, U8List(11, 12, 13, 14), U8List(0, 1, 2, 3, 4, 11, 12, 13, 14, 9))
}

func TestU8SliceReallocate(t *testing.T) {
	ConfirmReallocate := func(s *U8Slice, l, c int, r *U8Slice) {
		o := s.String()
		el := l
		if el > c {
			el = c
		}
		switch s.Reallocate(l, c); {
		case s == nil:				t.Fatalf("%v.Reallocate(%v, %v) created a nil value for Slice", o, l, c)
		case s.Cap() != c:			t.Fatalf("%v.Reallocate(%v, %v) capacity should be %v but is %v", o, l, c, c, s.Cap())
		case s.Len() != el:			t.Fatalf("%v.Reallocate(%v, %v) length should be %v but is %v", o, l, c, el, s.Len())
		case !r.Equal(s):			t.Fatalf("%v.Reallocate(%v, %v) should be %v but is %v", o, l, c, r, s)
		}
	}

	u := make(U8Slice, 0, 10)
	ConfirmReallocate(U8List(), 0, 10, &u)
	ConfirmReallocate(U8List(0, 1, 2, 3, 4), 3, 10, U8List(0, 1, 2))
	ConfirmReallocate(U8List(0, 1, 2, 3, 4), 5, 10, U8List(0, 1, 2, 3, 4))
	ConfirmReallocate(U8List(0, 1, 2, 3, 4), 10, 10, U8List(0, 1, 2, 3, 4, 0, 0, 0, 0, 0))
	ConfirmReallocate(U8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 1, 5, U8List(0))
	ConfirmReallocate(U8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 5, U8List(0, 1, 2, 3, 4))
	ConfirmReallocate(U8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, 5, U8List(0, 1, 2, 3, 4))
}

func TestU8SliceExtend(t *testing.T) {
	ConfirmExtend := func(s *U8Slice, n int, r *U8Slice) {
		c := s.Cap()
		s.Extend(n)
		switch {
		case s.Len() != r.Len():	t.Fatalf("Extend(%v) len should be %v but is %v", n, r.Len(), s.Len())
		case s.Cap() != c + n:		t.Fatalf("Extend(%v) cap should be %v but is %v", n, c + n, s.Cap())
		case !r.Equal(s):			t.Fatalf("Extend(%v) should be %v but is %v", n, r, s)
		}
	}

	ConfirmExtend(U8List(), 1, U8List(0))
	ConfirmExtend(U8List(), 2, U8List(0, 0))
}

func TestU8SliceExpand(t *testing.T) {
	ConfirmExpand := func(s *U8Slice, i, n int, r *U8Slice) {
		c := s.Cap()
		s.Expand(i, n)
		switch {
		case s.Len() != r.Len():	t.Fatalf("Expand(%v, %v) len should be %v but is %v", i, n, r.Len(), s.Len())
		case s.Cap() != c + n:		t.Fatalf("Expand(%v, %v) cap should be %v but is %v", i, n, c + n, s.Cap())
		case !r.Equal(s):			t.Fatalf("Expand(%v, %v) should be %v but is %v", i, n, r, s)
		}
	}

	ConfirmExpand(U8List(), -1, 1, U8List(0))
	ConfirmExpand(U8List(), 0, 1, U8List(0))
	ConfirmExpand(U8List(), 1, 1, U8List(0))
	ConfirmExpand(U8List(), 0, 2, U8List(0, 0))

	ConfirmExpand(U8List(0, 1, 2), -1, 2, U8List(0, 0, 0, 1, 2))
	ConfirmExpand(U8List(0, 1, 2), 0, 2, U8List(0, 0, 0, 1, 2))
	ConfirmExpand(U8List(0, 1, 2), 1, 2, U8List(0, 0, 0, 1, 2))
	ConfirmExpand(U8List(0, 1, 2), 2, 2, U8List(0, 1, 0, 0, 2))
	ConfirmExpand(U8List(0, 1, 2), 3, 2, U8List(0, 1, 2, 0, 0))
	ConfirmExpand(U8List(0, 1, 2), 4, 2, U8List(0, 1, 2, 0, 0))
}

func TestU8SliceDepth(t *testing.T) {
	ConfirmDepth := func(s *U8Slice, i int) {
		if x := s.Depth(); x != i {
			t.Fatalf("%v.Depth() should be %v but is %v", s, i, x)
		}
	}
	ConfirmDepth(U8List(0, 1), 0)
}

func TestU8SliceReverse(t *testing.T) {
	sxp := U8List(1, 2, 3, 4, 5)
	rxp := U8List(5, 4, 3, 2, 1)
	sxp.Reverse()
	if !rxp.Equal(sxp) {
		t.Fatalf("Reversal failed: %v", sxp)
	}
}

func TestU8SliceAppend(t *testing.T) {
	ConfirmAppend := func(s *U8Slice, v interface{}, r *U8Slice) {
		s.Append(v)
		if !r.Equal(s) {
			t.Fatalf("Append(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmAppend(U8List(), uint8(0), U8List(0))

	ConfirmAppend(U8List(), U8List(0), U8List(0))
	ConfirmAppend(U8List(), U8List(0, 1), U8List(0, 1))
	ConfirmAppend(U8List(0, 1, 2), U8List(3, 4), U8List(0, 1, 2, 3, 4))
}

func TestU8SlicePrepend(t *testing.T) {
	ConfirmPrepend := func(s *U8Slice, v interface{}, r *U8Slice) {
		if s.Prepend(v); !r.Equal(s) {
			t.Fatalf("Prepend(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmPrepend(U8List(), uint8(0), U8List(0))
	ConfirmPrepend(U8List(0), uint8(1), U8List(1, 0))

	ConfirmPrepend(U8List(), U8List(0), U8List(0))
	ConfirmPrepend(U8List(), U8List(0, 1), U8List(0, 1))
	ConfirmPrepend(U8List(0, 1, 2), U8List(3, 4), U8List(3, 4, 0, 1, 2))
}

func TestU8SliceRepeat(t *testing.T) {
	ConfirmRepeat := func(s *U8Slice, count int, r *U8Slice) {
		if x := s.Repeat(count); !x.Equal(r) {
			t.Fatalf("%v.Repeat(%v) should be %v but is %v", s, count, r, x)
		}
	}

	ConfirmRepeat(U8List(), 5, U8List())
	ConfirmRepeat(U8List(0), 1, U8List(0))
	ConfirmRepeat(U8List(0), 2, U8List(0, 0))
	ConfirmRepeat(U8List(0), 3, U8List(0, 0, 0))
	ConfirmRepeat(U8List(0), 4, U8List(0, 0, 0, 0))
	ConfirmRepeat(U8List(0), 5, U8List(0, 0, 0, 0, 0))
}

func TestU8SliceCar(t *testing.T) {
	ConfirmCar := func(s *U8Slice, r uint8) {
		n := s.Car().(uint8)
		if ok := n == r; !ok {
			t.Fatalf("head should be '%v' but is '%v'", r, n)
		}
	}
	ConfirmCar(U8List(1, 2, 3), 1)
}

func TestU8SliceCdr(t *testing.T) {
	ConfirmCdr := func(s, r *U8Slice) {
		if n := s.Cdr(); !n.Equal(r) {
			t.Fatalf("tail should be '%v' but is '%v'", r, n)
		}
	}
	ConfirmCdr(U8List(1, 2, 3), U8List(2, 3))
}

func TestU8SliceRplaca(t *testing.T) {
	ConfirmRplaca := func(s *U8Slice, v interface{}, r *U8Slice) {
		if s.Rplaca(v); !s.Equal(r) {
			t.Fatalf("slice should be '%v' but is '%v'", r, s)
		}
	}
	ConfirmRplaca(U8List(1, 2, 3, 4, 5), uint8(0), U8List(0, 2, 3, 4, 5))
}

func TestU8SliceRplacd(t *testing.T) {
	ConfirmRplacd := func(s *U8Slice, v interface{}, r *U8Slice) {
		if s.Rplacd(v); !s.Equal(r) {
			t.Fatalf("slice should be '%v' but is '%v'", r, s)
		}
	}
	ConfirmRplacd(U8List(1, 2, 3, 4, 5), nil, U8List(1))
	ConfirmRplacd(U8List(1, 2, 3, 4, 5), uint8(10), U8List(1, 10))
	ConfirmRplacd(U8List(1, 2, 3, 4, 5), U8List(5, 4, 3, 2), U8List(1, 5, 4, 3, 2))
	ConfirmRplacd(U8List(1, 2, 3, 4, 5, 6), U8List(2, 4, 8, 16), U8List(1, 2, 4, 8, 16))
}

func TestU8SliceSetIntersection(t *testing.T) {
	ConfirmSetIntersection := func(s, o, r *U8Slice) {
		x := s.SetIntersection(*o)
		x.Sort()
		if !r.Equal(x) {
			t.Fatalf("%v.SetIntersection(%v) should be %v but is %v", s, o, r, x)
		}
	}

	ConfirmSetIntersection(U8List(1, 2, 3), U8List(), U8List())
	ConfirmSetIntersection(U8List(1, 2, 3), U8List(1), U8List(1))
	ConfirmSetIntersection(U8List(1, 2, 3), U8List(1, 1), U8List(1))
	ConfirmSetIntersection(U8List(1, 2, 3), U8List(1, 2, 1), U8List(1, 2))
}

func TestU8SliceSetUnion(t *testing.T) {
	ConfirmSetUnion := func(s, o, r *U8Slice) {
		x := s.SetUnion(*o)
		x.Sort()
		if !r.Equal(x) {
			t.Fatalf("%v.SetUnion(%v) should be %v but is %v", s, o, r, x)
		}
	}

	ConfirmSetUnion(U8List(1, 2, 3), U8List(), U8List(1, 2, 3))
	ConfirmSetUnion(U8List(1, 2, 3), U8List(1), U8List(1, 2, 3))
	ConfirmSetUnion(U8List(1, 2, 3), U8List(1, 1), U8List(1, 2, 3))
	ConfirmSetUnion(U8List(1, 2, 3), U8List(1, 2, 1), U8List(1, 2, 3))
}

func TestU8SliceSetDifference(t *testing.T) {
	ConfirmSetUnion := func(s, o, r *U8Slice) {
		x := s.SetDifference(*o)
		x.Sort()
		if !r.Equal(x) {
			t.Fatalf("%v.SetUnion(%v) should be %v but is %v", s, o, r, x)
		}
	}

	ConfirmSetUnion(U8List(1, 2, 3), U8List(), U8List(1, 2, 3))
	ConfirmSetUnion(U8List(1, 2, 3), U8List(1), U8List(2, 3))
	ConfirmSetUnion(U8List(1, 2, 3), U8List(1, 1), U8List(2, 3))
	ConfirmSetUnion(U8List(1, 2, 3), U8List(1, 2, 1), U8List(3))
}

func TestU8SliceFind(t *testing.T) {
	ConfirmFind := func(s *U8Slice, v uint8, i int) {
		if x, ok := s.Find(v); !ok || x != i {
			t.Fatalf("%v.Find(%v) should be %v but is %v", s, v, i, x)
		}
	}

	ConfirmFind(U8List(0, 1, 2, 3, 4), 0, 0)
	ConfirmFind(U8List(0, 1, 2, 3, 4), 1, 1)
	ConfirmFind(U8List(0, 1, 2, 4, 3), 2, 2)
	ConfirmFind(U8List(0, 1, 2, 4, 3), 3, 4)
	ConfirmFind(U8List(0, 1, 2, 4, 3), 4, 3)
}

func TestU8SliceFindN(t *testing.T) {
	ConfirmFindN := func(s *U8Slice, v uint8, n int, i *ISlice) {
		if x := s.FindN(v, n); !x.Equal(i) {
			t.Fatalf("%v.Find(%v, %v) should be %v but is %v", s, v, n, i, x)
		}
	}

	ConfirmFindN(U8List(1, 0, 1, 0, 1), 2, 3, IList())
	ConfirmFindN(U8List(1, 0, 1, 0, 1), 1, 0, IList(0, 2, 4))
	ConfirmFindN(U8List(1, 0, 1, 0, 1), 1, 1, IList(0))
	ConfirmFindN(U8List(1, 0, 1, 0, 1), 1, 2, IList(0, 2))
	ConfirmFindN(U8List(1, 0, 1, 0, 1), 1, 3, IList(0, 2, 4))
	ConfirmFindN(U8List(1, 0, 1, 0, 1), 1, 4, IList(0, 2, 4))
}

func TestU8SliceKeepIf(t *testing.T) {
	ConfirmKeepIf := func(s *U8Slice, f interface{}, r *U8Slice) {
		if s.KeepIf(f); !r.Equal(s) {
			t.Fatalf("KeepIf(%v) should be %v but is %v", f, r, s)
		}
	}

	ConfirmKeepIf(U8List(0, 1, 0, 3, 0, 5), uint8(0), U8List(0, 0, 0))
	ConfirmKeepIf(U8List(0, 1, 0, 3, 0, 5), uint8(1), U8List(1))
	ConfirmKeepIf(U8List(0, 1, 0, 3, 0, 5), uint8(6), U8List())

	ConfirmKeepIf(U8List(0, 1, 0, 3, 0, 5), func(x interface{}) bool { return x == uint8(0) }, U8List(0, 0, 0))
	ConfirmKeepIf(U8List(0, 1, 0, 3, 0, 5), func(x interface{}) bool { return x == uint8(1) }, U8List(1))
	ConfirmKeepIf(U8List(0, 1, 0, 3, 0, 5), func(x interface{}) bool { return x == uint8(6) }, U8List())

	ConfirmKeepIf(U8List(0, 1, 0, 3, 0, 5), func(x uint8) bool { return x == uint8(0) }, U8List(0, 0, 0))
	ConfirmKeepIf(U8List(0, 1, 0, 3, 0, 5), func(x uint8) bool { return x == uint8(1) }, U8List(1))
	ConfirmKeepIf(U8List(0, 1, 0, 3, 0, 5), func(x uint8) bool { return x == uint8(6) }, U8List())
}

func TestU8SliceReverseEach(t *testing.T) {
	var count	uint8
	count = 9
	U8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9).ReverseEach(func(i interface{}) {
		if i != count {
			t.Fatalf("0: element %v erroneously reported as %v", count, i)
		}
		count--
	})

	U8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9).ReverseEach(func(index int, i interface{}) {
		if index != int(i.(uint8)) {
			t.Fatalf("1: element %v erroneously reported as %v", index, i)
		}
	})

	U8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9).ReverseEach(func(key, i interface{}) {
		if uint8(key.(int)) != i {
			t.Fatalf("2: element %v erroneously reported as %v", key, i)
		}
	})

	count = 9
	U8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9).ReverseEach(func(i uint8) {
		if i != count {
			t.Fatalf("3: element %v erroneously reported as %v", count, i)
		}
		count--
	})

	U8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9).ReverseEach(func(index int, i uint8) {
		if int(i) != index {
			t.Fatalf("4: element %v erroneously reported as %v", index, i)
		}
	})

	U8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9).ReverseEach(func(key interface{}, i uint8) {
		if key.(int) != int(i) {
			t.Fatalf("5: element %v erroneously reported as %v", key, i)
		}
	})
}

func TestU8SliceReplaceIf(t *testing.T) {
	ConfirmReplaceIf := func(s *U8Slice, f, v interface{}, r *U8Slice) {
		if s.ReplaceIf(f, v); !r.Equal(s) {
			t.Fatalf("ReplaceIf(%v, %v) should be %v but is %v", f, v, r, s)
		}
	}

	ConfirmReplaceIf(U8List(0, 1, 0, 3, 0, 5), uint8(0), uint8(1), U8List(1, 1, 1, 3, 1, 5))
	ConfirmReplaceIf(U8List(0, 1, 0, 3, 0, 5), uint8(1), uint8(0), U8List(0, 0, 0, 3, 0, 5))
	ConfirmReplaceIf(U8List(0, 1, 0, 3, 0, 5), uint8(6), uint8(0), U8List(0, 1, 0, 3, 0, 5))

	ConfirmReplaceIf(U8List(0, 1, 0, 3, 0, 5), func(x interface{}) bool { return x == uint8(0) }, uint8(1), U8List(1, 1, 1, 3, 1, 5))
	ConfirmReplaceIf(U8List(0, 1, 0, 3, 0, 5), func(x interface{}) bool { return x == uint8(1) }, uint8(0), U8List(0, 0, 0, 3, 0, 5))
	ConfirmReplaceIf(U8List(0, 1, 0, 3, 0, 5), func(x interface{}) bool { return x == uint8(6) }, uint8(0), U8List(0, 1, 0, 3, 0, 5))

	ConfirmReplaceIf(U8List(0, 1, 0, 3, 0, 5), func(x uint8) bool { return x == uint8(0) }, uint8(1), U8List(1, 1, 1, 3, 1, 5))
	ConfirmReplaceIf(U8List(0, 1, 0, 3, 0, 5), func(x uint8) bool { return x == uint8(1) }, uint8(0), U8List(0, 0, 0, 3, 0, 5))
	ConfirmReplaceIf(U8List(0, 1, 0, 3, 0, 5), func(x uint8) bool { return x == uint8(6) }, uint8(0), U8List(0, 1, 0, 3, 0, 5))
}

func TestU8SliceReplace(t *testing.T) {
	ConfirmReplace := func(s *U8Slice, v interface{}) {
		if s.Replace(v); !s.Equal(v) {
			t.Fatalf("Replace() should be %v but is %v", s, v)
		}
	}

	ConfirmReplace(U8List(0, 1, 2, 3, 4, 5), U8List(9, 8, 7, 6, 5))
	ConfirmReplace(U8List(0, 1, 2, 3, 4, 5), U8Slice{ 9, 8, 7, 6, 5 })
	ConfirmReplace(U8List(0, 1, 2, 3, 4, 5), &[]uint8{ 9, 8, 7, 6, 5 })
	ConfirmReplace(U8List(0, 1, 2, 3, 4, 5), []uint8{ 9, 8, 7, 6, 5 })
}

func TestU8SliceSelect(t *testing.T) {
	ConfirmSelect := func(s *U8Slice, f interface{}, r *U8Slice) {
		if x := s.Select(f); !r.Equal(x) {
			t.Fatalf("Select(%v) should be %v but is %v", f, r, s)
		}
	}

	ConfirmSelect(U8List(0, 1, 0, 3, 0, 5), uint8(0), U8List(0, 0, 0))
	ConfirmSelect(U8List(0, 1, 0, 3, 0, 5), uint8(1), U8List(1))
	ConfirmSelect(U8List(0, 1, 0, 3, 0, 5), uint8(6), U8List())

	ConfirmSelect(U8List(0, 1, 0, 3, 0, 5), func(x interface{}) bool { return x == uint8(0) }, U8List(0, 0, 0))
	ConfirmSelect(U8List(0, 1, 0, 3, 0, 5), func(x interface{}) bool { return x == uint8(1) }, U8List(1))
	ConfirmSelect(U8List(0, 1, 0, 3, 0, 5), func(x interface{}) bool { return x == uint8(6) }, U8List())

	ConfirmSelect(U8List(0, 1, 0, 3, 0, 5), func(x uint8) bool { return x == uint8(0) }, U8List(0, 0, 0))
	ConfirmSelect(U8List(0, 1, 0, 3, 0, 5), func(x uint8) bool { return x == uint8(1) }, U8List(1))
	ConfirmSelect(U8List(0, 1, 0, 3, 0, 5), func(x uint8) bool { return x == uint8(6) }, U8List())
}

func TestU8SliceUniq(t *testing.T) {
	ConfirmUniq := func(s, r *U8Slice) {
		if s.Uniq(); !r.Equal(s) {
			t.Fatalf("Uniq() should be %v but is %v", r, s)
		}
	}

	ConfirmUniq(U8List(0, 0, 0, 0, 0, 0), U8List(0))
	ConfirmUniq(U8List(0, 1, 0, 3, 0, 5), U8List(0, 1, 3, 5))
}

func TestU8SliceShuffle(t *testing.T) {
	ConfirmShuffle := func(s, r *U8Slice) {
		if s.Shuffle(); s.Equal(r) {
			t.Fatalf("%v.Shuffle() should change order of elements", s)
		}
		if s.Sort(); !s.Equal(r) {
			t.Fatalf("Shuffle() when sorted should be %v but is %v", r, s)
		}
	}

	ConfirmShuffle(U8List(0, 1, 2, 3, 4, 5), U8List(0, 1, 2, 3, 4, 5))
}

func TestU8SliceValuesAt(t *testing.T) {
	ConfirmValuesAt := func(s *U8Slice, i []int, r *U8Slice) {
		if x := s.ValuesAt(i...); !r.Equal(x) {
			t.Fatalf("%v.ValuesAt(%v) should be %v but is %v", s, i, r, x)
		}
	}

	ConfirmValuesAt(U8List(0, 1, 2, 3, 4, 5), []int{}, U8List())
	ConfirmValuesAt(U8List(0, 1, 2, 3, 4, 5), []int{ 0, 1 }, U8List(0, 1))
	ConfirmValuesAt(U8List(0, 1, 2, 3, 4, 5), []int{ 0, 3 }, U8List(0, 3))
	ConfirmValuesAt(U8List(0, 1, 2, 3, 4, 5), []int{ 0, 3, 4, 3 }, U8List(0, 3, 4, 3))
}

func TestU8SliceInsert(t *testing.T) {
	ConfirmInsert := func(s *U8Slice, n int, v interface{}, r *U8Slice) {
		if s.Insert(n, v); !r.Equal(s) {
			t.Fatalf("Insert(%v, %v) should be %v but is %v", n, v, r, s)
		}
	}

	ConfirmInsert(U8List(), 0, uint8(0), U8List(0))
	ConfirmInsert(U8List(), 0, U8List(0), U8List(0))
	ConfirmInsert(U8List(), 0, U8List(0, 1), U8List(0, 1))

	ConfirmInsert(U8List(0), 0, uint8(1), U8List(1, 0))
	ConfirmInsert(U8List(0), 0, U8List(1), U8List(1, 0))
	ConfirmInsert(U8List(0), 1, uint8(1), U8List(0, 1))
	ConfirmInsert(U8List(0), 1, U8List(1), U8List(0, 1))

	ConfirmInsert(U8List(0, 1, 2), 0, uint8(3), U8List(3, 0, 1, 2))
	ConfirmInsert(U8List(0, 1, 2), 1, uint8(3), U8List(0, 3, 1, 2))
	ConfirmInsert(U8List(0, 1, 2), 2, uint8(3), U8List(0, 1, 3, 2))
	ConfirmInsert(U8List(0, 1, 2), 3, uint8(3), U8List(0, 1, 2, 3))

	ConfirmInsert(U8List(0, 1, 2), 0, U8List(3, 4), U8List(3, 4, 0, 1, 2))
	ConfirmInsert(U8List(0, 1, 2), 1, U8List(3, 4), U8List(0, 3, 4, 1, 2))
	ConfirmInsert(U8List(0, 1, 2), 2, U8List(3, 4), U8List(0, 1, 3, 4, 2))
	ConfirmInsert(U8List(0, 1, 2), 3, U8List(3, 4), U8List(0, 1, 2, 3, 4))
}