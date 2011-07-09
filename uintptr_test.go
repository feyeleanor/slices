package slices

import "testing"

func TestASliceString(t *testing.T) {
	ConfirmString := func(s *ASlice, r string) {
		if x := s.String(); x != r {
			t.Fatalf("%v erroneously serialised as '%v'", r, x)
		}
	}

	ConfirmString(AList(), "()")
	ConfirmString(AList(0), "(0)")
	ConfirmString(AList(0, 1), "(0 1)")
}

func TestASliceLen(t *testing.T) {
	ConfirmLength := func(s *ASlice, i int) {
		if x := s.Len(); x != i {
			t.Fatalf("%v.Len() should be %v but is %v", s, i, x)
		}
	}
	
	ConfirmLength(AList(0), 1)
	ConfirmLength(AList(0, 1), 2)
}

func TestASliceSwap(t *testing.T) {
	ConfirmSwap := func(s *ASlice, i, j int, r *ASlice) {
		if s.Swap(i, j); !r.Equal(s) {
			t.Fatalf("Swap(%v, %v) should be %v but is %v", i, j, r, s)
		}
	}
	ConfirmSwap(AList(0, 1, 2), 0, 1, AList(1, 0, 2))
	ConfirmSwap(AList(0, 1, 2), 0, 2, AList(2, 1, 0))
}

func TestASliceSort(t *testing.T) {
	ConfirmSort := func(s, r *ASlice) {
		if s.Sort(); !r.Equal(s) {
			t.Fatalf("Sort() should be %v but is %v", r, s)
		}
	}

	ConfirmSort(AList(3, 2, 1, 4, 5, 0), AList(0, 1, 2, 3, 4, 5))
}

func TestASliceCompare(t *testing.T) {
	ConfirmCompare := func(s *ASlice, i, j, r int) {
		if x := s.Compare(i, j); x != r {
			t.Fatalf("Compare(%v, %v) should be %v but is %v", i, j, r, x)
		}
	}

	ConfirmCompare(AList(0, 1), 0, 0, IS_SAME_AS)
	ConfirmCompare(AList(0, 1), 0, 1, IS_LESS_THAN)
	ConfirmCompare(AList(0, 1), 1, 0, IS_GREATER_THAN)
}

func TestASliceZeroCompare(t *testing.T) {
	ConfirmCompare := func(s *ASlice, i, r int) {
		if x := s.ZeroCompare(i); x != r {
			t.Fatalf("ZeroCompare(%v) should be %v but is %v", i, r, x)
		}
	}

	ConfirmCompare(AList(1, 0, 2), 0, IS_LESS_THAN)
	ConfirmCompare(AList(1, 0, 2), 1, IS_SAME_AS)
	ConfirmCompare(AList(1, 0, 3), 2, IS_LESS_THAN)
}

func TestASliceCut(t *testing.T) {
	ConfirmCut := func(s *ASlice, start, end int, r *ASlice) {
		if s.Cut(start, end); !r.Equal(s) {
			t.Fatalf("Cut(%v, %v) should be %v but is %v", start, end, r, s)
		}
	}

	ConfirmCut(AList(0, 1, 2, 3, 4, 5), 0, 1, AList(1, 2, 3, 4, 5))
	ConfirmCut(AList(0, 1, 2, 3, 4, 5), 1, 2, AList(0, 2, 3, 4, 5))
	ConfirmCut(AList(0, 1, 2, 3, 4, 5), 2, 3, AList(0, 1, 3, 4, 5))
	ConfirmCut(AList(0, 1, 2, 3, 4, 5), 3, 4, AList(0, 1, 2, 4, 5))
	ConfirmCut(AList(0, 1, 2, 3, 4, 5), 4, 5, AList(0, 1, 2, 3, 5))
	ConfirmCut(AList(0, 1, 2, 3, 4, 5), 5, 6, AList(0, 1, 2, 3, 4))

	ConfirmCut(AList(0, 1, 2, 3, 4, 5), -1, 1, AList(1, 2, 3, 4, 5))
	ConfirmCut(AList(0, 1, 2, 3, 4, 5), 0, 2, AList(2, 3, 4, 5))
	ConfirmCut(AList(0, 1, 2, 3, 4, 5), 1, 3, AList(0, 3, 4, 5))
	ConfirmCut(AList(0, 1, 2, 3, 4, 5), 2, 4, AList(0, 1, 4, 5))
	ConfirmCut(AList(0, 1, 2, 3, 4, 5), 3, 5, AList(0, 1, 2, 5))
	ConfirmCut(AList(0, 1, 2, 3, 4, 5), 4, 6, AList(0, 1, 2, 3))
	ConfirmCut(AList(0, 1, 2, 3, 4, 5), 5, 7, AList(0, 1, 2, 3, 4))
}

