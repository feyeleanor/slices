package slices

import "testing"

func TestI8Slice(t *testing.T) {
	sxp := I8List(1)
	switch {
	case sxp.Len() != 1:			t.Fatalf("I8List(1) should allocate 1 cells, not %v cells", sxp.Len())
	case sxp.I8At(0) != 1:			t.Fatalf("I8List(1) element 0 should be 1 and not %v", sxp.I8At(0))
	}

	sxp = I8List(1, 2)
	switch {
	case sxp.Len() != 2:			t.Fatalf("I8List(1 2) should allocate 2 cells, not %v cells", sxp.Len())
	case sxp.I8At(0) != 1:			t.Fatalf("I8List(1 2) element 0 should be 1 and not %v", sxp.I8At(0))
	case sxp.I8At(1) != 2:			t.Fatalf("I8List(1 2) element 1 should be 2 and not %v", sxp.I8At(1))
	}

	sxp = I8List(1, 2, 3)
	switch {
	case sxp.Len() != 3:			t.Fatalf("I8List(1 2 3) should allocate 3 cells, not %v cells", sxp.Len())
	case sxp.I8At(0) != 1:			t.Fatalf("I8List(1 2 3) element 0 should be 1 and not %v", sxp.I8At(0))
	case sxp.I8At(1) != 2:			t.Fatalf("I8List(1 2 3) element 1 should be 2 and not %v", sxp.I8At(1))
	case sxp.I8At(2) != 3:			t.Fatalf("I8List(1 2 3) element 2 should be 3 and not %v", sxp.I8At(2))
	}
}

func TestI8SliceString(t *testing.T) {
	ConfirmString := func(s *I8Slice, r string) {
		if x := s.String(); x != r {
			t.Fatalf("%v erroneously serialised as '%v'", r, x)
		}
	}

	ConfirmString(I8List(), "()")
	ConfirmString(I8List(0), "(0)")
	ConfirmString(I8List(0, 1), "(0 1)")
}

func TestI8SliceLen(t *testing.T) {
	ConfirmLength := func(s *I8Slice, i int) {
		if x := s.Len(); x != i {
			t.Fatalf("%v.Len() should be %v but is %v", s, i, x)
		}
	}
	
	ConfirmLength(I8List(0), 1)
	ConfirmLength(I8List(0, 1), 2)
}

func TestI8SliceSwap(t *testing.T) {
	ConfirmSwap := func(s *I8Slice, i, j int, r *I8Slice) {
		if s.Swap(i, j); !r.Equal(s) {
			t.Fatalf("Swap(%v, %v) should be %v but is %v", i, j, r, s)
		}
	}
	ConfirmSwap(I8List(0, 1, 2), 0, 1, I8List(1, 0, 2))
	ConfirmSwap(I8List(0, 1, 2), 0, 2, I8List(2, 1, 0))
}

func TestI8SliceSort(t *testing.T) {
	ConfirmSort := func(s, r *I8Slice) {
		if s.Sort(); !r.Equal(s) {
			t.Fatalf("Sort() should be %v but is %v", r, s)
		}
	}

	ConfirmSort(I8List(3, 2, 1, 4, 5, 0), I8List(0, 1, 2, 3, 4, 5))
}

func TestI8SliceCompare(t *testing.T) {
	ConfirmCompare := func(s *I8Slice, i, j, r int) {
		if x := s.Compare(i, j); x != r {
			t.Fatalf("Compare(%v, %v) should be %v but is %v", i, j, r, x)
		}
	}

	ConfirmCompare(I8List(0, 1), 0, 0, IS_SAME_AS)
	ConfirmCompare(I8List(0, 1), 0, 1, IS_LESS_THAN)
	ConfirmCompare(I8List(0, 1), 1, 0, IS_GREATER_THAN)
}

func TestI8SliceZeroCompare(t *testing.T) {
	ConfirmCompare := func(s *I8Slice, i, r int) {
		if x := s.ZeroCompare(i); x != r {
			t.Fatalf("ZeroCompare(%v) should be %v but is %v", i, r, x)
		}
	}

	ConfirmCompare(I8List(0, -1, 1), 0, IS_SAME_AS)
	ConfirmCompare(I8List(0, -1, 1), 1, IS_GREATER_THAN)
	ConfirmCompare(I8List(0, -1, 1), 2, IS_LESS_THAN)
}

