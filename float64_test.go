package slices

import "testing"

func TestF64Slice(t *testing.T) {
	sxp := F64List(1)
	switch {
	case sxp.Len() != 1:			t.Fatalf("F64List(1) should allocate 1 cells, not %v cells", sxp.Len())
	case sxp.At(0) != float64(1):	t.Fatalf("F64List(1) element 0 should be 1 and not %v", sxp.At(0))
	}

	sxp = F64List(1, 2)
	switch {
	case sxp.Len() != 2:			t.Fatalf("F64List(1 2) should allocate 2 cells, not %v cells", sxp.Len())
	case sxp.At(0) != float64(1):	t.Fatalf("F64List(1 2) element 0 should be 1 and not %v", sxp.At(0))
	case sxp.At(1) != float64(2):	t.Fatalf("F64List(1 2) element 1 should be 2 and not %v", sxp.At(1))
	}

	sxp = F64List(1, 2, 3)
	switch {
	case sxp.Len() != 3:			t.Fatalf("F64List(1 2 3) should allocate 3 cells, not %v cells", sxp.Len())
	case sxp.At(0) != float64(1):	t.Fatalf("F64List(1 2 3) element 0 should be 1 and not %v", sxp.At(0))
	case sxp.At(1) != float64(2):	t.Fatalf("F64List(1 2 3) element 1 should be 2 and not %v", sxp.At(1))
	case sxp.At(2) != float64(3):	t.Fatalf("F64List(1 2 3) element 2 should be 3 and not %v", sxp.At(2))
	}
}

func TestF64SliceString(t *testing.T) {
	ConfirmString := func(s *F64Slice, r string) {
		if x := s.String(); x != r {
			t.Fatalf("%v erroneously serialised as '%v'", r, x)
		}
	}

	ConfirmString(F64List(), "()")
	ConfirmString(F64List(0), "(0)")
	ConfirmString(F64List(0, 1), "(0 1)")
}

func TestF64SliceLen(t *testing.T) {
	ConfirmLength := func(s *F64Slice, i int) {
		if x := s.Len(); x != i {
			t.Fatalf("%v.Len() should be %v but is %v", s, i, x)
		}
	}
	
	ConfirmLength(F64List(0), 1)
	ConfirmLength(F64List(0, 1), 2)
}

func TestF64SliceSwap(t *testing.T) {
	ConfirmSwap := func(s *F64Slice, i, j int, r *F64Slice) {
		if s.Swap(i, j); !r.Equal(s) {
			t.Fatalf("Swap(%v, %v) should be %v but is %v", i, j, r, s)
		}
	}
	ConfirmSwap(F64List(0, 1, 2), 0, 1, F64List(1, 0, 2))
	ConfirmSwap(F64List(0, 1, 2), 0, 2, F64List(2, 1, 0))
}

func TestF64SliceCompare(t *testing.T) {
	ConfirmCompare := func(s *F64Slice, i, j, r int) {
		if x := s.Compare(i, j); x != r {
			t.Fatalf("Compare(%v, %v) should be %v but is %v", i, j, r, x)
		}
	}

	ConfirmCompare(F64List(0, 1), 0, 0, IS_SAME_AS)
	ConfirmCompare(F64List(0, 1), 0, 1, IS_LESS_THAN)
	ConfirmCompare(F64List(0, 1), 1, 0, IS_GREATER_THAN)
}

func TestF64SliceZeroCompare(t *testing.T) {
	ConfirmCompare := func(s *F64Slice, i, r int) {
		if x := s.ZeroCompare(i); x != r {
			t.Fatalf("ZeroCompare(%v) should be %v but is %v", i, r, x)
		}
	}

	ConfirmCompare(F64List(0, -1, 1), 0, IS_SAME_AS)
	ConfirmCompare(F64List(0, -1, 1), 1, IS_GREATER_THAN)
	ConfirmCompare(F64List(0, -1, 1), 2, IS_LESS_THAN)
}

