package slices

import "testing"

func TestU32Slice(t *testing.T) {
	sxp := U32List(1)
	switch {
	case sxp.Len() != 1:			t.Fatalf("U32List(1) should allocate 1 cells, not %v cells", sxp.Len())
	case sxp.U32At(0) != 1:			t.Fatalf("U32List(1) element 0 should be 1 and not %v", sxp.U32At(0))
	}

	sxp = U32List(1, 2)
	switch {
	case sxp.Len() != 2:			t.Fatalf("U32List(1 2) should allocate 2 cells, not %v cells", sxp.Len())
	case sxp.U32At(0) != 1:			t.Fatalf("U32List(1 2) element 0 should be 1 and not %v", sxp.U32At(0))
	case sxp.U32At(1) != 2:			t.Fatalf("U32List(1 2) element 1 should be 2 and not %v", sxp.U32At(1))
	}

	sxp = U32List(1, 2, 3)
	switch {
	case sxp.Len() != 3:			t.Fatalf("U32List(1 2 3) should allocate 3 cells, not %v cells", sxp.Len())
	case sxp.U32At(0) != 1:			t.Fatalf("U32List(1 2 3) element 0 should be 1 and not %v", sxp.U32At(0))
	case sxp.U32At(1) != 2:			t.Fatalf("U32List(1 2 3) element 1 should be 2 and not %v", sxp.U32At(1))
	case sxp.U32At(2) != 3:			t.Fatalf("U32List(1 2 3) element 2 should be 3 and not %v", sxp.U32At(2))
	}
}

func TestU32SliceString(t *testing.T) {
	ConfirmString := func(s *U32Slice, r string) {
		if x := s.String(); x != r {
			t.Fatalf("%v erroneously serialised as '%v'", r, x)
		}
	}

	ConfirmString(U32List(), "()")
	ConfirmString(U32List(0), "(0)")
	ConfirmString(U32List(0, 1), "(0 1)")
}

func TestU32SliceLen(t *testing.T) {
	ConfirmLength := func(s *U32Slice, i int) {
		if x := s.Len(); x != i {
			t.Fatalf("%v.Len() should be %v but is %v", s, i, x)
		}
	}
	
	ConfirmLength(U32List(0), 1)
	ConfirmLength(U32List(0, 1), 2)
}

func TestU32SliceSwap(t *testing.T) {
	ConfirmSwap := func(s *U32Slice, i, j int, r *U32Slice) {
		if s.Swap(i, j); !r.Equal(s) {
			t.Fatalf("Swap(%v, %v) should be %v but is %v", i, j, r, s)
		}
	}
	ConfirmSwap(U32List(0, 1, 2), 0, 1, U32List(1, 0, 2))
	ConfirmSwap(U32List(0, 1, 2), 0, 2, U32List(2, 1, 0))
}

func TestU32SliceCompare(t *testing.T) {
	ConfirmCompare := func(s *U32Slice, i, j, r int) {
		if x := s.Compare(i, j); x != r {
			t.Fatalf("Compare(%v, %v) should be %v but is %v", i, j, r, x)
		}
	}

	ConfirmCompare(U32List(0, 1), 0, 0, IS_SAME_AS)
	ConfirmCompare(U32List(0, 1), 0, 1, IS_LESS_THAN)
	ConfirmCompare(U32List(0, 1), 1, 0, IS_GREATER_THAN)
}

func TestU32SliceZeroCompare(t *testing.T) {
	ConfirmCompare := func(s *U32Slice, i, r int) {
		if x := s.ZeroCompare(i); x != r {
			t.Fatalf("ZeroCompare(%v) should be %v but is %v", i, r, x)
		}
	}

	ConfirmCompare(U32List(0, 1, 2), 0, IS_SAME_AS)
	ConfirmCompare(U32List(0, 1, 2), 1, IS_LESS_THAN)
	ConfirmCompare(U32List(0, 1, 2), 2, IS_LESS_THAN)
}