func TestI8SliceCut(t *testing.T) {
	ConfirmCut := func(s *I8Slice, start, end int, r *I8Slice) {
		if s.Cut(start, end); !r.Equal(s) {
			t.Fatalf("Cut(%v, %v) should be %v but is %v", start, end, r, s)
		}
	}

	ConfirmCut(I8List(0, 1, 2, 3, 4, 5), 0, 1, I8List(1, 2, 3, 4, 5))
	ConfirmCut(I8List(0, 1, 2, 3, 4, 5), 1, 2, I8List(0, 2, 3, 4, 5))
	ConfirmCut(I8List(0, 1, 2, 3, 4, 5), 2, 3, I8List(0, 1, 3, 4, 5))
	ConfirmCut(I8List(0, 1, 2, 3, 4, 5), 3, 4, I8List(0, 1, 2, 4, 5))
	ConfirmCut(I8List(0, 1, 2, 3, 4, 5), 4, 5, I8List(0, 1, 2, 3, 5))
	ConfirmCut(I8List(0, 1, 2, 3, 4, 5), 5, 6, I8List(0, 1, 2, 3, 4))

	ConfirmCut(I8List(0, 1, 2, 3, 4, 5), -1, 1, I8List(1, 2, 3, 4, 5))
	ConfirmCut(I8List(0, 1, 2, 3, 4, 5), 0, 2, I8List(2, 3, 4, 5))
	ConfirmCut(I8List(0, 1, 2, 3, 4, 5), 1, 3, I8List(0, 3, 4, 5))
	ConfirmCut(I8List(0, 1, 2, 3, 4, 5), 2, 4, I8List(0, 1, 4, 5))
	ConfirmCut(I8List(0, 1, 2, 3, 4, 5), 3, 5, I8List(0, 1, 2, 5))
	ConfirmCut(I8List(0, 1, 2, 3, 4, 5), 4, 6, I8List(0, 1, 2, 3))
	ConfirmCut(I8List(0, 1, 2, 3, 4, 5), 5, 7, I8List(0, 1, 2, 3, 4))
}

func TestI8SliceTrim(t *testing.T) {
	ConfirmTrim := func(s *I8Slice, start, end int, r *I8Slice) {
		if s.Trim(start, end); !r.Equal(s) {
			t.Fatalf("Trim(%v, %v) should be %v but is %v", start, end, r, s)
		}
	}

	ConfirmTrim(I8List(0, 1, 2, 3, 4, 5), 0, 1, I8List(0))
	ConfirmTrim(I8List(0, 1, 2, 3, 4, 5), 1, 2, I8List(1))
	ConfirmTrim(I8List(0, 1, 2, 3, 4, 5), 2, 3, I8List(2))
	ConfirmTrim(I8List(0, 1, 2, 3, 4, 5), 3, 4, I8List(3))
	ConfirmTrim(I8List(0, 1, 2, 3, 4, 5), 4, 5, I8List(4))
	ConfirmTrim(I8List(0, 1, 2, 3, 4, 5), 5, 6, I8List(5))

	ConfirmTrim(I8List(0, 1, 2, 3, 4, 5), -1, 1, I8List(0))
	ConfirmTrim(I8List(0, 1, 2, 3, 4, 5), 0, 2, I8List(0, 1))
	ConfirmTrim(I8List(0, 1, 2, 3, 4, 5), 1, 3, I8List(1, 2))
	ConfirmTrim(I8List(0, 1, 2, 3, 4, 5), 2, 4, I8List(2, 3))
	ConfirmTrim(I8List(0, 1, 2, 3, 4, 5), 3, 5, I8List(3, 4))
	ConfirmTrim(I8List(0, 1, 2, 3, 4, 5), 4, 6, I8List(4, 5))
	ConfirmTrim(I8List(0, 1, 2, 3, 4, 5), 5, 7, I8List(5))
}

func TestI8SliceDelete(t *testing.T) {
	ConfirmCut := func(s *I8Slice, index int, r *I8Slice) {
		if s.Delete(index); !r.Equal(s) {
			t.Fatalf("Delete(%v) should be %v but is %v", index, r, s)
		}
	}

	ConfirmCut(I8List(0, 1, 2, 3, 4, 5), -1, I8List(0, 1, 2, 3, 4, 5))
	ConfirmCut(I8List(0, 1, 2, 3, 4, 5), 0, I8List(1, 2, 3, 4, 5))
	ConfirmCut(I8List(0, 1, 2, 3, 4, 5), 1, I8List(0, 2, 3, 4, 5))
	ConfirmCut(I8List(0, 1, 2, 3, 4, 5), 2, I8List(0, 1, 3, 4, 5))
	ConfirmCut(I8List(0, 1, 2, 3, 4, 5), 3, I8List(0, 1, 2, 4, 5))
	ConfirmCut(I8List(0, 1, 2, 3, 4, 5), 4, I8List(0, 1, 2, 3, 5))
	ConfirmCut(I8List(0, 1, 2, 3, 4, 5), 5, I8List(0, 1, 2, 3, 4))
	ConfirmCut(I8List(0, 1, 2, 3, 4, 5), 6, I8List(0, 1, 2, 3, 4, 5))
}

