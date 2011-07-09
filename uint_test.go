package slices

import "testing"

func TestUSliceString(t *testing.T) {
	ConfirmString := func(s *USlice, r string) {
		if x := s.String(); x != r {
			t.Fatalf("%v erroneously serialised as '%v'", r, x)
		}
	}

	ConfirmString(UList(), "()")
	ConfirmString(UList(0), "(0)")
	ConfirmString(UList(0, 1), "(0 1)")
}

func TestUSliceLen(t *testing.T) {
	ConfirmLength := func(s *USlice, i int) {
		if x := s.Len(); x != i {
			t.Fatalf("%v.Len() should be %v but is %v", s, i, x)
		}
	}
	
	ConfirmLength(UList(0), 1)
	ConfirmLength(UList(0, 1), 2)
}

func TestUSliceSwap(t *testing.T) {
	ConfirmSwap := func(s *USlice, i, j int, r *USlice) {
		if s.Swap(i, j); !r.Equal(s) {
			t.Fatalf("Swap(%v, %v) should be %v but is %v", i, j, r, s)
		}
	}
	ConfirmSwap(UList(0, 1, 2), 0, 1, UList(1, 0, 2))
	ConfirmSwap(UList(0, 1, 2), 0, 2, UList(2, 1, 0))
}

func TestUSliceSort(t *testing.T) {
	ConfirmSort := func(s, r *USlice) {
		if s.Sort(); !r.Equal(s) {
			t.Fatalf("Sort() should be %v but is %v", r, s)
		}
	}

	ConfirmSort(UList(3, 2, 1, 4, 5, 0), UList(0, 1, 2, 3, 4, 5))
}

func TestUSliceCompare(t *testing.T) {
	ConfirmCompare := func(s *USlice, i, j, r int) {
		if x := s.Compare(i, j); x != r {
			t.Fatalf("Compare(%v, %v) should be %v but is %v", i, j, r, x)
		}
	}

	ConfirmCompare(UList(0, 1), 0, 0, IS_SAME_AS)
	ConfirmCompare(UList(0, 1), 0, 1, IS_LESS_THAN)
	ConfirmCompare(UList(0, 1), 1, 0, IS_GREATER_THAN)
}

func TestUSliceZeroCompare(t *testing.T) {
	ConfirmCompare := func(s *USlice, i, r int) {
		if x := s.ZeroCompare(i); x != r {
			t.Fatalf("ZeroCompare(%v) should be %v but is %v", i, r, x)
		}
	}

	ConfirmCompare(UList(1, 0, 2), 0, IS_LESS_THAN)
	ConfirmCompare(UList(1, 0, 2), 1, IS_SAME_AS)
	ConfirmCompare(UList(1, 0, 3), 2, IS_LESS_THAN)
}

func TestUSliceCut(t *testing.T) {
	ConfirmCut := func(s *USlice, start, end int, r *USlice) {
		if s.Cut(start, end); !r.Equal(s) {
			t.Fatalf("Cut(%v, %v) should be %v but is %v", start, end, r, s)
		}
	}

	ConfirmCut(UList(0, 1, 2, 3, 4, 5), 0, 1, UList(1, 2, 3, 4, 5))
	ConfirmCut(UList(0, 1, 2, 3, 4, 5), 1, 2, UList(0, 2, 3, 4, 5))
	ConfirmCut(UList(0, 1, 2, 3, 4, 5), 2, 3, UList(0, 1, 3, 4, 5))
	ConfirmCut(UList(0, 1, 2, 3, 4, 5), 3, 4, UList(0, 1, 2, 4, 5))
	ConfirmCut(UList(0, 1, 2, 3, 4, 5), 4, 5, UList(0, 1, 2, 3, 5))
	ConfirmCut(UList(0, 1, 2, 3, 4, 5), 5, 6, UList(0, 1, 2, 3, 4))

	ConfirmCut(UList(0, 1, 2, 3, 4, 5), -1, 1, UList(1, 2, 3, 4, 5))
	ConfirmCut(UList(0, 1, 2, 3, 4, 5), 0, 2, UList(2, 3, 4, 5))
	ConfirmCut(UList(0, 1, 2, 3, 4, 5), 1, 3, UList(0, 3, 4, 5))
	ConfirmCut(UList(0, 1, 2, 3, 4, 5), 2, 4, UList(0, 1, 4, 5))
	ConfirmCut(UList(0, 1, 2, 3, 4, 5), 3, 5, UList(0, 1, 2, 5))
	ConfirmCut(UList(0, 1, 2, 3, 4, 5), 4, 6, UList(0, 1, 2, 3))
	ConfirmCut(UList(0, 1, 2, 3, 4, 5), 5, 7, UList(0, 1, 2, 3, 4))
}