func TestF64SliceCut(t *testing.T) {
	ConfirmCut := func(s *F64Slice, start, end int, r *F64Slice) {
		if s.Cut(start, end); !r.Equal(s) {
			t.Fatalf("Cut(%v, %v) should be %v but is %v", start, end, r, s)
		}
	}

	ConfirmCut(F64List(0, 1, 2, 3, 4, 5), 0, 1, F64List(1, 2, 3, 4, 5))
	ConfirmCut(F64List(0, 1, 2, 3, 4, 5), 1, 2, F64List(0, 2, 3, 4, 5))
	ConfirmCut(F64List(0, 1, 2, 3, 4, 5), 2, 3, F64List(0, 1, 3, 4, 5))
	ConfirmCut(F64List(0, 1, 2, 3, 4, 5), 3, 4, F64List(0, 1, 2, 4, 5))
	ConfirmCut(F64List(0, 1, 2, 3, 4, 5), 4, 5, F64List(0, 1, 2, 3, 5))
	ConfirmCut(F64List(0, 1, 2, 3, 4, 5), 5, 6, F64List(0, 1, 2, 3, 4))

	ConfirmCut(F64List(0, 1, 2, 3, 4, 5), -1, 1, F64List(1, 2, 3, 4, 5))
	ConfirmCut(F64List(0, 1, 2, 3, 4, 5), 0, 2, F64List(2, 3, 4, 5))
	ConfirmCut(F64List(0, 1, 2, 3, 4, 5), 1, 3, F64List(0, 3, 4, 5))
	ConfirmCut(F64List(0, 1, 2, 3, 4, 5), 2, 4, F64List(0, 1, 4, 5))
	ConfirmCut(F64List(0, 1, 2, 3, 4, 5), 3, 5, F64List(0, 1, 2, 5))
	ConfirmCut(F64List(0, 1, 2, 3, 4, 5), 4, 6, F64List(0, 1, 2, 3))
	ConfirmCut(F64List(0, 1, 2, 3, 4, 5), 5, 7, F64List(0, 1, 2, 3, 4))
}

func TestF64SliceTrim(t *testing.T) {
	ConfirmTrim := func(s *F64Slice, start, end int, r *F64Slice) {
		if s.Trim(start, end); !r.Equal(s) {
			t.Fatalf("Cut(%v, %v) should be %v but is %v", start, end, r, s)
		}
	}

	ConfirmTrim(F64List(0, 1, 2, 3, 4, 5), 0, 1, F64List(0))
	ConfirmTrim(F64List(0, 1, 2, 3, 4, 5), 1, 2, F64List(1))
	ConfirmTrim(F64List(0, 1, 2, 3, 4, 5), 2, 3, F64List(2))
	ConfirmTrim(F64List(0, 1, 2, 3, 4, 5), 3, 4, F64List(3))
	ConfirmTrim(F64List(0, 1, 2, 3, 4, 5), 4, 5, F64List(4))
	ConfirmTrim(F64List(0, 1, 2, 3, 4, 5), 5, 6, F64List(5))

	ConfirmTrim(F64List(0, 1, 2, 3, 4, 5), -1, 1, F64List(0))
	ConfirmTrim(F64List(0, 1, 2, 3, 4, 5), 0, 2, F64List(0, 1))
	ConfirmTrim(F64List(0, 1, 2, 3, 4, 5), 1, 3, F64List(1, 2))
	ConfirmTrim(F64List(0, 1, 2, 3, 4, 5), 2, 4, F64List(2, 3))
	ConfirmTrim(F64List(0, 1, 2, 3, 4, 5), 3, 5, F64List(3, 4))
	ConfirmTrim(F64List(0, 1, 2, 3, 4, 5), 4, 6, F64List(4, 5))
	ConfirmTrim(F64List(0, 1, 2, 3, 4, 5), 5, 7, F64List(5))
}

func TestF64SliceDelete(t *testing.T) {
	ConfirmCut := func(s *F64Slice, index int, r *F64Slice) {
		if s.Delete(index); !r.Equal(s) {
			t.Fatalf("Delete(%v) should be %v but is %v", index, r, s)
		}
	}

	ConfirmCut(F64List(0, 1, 2, 3, 4, 5), -1, F64List(0, 1, 2, 3, 4, 5))
	ConfirmCut(F64List(0, 1, 2, 3, 4, 5), 0, F64List(1, 2, 3, 4, 5))
	ConfirmCut(F64List(0, 1, 2, 3, 4, 5), 1, F64List(0, 2, 3, 4, 5))
	ConfirmCut(F64List(0, 1, 2, 3, 4, 5), 2, F64List(0, 1, 3, 4, 5))
	ConfirmCut(F64List(0, 1, 2, 3, 4, 5), 3, F64List(0, 1, 2, 4, 5))
	ConfirmCut(F64List(0, 1, 2, 3, 4, 5), 4, F64List(0, 1, 2, 3, 5))
	ConfirmCut(F64List(0, 1, 2, 3, 4, 5), 5, F64List(0, 1, 2, 3, 4))
	ConfirmCut(F64List(0, 1, 2, 3, 4, 5), 6, F64List(0, 1, 2, 3, 4, 5))
}

