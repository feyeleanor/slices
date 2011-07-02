package slices

import "testing"

func TestISlice(t *testing.T) {
	sxp := IList(1)
	switch {
	case sxp.Len() != 1:			t.Fatalf("IList(1) should allocate 1 cells, not %v cells", sxp.Len())
	case sxp.At(0) != 1:			t.Fatalf("IList(1) element 0 should be 1 and not %v", sxp.At(0))
	}

	sxp = IList(1, 2)
	switch {
	case sxp.Len() != 2:			t.Fatalf("IList(1 2) should allocate 2 cells, not %v cells", sxp.Len())
	case sxp.At(0) != 1:			t.Fatalf("IList(1 2) element 0 should be 1 and not %v", sxp.At(0))
	case sxp.At(1) != 2:			t.Fatalf("IList(1 2) element 1 should be 2 and not %v", sxp.At(1))
	}

	sxp = IList(1, 2, 3)
	switch {
	case sxp.Len() != 3:			t.Fatalf("IList(1 2 3) should allocate 3 cells, not %v cells", sxp.Len())
	case sxp.At(0) != 1:			t.Fatalf("IList(1 2 3) element 0 should be 1 and not %v", sxp.At(0))
	case sxp.At(1) != 2:			t.Fatalf("IList(1 2 3) element 1 should be 2 and not %v", sxp.At(1))
	case sxp.At(2) != 3:			t.Fatalf("IList(1 2 3) element 2 should be 3 and not %v", sxp.At(2))
	}
}

func TestISliceString(t *testing.T) {
	ConfirmString := func(s *ISlice, r string) {
		if x := s.String(); x != r {
			t.Fatalf("%v erroneously serialised as '%v'", r, x)
		}
	}

	ConfirmString(IList(), "()")
	ConfirmString(IList(0), "(0)")
	ConfirmString(IList(0, 1), "(0 1)")
}

func TestISliceLen(t *testing.T) {
	ConfirmLength := func(s *ISlice, i int) {
		if x := s.Len(); x != i {
			t.Fatalf("%v.Len() should be %v but is %v", s, i, x)
		}
	}
	
	ConfirmLength(IList(0), 1)
	ConfirmLength(IList(0, 1), 2)
}

func TestISliceSwap(t *testing.T) {
	ConfirmSwap := func(s *ISlice, i, j int, r *ISlice) {
		if s.Swap(i, j); !r.Equal(s) {
			t.Fatalf("Swap(%v, %v) should be %v but is %v", i, j, r, s)
		}
	}
	ConfirmSwap(IList(0, 1, 2), 0, 1, IList(1, 0, 2))
	ConfirmSwap(IList(0, 1, 2), 0, 2, IList(2, 1, 0))
}

func TestISliceCompare(t *testing.T) {
	ConfirmCompare := func(s *ISlice, i, j, r int) {
		if x := s.Compare(i, j); x != r {
			t.Fatalf("Compare(%v, %v) should be %v but is %v", i, j, r, x)
		}
	}

	ConfirmCompare(IList(0, 1), 0, 0, IS_SAME_AS)
	ConfirmCompare(IList(0, 1), 0, 1, IS_LESS_THAN)
	ConfirmCompare(IList(0, 1), 1, 0, IS_GREATER_THAN)
}

func TestISliceZeroCompare(t *testing.T) {
	ConfirmCompare := func(s *ISlice, i, r int) {
		if x := s.ZeroCompare(i); x != r {
			t.Fatalf("ZeroCompare(%v) should be %v but is %v", i, r, x)
		}
	}

	ConfirmCompare(IList(0, -1, 1), 0, IS_SAME_AS)
	ConfirmCompare(IList(0, -1, 1), 1, IS_GREATER_THAN)
	ConfirmCompare(IList(0, -1, 1), 2, IS_LESS_THAN)
}