func TestUSliceTrim(t *testing.T) {
	ConfirmTrim := func(s *USlice, start, end int, r *USlice) {
		if s.Trim(start, end); !r.Equal(s) {
			t.Fatalf("Trim(%v, %v) should be %v but is %v", start, end, r, s)
		}
	}

	ConfirmTrim(UList(0, 1, 2, 3, 4, 5), 0, 1, UList(0))
	ConfirmTrim(UList(0, 1, 2, 3, 4, 5), 1, 2, UList(1))
	ConfirmTrim(UList(0, 1, 2, 3, 4, 5), 2, 3, UList(2))
	ConfirmTrim(UList(0, 1, 2, 3, 4, 5), 3, 4, UList(3))
	ConfirmTrim(UList(0, 1, 2, 3, 4, 5), 4, 5, UList(4))
	ConfirmTrim(UList(0, 1, 2, 3, 4, 5), 5, 6, UList(5))

	ConfirmTrim(UList(0, 1, 2, 3, 4, 5), -1, 1, UList(0))
	ConfirmTrim(UList(0, 1, 2, 3, 4, 5), 0, 2, UList(0, 1))
	ConfirmTrim(UList(0, 1, 2, 3, 4, 5), 1, 3, UList(1, 2))
	ConfirmTrim(UList(0, 1, 2, 3, 4, 5), 2, 4, UList(2, 3))
	ConfirmTrim(UList(0, 1, 2, 3, 4, 5), 3, 5, UList(3, 4))
	ConfirmTrim(UList(0, 1, 2, 3, 4, 5), 4, 6, UList(4, 5))
	ConfirmTrim(UList(0, 1, 2, 3, 4, 5), 5, 7, UList(5))
}

func TestUSliceDelete(t *testing.T) {
	ConfirmDelete := func(s *USlice, index int, r *USlice) {
		if s.Delete(index); !r.Equal(s) {
			t.Fatalf("Delete(%v) should be %v but is %v", index, r, s)
		}
	}

	ConfirmDelete(UList(0, 1, 2, 3, 4, 5), -1, UList(0, 1, 2, 3, 4, 5))
	ConfirmDelete(UList(0, 1, 2, 3, 4, 5), 0, UList(1, 2, 3, 4, 5))
	ConfirmDelete(UList(0, 1, 2, 3, 4, 5), 1, UList(0, 2, 3, 4, 5))
	ConfirmDelete(UList(0, 1, 2, 3, 4, 5), 2, UList(0, 1, 3, 4, 5))
	ConfirmDelete(UList(0, 1, 2, 3, 4, 5), 3, UList(0, 1, 2, 4, 5))
	ConfirmDelete(UList(0, 1, 2, 3, 4, 5), 4, UList(0, 1, 2, 3, 5))
	ConfirmDelete(UList(0, 1, 2, 3, 4, 5), 5, UList(0, 1, 2, 3, 4))
	ConfirmDelete(UList(0, 1, 2, 3, 4, 5), 6, UList(0, 1, 2, 3, 4, 5))
}

func TestUSliceDeleteIf(t *testing.T) {
	ConfirmDeleteIf := func(s *USlice, f interface{}, r *USlice) {
		if s.DeleteIf(f); !r.Equal(s) {
			t.Fatalf("DeleteIf(%v) should be %v but is %v", f, r, s)
		}
	}

	ConfirmDeleteIf(UList(0, 1, 0, 3, 0, 5), uint(0), UList(1, 3, 5))
	ConfirmDeleteIf(UList(0, 1, 0, 3, 0, 5), uint(1), UList(0, 0, 3, 0, 5))
	ConfirmDeleteIf(UList(0, 1, 0, 3, 0, 5), uint(6), UList(0, 1, 0, 3, 0, 5))

	ConfirmDeleteIf(UList(0, 1, 0, 3, 0, 5), func(x interface{}) bool { return x == uint(0) }, UList(1, 3, 5))
	ConfirmDeleteIf(UList(0, 1, 0, 3, 0, 5), func(x interface{}) bool { return x == uint(1) }, UList(0, 0, 3, 0, 5))
	ConfirmDeleteIf(UList(0, 1, 0, 3, 0, 5), func(x interface{}) bool { return x == uint(6) }, UList(0, 1, 0, 3, 0, 5))

	ConfirmDeleteIf(UList(0, 1, 0, 3, 0, 5), func(x uint) bool { return x == uint(0) }, UList(1, 3, 5))
	ConfirmDeleteIf(UList(0, 1, 0, 3, 0, 5), func(x uint) bool { return x == uint(1) }, UList(0, 0, 3, 0, 5))
	ConfirmDeleteIf(UList(0, 1, 0, 3, 0, 5), func(x uint) bool { return x == uint(6) }, UList(0, 1, 0, 3, 0, 5))
}

