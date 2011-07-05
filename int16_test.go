package slices

import "testing"

func TestI16Slice(t *testing.T) {
	sxp := I16List(1)
	switch {
	case sxp.Len() != 1:			t.Fatalf("I16List(1) should allocate 1 cells, not %v cells", sxp.Len())
	case sxp.I16At(0) != 1:			t.Fatalf("I16List(1) element 0 should be 1 and not %v", sxp.I16At(0))
	}

	sxp = I16List(1, 2)
	switch {
	case sxp.Len() != 2:			t.Fatalf("I16List(1 2) should allocate 2 cells, not %v cells", sxp.Len())
	case sxp.I16At(0) != 1:			t.Fatalf("I16List(1 2) element 0 should be 1 and not %v", sxp.I16At(0))
	case sxp.I16At(1) != 2:			t.Fatalf("I16List(1 2) element 1 should be 2 and not %v", sxp.I16At(1))
	}

	sxp = I16List(1, 2, 3)
	switch {
	case sxp.Len() != 3:			t.Fatalf("I16List(1 2 3) should allocate 3 cells, not %v cells", sxp.Len())
	case sxp.I16At(0) != 1:			t.Fatalf("I16List(1 2 3) element 0 should be 1 and not %v", sxp.I16At(0))
	case sxp.I16At(1) != 2:			t.Fatalf("I16List(1 2 3) element 1 should be 2 and not %v", sxp.I16At(1))
	case sxp.I16At(2) != 3:			t.Fatalf("I16List(1 2 3) element 2 should be 3 and not %v", sxp.I16At(2))
	}
}

func TestI16SliceString(t *testing.T) {
	ConfirmString := func(s *I16Slice, r string) {
		if x := s.String(); x != r {
			t.Fatalf("%v erroneously serialised as '%v'", r, x)
		}
	}

	ConfirmString(I16List(), "()")
	ConfirmString(I16List(0), "(0)")
	ConfirmString(I16List(0, 1), "(0 1)")
}

func TestI16SliceLen(t *testing.T) {
	ConfirmLength := func(s *I16Slice, i int) {
		if x := s.Len(); x != i {
			t.Fatalf("%v.Len() should be %v but is %v", s, i, x)
		}
	}
	
	ConfirmLength(I16List(0), 1)
	ConfirmLength(I16List(0, 1), 2)
}

func TestI16SliceSwap(t *testing.T) {
	ConfirmSwap := func(s *I16Slice, i, j int, r *I16Slice) {
		if s.Swap(i, j); !r.Equal(s) {
			t.Fatalf("Swap(%v, %v) should be %v but is %v", i, j, r, s)
		}
	}
	ConfirmSwap(I16List(0, 1, 2), 0, 1, I16List(1, 0, 2))
	ConfirmSwap(I16List(0, 1, 2), 0, 2, I16List(2, 1, 0))
}

func TestI16SliceSort(t *testing.T) {
	ConfirmSort := func(s, r *I16Slice) {
		if s.Sort(); !r.Equal(s) {
			t.Fatalf("Sort() should be %v but is %v", r, s)
		}
	}

	ConfirmSort(I16List(3, 2, 1, 4, 5, 0), I16List(0, 1, 2, 3, 4, 5))
}

func TestI16SliceCompare(t *testing.T) {
	ConfirmCompare := func(s *I16Slice, i, j, r int) {
		if x := s.Compare(i, j); x != r {
			t.Fatalf("Compare(%v, %v) should be %v but is %v", i, j, r, x)
		}
	}

	ConfirmCompare(I16List(0, 1), 0, 0, IS_SAME_AS)
	ConfirmCompare(I16List(0, 1), 0, 1, IS_LESS_THAN)
	ConfirmCompare(I16List(0, 1), 1, 0, IS_GREATER_THAN)
}

func TestI16SliceZeroCompare(t *testing.T) {
	ConfirmCompare := func(s *I16Slice, i, r int) {
		if x := s.ZeroCompare(i); x != r {
			t.Fatalf("ZeroCompare(%v) should be %v but is %v", i, r, x)
		}
	}

	ConfirmCompare(I16List(0, -1, 1), 0, IS_SAME_AS)
	ConfirmCompare(I16List(0, -1, 1), 1, IS_GREATER_THAN)
	ConfirmCompare(I16List(0, -1, 1), 2, IS_LESS_THAN)
}