func TestISliceCut(t *testing.T) {
	ConfirmCut := func(s *ISlice, start, end int, r *ISlice) {
		if s.Cut(start, end); !r.Equal(s) {
			t.Fatalf("Cut(%v, %v) should be %v but is %v", start, end, r, s)
		}
	}

	ConfirmCut(IList(0, 1, 2, 3, 4, 5), 0, 1, IList(1, 2, 3, 4, 5))
	ConfirmCut(IList(0, 1, 2, 3, 4, 5), 1, 2, IList(0, 2, 3, 4, 5))
	ConfirmCut(IList(0, 1, 2, 3, 4, 5), 2, 3, IList(0, 1, 3, 4, 5))
	ConfirmCut(IList(0, 1, 2, 3, 4, 5), 3, 4, IList(0, 1, 2, 4, 5))
	ConfirmCut(IList(0, 1, 2, 3, 4, 5), 4, 5, IList(0, 1, 2, 3, 5))
	ConfirmCut(IList(0, 1, 2, 3, 4, 5), 5, 6, IList(0, 1, 2, 3, 4))

	ConfirmCut(IList(0, 1, 2, 3, 4, 5), -1, 1, IList(1, 2, 3, 4, 5))
	ConfirmCut(IList(0, 1, 2, 3, 4, 5), 0, 2, IList(2, 3, 4, 5))
	ConfirmCut(IList(0, 1, 2, 3, 4, 5), 1, 3, IList(0, 3, 4, 5))
	ConfirmCut(IList(0, 1, 2, 3, 4, 5), 2, 4, IList(0, 1, 4, 5))
	ConfirmCut(IList(0, 1, 2, 3, 4, 5), 3, 5, IList(0, 1, 2, 5))
	ConfirmCut(IList(0, 1, 2, 3, 4, 5), 4, 6, IList(0, 1, 2, 3))
	ConfirmCut(IList(0, 1, 2, 3, 4, 5), 5, 7, IList(0, 1, 2, 3, 4))
}

func TestISliceTrim(t *testing.T) {
	ConfirmTrim := func(s *ISlice, start, end int, r *ISlice) {
		if s.Trim(start, end); !r.Equal(s) {
			t.Fatalf("Cut(%v, %v) should be %v but is %v", start, end, r, s)
		}
	}

	ConfirmTrim(IList(0, 1, 2, 3, 4, 5), 0, 1, IList(0))
	ConfirmTrim(IList(0, 1, 2, 3, 4, 5), 1, 2, IList(1))
	ConfirmTrim(IList(0, 1, 2, 3, 4, 5), 2, 3, IList(2))
	ConfirmTrim(IList(0, 1, 2, 3, 4, 5), 3, 4, IList(3))
	ConfirmTrim(IList(0, 1, 2, 3, 4, 5), 4, 5, IList(4))
	ConfirmTrim(IList(0, 1, 2, 3, 4, 5), 5, 6, IList(5))

	ConfirmTrim(IList(0, 1, 2, 3, 4, 5), -1, 1, IList(0))
	ConfirmTrim(IList(0, 1, 2, 3, 4, 5), 0, 2, IList(0, 1))
	ConfirmTrim(IList(0, 1, 2, 3, 4, 5), 1, 3, IList(1, 2))
	ConfirmTrim(IList(0, 1, 2, 3, 4, 5), 2, 4, IList(2, 3))
	ConfirmTrim(IList(0, 1, 2, 3, 4, 5), 3, 5, IList(3, 4))
	ConfirmTrim(IList(0, 1, 2, 3, 4, 5), 4, 6, IList(4, 5))
	ConfirmTrim(IList(0, 1, 2, 3, 4, 5), 5, 7, IList(5))
}