func TestASliceTrim(t *testing.T) {
	ConfirmTrim := func(s *ASlice, start, end int, r *ASlice) {
		if s.Trim(start, end); !r.Equal(s) {
			t.Fatalf("Trim(%v, %v) should be %v but is %v", start, end, r, s)
		}
	}

	ConfirmTrim(AList(0, 1, 2, 3, 4, 5), 0, 1, AList(0))
	ConfirmTrim(AList(0, 1, 2, 3, 4, 5), 1, 2, AList(1))
	ConfirmTrim(AList(0, 1, 2, 3, 4, 5), 2, 3, AList(2))
	ConfirmTrim(AList(0, 1, 2, 3, 4, 5), 3, 4, AList(3))
	ConfirmTrim(AList(0, 1, 2, 3, 4, 5), 4, 5, AList(4))
	ConfirmTrim(AList(0, 1, 2, 3, 4, 5), 5, 6, AList(5))

	ConfirmTrim(AList(0, 1, 2, 3, 4, 5), -1, 1, AList(0))
	ConfirmTrim(AList(0, 1, 2, 3, 4, 5), 0, 2, AList(0, 1))
	ConfirmTrim(AList(0, 1, 2, 3, 4, 5), 1, 3, AList(1, 2))
	ConfirmTrim(AList(0, 1, 2, 3, 4, 5), 2, 4, AList(2, 3))
	ConfirmTrim(AList(0, 1, 2, 3, 4, 5), 3, 5, AList(3, 4))
	ConfirmTrim(AList(0, 1, 2, 3, 4, 5), 4, 6, AList(4, 5))
	ConfirmTrim(AList(0, 1, 2, 3, 4, 5), 5, 7, AList(5))
}

func TestASliceDelete(t *testing.T) {
	ConfirmDelete := func(s *ASlice, index int, r *ASlice) {
		if s.Delete(index); !r.Equal(s) {
			t.Fatalf("Delete(%v) should be %v but is %v", index, r, s)
		}
	}

	ConfirmDelete(AList(0, 1, 2, 3, 4, 5), -1, AList(0, 1, 2, 3, 4, 5))
	ConfirmDelete(AList(0, 1, 2, 3, 4, 5), 0, AList(1, 2, 3, 4, 5))
	ConfirmDelete(AList(0, 1, 2, 3, 4, 5), 1, AList(0, 2, 3, 4, 5))
	ConfirmDelete(AList(0, 1, 2, 3, 4, 5), 2, AList(0, 1, 3, 4, 5))
	ConfirmDelete(AList(0, 1, 2, 3, 4, 5), 3, AList(0, 1, 2, 4, 5))
	ConfirmDelete(AList(0, 1, 2, 3, 4, 5), 4, AList(0, 1, 2, 3, 5))
	ConfirmDelete(AList(0, 1, 2, 3, 4, 5), 5, AList(0, 1, 2, 3, 4))
	ConfirmDelete(AList(0, 1, 2, 3, 4, 5), 6, AList(0, 1, 2, 3, 4, 5))
}