func TestI16SliceCut(t *testing.T) {
	ConfirmCut := func(s *I16Slice, start, end int, r *I16Slice) {
		if s.Cut(start, end); !r.Equal(s) {
			t.Fatalf("Cut(%v, %v) should be %v but is %v", start, end, r, s)
		}
	}

	ConfirmCut(I16List(0, 1, 2, 3, 4, 5), 0, 1, I16List(1, 2, 3, 4, 5))
	ConfirmCut(I16List(0, 1, 2, 3, 4, 5), 1, 2, I16List(0, 2, 3, 4, 5))
	ConfirmCut(I16List(0, 1, 2, 3, 4, 5), 2, 3, I16List(0, 1, 3, 4, 5))
	ConfirmCut(I16List(0, 1, 2, 3, 4, 5), 3, 4, I16List(0, 1, 2, 4, 5))
	ConfirmCut(I16List(0, 1, 2, 3, 4, 5), 4, 5, I16List(0, 1, 2, 3, 5))
	ConfirmCut(I16List(0, 1, 2, 3, 4, 5), 5, 6, I16List(0, 1, 2, 3, 4))

	ConfirmCut(I16List(0, 1, 2, 3, 4, 5), -1, 1, I16List(1, 2, 3, 4, 5))
	ConfirmCut(I16List(0, 1, 2, 3, 4, 5), 0, 2, I16List(2, 3, 4, 5))
	ConfirmCut(I16List(0, 1, 2, 3, 4, 5), 1, 3, I16List(0, 3, 4, 5))
	ConfirmCut(I16List(0, 1, 2, 3, 4, 5), 2, 4, I16List(0, 1, 4, 5))
	ConfirmCut(I16List(0, 1, 2, 3, 4, 5), 3, 5, I16List(0, 1, 2, 5))
	ConfirmCut(I16List(0, 1, 2, 3, 4, 5), 4, 6, I16List(0, 1, 2, 3))
	ConfirmCut(I16List(0, 1, 2, 3, 4, 5), 5, 7, I16List(0, 1, 2, 3, 4))
}

func TestI16SliceTrim(t *testing.T) {
	ConfirmTrim := func(s *I16Slice, start, end int, r *I16Slice) {
		if s.Trim(start, end); !r.Equal(s) {
			t.Fatalf("Trim(%v, %v) should be %v but is %v", start, end, r, s)
		}
	}

	ConfirmTrim(I16List(0, 1, 2, 3, 4, 5), 0, 1, I16List(0))
	ConfirmTrim(I16List(0, 1, 2, 3, 4, 5), 1, 2, I16List(1))
	ConfirmTrim(I16List(0, 1, 2, 3, 4, 5), 2, 3, I16List(2))
	ConfirmTrim(I16List(0, 1, 2, 3, 4, 5), 3, 4, I16List(3))
	ConfirmTrim(I16List(0, 1, 2, 3, 4, 5), 4, 5, I16List(4))
	ConfirmTrim(I16List(0, 1, 2, 3, 4, 5), 5, 6, I16List(5))

	ConfirmTrim(I16List(0, 1, 2, 3, 4, 5), -1, 1, I16List(0))
	ConfirmTrim(I16List(0, 1, 2, 3, 4, 5), 0, 2, I16List(0, 1))
	ConfirmTrim(I16List(0, 1, 2, 3, 4, 5), 1, 3, I16List(1, 2))
	ConfirmTrim(I16List(0, 1, 2, 3, 4, 5), 2, 4, I16List(2, 3))
	ConfirmTrim(I16List(0, 1, 2, 3, 4, 5), 3, 5, I16List(3, 4))
	ConfirmTrim(I16List(0, 1, 2, 3, 4, 5), 4, 6, I16List(4, 5))
	ConfirmTrim(I16List(0, 1, 2, 3, 4, 5), 5, 7, I16List(5))
}