func TestI8SliceEach(t *testing.T) {
	var count	int8
	I8List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).Each(func(i interface{}) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})
}

func TestI8SliceEachWithIndex(t *testing.T) {
	I8List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).EachWithIndex(func(index int, i interface{}) {
		if i != int8(index) {
			t.Fatalf("element %v erroneously reported as %v", index, i)
		}
	})
}

func TestI8SliceEachWithKey(t *testing.T) {
	I8List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).EachWithKey(func(key, i interface{}) {
		if i != int8(key.(int)) {
			t.Fatalf("element %v erroneously reported as %v", key, i)
		}
	})
}

func TestI8SliceI8Each(t *testing.T) {
	var count	int8
	I8List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).I8Each(func(i int8) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})
}

func TestI8SliceI8EachWithIndex(t *testing.T) {
	I8List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).I8EachWithIndex(func(index int, i int8) {
		if i != int8(index) {
			t.Fatalf("element %v erroneously reported as %v", index, i)
		}
	})
}

func TestI8SliceI8EachWithKey(t *testing.T) {
	c := I8List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9)
	c.I8EachWithKey(func(key interface{}, i int8) {
		if i != int8(key.(int)) {
			t.Fatalf("element %v erroneously reported as %v", key, i)
		}
	})
}

func TestI8SliceBlockCopy(t *testing.T) {
	ConfirmBlockCopy := func(s *I8Slice, destination, source, count int, r *I8Slice) {
		s.BlockCopy(destination, source, count)
		if !r.Equal(s) {
			t.Fatalf("BlockCopy(%v, %v, %v) should be %v but is %v", destination, source, count, r, s)
		}
	}

	ConfirmBlockCopy(I8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, 0, 4, I8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockCopy(I8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 9, 9, 4, I8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockCopy(I8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 2, 4, I8List(0, 1, 2, 3, 4, 2, 3, 4, 5, 9))
	ConfirmBlockCopy(I8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 2, 5, 4, I8List(0, 1, 5, 6, 7, 8, 6, 7, 8, 9))
}

func TestI8SliceBlockClear(t *testing.T) {
	ConfirmBlockClear := func(s *I8Slice, start, count int, r *I8Slice) {
		s.BlockClear(start, count)
		if !r.Equal(s) {
			t.Fatalf("BlockClear(%v, %v) should be %v but is %v", start, count, r, s)
		}
	}

	ConfirmBlockClear(I8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, 4, I8List(0, 0, 0, 0, 4, 5, 6, 7, 8, 9))
	ConfirmBlockClear(I8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, 4, I8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockClear(I8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 4, I8List(0, 1, 2, 3, 4, 0, 0, 0, 0, 9))
}

func TestI8SliceOverwrite(t *testing.T) {
	ConfirmOverwrite := func(s *I8Slice, offset int, v, r *I8Slice) {
		s.Overwrite(offset, *v)
		if !r.Equal(s) {
			t.Fatalf("Overwrite(%v, %v) should be %v but is %v", offset, v, r, s)
		}
	}

	ConfirmOverwrite(I8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, I8List(10, 9, 8, 7), I8List(10, 9, 8, 7, 4, 5, 6, 7, 8, 9))
	ConfirmOverwrite(I8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, I8List(10, 9, 8, 7), I8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmOverwrite(I8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, I8List(11, 12, 13, 14), I8List(0, 1, 2, 3, 4, 11, 12, 13, 14, 9))
}

func TestI8SliceReallocate(t *testing.T) {
	ConfirmReallocate := func(s *I8Slice, l, c int, r *I8Slice) {
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

	ConfirmReallocate(I8List(), 0, 10, I8List())
	ConfirmReallocate(I8List(0, 1, 2, 3, 4), 3, 10, I8List(0, 1, 2))
	ConfirmReallocate(I8List(0, 1, 2, 3, 4), 5, 10, I8List(0, 1, 2, 3, 4))
	ConfirmReallocate(I8List(0, 1, 2, 3, 4), 10, 10, I8List(0, 1, 2, 3, 4, 0, 0, 0, 0, 0))
	ConfirmReallocate(I8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 1, 5, I8List(0))
	ConfirmReallocate(I8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 5, I8List(0, 1, 2, 3, 4))
	ConfirmReallocate(I8List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, 5, I8List(0, 1, 2, 3, 4))
}

func TestI8SliceExtend(t *testing.T) {
	ConfirmExtend := func(s *I8Slice, n int, r *I8Slice) {
		c := s.Cap()
		s.Extend(n)
		switch {
		case s.Len() != r.Len():	t.Fatalf("Extend(%v) len should be %v but is %v", n, r.Len(), s.Len())
		case s.Cap() != c + n:		t.Fatalf("Extend(%v) cap should be %v but is %v", n, c + n, s.Cap())
		case !r.Equal(s):			t.Fatalf("Extend(%v) should be %v but is %v", n, r, s)
		}
	}

	ConfirmExtend(I8List(), 1, I8List(0))
	ConfirmExtend(I8List(), 2, I8List(0, 0))
}

func TestI8SliceExpand(t *testing.T) {
	ConfirmExpand := func(s *I8Slice, i, n int, r *I8Slice) {
		c := s.Cap()
		s.Expand(i, n)
		switch {
		case s.Len() != r.Len():	t.Fatalf("Expand(%v, %v) len should be %v but is %v", i, n, r.Len(), s.Len())
		case s.Cap() != c + n:		t.Fatalf("Expand(%v, %v) cap should be %v but is %v", i, n, c + n, s.Cap())
		case !r.Equal(s):			t.Fatalf("Expand(%v, %v) should be %v but is %v", i, n, r, s)
		}
	}

	ConfirmExpand(I8List(), -1, 1, I8List(0))
	ConfirmExpand(I8List(), 0, 1, I8List(0))
	ConfirmExpand(I8List(), 1, 1, I8List(0))
	ConfirmExpand(I8List(), 0, 2, I8List(0, 0))

	ConfirmExpand(I8List(0, 1, 2), -1, 2, I8List(0, 0, 0, 1, 2))
	ConfirmExpand(I8List(0, 1, 2), 0, 2, I8List(0, 0, 0, 1, 2))
	ConfirmExpand(I8List(0, 1, 2), 1, 2, I8List(0, 0, 0, 1, 2))
	ConfirmExpand(I8List(0, 1, 2), 2, 2, I8List(0, 1, 0, 0, 2))
	ConfirmExpand(I8List(0, 1, 2), 3, 2, I8List(0, 1, 2, 0, 0))
	ConfirmExpand(I8List(0, 1, 2), 4, 2, I8List(0, 1, 2, 0, 0))
}

func TestI8SliceDepth(t *testing.T) {
	ConfirmDepth := func(s *I8Slice, i int) {
		if x := s.Depth(); x != i {
			t.Fatalf("%v.Depth() should be %v but is %v", s, i, x)
		}
	}
	ConfirmDepth(I8List(0, 1), 0)
}

func TestI8SliceReverse(t *testing.T) {
	sxp := I8List(1, 2, 3, 4, 5)
	rxp := I8List(5, 4, 3, 2, 1)
	sxp.Reverse()
	if !rxp.Equal(sxp) {
		t.Fatalf("Reversal failed: %v", sxp)
	}
}

func TestI8SliceAppend(t *testing.T) {
	ConfirmAppend := func(s *I8Slice, v interface{}, r *I8Slice) {
		s.Append(v)
		if !r.Equal(s) {
			t.Fatalf("Append(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmAppend(I8List(), int8(0), I8List(0))
}

func TestI8SliceAppendSlice(t *testing.T) {
	ConfirmAppendSlice := func(s, v, r *I8Slice) {
		s.AppendSlice(*v)
		if !r.Equal(s) {
			t.Fatalf("AppendSlice(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmAppendSlice(I8List(), I8List(0), I8List(0))
	ConfirmAppendSlice(I8List(), I8List(0, 1), I8List(0, 1))
	ConfirmAppendSlice(I8List(0, 1, 2), I8List(3, 4), I8List(0, 1, 2, 3, 4))
}

func TestI8SlicePrepend(t *testing.T) {
	ConfirmPrepend := func(s *I8Slice, v interface{}, r *I8Slice) {
		if s.Prepend(v); !r.Equal(s) {
			t.Fatalf("Prepend(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmPrepend(I8List(), int8(0), I8List(0))
	ConfirmPrepend(I8List(0), int8(1), I8List(1, 0))
}

func TestI8SlicePrependSlice(t *testing.T) {
	ConfirmPrependSlice := func(s, v, r *I8Slice) {
		if s.PrependSlice(*v); !r.Equal(s) {
			t.Fatalf("PrependSlice(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmPrependSlice(I8List(), I8List(0), I8List(0))
	ConfirmPrependSlice(I8List(), I8List(0, 1), I8List(0, 1))
	ConfirmPrependSlice(I8List(0, 1, 2), I8List(3, 4), I8List(3, 4, 0, 1, 2))
}

func TestI8SliceRepeat(t *testing.T) {
	ConfirmRepeat := func(s *I8Slice, count int, r *I8Slice) {
		if x := s.Repeat(count); !x.Equal(r) {
			t.Fatalf("%v.Repeat(%v) should be %v but is %v", s, count, r, x)
		}
	}

	ConfirmRepeat(I8List(), 5, I8List())
	ConfirmRepeat(I8List(0), 1, I8List(0))
	ConfirmRepeat(I8List(0), 2, I8List(0, 0))
	ConfirmRepeat(I8List(0), 3, I8List(0, 0, 0))
	ConfirmRepeat(I8List(0), 4, I8List(0, 0, 0, 0))
	ConfirmRepeat(I8List(0), 5, I8List(0, 0, 0, 0, 0))
}

func TestI8SliceCar(t *testing.T) {
	ConfirmCar := func(s *I8Slice, r int8) {
		n := s.Car().(int8)
		if ok := n == r; !ok {
			t.Fatalf("head should be '%v' but is '%v'", r, n)
		}
	}
	ConfirmCar(I8List(1, 2, 3), 1)
}

func TestI8SliceCdr(t *testing.T) {
	ConfirmCdr := func(s, r *I8Slice) {
		if n := s.Cdr(); !n.Equal(r) {
			t.Fatalf("tail should be '%v' but is '%v'", r, n)
		}
	}
	ConfirmCdr(I8List(1, 2, 3), I8List(2, 3))
}

func TestI8SliceRplaca(t *testing.T) {
	ConfirmRplaca := func(s *I8Slice, v interface{}, r *I8Slice) {
		if s.Rplaca(v); !s.Equal(r) {
			t.Fatalf("slice should be '%v' but is '%v'", r, s)
		}
	}
	ConfirmRplaca(I8List(1, 2, 3, 4, 5), int8(0), I8List(0, 2, 3, 4, 5))
}

func TestI8SliceRplacd(t *testing.T) {
	ConfirmRplacd := func(s *I8Slice, v interface{}, r *I8Slice) {
		if s.Rplacd(v); !s.Equal(r) {
			t.Fatalf("slice should be '%v' but is '%v'", r, s)
		}
	}
	ConfirmRplacd(I8List(1, 2, 3, 4, 5), nil, I8List(1))
	ConfirmRplacd(I8List(1, 2, 3, 4, 5), int8(10), I8List(1, 10))
	ConfirmRplacd(I8List(1, 2, 3, 4, 5), I8List(5, 4, 3, 2), I8List(1, 5, 4, 3, 2))
	ConfirmRplacd(I8List(1, 2, 3, 4, 5, 6), I8List(2, 4, 8, 16), I8List(1, 2, 4, 8, 16))
}

func TestI8SliceSetIntersection(t *testing.T) {
	ConfirmSetIntersection := func(s, o, r *I8Slice) {
		x := s.SetIntersection(*o)
		x.Sort()
		if !r.Equal(x) {
			t.Fatalf("%v.SetIntersection(%v) should be %v but is %v", s, o, r, x)
		}
	}

	ConfirmSetIntersection(I8List(1, 2, 3), I8List(), I8List())
	ConfirmSetIntersection(I8List(1, 2, 3), I8List(1), I8List(1))
	ConfirmSetIntersection(I8List(1, 2, 3), I8List(1, 1), I8List(1))
	ConfirmSetIntersection(I8List(1, 2, 3), I8List(1, 2, 1), I8List(1, 2))
}

func TestI8SliceSetUnion(t *testing.T) {
	ConfirmSetUnion := func(s, o, r *I8Slice) {
		x := s.SetUnion(*o)
		x.Sort()
		if !r.Equal(x) {
			t.Fatalf("%v.SetUnion(%v) should be %v but is %v", s, o, r, x)
		}
	}

	ConfirmSetUnion(I8List(1, 2, 3), I8List(), I8List(1, 2, 3))
	ConfirmSetUnion(I8List(1, 2, 3), I8List(1), I8List(1, 2, 3))
	ConfirmSetUnion(I8List(1, 2, 3), I8List(1, 1), I8List(1, 2, 3))
	ConfirmSetUnion(I8List(1, 2, 3), I8List(1, 2, 1), I8List(1, 2, 3))
}