func TestF64SliceEach(t *testing.T) {
	c := F64List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9)
	count := 0
	c.Each(func(i interface{}) {
		if i != float64(count) {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})
}

func TestF64SliceEachWithIndex(t *testing.T) {
	c := F64List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9)
	c.EachWithIndex(func(index int, i interface{}) {
		if i != float64(index) {
			t.Fatalf("element %v erroneously reported as %v", index, i)
		}
	})
}

func TestF64SliceEachWithKey(t *testing.T) {
	c := F64List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9)
	c.EachWithKey(func(key, i interface{}) {
		if i != float64(key.(int)) {
			t.Fatalf("element %v erroneously reported as %v", key, i)
		}
	})
}

func TestF64SliceEachF64(t *testing.T) {
	c := F64List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9)
	count := 0
	c.F64Each(func(f float64) {
		if f != float64(count) {
			t.Fatalf("element %v erroneously reported as %v", count, f)
		}
		count++
	})
}

func TestF64SliceEachF32WithIndex(t *testing.T) {
	c := F64List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9)
	c.F64EachWithIndex(func(index int, f float64) {
		if f != float64(index) {
			t.Fatalf("element %v erroneously reported as %v", index, f)
		}
	})
}

func TestF64SliceEachIntWithKey(t *testing.T) {
	c := F64List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9)
	c.F64EachWithKey(func(key interface{}, f float64) {
		if f != float64(key.(int)) {
			t.Fatalf("element %v erroneously reported as %v", key, f)
		}
	})
}