func TestUSliceEach(t *testing.T) {
	var	count	uint
	UList(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).Each(func(i interface{}) {
		if i != uint(count) {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})

	UList(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).Each(func(index int, i interface{}) {
		if i != uint(index) {
			t.Fatalf("element %v erroneously reported as %v", index, i)
		}
	})

	UList(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).Each(func(key, i interface{}) {
		if i != uint(key.(int)) {
			t.Fatalf("element %v erroneously reported as %v", key, i)
		}
	})

	count = 0
	UList(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).Each(func(i uint) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})

	UList(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).Each(func(index int, i uint) {
		if i != uint(index) {
			t.Fatalf("element %v erroneously reported as %v", index, i)
		}
	})

	UList(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).Each(func(key interface{}, i uint) {
		if i != uint(key.(int)) {
			t.Fatalf("element %v erroneously reported as %v", key, i)
		}
	})
}

func TestUSliceBlockCopy(t *testing.T) {
	ConfirmBlockCopy := func(s *USlice, destination, source, count int, r *USlice) {
		s.BlockCopy(destination, source, count)
		if !r.Equal(s) {
			t.Fatalf("BlockCopy(%v, %v, %v) should be %v but is %v", destination, source, count, r, s)
		}
	}

	ConfirmBlockCopy(UList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, 0, 4, UList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockCopy(UList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 9, 9, 4, UList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockCopy(UList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 2, 4, UList(0, 1, 2, 3, 4, 2, 3, 4, 5, 9))
	ConfirmBlockCopy(UList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 2, 5, 4, UList(0, 1, 5, 6, 7, 8, 6, 7, 8, 9))
}

func TestUSliceBlockClear(t *testing.T) {
	ConfirmBlockClear := func(s *USlice, start, count int, r *USlice) {
		s.BlockClear(start, count)
		if !r.Equal(s) {
			t.Fatalf("BlockClear(%v, %v) should be %v but is %v", start, count, r, s)
		}
	}

	ConfirmBlockClear(UList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, 4, UList(0, 0, 0, 0, 4, 5, 6, 7, 8, 9))
	ConfirmBlockClear(UList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, 4, UList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockClear(UList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 4, UList(0, 1, 2, 3, 4, 0, 0, 0, 0, 9))
}

func TestUSliceOverwrite(t *testing.T) {
	ConfirmOverwrite := func(s *USlice, offset int, v, r *USlice) {
		s.Overwrite(offset, *v)
		if !r.Equal(s) {
			t.Fatalf("Overwrite(%v, %v) should be %v but is %v", offset, v, r, s)
		}
	}

	ConfirmOverwrite(UList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, UList(10, 9, 8, 7), UList(10, 9, 8, 7, 4, 5, 6, 7, 8, 9))
	ConfirmOverwrite(UList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, UList(10, 9, 8, 7), UList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmOverwrite(UList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, UList(11, 12, 13, 14), UList(0, 1, 2, 3, 4, 11, 12, 13, 14, 9))
}

func TestUSliceReallocate(t *testing.T) {
	ConfirmReallocate := func(s *USlice, l, c int, r *USlice) {
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

	ConfirmReallocate(UList(), 0, 10, UList())
	ConfirmReallocate(UList(0, 1, 2, 3, 4), 3, 10, UList(0, 1, 2))
	ConfirmReallocate(UList(0, 1, 2, 3, 4), 5, 10, UList(0, 1, 2, 3, 4))
	ConfirmReallocate(UList(0, 1, 2, 3, 4), 10, 10, UList(0, 1, 2, 3, 4, 0, 0, 0, 0, 0))
	ConfirmReallocate(UList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 1, 5, UList(0))
	ConfirmReallocate(UList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 5, UList(0, 1, 2, 3, 4))
	ConfirmReallocate(UList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, 5, UList(0, 1, 2, 3, 4))
}

func TestUSliceExtend(t *testing.T) {
	ConfirmExtend := func(s *USlice, n int, r *USlice) {
		c := s.Cap()
		s.Extend(n)
		switch {
		case s.Len() != r.Len():	t.Fatalf("Extend(%v) len should be %v but is %v", n, r.Len(), s.Len())
		case s.Cap() != c + n:		t.Fatalf("Extend(%v) cap should be %v but is %v", n, c + n, s.Cap())
		case !r.Equal(s):			t.Fatalf("Extend(%v) should be %v but is %v", n, r, s)
		}
	}

	ConfirmExtend(UList(), 1, UList(0))
	ConfirmExtend(UList(), 2, UList(0, 0))
}

func TestUSliceExpand(t *testing.T) {
	ConfirmExpand := func(s *USlice, i, n int, r *USlice) {
		c := s.Cap()
		s.Expand(i, n)
		switch {
		case s.Len() != r.Len():	t.Fatalf("Expand(%v, %v) len should be %v but is %v", i, n, r.Len(), s.Len())
		case s.Cap() != c + n:		t.Fatalf("Expand(%v, %v) cap should be %v but is %v", i, n, c + n, s.Cap())
		case !r.Equal(s):			t.Fatalf("Expand(%v, %v) should be %v but is %v", i, n, r, s)
		}
	}

	ConfirmExpand(UList(), -1, 1, UList(0))
	ConfirmExpand(UList(), 0, 1, UList(0))
	ConfirmExpand(UList(), 1, 1, UList(0))
	ConfirmExpand(UList(), 0, 2, UList(0, 0))

	ConfirmExpand(UList(0, 1, 2), -1, 2, UList(0, 0, 0, 1, 2))
	ConfirmExpand(UList(0, 1, 2), 0, 2, UList(0, 0, 0, 1, 2))
	ConfirmExpand(UList(0, 1, 2), 1, 2, UList(0, 0, 0, 1, 2))
	ConfirmExpand(UList(0, 1, 2), 2, 2, UList(0, 1, 0, 0, 2))
	ConfirmExpand(UList(0, 1, 2), 3, 2, UList(0, 1, 2, 0, 0))
	ConfirmExpand(UList(0, 1, 2), 4, 2, UList(0, 1, 2, 0, 0))
}

func TestUSliceDepth(t *testing.T) {
	ConfirmDepth := func(s *USlice, i int) {
		if x := s.Depth(); x != i {
			t.Fatalf("%v.Depth() should be %v but is %v", s, i, x)
		}
	}
	ConfirmDepth(UList(0, 1), 0)
}

func TestUSliceReverse(t *testing.T) {
	sxp := UList(1, 2, 3, 4, 5)
	rxp := UList(5, 4, 3, 2, 1)
	sxp.Reverse()
	if !rxp.Equal(sxp) {
		t.Fatalf("Reversal failed: %v", sxp)
	}
}

func TestUSliceAppend(t *testing.T) {
	ConfirmAppend := func(s *USlice, v interface{}, r *USlice) {
		s.Append(v)
		if !r.Equal(s) {
			t.Fatalf("Append(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmAppend(UList(), uint(0), UList(0))

	ConfirmAppend(UList(), UList(0), UList(0))
	ConfirmAppend(UList(), UList(0, 1), UList(0, 1))
	ConfirmAppend(UList(0, 1, 2), UList(3, 4), UList(0, 1, 2, 3, 4))
}

func TestUSlicePrepend(t *testing.T) {
	ConfirmPrepend := func(s *USlice, v interface{}, r *USlice) {
		if s.Prepend(v); !r.Equal(s) {
			t.Fatalf("Prepend(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmPrepend(UList(), uint(0), UList(0))
	ConfirmPrepend(UList(0), uint(1), UList(1, 0))

	ConfirmPrepend(UList(), UList(0), UList(0))
	ConfirmPrepend(UList(), UList(0, 1), UList(0, 1))
	ConfirmPrepend(UList(0, 1, 2), UList(3, 4), UList(3, 4, 0, 1, 2))
}

func TestUSliceRepeat(t *testing.T) {
	ConfirmRepeat := func(s *USlice, count int, r *USlice) {
		if x := s.Repeat(count); !x.Equal(r) {
			t.Fatalf("%v.Repeat(%v) should be %v but is %v", s, count, r, x)
		}
	}

	ConfirmRepeat(UList(), 5, UList())
	ConfirmRepeat(UList(0), 1, UList(0))
	ConfirmRepeat(UList(0), 2, UList(0, 0))
	ConfirmRepeat(UList(0), 3, UList(0, 0, 0))
	ConfirmRepeat(UList(0), 4, UList(0, 0, 0, 0))
	ConfirmRepeat(UList(0), 5, UList(0, 0, 0, 0, 0))
}

func TestUSliceCar(t *testing.T) {
	ConfirmCar := func(s *USlice, r uint) {
		n := s.Car()
		if ok := n == r; !ok {
			t.Fatalf("head should be '%v' but is '%v'", r, n)
		}
	}
	ConfirmCar(UList(1, 2, 3), 1)
}

func TestUSliceCdr(t *testing.T) {
	ConfirmCdr := func(s, r *USlice) {
		if n := s.Cdr(); !n.Equal(r) {
			t.Fatalf("tail should be '%v' but is '%v'", r, n)
		}
	}
	ConfirmCdr(UList(1, 2, 3), UList(2, 3))
}

func TestUSliceRplaca(t *testing.T) {
	ConfirmRplaca := func(s *USlice, v interface{}, r *USlice) {
		if s.Rplaca(v); !s.Equal(r) {
			t.Fatalf("slice should be '%v' but is '%v'", r, s)
		}
	}
	ConfirmRplaca(UList(1, 2, 3, 4, 5), uint(0), UList(0, 2, 3, 4, 5))
}

func TestUSliceRplacd(t *testing.T) {
	ConfirmRplacd := func(s *USlice, v interface{}, r *USlice) {
		if s.Rplacd(v); !s.Equal(r) {
			t.Fatalf("slice should be '%v' but is '%v'", r, s)
		}
	}
	ConfirmRplacd(UList(1, 2, 3, 4, 5), nil, UList(1))
	ConfirmRplacd(UList(1, 2, 3, 4, 5), uint(10), UList(1, 10))
	ConfirmRplacd(UList(1, 2, 3, 4, 5), UList(5, 4, 3, 2), UList(1, 5, 4, 3, 2))
	ConfirmRplacd(UList(1, 2, 3, 4, 5, 6), UList(2, 4, 8, 16), UList(1, 2, 4, 8, 16))
}

func TestUSliceSetIntersection(t *testing.T) {
	ConfirmSetIntersection := func(s, o, r *USlice) {
		x := s.SetIntersection(*o)
		x.Sort()
		if !r.Equal(x) {
			t.Fatalf("%v.SetIntersection(%v) should be %v but is %v", s, o, r, x)
		}
	}

	ConfirmSetIntersection(UList(1, 2, 3), UList(), UList())
	ConfirmSetIntersection(UList(1, 2, 3), UList(1), UList(1))
	ConfirmSetIntersection(UList(1, 2, 3), UList(1, 1), UList(1))
	ConfirmSetIntersection(UList(1, 2, 3), UList(1, 2, 1), UList(1, 2))
}

func TestUSliceSetUnion(t *testing.T) {
	ConfirmSetUnion := func(s, o, r *USlice) {
		x := s.SetUnion(*o)
		x.Sort()
		if !r.Equal(x) {
			t.Fatalf("%v.SetUnion(%v) should be %v but is %v", s, o, r, x)
		}
	}

	ConfirmSetUnion(UList(1, 2, 3), UList(), UList(1, 2, 3))
	ConfirmSetUnion(UList(1, 2, 3), UList(1), UList(1, 2, 3))
	ConfirmSetUnion(UList(1, 2, 3), UList(1, 1), UList(1, 2, 3))
	ConfirmSetUnion(UList(1, 2, 3), UList(1, 2, 1), UList(1, 2, 3))
}

func TestUSliceSetDifference(t *testing.T) {
	ConfirmSetUnion := func(s, o, r *USlice) {
		x := s.SetDifference(*o)
		x.Sort()
		if !r.Equal(x) {
			t.Fatalf("%v.SetUnion(%v) should be %v but is %v", s, o, r, x)
		}
	}

	ConfirmSetUnion(UList(1, 2, 3), UList(), UList(1, 2, 3))
	ConfirmSetUnion(UList(1, 2, 3), UList(1), UList(2, 3))
	ConfirmSetUnion(UList(1, 2, 3), UList(1, 1), UList(2, 3))
	ConfirmSetUnion(UList(1, 2, 3), UList(1, 2, 1), UList(3))
}