func TestU32SliceCut(t *testing.T) {
	ConfirmCut := func(s *U32Slice, start, end int, r *U32Slice) {
		if s.Cut(start, end); !r.Equal(s) {
			t.Fatalf("Cut(%v, %v) should be %v but is %v", start, end, r, s)
		}
	}

	ConfirmCut(U32List(0, 1, 2, 3, 4, 5), 0, 1, U32List(1, 2, 3, 4, 5))
	ConfirmCut(U32List(0, 1, 2, 3, 4, 5), 1, 2, U32List(0, 2, 3, 4, 5))
	ConfirmCut(U32List(0, 1, 2, 3, 4, 5), 2, 3, U32List(0, 1, 3, 4, 5))
	ConfirmCut(U32List(0, 1, 2, 3, 4, 5), 3, 4, U32List(0, 1, 2, 4, 5))
	ConfirmCut(U32List(0, 1, 2, 3, 4, 5), 4, 5, U32List(0, 1, 2, 3, 5))
	ConfirmCut(U32List(0, 1, 2, 3, 4, 5), 5, 6, U32List(0, 1, 2, 3, 4))

	ConfirmCut(U32List(0, 1, 2, 3, 4, 5), -1, 1, U32List(1, 2, 3, 4, 5))
	ConfirmCut(U32List(0, 1, 2, 3, 4, 5), 0, 2, U32List(2, 3, 4, 5))
	ConfirmCut(U32List(0, 1, 2, 3, 4, 5), 1, 3, U32List(0, 3, 4, 5))
	ConfirmCut(U32List(0, 1, 2, 3, 4, 5), 2, 4, U32List(0, 1, 4, 5))
	ConfirmCut(U32List(0, 1, 2, 3, 4, 5), 3, 5, U32List(0, 1, 2, 5))
	ConfirmCut(U32List(0, 1, 2, 3, 4, 5), 4, 6, U32List(0, 1, 2, 3))
	ConfirmCut(U32List(0, 1, 2, 3, 4, 5), 5, 7, U32List(0, 1, 2, 3, 4))
}

func TestU32SliceTrim(t *testing.T) {
	ConfirmTrim := func(s *U32Slice, start, end int, r *U32Slice) {
		if s.Trim(start, end); !r.Equal(s) {
			t.Fatalf("Cut(%v, %v) should be %v but is %v", start, end, r, s)
		}
	}

	ConfirmTrim(U32List(0, 1, 2, 3, 4, 5), 0, 1, U32List(0))
	ConfirmTrim(U32List(0, 1, 2, 3, 4, 5), 1, 2, U32List(1))
	ConfirmTrim(U32List(0, 1, 2, 3, 4, 5), 2, 3, U32List(2))
	ConfirmTrim(U32List(0, 1, 2, 3, 4, 5), 3, 4, U32List(3))
	ConfirmTrim(U32List(0, 1, 2, 3, 4, 5), 4, 5, U32List(4))
	ConfirmTrim(U32List(0, 1, 2, 3, 4, 5), 5, 6, U32List(5))

	ConfirmTrim(U32List(0, 1, 2, 3, 4, 5), -1, 1, U32List(0))
	ConfirmTrim(U32List(0, 1, 2, 3, 4, 5), 0, 2, U32List(0, 1))
	ConfirmTrim(U32List(0, 1, 2, 3, 4, 5), 1, 3, U32List(1, 2))
	ConfirmTrim(U32List(0, 1, 2, 3, 4, 5), 2, 4, U32List(2, 3))
	ConfirmTrim(U32List(0, 1, 2, 3, 4, 5), 3, 5, U32List(3, 4))
	ConfirmTrim(U32List(0, 1, 2, 3, 4, 5), 4, 6, U32List(4, 5))
	ConfirmTrim(U32List(0, 1, 2, 3, 4, 5), 5, 7, U32List(5))
}