func TestF64SliceBlockCopy(t *testing.T) {
	ConfirmBlockCopy := func(s *F64Slice, destination, source, count int, r *F64Slice) {
		s.BlockCopy(destination, source, count)
		if !r.Equal(s) {
			t.Fatalf("BlockCopy(%v, %v, %v) should be %v but is %v", destination, source, count, r, s)
		}
	}

	ConfirmBlockCopy(F64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, 0, 4, F64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockCopy(F64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 9, 9, 4, F64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockCopy(F64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 2, 4, F64List(0, 1, 2, 3, 4, 2, 3, 4, 5, 9))
	ConfirmBlockCopy(F64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 2, 5, 4, F64List(0, 1, 5, 6, 7, 8, 6, 7, 8, 9))
}

func TestF64SliceBlockClear(t *testing.T) {
	ConfirmBlockClear := func(s *F64Slice, start, count int, r *F64Slice) {
		s.BlockClear(start, count)
		if !r.Equal(s) {
			t.Fatalf("BlockClear(%v, %v) should be %v but is %v", start, count, r, s)
		}
	}

	ConfirmBlockClear(F64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, 4, F64List(0, 0, 0, 0, 4, 5, 6, 7, 8, 9))
	ConfirmBlockClear(F64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, 4, F64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockClear(F64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 4, F64List(0, 1, 2, 3, 4, 0, 0, 0, 0, 9))
}

func TestF64SliceOverwrite(t *testing.T) {
	ConfirmOverwrite := func(s *F64Slice, offset int, v, r *F64Slice) {
		s.Overwrite(offset, *v)
		if !r.Equal(s) {
			t.Fatalf("Overwrite(%v, %v) should be %v but is %v", offset, v, r, s)
		}
	}

	ConfirmOverwrite(F64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, F64List(10, 9, 8, 7), F64List(10, 9, 8, 7, 4, 5, 6, 7, 8, 9))
	ConfirmOverwrite(F64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, F64List(10, 9, 8, 7), F64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmOverwrite(F64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, F64List(11, 12, 13, 14), F64List(0, 1, 2, 3, 4, 11, 12, 13, 14, 9))
}

func TestF64SliceReallocate(t *testing.T) {
	ConfirmReallocate := func(s *F64Slice, l, c int, r *F64Slice) {
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

	ConfirmReallocate(F64List(), 0, 10, F64List())
	ConfirmReallocate(F64List(0, 1, 2, 3, 4), 3, 10, F64List(0, 1, 2))
	ConfirmReallocate(F64List(0, 1, 2, 3, 4), 5, 10, F64List(0, 1, 2, 3, 4))
	ConfirmReallocate(F64List(0, 1, 2, 3, 4), 10, 10, F64List(0, 1, 2, 3, 4, 0, 0, 0, 0, 0))
	ConfirmReallocate(F64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 1, 5, F64List(0))
	ConfirmReallocate(F64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 5, F64List(0, 1, 2, 3, 4))
	ConfirmReallocate(F64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, 5, F64List(0, 1, 2, 3, 4))
}

func TestF64SliceExtend(t *testing.T) {
	ConfirmExtend := func(s *F64Slice, n int, r *F64Slice) {
		c := s.Cap()
		s.Extend(n)
		switch {
		case s.Len() != r.Len():	t.Fatalf("Extend(%v) len should be %v but is %v", n, r.Len(), s.Len())
		case s.Cap() != c + n:		t.Fatalf("Extend(%v) cap should be %v but is %v", n, c + n, s.Cap())
		case !r.Equal(s):			t.Fatalf("Extend(%v) should be %v but is %v", n, r, s)
		}
	}

	ConfirmExtend(F64List(), 1, F64List(0))
	ConfirmExtend(F64List(), 2, F64List(0, 0))
}

func TestF64SliceExpand(t *testing.T) {
	ConfirmExpand := func(s *F64Slice, i, n int, r *F64Slice) {
		c := s.Cap()
		s.Expand(i, n)
		switch {
		case s.Len() != r.Len():	t.Fatalf("Expand(%v, %v) len should be %v but is %v", i, n, r.Len(), s.Len())
		case s.Cap() != c + n:		t.Fatalf("Expand(%v, %v) cap should be %v but is %v", i, n, c + n, s.Cap())
		case !r.Equal(s):			t.Fatalf("Expand(%v, %v) should be %v but is %v", i, n, r, s)
		}
	}

	ConfirmExpand(F64List(), -1, 1, F64List(0))
	ConfirmExpand(F64List(), 0, 1, F64List(0))
	ConfirmExpand(F64List(), 1, 1, F64List(0))
	ConfirmExpand(F64List(), 0, 2, F64List(0, 0))

	ConfirmExpand(F64List(0, 1, 2), -1, 2, F64List(0, 0, 0, 1, 2))
	ConfirmExpand(F64List(0, 1, 2), 0, 2, F64List(0, 0, 0, 1, 2))
	ConfirmExpand(F64List(0, 1, 2), 1, 2, F64List(0, 0, 0, 1, 2))
	ConfirmExpand(F64List(0, 1, 2), 2, 2, F64List(0, 1, 0, 0, 2))
	ConfirmExpand(F64List(0, 1, 2), 3, 2, F64List(0, 1, 2, 0, 0))
	ConfirmExpand(F64List(0, 1, 2), 4, 2, F64List(0, 1, 2, 0, 0))
}

func TestF64SliceDepth(t *testing.T) {
	ConfirmDepth := func(s *F64Slice, i int) {
		if x := s.Depth(); x != i {
			t.Fatalf("%v.Depth() should be %v but is %v", s, i, x)
		}
	}
	ConfirmDepth(F64List(0, 1), 0)
}

func TestF64SliceReverse(t *testing.T) {
	sxp := F64List(1, 2, 3, 4, 5)
	rxp := F64List(5, 4, 3, 2, 1)
	sxp.Reverse()
	if !rxp.Equal(sxp) {
		t.Fatalf("Reversal failed: %v", sxp)
	}
}

func TestF64SliceAppend(t *testing.T) {
	ConfirmAppend := func(s *F64Slice, v interface{}, r *F64Slice) {
		s.Append(v)
		if !r.Equal(s) {
			t.Fatalf("Append(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmAppend(F64List(), float64(0), F64List(0))
}

func TestF64SliceAppendSlice(t *testing.T) {
	ConfirmAppendSlice := func(s, v, r *F64Slice) {
		s.AppendSlice(*v)
		if !r.Equal(s) {
			t.Fatalf("AppendSlice(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmAppendSlice(F64List(), F64List(0), F64List(0))
	ConfirmAppendSlice(F64List(), F64List(0, 1), F64List(0, 1))
	ConfirmAppendSlice(F64List(0, 1, 2), F64List(3, 4), F64List(0, 1, 2, 3, 4))
}

func TestF64SlicePrepend(t *testing.T) {
	ConfirmPrepend := func(s *F64Slice, v interface{}, r *F64Slice) {
		if s.Prepend(v); !r.Equal(s) {
			t.Fatalf("Prepend(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmPrepend(F64List(), float64(0), F64List(0))
	ConfirmPrepend(F64List(0), float64(1), F64List(1, 0))
}

func TestF64SlicePrependSlice(t *testing.T) {
	ConfirmPrependSlice := func(s, v, r *F64Slice) {
		if s.PrependSlice(*v); !r.Equal(s) {
			t.Fatalf("PrependSlice(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmPrependSlice(F64List(), F64List(0), F64List(0))
	ConfirmPrependSlice(F64List(), F64List(0, 1), F64List(0, 1))
	ConfirmPrependSlice(F64List(0, 1, 2), F64List(3, 4), F64List(3, 4, 0, 1, 2))
}

func TestF64SliceRepeat(t *testing.T) {
	ConfirmRepeat := func(s *F64Slice, count int, r *F64Slice) {
		if x := s.Repeat(count); !x.Equal(r) {
			t.Fatalf("%v.Repeat(%v) should be %v but is %v", s, count, r, x)
		}
	}

	ConfirmRepeat(F64List(), 5, F64List())
	ConfirmRepeat(F64List(0), 1, F64List(0))
	ConfirmRepeat(F64List(0), 2, F64List(0, 0))
	ConfirmRepeat(F64List(0), 3, F64List(0, 0, 0))
	ConfirmRepeat(F64List(0), 4, F64List(0, 0, 0, 0))
	ConfirmRepeat(F64List(0), 5, F64List(0, 0, 0, 0, 0))
}

func TestF64SliceCar(t *testing.T) {
	ConfirmCar := func(s *F64Slice, r float64) {
		n := s.Car()
		if ok := n == r; !ok {
			t.Fatalf("head should be '%v' but is '%v'", r, n)
		}
	}
	ConfirmCar(F64List(1, 2, 3), 1)
}

func TestF64SliceCdr(t *testing.T) {
	ConfirmCdr := func(s, r *F64Slice) {
		if n := s.Cdr(); !n.Equal(r) {
			t.Fatalf("tail should be '%v' but is '%v'", r, n)
		}
	}
	ConfirmCdr(F64List(1, 2, 3), F64List(2, 3))
}

func TestF64SliceRplaca(t *testing.T) {
	ConfirmRplaca := func(s *F64Slice, v interface{}, r *F64Slice) {
		if s.Rplaca(v); !s.Equal(r) {
			t.Fatalf("slice should be '%v' but is '%v'", r, s)
		}
	}
	ConfirmRplaca(F64List(1, 2, 3, 4, 5), float64(0), F64List(0, 2, 3, 4, 5))
}

func TestF64SliceRplacd(t *testing.T) {
	ConfirmRplacd := func(s *F64Slice, v interface{}, r *F64Slice) {
		if s.Rplacd(v); !s.Equal(r) {
			t.Fatalf("slice should be '%v' but is '%v'", r, s)
		}
	}
	ConfirmRplacd(F64List(1, 2, 3, 4, 5), nil, F64List(1))
	ConfirmRplacd(F64List(1, 2, 3, 4, 5), float64(10), F64List(1, 10))
	ConfirmRplacd(F64List(1, 2, 3, 4, 5), F64List(5, 4, 3, 2), F64List(1, 5, 4, 3, 2))
	ConfirmRplacd(F64List(1, 2, 3, 4, 5, 6), F64List(2, 4, 8, 16), F64List(1, 2, 4, 8, 16))
}