func TestASliceDeleteIf(t *testing.T) {
	ConfirmDeleteIf := func(s *ASlice, f interface{}, r *ASlice) {
		if s.DeleteIf(f); !r.Equal(s) {
			t.Fatalf("DeleteIf(%v) should be %v but is %v", f, r, s)
		}
	}

	ConfirmDeleteIf(AList(0, 1, 0, 3, 0, 5), uintptr(0), AList(1, 3, 5))
	ConfirmDeleteIf(AList(0, 1, 0, 3, 0, 5), uintptr(1), AList(0, 0, 3, 0, 5))
	ConfirmDeleteIf(AList(0, 1, 0, 3, 0, 5), uintptr(6), AList(0, 1, 0, 3, 0, 5))

	ConfirmDeleteIf(AList(0, 1, 0, 3, 0, 5), func(x interface{}) bool { return x == uintptr(0) }, AList(1, 3, 5))
	ConfirmDeleteIf(AList(0, 1, 0, 3, 0, 5), func(x interface{}) bool { return x == uintptr(1) }, AList(0, 0, 3, 0, 5))
	ConfirmDeleteIf(AList(0, 1, 0, 3, 0, 5), func(x interface{}) bool { return x == uintptr(6) }, AList(0, 1, 0, 3, 0, 5))

	ConfirmDeleteIf(AList(0, 1, 0, 3, 0, 5), func(x uintptr) bool { return x == uintptr(0) }, AList(1, 3, 5))
	ConfirmDeleteIf(AList(0, 1, 0, 3, 0, 5), func(x uintptr) bool { return x == uintptr(1) }, AList(0, 0, 3, 0, 5))
	ConfirmDeleteIf(AList(0, 1, 0, 3, 0, 5), func(x uintptr) bool { return x == uintptr(6) }, AList(0, 1, 0, 3, 0, 5))
}

func TestASliceEach(t *testing.T) {
	var	count	uintptr
	AList(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).Each(func(i interface{}) {
		if i != uintptr(count) {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})

	AList(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).Each(func(index int, i interface{}) {
		if i != uintptr(index) {
			t.Fatalf("element %v erroneously reported as %v", index, i)
		}
	})

	AList(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).Each(func(key, i interface{}) {
		if i != uintptr(key.(int)) {
			t.Fatalf("element %v erroneously reported as %v", key, i)
		}
	})

	count = 0
	AList(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).Each(func(i uintptr) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})

	AList(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).Each(func(index int, i uintptr) {
		if i != uintptr(index) {
			t.Fatalf("element %v erroneously reported as %v", index, i)
		}
	})

	AList(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).Each(func(key interface{}, i uintptr) {
		if i != uintptr(key.(int)) {
			t.Fatalf("element %v erroneously reported as %v", key, i)
		}
	})
}