func TestU32SliceDelete(t *testing.T) {
	ConfirmCut := func(s *U32Slice, index int, r *U32Slice) {
		if s.Delete(index); !r.Equal(s) {
			t.Fatalf("Delete(%v) should be %v but is %v", index, r, s)
		}
	}

	ConfirmCut(U32List(0, 1, 2, 3, 4, 5), -1, U32List(0, 1, 2, 3, 4, 5))
	ConfirmCut(U32List(0, 1, 2, 3, 4, 5), 0, U32List(1, 2, 3, 4, 5))
	ConfirmCut(U32List(0, 1, 2, 3, 4, 5), 1, U32List(0, 2, 3, 4, 5))
	ConfirmCut(U32List(0, 1, 2, 3, 4, 5), 2, U32List(0, 1, 3, 4, 5))
	ConfirmCut(U32List(0, 1, 2, 3, 4, 5), 3, U32List(0, 1, 2, 4, 5))
	ConfirmCut(U32List(0, 1, 2, 3, 4, 5), 4, U32List(0, 1, 2, 3, 5))
	ConfirmCut(U32List(0, 1, 2, 3, 4, 5), 5, U32List(0, 1, 2, 3, 4))
	ConfirmCut(U32List(0, 1, 2, 3, 4, 5), 6, U32List(0, 1, 2, 3, 4, 5))
}

func TestU32SliceEach(t *testing.T) {
	var count	uint32
	U32List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).Each(func(i interface{}) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})
}

func TestU32SliceEachWithIndex(t *testing.T) {
	U32List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).EachWithIndex(func(index int, i interface{}) {
		if i != uint32(index) {
			t.Fatalf("element %v erroneously reported as %v", index, i)
		}
	})
}

func TestU32SliceEachWithKey(t *testing.T) {
	U32List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).EachWithKey(func(key, i interface{}) {
		if i != uint32(key.(int)) {
			t.Fatalf("element %v erroneously reported as %v", key, i)
		}
	})
}

func TestU32SliceU32Each(t *testing.T) {
	var count	uint32
	U32List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).U32Each(func(i uint32) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})
}

func TestU32SliceU32EachWithIndex(t *testing.T) {
	U32List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).U32EachWithIndex(func(index int, i uint32) {
		if i != uint32(index) {
			t.Fatalf("element %v erroneously reported as %v", index, i)
		}
	})
}

func TestU32SliceU32EachWithKey(t *testing.T) {
	c := U32List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9)
	c.U32EachWithKey(func(key interface{}, i uint32) {
		if i != uint32(key.(int)) {
			t.Fatalf("element %v erroneously reported as %v", key, i)
		}
	})
}

