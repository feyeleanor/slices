package slices

import "testing"

func TestI32SliceString(t *testing.T) {
	ConfirmString := func(s *I32Slice, r string) {
		if x := s.String(); x != r {
			t.Fatalf("%v erroneously serialised as '%v'", r, x)
		}
	}

	ConfirmString(I32List(), "()")
	ConfirmString(I32List(0), "(0)")
	ConfirmString(I32List(0, 1), "(0 1)")
}

func TestI32SliceLen(t *testing.T) {
	ConfirmLength := func(s *I32Slice, i int) {
		if x := s.Len(); x != i {
			t.Fatalf("%v.Len() should be %v but is %v", s, i, x)
		}
	}
	
	ConfirmLength(I32List(0), 1)
	ConfirmLength(I32List(0, 1), 2)
}

func TestI32SliceSwap(t *testing.T) {
	ConfirmSwap := func(s *I32Slice, i, j int, r *I32Slice) {
		if s.Swap(i, j); !r.Equal(s) {
			t.Fatalf("Swap(%v, %v) should be %v but is %v", i, j, r, s)
		}
	}
	ConfirmSwap(I32List(0, 1, 2), 0, 1, I32List(1, 0, 2))
	ConfirmSwap(I32List(0, 1, 2), 0, 2, I32List(2, 1, 0))
}

func TestI32SliceSort(t *testing.T) {
	ConfirmSort := func(s, r *I32Slice) {
		if s.Sort(); !r.Equal(s) {
			t.Fatalf("Sort() should be %v but is %v", r, s)
		}
	}

	ConfirmSort(I32List(3, 2, 1, 4, 5, 0), I32List(0, 1, 2, 3, 4, 5))
}

func TestI32SliceCompare(t *testing.T) {
	ConfirmCompare := func(s *I32Slice, i, j, r int) {
		if x := s.Compare(i, j); x != r {
			t.Fatalf("Compare(%v, %v) should be %v but is %v", i, j, r, x)
		}
	}

	ConfirmCompare(I32List(0, 1), 0, 0, IS_SAME_AS)
	ConfirmCompare(I32List(0, 1), 0, 1, IS_LESS_THAN)
	ConfirmCompare(I32List(0, 1), 1, 0, IS_GREATER_THAN)
}

func TestI32SliceZeroCompare(t *testing.T) {
	ConfirmCompare := func(s *I32Slice, i, r int) {
		if x := s.ZeroCompare(i); x != r {
			t.Fatalf("ZeroCompare(%v) should be %v but is %v", i, r, x)
		}
	}

	ConfirmCompare(I32List(0, -1, 1), 0, IS_SAME_AS)
	ConfirmCompare(I32List(0, -1, 1), 1, IS_GREATER_THAN)
	ConfirmCompare(I32List(0, -1, 1), 2, IS_LESS_THAN)
}

func TestI32SliceCut(t *testing.T) {
	ConfirmCut := func(s *I32Slice, start, end int, r *I32Slice) {
		if s.Cut(start, end); !r.Equal(s) {
			t.Fatalf("Cut(%v, %v) should be %v but is %v", start, end, r, s)
		}
	}

	ConfirmCut(I32List(0, 1, 2, 3, 4, 5), 0, 1, I32List(1, 2, 3, 4, 5))
	ConfirmCut(I32List(0, 1, 2, 3, 4, 5), 1, 2, I32List(0, 2, 3, 4, 5))
	ConfirmCut(I32List(0, 1, 2, 3, 4, 5), 2, 3, I32List(0, 1, 3, 4, 5))
	ConfirmCut(I32List(0, 1, 2, 3, 4, 5), 3, 4, I32List(0, 1, 2, 4, 5))
	ConfirmCut(I32List(0, 1, 2, 3, 4, 5), 4, 5, I32List(0, 1, 2, 3, 5))
	ConfirmCut(I32List(0, 1, 2, 3, 4, 5), 5, 6, I32List(0, 1, 2, 3, 4))

	ConfirmCut(I32List(0, 1, 2, 3, 4, 5), -1, 1, I32List(1, 2, 3, 4, 5))
	ConfirmCut(I32List(0, 1, 2, 3, 4, 5), 0, 2, I32List(2, 3, 4, 5))
	ConfirmCut(I32List(0, 1, 2, 3, 4, 5), 1, 3, I32List(0, 3, 4, 5))
	ConfirmCut(I32List(0, 1, 2, 3, 4, 5), 2, 4, I32List(0, 1, 4, 5))
	ConfirmCut(I32List(0, 1, 2, 3, 4, 5), 3, 5, I32List(0, 1, 2, 5))
	ConfirmCut(I32List(0, 1, 2, 3, 4, 5), 4, 6, I32List(0, 1, 2, 3))
	ConfirmCut(I32List(0, 1, 2, 3, 4, 5), 5, 7, I32List(0, 1, 2, 3, 4))
}