func TestI16SliceDelete(t *testing.T) {
	ConfirmCut := func(s *I16Slice, index int, r *I16Slice) {
		if s.Delete(index); !r.Equal(s) {
			t.Fatalf("Delete(%v) should be %v but is %v", index, r, s)
		}
	}

	ConfirmCut(I16List(0, 1, 2, 3, 4, 5), -1, I16List(0, 1, 2, 3, 4, 5))
	ConfirmCut(I16List(0, 1, 2, 3, 4, 5), 0, I16List(1, 2, 3, 4, 5))
	ConfirmCut(I16List(0, 1, 2, 3, 4, 5), 1, I16List(0, 2, 3, 4, 5))
	ConfirmCut(I16List(0, 1, 2, 3, 4, 5), 2, I16List(0, 1, 3, 4, 5))
	ConfirmCut(I16List(0, 1, 2, 3, 4, 5), 3, I16List(0, 1, 2, 4, 5))
	ConfirmCut(I16List(0, 1, 2, 3, 4, 5), 4, I16List(0, 1, 2, 3, 5))
	ConfirmCut(I16List(0, 1, 2, 3, 4, 5), 5, I16List(0, 1, 2, 3, 4))
	ConfirmCut(I16List(0, 1, 2, 3, 4, 5), 6, I16List(0, 1, 2, 3, 4, 5))
}

func TestI16SliceEach(t *testing.T) {
	var count	int16
	I16List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).Each(func(i interface{}) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})
}

func TestI16SliceEachWithIndex(t *testing.T) {
	I16List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).EachWithIndex(func(index int, i interface{}) {
		if i != int16(index) {
			t.Fatalf("element %v erroneously reported as %v", index, i)
		}
	})
}

func TestI16SliceEachWithKey(t *testing.T) {
	I16List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).EachWithKey(func(key, i interface{}) {
		if i != int16(key.(int)) {
			t.Fatalf("element %v erroneously reported as %v", key, i)
		}
	})
}

func TestI16SliceI16Each(t *testing.T) {
	var count	int16
	I16List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).I16Each(func(i int16) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})
}

func TestI16SliceI16EachWithIndex(t *testing.T) {
	I16List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).I16EachWithIndex(func(index int, i int16) {
		if i != int16(index) {
			t.Fatalf("element %v erroneously reported as %v", index, i)
		}
	})
}

func TestI16SliceI16EachWithKey(t *testing.T) {
	c := I16List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9)
	c.I16EachWithKey(func(key interface{}, i int16) {
		if i != int16(key.(int)) {
			t.Fatalf("element %v erroneously reported as %v", key, i)
		}
	})
}