func TestISliceDelete(t *testing.T) {
	ConfirmCut := func(s *ISlice, index int, r *ISlice) {
		if s.Delete(index); !r.Equal(s) {
			t.Fatalf("Delete(%v) should be %v but is %v", index, r, s)
		}
	}

	ConfirmCut(IList(0, 1, 2, 3, 4, 5), -1, IList(0, 1, 2, 3, 4, 5))
	ConfirmCut(IList(0, 1, 2, 3, 4, 5), 0, IList(1, 2, 3, 4, 5))
	ConfirmCut(IList(0, 1, 2, 3, 4, 5), 1, IList(0, 2, 3, 4, 5))
	ConfirmCut(IList(0, 1, 2, 3, 4, 5), 2, IList(0, 1, 3, 4, 5))
	ConfirmCut(IList(0, 1, 2, 3, 4, 5), 3, IList(0, 1, 2, 4, 5))
	ConfirmCut(IList(0, 1, 2, 3, 4, 5), 4, IList(0, 1, 2, 3, 5))
	ConfirmCut(IList(0, 1, 2, 3, 4, 5), 5, IList(0, 1, 2, 3, 4))
	ConfirmCut(IList(0, 1, 2, 3, 4, 5), 6, IList(0, 1, 2, 3, 4, 5))
}

func TestISliceEach(t *testing.T) {
	c := IList(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9)
	count := 0
	c.Each(func(i interface{}) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})
}

func TestISliceEachWithIndex(t *testing.T) {
	c := IList(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9)
	c.EachWithIndex(func(index int, i interface{}) {
		if i != index {
			t.Fatalf("element %v erroneously reported as %v", index, i)
		}
	})
}

func TestISliceEachWithKey(t *testing.T) {
	c := IList(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9)
	c.EachWithKey(func(key, i interface{}) {
		if i != key {
			t.Fatalf("element %v erroneously reported as %v", key, i)
		}
	})
}

func TestISliceIEach(t *testing.T) {
	c := IList(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9)
	count := 0
	c.IEach(func(i int) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})
}

func TestISliceIEachWithIndex(t *testing.T) {
	c := IList(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9)
	c.IEachWithIndex(func(index int, i int) {
		if i != index {
			t.Fatalf("element %v erroneously reported as %v", index, i)
		}
	})
}

func TestISliceIEachWithKey(t *testing.T) {
	c := IList(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9)
	c.IEachWithKey(func(key interface{}, i int) {
		if i != key {
			t.Fatalf("element %v erroneously reported as %v", key, i)
		}
	})
}