func TestI32SliceTrim(t *testing.T) {
	ConfirmTrim := func(s *I32Slice, start, end int, r *I32Slice) {
		if s.Trim(start, end); !r.Equal(s) {
			t.Fatalf("Trim(%v, %v) should be %v but is %v", start, end, r, s)
		}
	}

	ConfirmTrim(I32List(0, 1, 2, 3, 4, 5), 0, 1, I32List(0))
	ConfirmTrim(I32List(0, 1, 2, 3, 4, 5), 1, 2, I32List(1))
	ConfirmTrim(I32List(0, 1, 2, 3, 4, 5), 2, 3, I32List(2))
	ConfirmTrim(I32List(0, 1, 2, 3, 4, 5), 3, 4, I32List(3))
	ConfirmTrim(I32List(0, 1, 2, 3, 4, 5), 4, 5, I32List(4))
	ConfirmTrim(I32List(0, 1, 2, 3, 4, 5), 5, 6, I32List(5))

	ConfirmTrim(I32List(0, 1, 2, 3, 4, 5), -1, 1, I32List(0))
	ConfirmTrim(I32List(0, 1, 2, 3, 4, 5), 0, 2, I32List(0, 1))
	ConfirmTrim(I32List(0, 1, 2, 3, 4, 5), 1, 3, I32List(1, 2))
	ConfirmTrim(I32List(0, 1, 2, 3, 4, 5), 2, 4, I32List(2, 3))
	ConfirmTrim(I32List(0, 1, 2, 3, 4, 5), 3, 5, I32List(3, 4))
	ConfirmTrim(I32List(0, 1, 2, 3, 4, 5), 4, 6, I32List(4, 5))
	ConfirmTrim(I32List(0, 1, 2, 3, 4, 5), 5, 7, I32List(5))
}

func TestI32SliceDelete(t *testing.T) {
	ConfirmDelete := func(s *I32Slice, index int, r *I32Slice) {
		if s.Delete(index); !r.Equal(s) {
			t.Fatalf("Delete(%v) should be %v but is %v", index, r, s)
		}
	}

	ConfirmDelete(I32List(0, 1, 2, 3, 4, 5), -1, I32List(0, 1, 2, 3, 4, 5))
	ConfirmDelete(I32List(0, 1, 2, 3, 4, 5), 0, I32List(1, 2, 3, 4, 5))
	ConfirmDelete(I32List(0, 1, 2, 3, 4, 5), 1, I32List(0, 2, 3, 4, 5))
	ConfirmDelete(I32List(0, 1, 2, 3, 4, 5), 2, I32List(0, 1, 3, 4, 5))
	ConfirmDelete(I32List(0, 1, 2, 3, 4, 5), 3, I32List(0, 1, 2, 4, 5))
	ConfirmDelete(I32List(0, 1, 2, 3, 4, 5), 4, I32List(0, 1, 2, 3, 5))
	ConfirmDelete(I32List(0, 1, 2, 3, 4, 5), 5, I32List(0, 1, 2, 3, 4))
	ConfirmDelete(I32List(0, 1, 2, 3, 4, 5), 6, I32List(0, 1, 2, 3, 4, 5))
}

func TestI32SliceDeleteIf(t *testing.T) {
	ConfirmDeleteIf := func(s *I32Slice, f interface{}, r *I32Slice) {
		if s.DeleteIf(f); !r.Equal(s) {
			t.Fatalf("DeleteIf(%v) should be %v but is %v", f, r, s)
		}
	}

	ConfirmDeleteIf(I32List(0, 1, 0, 3, 0, 5), int32(0), I32List(1, 3, 5))
	ConfirmDeleteIf(I32List(0, 1, 0, 3, 0, 5), int32(1), I32List(0, 0, 3, 0, 5))
	ConfirmDeleteIf(I32List(0, 1, 0, 3, 0, 5), int32(6), I32List(0, 1, 0, 3, 0, 5))

	ConfirmDeleteIf(I32List(0, 1, 0, 3, 0, 5), func(x interface{}) bool { return x == int32(0) }, I32List(1, 3, 5))
	ConfirmDeleteIf(I32List(0, 1, 0, 3, 0, 5), func(x interface{}) bool { return x == int32(1) }, I32List(0, 0, 3, 0, 5))
	ConfirmDeleteIf(I32List(0, 1, 0, 3, 0, 5), func(x interface{}) bool { return x == int32(6) }, I32List(0, 1, 0, 3, 0, 5))

	ConfirmDeleteIf(I32List(0, 1, 0, 3, 0, 5), func(x int32) bool { return x == int32(0) }, I32List(1, 3, 5))
	ConfirmDeleteIf(I32List(0, 1, 0, 3, 0, 5), func(x int32) bool { return x == int32(1) }, I32List(0, 0, 3, 0, 5))
	ConfirmDeleteIf(I32List(0, 1, 0, 3, 0, 5), func(x int32) bool { return x == int32(6) }, I32List(0, 1, 0, 3, 0, 5))
}