func TestI16SliceBlockCopy(t *testing.T) {
	ConfirmBlockCopy := func(s *I16Slice, destination, source, count int, r *I16Slice) {
		s.BlockCopy(destination, source, count)
		if !r.Equal(s) {
			t.Fatalf("BlockCopy(%v, %v, %v) should be %v but is %v", destination, source, count, r, s)
		}
	}

	ConfirmBlockCopy(I16List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, 0, 4, I16List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockCopy(I16List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 9, 9, 4, I16List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockCopy(I16List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 2, 4, I16List(0, 1, 2, 3, 4, 2, 3, 4, 5, 9))
	ConfirmBlockCopy(I16List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 2, 5, 4, I16List(0, 1, 5, 6, 7, 8, 6, 7, 8, 9))
}

func TestI16SliceBlockClear(t *testing.T) {
	ConfirmBlockClear := func(s *I16Slice, start, count int, r *I16Slice) {
		s.BlockClear(start, count)
		if !r.Equal(s) {
			t.Fatalf("BlockClear(%v, %v) should be %v but is %v", start, count, r, s)
		}
	}

	ConfirmBlockClear(I16List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, 4, I16List(0, 0, 0, 0, 4, 5, 6, 7, 8, 9))
	ConfirmBlockClear(I16List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, 4, I16List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockClear(I16List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 4, I16List(0, 1, 2, 3, 4, 0, 0, 0, 0, 9))
}

func TestI16SliceOverwrite(t *testing.T) {
	ConfirmOverwrite := func(s *I16Slice, offset int, v, r *I16Slice) {
		s.Overwrite(offset, *v)
		if !r.Equal(s) {
			t.Fatalf("Overwrite(%v, %v) should be %v but is %v", offset, v, r, s)
		}
	}

	ConfirmOverwrite(I16List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, I16List(10, 9, 8, 7), I16List(10, 9, 8, 7, 4, 5, 6, 7, 8, 9))
	ConfirmOverwrite(I16List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, I16List(10, 9, 8, 7), I16List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmOverwrite(I16List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, I16List(11, 12, 13, 14), I16List(0, 1, 2, 3, 4, 11, 12, 13, 14, 9))
}

func TestI16SliceReallocate(t *testing.T) {
	ConfirmReallocate := func(s *I16Slice, l, c int, r *I16Slice) {
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

	ConfirmReallocate(I16List(), 0, 10, I16List())
	ConfirmReallocate(I16List(0, 1, 2, 3, 4), 3, 10, I16List(0, 1, 2))
	ConfirmReallocate(I16List(0, 1, 2, 3, 4), 5, 10, I16List(0, 1, 2, 3, 4))
	ConfirmReallocate(I16List(0, 1, 2, 3, 4), 10, 10, I16List(0, 1, 2, 3, 4, 0, 0, 0, 0, 0))
	ConfirmReallocate(I16List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 1, 5, I16List(0))
	ConfirmReallocate(I16List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 5, I16List(0, 1, 2, 3, 4))
	ConfirmReallocate(I16List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, 5, I16List(0, 1, 2, 3, 4))
}

func TestI16SliceExtend(t *testing.T) {
	ConfirmExtend := func(s *I16Slice, n int, r *I16Slice) {
		c := s.Cap()
		s.Extend(n)
		switch {
		case s.Len() != r.Len():	t.Fatalf("Extend(%v) len should be %v but is %v", n, r.Len(), s.Len())
		case s.Cap() != c + n:		t.Fatalf("Extend(%v) cap should be %v but is %v", n, c + n, s.Cap())
		case !r.Equal(s):			t.Fatalf("Extend(%v) should be %v but is %v", n, r, s)
		}
	}

	ConfirmExtend(I16List(), 1, I16List(0))
	ConfirmExtend(I16List(), 2, I16List(0, 0))
}

func TestI16SliceExpand(t *testing.T) {
	ConfirmExpand := func(s *I16Slice, i, n int, r *I16Slice) {
		c := s.Cap()
		s.Expand(i, n)
		switch {
		case s.Len() != r.Len():	t.Fatalf("Expand(%v, %v) len should be %v but is %v", i, n, r.Len(), s.Len())
		case s.Cap() != c + n:		t.Fatalf("Expand(%v, %v) cap should be %v but is %v", i, n, c + n, s.Cap())
		case !r.Equal(s):			t.Fatalf("Expand(%v, %v) should be %v but is %v", i, n, r, s)
		}
	}

	ConfirmExpand(I16List(), -1, 1, I16List(0))
	ConfirmExpand(I16List(), 0, 1, I16List(0))
	ConfirmExpand(I16List(), 1, 1, I16List(0))
	ConfirmExpand(I16List(), 0, 2, I16List(0, 0))

	ConfirmExpand(I16List(0, 1, 2), -1, 2, I16List(0, 0, 0, 1, 2))
	ConfirmExpand(I16List(0, 1, 2), 0, 2, I16List(0, 0, 0, 1, 2))
	ConfirmExpand(I16List(0, 1, 2), 1, 2, I16List(0, 0, 0, 1, 2))
	ConfirmExpand(I16List(0, 1, 2), 2, 2, I16List(0, 1, 0, 0, 2))
	ConfirmExpand(I16List(0, 1, 2), 3, 2, I16List(0, 1, 2, 0, 0))
	ConfirmExpand(I16List(0, 1, 2), 4, 2, I16List(0, 1, 2, 0, 0))
}

func TestI16SliceDepth(t *testing.T) {
	ConfirmDepth := func(s *I16Slice, i int) {
		if x := s.Depth(); x != i {
			t.Fatalf("%v.Depth() should be %v but is %v", s, i, x)
		}
	}
	ConfirmDepth(I16List(0, 1), 0)
}

func TestI16SliceReverse(t *testing.T) {
	sxp := I16List(1, 2, 3, 4, 5)
	rxp := I16List(5, 4, 3, 2, 1)
	sxp.Reverse()
	if !rxp.Equal(sxp) {
		t.Fatalf("Reversal failed: %v", sxp)
	}
}

func TestI16SliceAppend(t *testing.T) {
	ConfirmAppend := func(s *I16Slice, v interface{}, r *I16Slice) {
		s.Append(v)
		if !r.Equal(s) {
			t.Fatalf("Append(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmAppend(I16List(), int16(0), I16List(0))
}

func TestI16SliceAppendSlice(t *testing.T) {
	ConfirmAppendSlice := func(s, v, r *I16Slice) {
		s.AppendSlice(*v)
		if !r.Equal(s) {
			t.Fatalf("AppendSlice(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmAppendSlice(I16List(), I16List(0), I16List(0))
	ConfirmAppendSlice(I16List(), I16List(0, 1), I16List(0, 1))
	ConfirmAppendSlice(I16List(0, 1, 2), I16List(3, 4), I16List(0, 1, 2, 3, 4))
}

func TestI16SlicePrepend(t *testing.T) {
	ConfirmPrepend := func(s *I16Slice, v interface{}, r *I16Slice) {
		if s.Prepend(v); !r.Equal(s) {
			t.Fatalf("Prepend(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmPrepend(I16List(), int16(0), I16List(0))
	ConfirmPrepend(I16List(0), int16(1), I16List(1, 0))
}

func TestI16SlicePrependSlice(t *testing.T) {
	ConfirmPrependSlice := func(s, v, r *I16Slice) {
		if s.PrependSlice(*v); !r.Equal(s) {
			t.Fatalf("PrependSlice(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmPrependSlice(I16List(), I16List(0), I16List(0))
	ConfirmPrependSlice(I16List(), I16List(0, 1), I16List(0, 1))
	ConfirmPrependSlice(I16List(0, 1, 2), I16List(3, 4), I16List(3, 4, 0, 1, 2))
}

func TestI16SliceRepeat(t *testing.T) {
	ConfirmRepeat := func(s *I16Slice, count int, r *I16Slice) {
		if x := s.Repeat(count); !x.Equal(r) {
			t.Fatalf("%v.Repeat(%v) should be %v but is %v", s, count, r, x)
		}
	}

	ConfirmRepeat(I16List(), 5, I16List())
	ConfirmRepeat(I16List(0), 1, I16List(0))
	ConfirmRepeat(I16List(0), 2, I16List(0, 0))
	ConfirmRepeat(I16List(0), 3, I16List(0, 0, 0))
	ConfirmRepeat(I16List(0), 4, I16List(0, 0, 0, 0))
	ConfirmRepeat(I16List(0), 5, I16List(0, 0, 0, 0, 0))
}

func TestI16SliceCar(t *testing.T) {
	ConfirmCar := func(s *I16Slice, r int16) {
		n := s.Car().(int16)
		if ok := n == r; !ok {
			t.Fatalf("head should be '%v' but is '%v'", r, n)
		}
	}
	ConfirmCar(I16List(1, 2, 3), 1)
}

func TestI16SliceCdr(t *testing.T) {
	ConfirmCdr := func(s, r *I16Slice) {
		if n := s.Cdr(); !n.Equal(r) {
			t.Fatalf("tail should be '%v' but is '%v'", r, n)
		}
	}
	ConfirmCdr(I16List(1, 2, 3), I16List(2, 3))
}

func TestI16SliceRplaca(t *testing.T) {
	ConfirmRplaca := func(s *I16Slice, v interface{}, r *I16Slice) {
		if s.Rplaca(v); !s.Equal(r) {
			t.Fatalf("slice should be '%v' but is '%v'", r, s)
		}
	}
	ConfirmRplaca(I16List(1, 2, 3, 4, 5), int16(0), I16List(0, 2, 3, 4, 5))
}

func TestI16SliceRplacd(t *testing.T) {
	ConfirmRplacd := func(s *I16Slice, v interface{}, r *I16Slice) {
		if s.Rplacd(v); !s.Equal(r) {
			t.Fatalf("slice should be '%v' but is '%v'", r, s)
		}
	}
	ConfirmRplacd(I16List(1, 2, 3, 4, 5), nil, I16List(1))
	ConfirmRplacd(I16List(1, 2, 3, 4, 5), int16(10), I16List(1, 10))
	ConfirmRplacd(I16List(1, 2, 3, 4, 5), I16List(5, 4, 3, 2), I16List(1, 5, 4, 3, 2))
	ConfirmRplacd(I16List(1, 2, 3, 4, 5, 6), I16List(2, 4, 8, 16), I16List(1, 2, 4, 8, 16))
}