func TestASliceBlockCopy(t *testing.T) {
	ConfirmBlockCopy := func(s *ASlice, destination, source, count int, r *ASlice) {
		s.BlockCopy(destination, source, count)
		if !r.Equal(s) {
			t.Fatalf("BlockCopy(%v, %v, %v) should be %v but is %v", destination, source, count, r, s)
		}
	}

	ConfirmBlockCopy(AList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, 0, 4, AList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockCopy(AList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 9, 9, 4, AList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockCopy(AList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 2, 4, AList(0, 1, 2, 3, 4, 2, 3, 4, 5, 9))
	ConfirmBlockCopy(AList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 2, 5, 4, AList(0, 1, 5, 6, 7, 8, 6, 7, 8, 9))
}

func TestASliceBlockClear(t *testing.T) {
	ConfirmBlockClear := func(s *ASlice, start, count int, r *ASlice) {
		s.BlockClear(start, count)
		if !r.Equal(s) {
			t.Fatalf("BlockClear(%v, %v) should be %v but is %v", start, count, r, s)
		}
	}

	ConfirmBlockClear(AList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, 4, AList(0, 0, 0, 0, 4, 5, 6, 7, 8, 9))
	ConfirmBlockClear(AList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, 4, AList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockClear(AList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 4, AList(0, 1, 2, 3, 4, 0, 0, 0, 0, 9))
}

func TestASliceOverwrite(t *testing.T) {
	ConfirmOverwrite := func(s *ASlice, offset int, v, r *ASlice) {
		s.Overwrite(offset, *v)
		if !r.Equal(s) {
			t.Fatalf("Overwrite(%v, %v) should be %v but is %v", offset, v, r, s)
		}
	}

	ConfirmOverwrite(AList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, AList(10, 9, 8, 7), AList(10, 9, 8, 7, 4, 5, 6, 7, 8, 9))
	ConfirmOverwrite(AList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, AList(10, 9, 8, 7), AList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmOverwrite(AList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, AList(11, 12, 13, 14), AList(0, 1, 2, 3, 4, 11, 12, 13, 14, 9))
}

func TestASliceReallocate(t *testing.T) {
	ConfirmReallocate := func(s *ASlice, l, c int, r *ASlice) {
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

	ConfirmReallocate(AList(), 0, 10, AList())
	ConfirmReallocate(AList(0, 1, 2, 3, 4), 3, 10, AList(0, 1, 2))
	ConfirmReallocate(AList(0, 1, 2, 3, 4), 5, 10, AList(0, 1, 2, 3, 4))
	ConfirmReallocate(AList(0, 1, 2, 3, 4), 10, 10, AList(0, 1, 2, 3, 4, 0, 0, 0, 0, 0))
	ConfirmReallocate(AList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 1, 5, AList(0))
	ConfirmReallocate(AList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 5, AList(0, 1, 2, 3, 4))
	ConfirmReallocate(AList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, 5, AList(0, 1, 2, 3, 4))
}

func TestASliceExtend(t *testing.T) {
	ConfirmExtend := func(s *ASlice, n int, r *ASlice) {
		c := s.Cap()
		s.Extend(n)
		switch {
		case s.Len() != r.Len():	t.Fatalf("Extend(%v) len should be %v but is %v", n, r.Len(), s.Len())
		case s.Cap() != c + n:		t.Fatalf("Extend(%v) cap should be %v but is %v", n, c + n, s.Cap())
		case !r.Equal(s):			t.Fatalf("Extend(%v) should be %v but is %v", n, r, s)
		}
	}

	ConfirmExtend(AList(), 1, AList(0))
	ConfirmExtend(AList(), 2, AList(0, 0))
}

func TestASliceExpand(t *testing.T) {
	ConfirmExpand := func(s *ASlice, i, n int, r *ASlice) {
		c := s.Cap()
		s.Expand(i, n)
		switch {
		case s.Len() != r.Len():	t.Fatalf("Expand(%v, %v) len should be %v but is %v", i, n, r.Len(), s.Len())
		case s.Cap() != c + n:		t.Fatalf("Expand(%v, %v) cap should be %v but is %v", i, n, c + n, s.Cap())
		case !r.Equal(s):			t.Fatalf("Expand(%v, %v) should be %v but is %v", i, n, r, s)
		}
	}

	ConfirmExpand(AList(), -1, 1, AList(0))
	ConfirmExpand(AList(), 0, 1, AList(0))
	ConfirmExpand(AList(), 1, 1, AList(0))
	ConfirmExpand(AList(), 0, 2, AList(0, 0))

	ConfirmExpand(AList(0, 1, 2), -1, 2, AList(0, 0, 0, 1, 2))
	ConfirmExpand(AList(0, 1, 2), 0, 2, AList(0, 0, 0, 1, 2))
	ConfirmExpand(AList(0, 1, 2), 1, 2, AList(0, 0, 0, 1, 2))
	ConfirmExpand(AList(0, 1, 2), 2, 2, AList(0, 1, 0, 0, 2))
	ConfirmExpand(AList(0, 1, 2), 3, 2, AList(0, 1, 2, 0, 0))
	ConfirmExpand(AList(0, 1, 2), 4, 2, AList(0, 1, 2, 0, 0))
}

func TestASliceDepth(t *testing.T) {
	ConfirmDepth := func(s *ASlice, i int) {
		if x := s.Depth(); x != i {
			t.Fatalf("%v.Depth() should be %v but is %v", s, i, x)
		}
	}
	ConfirmDepth(AList(0, 1), 0)
}

func TestASliceReverse(t *testing.T) {
	sxp := AList(1, 2, 3, 4, 5)
	rxp := AList(5, 4, 3, 2, 1)
	sxp.Reverse()
	if !rxp.Equal(sxp) {
		t.Fatalf("Reversal failed: %v", sxp)
	}
}

func TestASliceAppend(t *testing.T) {
	ConfirmAppend := func(s *ASlice, v interface{}, r *ASlice) {
		s.Append(v)
		if !r.Equal(s) {
			t.Fatalf("Append(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmAppend(AList(), uintptr(0), AList(0))

	ConfirmAppend(AList(), AList(0), AList(0))
	ConfirmAppend(AList(), AList(0, 1), AList(0, 1))
	ConfirmAppend(AList(0, 1, 2), AList(3, 4), AList(0, 1, 2, 3, 4))
}

func TestASlicePrepend(t *testing.T) {
	ConfirmPrepend := func(s *ASlice, v interface{}, r *ASlice) {
		if s.Prepend(v); !r.Equal(s) {
			t.Fatalf("Prepend(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmPrepend(AList(), uintptr(0), AList(0))
	ConfirmPrepend(AList(0), uintptr(1), AList(1, 0))

	ConfirmPrepend(AList(), AList(0), AList(0))
	ConfirmPrepend(AList(), AList(0, 1), AList(0, 1))
	ConfirmPrepend(AList(0, 1, 2), AList(3, 4), AList(3, 4, 0, 1, 2))
}

func TestASliceRepeat(t *testing.T) {
	ConfirmRepeat := func(s *ASlice, count int, r *ASlice) {
		if x := s.Repeat(count); !x.Equal(r) {
			t.Fatalf("%v.Repeat(%v) should be %v but is %v", s, count, r, x)
		}
	}

	ConfirmRepeat(AList(), 5, AList())
	ConfirmRepeat(AList(0), 1, AList(0))
	ConfirmRepeat(AList(0), 2, AList(0, 0))
	ConfirmRepeat(AList(0), 3, AList(0, 0, 0))
	ConfirmRepeat(AList(0), 4, AList(0, 0, 0, 0))
	ConfirmRepeat(AList(0), 5, AList(0, 0, 0, 0, 0))
}

func TestASliceCar(t *testing.T) {
	ConfirmCar := func(s *ASlice, r uintptr) {
		n := s.Car()
		if ok := n == r; !ok {
			t.Fatalf("head should be '%v' but is '%v'", r, n)
		}
	}
	ConfirmCar(AList(1, 2, 3), 1)
}

func TestASliceCdr(t *testing.T) {
	ConfirmCdr := func(s, r *ASlice) {
		if n := s.Cdr(); !n.Equal(r) {
			t.Fatalf("tail should be '%v' but is '%v'", r, n)
		}
	}
	ConfirmCdr(AList(1, 2, 3), AList(2, 3))
}

func TestASliceRplaca(t *testing.T) {
	ConfirmRplaca := func(s *ASlice, v interface{}, r *ASlice) {
		if s.Rplaca(v); !s.Equal(r) {
			t.Fatalf("slice should be '%v' but is '%v'", r, s)
		}
	}
	ConfirmRplaca(AList(1, 2, 3, 4, 5), uintptr(0), AList(0, 2, 3, 4, 5))
}

func TestASliceRplacd(t *testing.T) {
	ConfirmRplacd := func(s *ASlice, v interface{}, r *ASlice) {
		if s.Rplacd(v); !s.Equal(r) {
			t.Fatalf("slice should be '%v' but is '%v'", r, s)
		}
	}
	ConfirmRplacd(AList(1, 2, 3, 4, 5), nil, AList(1))
	ConfirmRplacd(AList(1, 2, 3, 4, 5), uintptr(10), AList(1, 10))
	ConfirmRplacd(AList(1, 2, 3, 4, 5), AList(5, 4, 3, 2), AList(1, 5, 4, 3, 2))
	ConfirmRplacd(AList(1, 2, 3, 4, 5, 6), AList(2, 4, 8, 16), AList(1, 2, 4, 8, 16))
}

func TestASliceSetIntersection(t *testing.T) {
	ConfirmSetIntersection := func(s, o, r *ASlice) {
		x := s.SetIntersection(*o)
		x.Sort()
		if !r.Equal(x) {
			t.Fatalf("%v.SetIntersection(%v) should be %v but is %v", s, o, r, x)
		}
	}

	ConfirmSetIntersection(AList(1, 2, 3), AList(), AList())
	ConfirmSetIntersection(AList(1, 2, 3), AList(1), AList(1))
	ConfirmSetIntersection(AList(1, 2, 3), AList(1, 1), AList(1))
	ConfirmSetIntersection(AList(1, 2, 3), AList(1, 2, 1), AList(1, 2))
}

func TestASliceSetUnion(t *testing.T) {
	ConfirmSetUnion := func(s, o, r *ASlice) {
		x := s.SetUnion(*o)
		x.Sort()
		if !r.Equal(x) {
			t.Fatalf("%v.SetUnion(%v) should be %v but is %v", s, o, r, x)
		}
	}

	ConfirmSetUnion(AList(1, 2, 3), AList(), AList(1, 2, 3))
	ConfirmSetUnion(AList(1, 2, 3), AList(1), AList(1, 2, 3))
	ConfirmSetUnion(AList(1, 2, 3), AList(1, 1), AList(1, 2, 3))
	ConfirmSetUnion(AList(1, 2, 3), AList(1, 2, 1), AList(1, 2, 3))
}

func TestASliceSetDifference(t *testing.T) {
	ConfirmSetUnion := func(s, o, r *ASlice) {
		x := s.SetDifference(*o)
		x.Sort()
		if !r.Equal(x) {
			t.Fatalf("%v.SetUnion(%v) should be %v but is %v", s, o, r, x)
		}
	}

	ConfirmSetUnion(AList(1, 2, 3), AList(), AList(1, 2, 3))
	ConfirmSetUnion(AList(1, 2, 3), AList(1), AList(2, 3))
	ConfirmSetUnion(AList(1, 2, 3), AList(1, 1), AList(2, 3))
	ConfirmSetUnion(AList(1, 2, 3), AList(1, 2, 1), AList(3))
}

func TestASliceFind(t *testing.T) {
	ConfirmFind := func(s *ASlice, v uintptr, i int) {
		if x, ok := s.Find(v); !ok || x != i {
			t.Fatalf("%v.Find(%v) should be %v but is %v", s, v, i, x)
		}
	}

	ConfirmFind(AList(0, 1, 2, 3, 4), 0, 0)
	ConfirmFind(AList(0, 1, 2, 3, 4), 1, 1)
	ConfirmFind(AList(0, 1, 2, 4, 3), 2, 2)
	ConfirmFind(AList(0, 1, 2, 4, 3), 3, 4)
	ConfirmFind(AList(0, 1, 2, 4, 3), 4, 3)
}

func TestASliceFindN(t *testing.T) {
	ConfirmFindN := func(s *ASlice, v uintptr, n int, i *ISlice) {
		if x := s.FindN(v, n); !x.Equal(i) {
			t.Fatalf("%v.Find(%v, %v) should be %v but is %v", s, v, n, i, x)
		}
	}

	ConfirmFindN(AList(1, 0, 1, 0, 1), 2, 3, IList())
	ConfirmFindN(AList(1, 0, 1, 0, 1), 1, 0, IList(0, 2, 4))
	ConfirmFindN(AList(1, 0, 1, 0, 1), 1, 1, IList(0))
	ConfirmFindN(AList(1, 0, 1, 0, 1), 1, 2, IList(0, 2))
	ConfirmFindN(AList(1, 0, 1, 0, 1), 1, 3, IList(0, 2, 4))
	ConfirmFindN(AList(1, 0, 1, 0, 1), 1, 4, IList(0, 2, 4))
}