func TestI32SliceEach(t *testing.T) {
	var count	int32
	I32List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).Each(func(i interface{}) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})

	I32List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).Each(func(index int, i interface{}) {
		if i != int32(index) {
			t.Fatalf("element %v erroneously reported as %v", index, i)
		}
	})

	I32List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).Each(func(key, i interface{}) {
		if i != int32(key.(int)) {
			t.Fatalf("element %v erroneously reported as %v", key, i)
		}
	})

	count = 0
	I32List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).Each(func(i int32) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})

	I32List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).Each(func(index int, i int32) {
		if i != int32(index) {
			t.Fatalf("element %v erroneously reported as %v", index, i)
		}
	})

	I32List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).Each(func(key interface{}, i int32) {
		if i != int32(key.(int)) {
			t.Fatalf("element %v erroneously reported as %v", key, i)
		}
	})
}

func TestI32SliceBlockCopy(t *testing.T) {
	ConfirmBlockCopy := func(s *I32Slice, destination, source, count int, r *I32Slice) {
		s.BlockCopy(destination, source, count)
		if !r.Equal(s) {
			t.Fatalf("BlockCopy(%v, %v, %v) should be %v but is %v", destination, source, count, r, s)
		}
	}

	ConfirmBlockCopy(I32List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, 0, 4, I32List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockCopy(I32List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 9, 9, 4, I32List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockCopy(I32List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 2, 4, I32List(0, 1, 2, 3, 4, 2, 3, 4, 5, 9))
	ConfirmBlockCopy(I32List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 2, 5, 4, I32List(0, 1, 5, 6, 7, 8, 6, 7, 8, 9))
}

func TestI32SliceBlockClear(t *testing.T) {
	ConfirmBlockClear := func(s *I32Slice, start, count int, r *I32Slice) {
		s.BlockClear(start, count)
		if !r.Equal(s) {
			t.Fatalf("BlockClear(%v, %v) should be %v but is %v", start, count, r, s)
		}
	}

	ConfirmBlockClear(I32List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, 4, I32List(0, 0, 0, 0, 4, 5, 6, 7, 8, 9))
	ConfirmBlockClear(I32List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, 4, I32List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockClear(I32List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 4, I32List(0, 1, 2, 3, 4, 0, 0, 0, 0, 9))
}

func TestI32SliceOverwrite(t *testing.T) {
	ConfirmOverwrite := func(s *I32Slice, offset int, v, r *I32Slice) {
		s.Overwrite(offset, *v)
		if !r.Equal(s) {
			t.Fatalf("Overwrite(%v, %v) should be %v but is %v", offset, v, r, s)
		}
	}

	ConfirmOverwrite(I32List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, I32List(10, 9, 8, 7), I32List(10, 9, 8, 7, 4, 5, 6, 7, 8, 9))
	ConfirmOverwrite(I32List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, I32List(10, 9, 8, 7), I32List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmOverwrite(I32List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, I32List(11, 12, 13, 14), I32List(0, 1, 2, 3, 4, 11, 12, 13, 14, 9))
}

func TestI32SliceReallocate(t *testing.T) {
	ConfirmReallocate := func(s *I32Slice, l, c int, r *I32Slice) {
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

	ConfirmReallocate(I32List(), 0, 10, I32List())
	ConfirmReallocate(I32List(0, 1, 2, 3, 4), 3, 10, I32List(0, 1, 2))
	ConfirmReallocate(I32List(0, 1, 2, 3, 4), 5, 10, I32List(0, 1, 2, 3, 4))
	ConfirmReallocate(I32List(0, 1, 2, 3, 4), 10, 10, I32List(0, 1, 2, 3, 4, 0, 0, 0, 0, 0))
	ConfirmReallocate(I32List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 1, 5, I32List(0))
	ConfirmReallocate(I32List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 5, I32List(0, 1, 2, 3, 4))
	ConfirmReallocate(I32List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, 5, I32List(0, 1, 2, 3, 4))
}

func TestI32SliceExtend(t *testing.T) {
	ConfirmExtend := func(s *I32Slice, n int, r *I32Slice) {
		c := s.Cap()
		s.Extend(n)
		switch {
		case s.Len() != r.Len():	t.Fatalf("Extend(%v) len should be %v but is %v", n, r.Len(), s.Len())
		case s.Cap() != c + n:		t.Fatalf("Extend(%v) cap should be %v but is %v", n, c + n, s.Cap())
		case !r.Equal(s):			t.Fatalf("Extend(%v) should be %v but is %v", n, r, s)
		}
	}

	ConfirmExtend(I32List(), 1, I32List(0))
	ConfirmExtend(I32List(), 2, I32List(0, 0))
}

func TestI32SliceExpand(t *testing.T) {
	ConfirmExpand := func(s *I32Slice, i, n int, r *I32Slice) {
		c := s.Cap()
		s.Expand(i, n)
		switch {
		case s.Len() != r.Len():	t.Fatalf("Expand(%v, %v) len should be %v but is %v", i, n, r.Len(), s.Len())
		case s.Cap() != c + n:		t.Fatalf("Expand(%v, %v) cap should be %v but is %v", i, n, c + n, s.Cap())
		case !r.Equal(s):			t.Fatalf("Expand(%v, %v) should be %v but is %v", i, n, r, s)
		}
	}

	ConfirmExpand(I32List(), -1, 1, I32List(0))
	ConfirmExpand(I32List(), 0, 1, I32List(0))
	ConfirmExpand(I32List(), 1, 1, I32List(0))
	ConfirmExpand(I32List(), 0, 2, I32List(0, 0))

	ConfirmExpand(I32List(0, 1, 2), -1, 2, I32List(0, 0, 0, 1, 2))
	ConfirmExpand(I32List(0, 1, 2), 0, 2, I32List(0, 0, 0, 1, 2))
	ConfirmExpand(I32List(0, 1, 2), 1, 2, I32List(0, 0, 0, 1, 2))
	ConfirmExpand(I32List(0, 1, 2), 2, 2, I32List(0, 1, 0, 0, 2))
	ConfirmExpand(I32List(0, 1, 2), 3, 2, I32List(0, 1, 2, 0, 0))
	ConfirmExpand(I32List(0, 1, 2), 4, 2, I32List(0, 1, 2, 0, 0))
}

func TestI32SliceDepth(t *testing.T) {
	ConfirmDepth := func(s *I32Slice, i int) {
		if x := s.Depth(); x != i {
			t.Fatalf("%v.Depth() should be %v but is %v", s, i, x)
		}
	}
	ConfirmDepth(I32List(0, 1), 0)
}

func TestI32SliceReverse(t *testing.T) {
	sxp := I32List(1, 2, 3, 4, 5)
	rxp := I32List(5, 4, 3, 2, 1)
	sxp.Reverse()
	if !rxp.Equal(sxp) {
		t.Fatalf("Reversal failed: %v", sxp)
	}
}

func TestI32SliceAppend(t *testing.T) {
	ConfirmAppend := func(s *I32Slice, v interface{}, r *I32Slice) {
		s.Append(v)
		if !r.Equal(s) {
			t.Fatalf("Append(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmAppend(I32List(), int32(0), I32List(0))

	ConfirmAppend(I32List(), I32List(0), I32List(0))
	ConfirmAppend(I32List(), I32List(0, 1), I32List(0, 1))
	ConfirmAppend(I32List(0, 1, 2), I32List(3, 4), I32List(0, 1, 2, 3, 4))
}

func TestI32SlicePrepend(t *testing.T) {
	ConfirmPrepend := func(s *I32Slice, v interface{}, r *I32Slice) {
		if s.Prepend(v); !r.Equal(s) {
			t.Fatalf("Prepend(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmPrepend(I32List(), int32(0), I32List(0))
	ConfirmPrepend(I32List(0), int32(1), I32List(1, 0))

	ConfirmPrepend(I32List(), I32List(0), I32List(0))
	ConfirmPrepend(I32List(), I32List(0, 1), I32List(0, 1))
	ConfirmPrepend(I32List(0, 1, 2), I32List(3, 4), I32List(3, 4, 0, 1, 2))
}

func TestI32SliceRepeat(t *testing.T) {
	ConfirmRepeat := func(s *I32Slice, count int, r *I32Slice) {
		if x := s.Repeat(count); !x.Equal(r) {
			t.Fatalf("%v.Repeat(%v) should be %v but is %v", s, count, r, x)
		}
	}

	ConfirmRepeat(I32List(), 5, I32List())
	ConfirmRepeat(I32List(0), 1, I32List(0))
	ConfirmRepeat(I32List(0), 2, I32List(0, 0))
	ConfirmRepeat(I32List(0), 3, I32List(0, 0, 0))
	ConfirmRepeat(I32List(0), 4, I32List(0, 0, 0, 0))
	ConfirmRepeat(I32List(0), 5, I32List(0, 0, 0, 0, 0))
}

func TestI32SliceCar(t *testing.T) {
	ConfirmCar := func(s *I32Slice, r int32) {
		n := s.Car().(int32)
		if ok := n == r; !ok {
			t.Fatalf("head should be '%v' but is '%v'", r, n)
		}
	}
	ConfirmCar(I32List(1, 2, 3), 1)
}

func TestI32SliceCdr(t *testing.T) {
	ConfirmCdr := func(s, r *I32Slice) {
		if n := s.Cdr(); !n.Equal(r) {
			t.Fatalf("tail should be '%v' but is '%v'", r, n)
		}
	}
	ConfirmCdr(I32List(1, 2, 3), I32List(2, 3))
}

func TestI32SliceRplaca(t *testing.T) {
	ConfirmRplaca := func(s *I32Slice, v interface{}, r *I32Slice) {
		if s.Rplaca(v); !s.Equal(r) {
			t.Fatalf("slice should be '%v' but is '%v'", r, s)
		}
	}
	ConfirmRplaca(I32List(1, 2, 3, 4, 5), int32(0), I32List(0, 2, 3, 4, 5))
}

func TestI32SliceRplacd(t *testing.T) {
	ConfirmRplacd := func(s *I32Slice, v interface{}, r *I32Slice) {
		if s.Rplacd(v); !s.Equal(r) {
			t.Fatalf("slice should be '%v' but is '%v'", r, s)
		}
	}
	ConfirmRplacd(I32List(1, 2, 3, 4, 5), nil, I32List(1))
	ConfirmRplacd(I32List(1, 2, 3, 4, 5), int32(10), I32List(1, 10))
	ConfirmRplacd(I32List(1, 2, 3, 4, 5), I32List(5, 4, 3, 2), I32List(1, 5, 4, 3, 2))
	ConfirmRplacd(I32List(1, 2, 3, 4, 5, 6), I32List(2, 4, 8, 16), I32List(1, 2, 4, 8, 16))
}

func TestI32SliceSetIntersection(t *testing.T) {
	ConfirmSetIntersection := func(s, o, r *I32Slice) {
		x := s.SetIntersection(*o)
		x.Sort()
		if !r.Equal(x) {
			t.Fatalf("%v.SetIntersection(%v) should be %v but is %v", s, o, r, x)
		}
	}

	ConfirmSetIntersection(I32List(1, 2, 3), I32List(), I32List())
	ConfirmSetIntersection(I32List(1, 2, 3), I32List(1), I32List(1))
	ConfirmSetIntersection(I32List(1, 2, 3), I32List(1, 1), I32List(1))
	ConfirmSetIntersection(I32List(1, 2, 3), I32List(1, 2, 1), I32List(1, 2))
}

func TestI32SliceSetUnion(t *testing.T) {
	ConfirmSetUnion := func(s, o, r *I32Slice) {
		x := s.SetUnion(*o)
		x.Sort()
		if !r.Equal(x) {
			t.Fatalf("%v.SetUnion(%v) should be %v but is %v", s, o, r, x)
		}
	}

	ConfirmSetUnion(I32List(1, 2, 3), I32List(), I32List(1, 2, 3))
	ConfirmSetUnion(I32List(1, 2, 3), I32List(1), I32List(1, 2, 3))
	ConfirmSetUnion(I32List(1, 2, 3), I32List(1, 1), I32List(1, 2, 3))
	ConfirmSetUnion(I32List(1, 2, 3), I32List(1, 2, 1), I32List(1, 2, 3))
}

func TestI32SliceSetDifference(t *testing.T) {
	ConfirmSetUnion := func(s, o, r *I32Slice) {
		x := s.SetDifference(*o)
		x.Sort()
		if !r.Equal(x) {
			t.Fatalf("%v.SetUnion(%v) should be %v but is %v", s, o, r, x)
		}
	}

	ConfirmSetUnion(I32List(1, 2, 3), I32List(), I32List(1, 2, 3))
	ConfirmSetUnion(I32List(1, 2, 3), I32List(1), I32List(2, 3))
	ConfirmSetUnion(I32List(1, 2, 3), I32List(1, 1), I32List(2, 3))
	ConfirmSetUnion(I32List(1, 2, 3), I32List(1, 2, 1), I32List(3))
}