func TestU32SliceBlockCopy(t *testing.T) {
	ConfirmBlockCopy := func(s *U32Slice, destination, source, count int, r *U32Slice) {
		s.BlockCopy(destination, source, count)
		if !r.Equal(s) {
			t.Fatalf("BlockCopy(%v, %v, %v) should be %v but is %v", destination, source, count, r, s)
		}
	}

	ConfirmBlockCopy(U32List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, 0, 4, U32List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockCopy(U32List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 9, 9, 4, U32List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockCopy(U32List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 2, 4, U32List(0, 1, 2, 3, 4, 2, 3, 4, 5, 9))
	ConfirmBlockCopy(U32List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 2, 5, 4, U32List(0, 1, 5, 6, 7, 8, 6, 7, 8, 9))
}

func TestU32SliceBlockClear(t *testing.T) {
	ConfirmBlockClear := func(s *U32Slice, start, count int, r *U32Slice) {
		s.BlockClear(start, count)
		if !r.Equal(s) {
			t.Fatalf("BlockClear(%v, %v) should be %v but is %v", start, count, r, s)
		}
	}

	ConfirmBlockClear(U32List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, 4, U32List(0, 0, 0, 0, 4, 5, 6, 7, 8, 9))
	ConfirmBlockClear(U32List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, 4, U32List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockClear(U32List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 4, U32List(0, 1, 2, 3, 4, 0, 0, 0, 0, 9))
}

func TestU32SliceOverwrite(t *testing.T) {
	ConfirmOverwrite := func(s *U32Slice, offset int, v, r *U32Slice) {
		s.Overwrite(offset, *v)
		if !r.Equal(s) {
			t.Fatalf("Overwrite(%v, %v) should be %v but is %v", offset, v, r, s)
		}
	}

	ConfirmOverwrite(U32List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, U32List(10, 9, 8, 7), U32List(10, 9, 8, 7, 4, 5, 6, 7, 8, 9))
	ConfirmOverwrite(U32List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, U32List(10, 9, 8, 7), U32List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmOverwrite(U32List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, U32List(11, 12, 13, 14), U32List(0, 1, 2, 3, 4, 11, 12, 13, 14, 9))
}

func TestU32SliceReallocate(t *testing.T) {
	ConfirmReallocate := func(s *U32Slice, l, c int, r *U32Slice) {
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

	ConfirmReallocate(U32List(), 0, 10, U32List())
	ConfirmReallocate(U32List(0, 1, 2, 3, 4), 3, 10, U32List(0, 1, 2))
	ConfirmReallocate(U32List(0, 1, 2, 3, 4), 5, 10, U32List(0, 1, 2, 3, 4))
	ConfirmReallocate(U32List(0, 1, 2, 3, 4), 10, 10, U32List(0, 1, 2, 3, 4, 0, 0, 0, 0, 0))
	ConfirmReallocate(U32List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 1, 5, U32List(0))
	ConfirmReallocate(U32List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 5, U32List(0, 1, 2, 3, 4))
	ConfirmReallocate(U32List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, 5, U32List(0, 1, 2, 3, 4))
}

func TestU32SliceExtend(t *testing.T) {
	ConfirmExtend := func(s *U32Slice, n int, r *U32Slice) {
		c := s.Cap()
		s.Extend(n)
		switch {
		case s.Len() != r.Len():	t.Fatalf("Extend(%v) len should be %v but is %v", n, r.Len(), s.Len())
		case s.Cap() != c + n:		t.Fatalf("Extend(%v) cap should be %v but is %v", n, c + n, s.Cap())
		case !r.Equal(s):			t.Fatalf("Extend(%v) should be %v but is %v", n, r, s)
		}
	}

	ConfirmExtend(U32List(), 1, U32List(0))
	ConfirmExtend(U32List(), 2, U32List(0, 0))
}

func TestU32SliceExpand(t *testing.T) {
	ConfirmExpand := func(s *U32Slice, i, n int, r *U32Slice) {
		c := s.Cap()
		s.Expand(i, n)
		switch {
		case s.Len() != r.Len():	t.Fatalf("Expand(%v, %v) len should be %v but is %v", i, n, r.Len(), s.Len())
		case s.Cap() != c + n:		t.Fatalf("Expand(%v, %v) cap should be %v but is %v", i, n, c + n, s.Cap())
		case !r.Equal(s):			t.Fatalf("Expand(%v, %v) should be %v but is %v", i, n, r, s)
		}
	}

	ConfirmExpand(U32List(), -1, 1, U32List(0))
	ConfirmExpand(U32List(), 0, 1, U32List(0))
	ConfirmExpand(U32List(), 1, 1, U32List(0))
	ConfirmExpand(U32List(), 0, 2, U32List(0, 0))

	ConfirmExpand(U32List(0, 1, 2), -1, 2, U32List(0, 0, 0, 1, 2))
	ConfirmExpand(U32List(0, 1, 2), 0, 2, U32List(0, 0, 0, 1, 2))
	ConfirmExpand(U32List(0, 1, 2), 1, 2, U32List(0, 0, 0, 1, 2))
	ConfirmExpand(U32List(0, 1, 2), 2, 2, U32List(0, 1, 0, 0, 2))
	ConfirmExpand(U32List(0, 1, 2), 3, 2, U32List(0, 1, 2, 0, 0))
	ConfirmExpand(U32List(0, 1, 2), 4, 2, U32List(0, 1, 2, 0, 0))
}

func TestU32SliceDepth(t *testing.T) {
	ConfirmDepth := func(s *U32Slice, i int) {
		if x := s.Depth(); x != i {
			t.Fatalf("%v.Depth() should be %v but is %v", s, i, x)
		}
	}
	ConfirmDepth(U32List(0, 1), 0)
}

func TestU32SliceReverse(t *testing.T) {
	sxp := U32List(1, 2, 3, 4, 5)
	rxp := U32List(5, 4, 3, 2, 1)
	sxp.Reverse()
	if !rxp.Equal(sxp) {
		t.Fatalf("Reversal failed: %v", sxp)
	}
}

func TestU32SliceAppend(t *testing.T) {
	ConfirmAppend := func(s *U32Slice, v interface{}, r *U32Slice) {
		s.Append(v)
		if !r.Equal(s) {
			t.Fatalf("Append(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmAppend(U32List(), uint32(0), U32List(0))
}

func TestU32SliceAppendSlice(t *testing.T) {
	ConfirmAppendSlice := func(s, v, r *U32Slice) {
		s.AppendSlice(*v)
		if !r.Equal(s) {
			t.Fatalf("AppendSlice(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmAppendSlice(U32List(), U32List(0), U32List(0))
	ConfirmAppendSlice(U32List(), U32List(0, 1), U32List(0, 1))
	ConfirmAppendSlice(U32List(0, 1, 2), U32List(3, 4), U32List(0, 1, 2, 3, 4))
}

func TestU32SlicePrepend(t *testing.T) {
	ConfirmPrepend := func(s *U32Slice, v interface{}, r *U32Slice) {
		if s.Prepend(v); !r.Equal(s) {
			t.Fatalf("Prepend(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmPrepend(U32List(), uint32(0), U32List(0))
	ConfirmPrepend(U32List(0), uint32(1), U32List(1, 0))
}

func TestU32SlicePrependSlice(t *testing.T) {
	ConfirmPrependSlice := func(s, v, r *U32Slice) {
		if s.PrependSlice(*v); !r.Equal(s) {
			t.Fatalf("PrependSlice(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmPrependSlice(U32List(), U32List(0), U32List(0))
	ConfirmPrependSlice(U32List(), U32List(0, 1), U32List(0, 1))
	ConfirmPrependSlice(U32List(0, 1, 2), U32List(3, 4), U32List(3, 4, 0, 1, 2))
}

func TestU32SliceRepeat(t *testing.T) {
	ConfirmRepeat := func(s *U32Slice, count int, r *U32Slice) {
		if x := s.Repeat(count); !x.Equal(r) {
			t.Fatalf("%v.Repeat(%v) should be %v but is %v", s, count, r, x)
		}
	}

	ConfirmRepeat(U32List(), 5, U32List())
	ConfirmRepeat(U32List(0), 1, U32List(0))
	ConfirmRepeat(U32List(0), 2, U32List(0, 0))
	ConfirmRepeat(U32List(0), 3, U32List(0, 0, 0))
	ConfirmRepeat(U32List(0), 4, U32List(0, 0, 0, 0))
	ConfirmRepeat(U32List(0), 5, U32List(0, 0, 0, 0, 0))
}

func TestU32SliceCar(t *testing.T) {
	ConfirmCar := func(s *U32Slice, r uint32) {
		n := s.Car().(uint32)
		if ok := n == r; !ok {
			t.Fatalf("head should be '%v' but is '%v'", r, n)
		}
	}
	ConfirmCar(U32List(1, 2, 3), 1)
}

func TestU32SliceCdr(t *testing.T) {
	ConfirmCdr := func(s, r *U32Slice) {
		if n := s.Cdr(); !n.Equal(r) {
			t.Fatalf("tail should be '%v' but is '%v'", r, n)
		}
	}
	ConfirmCdr(U32List(1, 2, 3), U32List(2, 3))
}

func TestU32SliceRplaca(t *testing.T) {
	ConfirmRplaca := func(s *U32Slice, v interface{}, r *U32Slice) {
		if s.Rplaca(v); !s.Equal(r) {
			t.Fatalf("slice should be '%v' but is '%v'", r, s)
		}
	}
	ConfirmRplaca(U32List(1, 2, 3, 4, 5), uint32(0), U32List(0, 2, 3, 4, 5))
}

func TestU32SliceRplacd(t *testing.T) {
	ConfirmRplacd := func(s *U32Slice, v interface{}, r *U32Slice) {
		if s.Rplacd(v); !s.Equal(r) {
			t.Fatalf("slice should be '%v' but is '%v'", r, s)
		}
	}
	ConfirmRplacd(U32List(1, 2, 3, 4, 5), nil, U32List(1))
	ConfirmRplacd(U32List(1, 2, 3, 4, 5), uint32(10), U32List(1, 10))
	ConfirmRplacd(U32List(1, 2, 3, 4, 5), U32List(5, 4, 3, 2), U32List(1, 5, 4, 3, 2))
	ConfirmRplacd(U32List(1, 2, 3, 4, 5, 6), U32List(2, 4, 8, 16), U32List(1, 2, 4, 8, 16))
}