func TestISliceBlockCopy(t *testing.T) {
	ConfirmBlockCopy := func(s *ISlice, destination, source, count int, r *ISlice) {
		s.BlockCopy(destination, source, count)
		if !r.Equal(s) {
			t.Fatalf("BlockCopy(%v, %v, %v) should be %v but is %v", destination, source, count, r, s)
		}
	}

	ConfirmBlockCopy(IList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, 0, 4, IList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockCopy(IList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 9, 9, 4, IList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockCopy(IList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 2, 4, IList(0, 1, 2, 3, 4, 2, 3, 4, 5, 9))
	ConfirmBlockCopy(IList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 2, 5, 4, IList(0, 1, 5, 6, 7, 8, 6, 7, 8, 9))
}

func TestISliceBlockClear(t *testing.T) {
	ConfirmBlockClear := func(s *ISlice, start, count int, r *ISlice) {
		s.BlockClear(start, count)
		if !r.Equal(s) {
			t.Fatalf("BlockClear(%v, %v) should be %v but is %v", start, count, r, s)
		}
	}

	ConfirmBlockClear(IList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, 4, IList(0, 0, 0, 0, 4, 5, 6, 7, 8, 9))
	ConfirmBlockClear(IList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, 4, IList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockClear(IList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 4, IList(0, 1, 2, 3, 4, 0, 0, 0, 0, 9))
}

func TestISliceOverwrite(t *testing.T) {
	ConfirmOverwrite := func(s *ISlice, offset int, v, r *ISlice) {
		s.Overwrite(offset, *v)
		if !r.Equal(s) {
			t.Fatalf("Overwrite(%v, %v) should be %v but is %v", offset, v, r, s)
		}
	}

	ConfirmOverwrite(IList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, IList(10, 9, 8, 7), IList(10, 9, 8, 7, 4, 5, 6, 7, 8, 9))
	ConfirmOverwrite(IList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, IList(10, 9, 8, 7), IList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmOverwrite(IList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, IList(11, 12, 13, 14), IList(0, 1, 2, 3, 4, 11, 12, 13, 14, 9))
}

func TestISliceReallocate(t *testing.T) {
	ConfirmReallocate := func(s *ISlice, l, c int, r *ISlice) {
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

	ConfirmReallocate(IList(), 0, 10, IList())
	ConfirmReallocate(IList(0, 1, 2, 3, 4), 3, 10, IList(0, 1, 2))
	ConfirmReallocate(IList(0, 1, 2, 3, 4), 5, 10, IList(0, 1, 2, 3, 4))
	ConfirmReallocate(IList(0, 1, 2, 3, 4), 10, 10, IList(0, 1, 2, 3, 4, 0, 0, 0, 0, 0))
	ConfirmReallocate(IList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 1, 5, IList(0))
	ConfirmReallocate(IList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 5, IList(0, 1, 2, 3, 4))
	ConfirmReallocate(IList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, 5, IList(0, 1, 2, 3, 4))
}

func TestISliceExtend(t *testing.T) {
	ConfirmExtend := func(s *ISlice, n int, r *ISlice) {
		c := s.Cap()
		s.Extend(n)
		switch {
		case s.Len() != r.Len():	t.Fatalf("Extend(%v) len should be %v but is %v", n, r.Len(), s.Len())
		case s.Cap() != c + n:		t.Fatalf("Extend(%v) cap should be %v but is %v", n, c + n, s.Cap())
		case !r.Equal(s):			t.Fatalf("Extend(%v) should be %v but is %v", n, r, s)
		}
	}

	ConfirmExtend(IList(), 1, IList(0))
	ConfirmExtend(IList(), 2, IList(0, 0))
}

func TestISliceExpand(t *testing.T) {
	ConfirmExpand := func(s *ISlice, i, n int, r *ISlice) {
		c := s.Cap()
		s.Expand(i, n)
		switch {
		case s.Len() != r.Len():	t.Fatalf("Expand(%v, %v) len should be %v but is %v", i, n, r.Len(), s.Len())
		case s.Cap() != c + n:		t.Fatalf("Expand(%v, %v) cap should be %v but is %v", i, n, c + n, s.Cap())
		case !r.Equal(s):			t.Fatalf("Expand(%v, %v) should be %v but is %v", i, n, r, s)
		}
	}

	ConfirmExpand(IList(), -1, 1, IList(0))
	ConfirmExpand(IList(), 0, 1, IList(0))
	ConfirmExpand(IList(), 1, 1, IList(0))
	ConfirmExpand(IList(), 0, 2, IList(0, 0))

	ConfirmExpand(IList(0, 1, 2), -1, 2, IList(0, 0, 0, 1, 2))
	ConfirmExpand(IList(0, 1, 2), 0, 2, IList(0, 0, 0, 1, 2))
	ConfirmExpand(IList(0, 1, 2), 1, 2, IList(0, 0, 0, 1, 2))
	ConfirmExpand(IList(0, 1, 2), 2, 2, IList(0, 1, 0, 0, 2))
	ConfirmExpand(IList(0, 1, 2), 3, 2, IList(0, 1, 2, 0, 0))
	ConfirmExpand(IList(0, 1, 2), 4, 2, IList(0, 1, 2, 0, 0))
}

func TestISliceDepth(t *testing.T) {
	ConfirmDepth := func(s *ISlice, i int) {
		if x := s.Depth(); x != i {
			t.Fatalf("%v.Depth() should be %v but is %v", s, i, x)
		}
	}
	ConfirmDepth(IList(0, 1), 0)
}

func TestISliceReverse(t *testing.T) {
	sxp := IList(1, 2, 3, 4, 5)
	rxp := IList(5, 4, 3, 2, 1)
	sxp.Reverse()
	if !rxp.Equal(sxp) {
		t.Fatalf("Reversal failed: %v", sxp)
	}
}

func TestISliceAppend(t *testing.T) {
	ConfirmAppend := func(s *ISlice, v interface{}, r *ISlice) {
		s.Append(v)
		if !r.Equal(s) {
			t.Fatalf("Append(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmAppend(IList(), 0, IList(0))
}

func TestISliceAppendSlice(t *testing.T) {
	ConfirmAppendSlice := func(s, v, r *ISlice) {
		s.AppendSlice(*v)
		if !r.Equal(s) {
			t.Fatalf("AppendSlice(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmAppendSlice(IList(), IList(0), IList(0))
	ConfirmAppendSlice(IList(), IList(0, 1), IList(0, 1))
	ConfirmAppendSlice(IList(0, 1, 2), IList(3, 4), IList(0, 1, 2, 3, 4))
}

func TestISlicePrepend(t *testing.T) {
	ConfirmPrepend := func(s *ISlice, v interface{}, r *ISlice) {
		if s.Prepend(v); !r.Equal(s) {
			t.Fatalf("Prepend(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmPrepend(IList(), 0, IList(0))
	ConfirmPrepend(IList(0), 1, IList(1, 0))
}

func TestISlicePrependSlice(t *testing.T) {
	ConfirmPrependSlice := func(s, v, r *ISlice) {
		if s.PrependSlice(*v); !r.Equal(s) {
			t.Fatalf("PrependSlice(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmPrependSlice(IList(), IList(0), IList(0))
	ConfirmPrependSlice(IList(), IList(0, 1), IList(0, 1))
	ConfirmPrependSlice(IList(0, 1, 2), IList(3, 4), IList(3, 4, 0, 1, 2))
}

func TestISliceRepeat(t *testing.T) {
	ConfirmRepeat := func(s *ISlice, count int, r *ISlice) {
		if x := s.Repeat(count); !x.Equal(r) {
			t.Fatalf("%v.Repeat(%v) should be %v but is %v", s, count, r, x)
		}
	}

	ConfirmRepeat(IList(), 5, IList())
	ConfirmRepeat(IList(0), 1, IList(0))
	ConfirmRepeat(IList(0), 2, IList(0, 0))
	ConfirmRepeat(IList(0), 3, IList(0, 0, 0))
	ConfirmRepeat(IList(0), 4, IList(0, 0, 0, 0))
	ConfirmRepeat(IList(0), 5, IList(0, 0, 0, 0, 0))
}

func TestISliceCar(t *testing.T) {
	ConfirmCar := func(s *ISlice, r int) {
		n := s.Car()
		if ok := n == r; !ok {
			t.Fatalf("head should be '%v' but is '%v'", r, n)
		}
	}
	ConfirmCar(IList(1, 2, 3), 1)
}

func TestISliceCdr(t *testing.T) {
	ConfirmCdr := func(s, r *ISlice) {
		if n := s.Cdr(); !n.Equal(r) {
			t.Fatalf("tail should be '%v' but is '%v'", r, n)
		}
	}
	ConfirmCdr(IList(1, 2, 3), IList(2, 3))
}

func TestISliceRplaca(t *testing.T) {
	ConfirmRplaca := func(s *ISlice, v interface{}, r *ISlice) {
		if s.Rplaca(v); !s.Equal(r) {
			t.Fatalf("slice should be '%v' but is '%v'", r, s)
		}
	}
	ConfirmRplaca(IList(1, 2, 3, 4, 5), 0, IList(0, 2, 3, 4, 5))
}

func TestISliceRplacd(t *testing.T) {
	ConfirmRplacd := func(s *ISlice, v interface{}, r *ISlice) {
		if s.Rplacd(v); !s.Equal(r) {
			t.Fatalf("slice should be '%v' but is '%v'", r, s)
		}
	}
	ConfirmRplacd(IList(1, 2, 3, 4, 5), nil, IList(1))
	ConfirmRplacd(IList(1, 2, 3, 4, 5), 10, IList(1, 10))
	ConfirmRplacd(IList(1, 2, 3, 4, 5), IList(5, 4, 3, 2), IList(1, 5, 4, 3, 2))
	ConfirmRplacd(IList(1, 2, 3, 4, 5), IList(2, 4, 8, 16, 32), IList(1, 2, 4, 8, 16, 32))
}