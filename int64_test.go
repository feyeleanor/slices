package slices

import "testing"

func TestI64Slice(t *testing.T) {
	sxp := I64List(1)
	switch {
	case sxp.Len() != 1:			t.Fatalf("I64List(1) should allocate 1 cells, not %v cells", sxp.Len())
	case sxp.I64At(0) != 1:			t.Fatalf("I64List(1) element 0 should be 1 and not %v", sxp.I64At(0))
	}

	sxp = I64List(1, 2)
	switch {
	case sxp.Len() != 2:			t.Fatalf("I64List(1 2) should allocate 2 cells, not %v cells", sxp.Len())
	case sxp.I64At(0) != 1:			t.Fatalf("I64List(1 2) element 0 should be 1 and not %v", sxp.I64At(0))
	case sxp.I64At(1) != 2:			t.Fatalf("I64List(1 2) element 1 should be 2 and not %v", sxp.I64At(1))
	}

	sxp = I64List(1, 2, 3)
	switch {
	case sxp.Len() != 3:			t.Fatalf("I64List(1 2 3) should allocate 3 cells, not %v cells", sxp.Len())
	case sxp.I64At(0) != 1:			t.Fatalf("I64List(1 2 3) element 0 should be 1 and not %v", sxp.I64At(0))
	case sxp.I64At(1) != 2:			t.Fatalf("I64List(1 2 3) element 1 should be 2 and not %v", sxp.I64At(1))
	case sxp.I64At(2) != 3:			t.Fatalf("I64List(1 2 3) element 2 should be 3 and not %v", sxp.I64At(2))
	}
}

func TestI64SliceString(t *testing.T) {
	ConfirmString := func(s *I64Slice, r string) {
		if x := s.String(); x != r {
			t.Fatalf("%v erroneously serialised as '%v'", r, x)
		}
	}

	ConfirmString(I64List(), "()")
	ConfirmString(I64List(0), "(0)")
	ConfirmString(I64List(0, 1), "(0 1)")
}

func TestI64SliceLen(t *testing.T) {
	ConfirmLength := func(s *I64Slice, i int) {
		if x := s.Len(); x != i {
			t.Fatalf("%v.Len() should be %v but is %v", s, i, x)
		}
	}
	
	ConfirmLength(I64List(0), 1)
	ConfirmLength(I64List(0, 1), 2)
}

func TestI64SliceSwap(t *testing.T) {
	ConfirmSwap := func(s *I64Slice, i, j int, r *I64Slice) {
		if s.Swap(i, j); !r.Equal(s) {
			t.Fatalf("Swap(%v, %v) should be %v but is %v", i, j, r, s)
		}
	}
	ConfirmSwap(I64List(0, 1, 2), 0, 1, I64List(1, 0, 2))
	ConfirmSwap(I64List(0, 1, 2), 0, 2, I64List(2, 1, 0))
}

func TestI64SliceSort(t *testing.T) {
	ConfirmSort := func(s, r *I64Slice) {
		if s.Sort(); !r.Equal(s) {
			t.Fatalf("Sort() should be %v but is %v", r, s)
		}
	}

	ConfirmSort(I64List(3, 2, 1, 4, 5, 0), I64List(0, 1, 2, 3, 4, 5))
}

func TestI64SliceCompare(t *testing.T) {
	ConfirmCompare := func(s *I64Slice, i, j, r int) {
		if x := s.Compare(i, j); x != r {
			t.Fatalf("Compare(%v, %v) should be %v but is %v", i, j, r, x)
		}
	}

	ConfirmCompare(I64List(0, 1), 0, 0, IS_SAME_AS)
	ConfirmCompare(I64List(0, 1), 0, 1, IS_LESS_THAN)
	ConfirmCompare(I64List(0, 1), 1, 0, IS_GREATER_THAN)
}

func TestI64SliceZeroCompare(t *testing.T) {
	ConfirmCompare := func(s *I64Slice, i, r int) {
		if x := s.ZeroCompare(i); x != r {
			t.Fatalf("ZeroCompare(%v) should be %v but is %v", i, r, x)
		}
	}

	ConfirmCompare(I64List(0, -1, 1), 0, IS_SAME_AS)
	ConfirmCompare(I64List(0, -1, 1), 1, IS_GREATER_THAN)
	ConfirmCompare(I64List(0, -1, 1), 2, IS_LESS_THAN)
}

func TestI64SliceCut(t *testing.T) {
	ConfirmCut := func(s *I64Slice, start, end int, r *I64Slice) {
		if s.Cut(start, end); !r.Equal(s) {
			t.Fatalf("Cut(%v, %v) should be %v but is %v", start, end, r, s)
		}
	}

	ConfirmCut(I64List(0, 1, 2, 3, 4, 5), 0, 1, I64List(1, 2, 3, 4, 5))
	ConfirmCut(I64List(0, 1, 2, 3, 4, 5), 1, 2, I64List(0, 2, 3, 4, 5))
	ConfirmCut(I64List(0, 1, 2, 3, 4, 5), 2, 3, I64List(0, 1, 3, 4, 5))
	ConfirmCut(I64List(0, 1, 2, 3, 4, 5), 3, 4, I64List(0, 1, 2, 4, 5))
	ConfirmCut(I64List(0, 1, 2, 3, 4, 5), 4, 5, I64List(0, 1, 2, 3, 5))
	ConfirmCut(I64List(0, 1, 2, 3, 4, 5), 5, 6, I64List(0, 1, 2, 3, 4))

	ConfirmCut(I64List(0, 1, 2, 3, 4, 5), -1, 1, I64List(1, 2, 3, 4, 5))
	ConfirmCut(I64List(0, 1, 2, 3, 4, 5), 0, 2, I64List(2, 3, 4, 5))
	ConfirmCut(I64List(0, 1, 2, 3, 4, 5), 1, 3, I64List(0, 3, 4, 5))
	ConfirmCut(I64List(0, 1, 2, 3, 4, 5), 2, 4, I64List(0, 1, 4, 5))
	ConfirmCut(I64List(0, 1, 2, 3, 4, 5), 3, 5, I64List(0, 1, 2, 5))
	ConfirmCut(I64List(0, 1, 2, 3, 4, 5), 4, 6, I64List(0, 1, 2, 3))
	ConfirmCut(I64List(0, 1, 2, 3, 4, 5), 5, 7, I64List(0, 1, 2, 3, 4))
}

func TestI64SliceTrim(t *testing.T) {
	ConfirmTrim := func(s *I64Slice, start, end int, r *I64Slice) {
		if s.Trim(start, end); !r.Equal(s) {
			t.Fatalf("Trim(%v, %v) should be %v but is %v", start, end, r, s)
		}
	}

	ConfirmTrim(I64List(0, 1, 2, 3, 4, 5), 0, 1, I64List(0))
	ConfirmTrim(I64List(0, 1, 2, 3, 4, 5), 1, 2, I64List(1))
	ConfirmTrim(I64List(0, 1, 2, 3, 4, 5), 2, 3, I64List(2))
	ConfirmTrim(I64List(0, 1, 2, 3, 4, 5), 3, 4, I64List(3))
	ConfirmTrim(I64List(0, 1, 2, 3, 4, 5), 4, 5, I64List(4))
	ConfirmTrim(I64List(0, 1, 2, 3, 4, 5), 5, 6, I64List(5))

	ConfirmTrim(I64List(0, 1, 2, 3, 4, 5), -1, 1, I64List(0))
	ConfirmTrim(I64List(0, 1, 2, 3, 4, 5), 0, 2, I64List(0, 1))
	ConfirmTrim(I64List(0, 1, 2, 3, 4, 5), 1, 3, I64List(1, 2))
	ConfirmTrim(I64List(0, 1, 2, 3, 4, 5), 2, 4, I64List(2, 3))
	ConfirmTrim(I64List(0, 1, 2, 3, 4, 5), 3, 5, I64List(3, 4))
	ConfirmTrim(I64List(0, 1, 2, 3, 4, 5), 4, 6, I64List(4, 5))
	ConfirmTrim(I64List(0, 1, 2, 3, 4, 5), 5, 7, I64List(5))
}

func TestI64SliceDelete(t *testing.T) {
	ConfirmDelete := func(s *I64Slice, index int, r *I64Slice) {
		if s.Delete(index); !r.Equal(s) {
			t.Fatalf("Delete(%v) should be %v but is %v", index, r, s)
		}
	}

	ConfirmDelete(I64List(0, 1, 2, 3, 4, 5), -1, I64List(0, 1, 2, 3, 4, 5))
	ConfirmDelete(I64List(0, 1, 2, 3, 4, 5), 0, I64List(1, 2, 3, 4, 5))
	ConfirmDelete(I64List(0, 1, 2, 3, 4, 5), 1, I64List(0, 2, 3, 4, 5))
	ConfirmDelete(I64List(0, 1, 2, 3, 4, 5), 2, I64List(0, 1, 3, 4, 5))
	ConfirmDelete(I64List(0, 1, 2, 3, 4, 5), 3, I64List(0, 1, 2, 4, 5))
	ConfirmDelete(I64List(0, 1, 2, 3, 4, 5), 4, I64List(0, 1, 2, 3, 5))
	ConfirmDelete(I64List(0, 1, 2, 3, 4, 5), 5, I64List(0, 1, 2, 3, 4))
	ConfirmDelete(I64List(0, 1, 2, 3, 4, 5), 6, I64List(0, 1, 2, 3, 4, 5))
}

func TestI64SliceDeleteAll(t *testing.T) {
	ConfirmDeleteAll := func(s *I64Slice, v interface{}, r *I64Slice) {
		if s.DeleteAll(v); !r.Equal(s) {
			t.Fatalf("DeleteAll(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmDeleteAll(I64List(0, 1, 0, 3, 0, 5), int64(0), I64List(1, 3, 5))
	ConfirmDeleteAll(I64List(0, 1, 0, 3, 0, 5), int64(1), I64List(0, 0, 3, 0, 5))
	ConfirmDeleteAll(I64List(0, 1, 0, 3, 0, 5), int64(6), I64List(0, 1, 0, 3, 0, 5))
}

func TestI64SliceI64DeleteAll(t *testing.T) {
	ConfirmDeleteAll := func(s *I64Slice, v int64, r *I64Slice) {
		if s.I64DeleteAll(v); !r.Equal(s) {
			t.Fatalf("DeleteAll(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmDeleteAll(I64List(0, 1, 0, 3, 0, 5), int64(0), I64List(1, 3, 5))
	ConfirmDeleteAll(I64List(0, 1, 0, 3, 0, 5), int64(1), I64List(0, 0, 3, 0, 5))
	ConfirmDeleteAll(I64List(0, 1, 0, 3, 0, 5), int64(6), I64List(0, 1, 0, 3, 0, 5))
}

func TestI64SliceDeleteIf(t *testing.T) {
	ConfirmDeleteIf := func(s, r *I64Slice, f func(interface{}) bool) {
		if s.DeleteIf(f); !r.Equal(s) {
			t.Fatalf("DeleteIf(%v) should be %v but is %v", f, r, s)
		}
	}

	ConfirmDeleteIf(I64List(0, 1, 0, 3, 0, 5), I64List(1, 3, 5), func(x interface{}) bool { return x == int64(0) })
	ConfirmDeleteIf(I64List(0, 1, 0, 3, 0, 5), I64List(0, 0, 3, 0, 5), func(x interface{}) bool { return x == int64(1) })
	ConfirmDeleteIf(I64List(0, 1, 0, 3, 0, 5), I64List(0, 1, 0, 3, 0, 5), func(x interface{}) bool { return x == int64(6) })
}

func TestI64SliceI64DeleteIf(t *testing.T) {
	ConfirmDeleteIf := func(s, r *I64Slice, f func(int64) bool) {
		if s.I64DeleteIf(f); !r.Equal(s) {
			t.Fatalf("DeleteIf(%v) should be %v but is %v", f, r, s)
		}
	}

	ConfirmDeleteIf(I64List(0, 1, 0, 3, 0, 5), I64List(1, 3, 5), func(x int64) bool { return x == int64(0) })
	ConfirmDeleteIf(I64List(0, 1, 0, 3, 0, 5), I64List(0, 0, 3, 0, 5), func(x int64) bool { return x == int64(1) })
	ConfirmDeleteIf(I64List(0, 1, 0, 3, 0, 5), I64List(0, 1, 0, 3, 0, 5), func(x int64) bool { return x == int64(6) })
}

func TestI64SliceEach(t *testing.T) {
	var count	int64
	I64List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).Each(func(i interface{}) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})
}

func TestI64SliceEachWithIndex(t *testing.T) {
	I64List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).EachWithIndex(func(index int, i interface{}) {
		if i != int64(index) {
			t.Fatalf("element %v erroneously reported as %v", index, i)
		}
	})
}

func TestI64SliceEachWithKey(t *testing.T) {
	I64List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).EachWithKey(func(key, i interface{}) {
		if i != int64(key.(int)) {
			t.Fatalf("element %v erroneously reported as %v", key, i)
		}
	})
}

func TestI64SliceI64Each(t *testing.T) {
	var count	int64
	I64List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).I64Each(func(i int64) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})
}

func TestI64SliceI64EachWithIndex(t *testing.T) {
	I64List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).I64EachWithIndex(func(index int, i int64) {
		if i != int64(index) {
			t.Fatalf("element %v erroneously reported as %v", index, i)
		}
	})
}

func TestI64SliceI64EachWithKey(t *testing.T) {
	c := I64List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9)
	c.I64EachWithKey(func(key interface{}, i int64) {
		if i != int64(key.(int)) {
			t.Fatalf("element %v erroneously reported as %v", key, i)
		}
	})
}

func TestI64SliceBlockCopy(t *testing.T) {
	ConfirmBlockCopy := func(s *I64Slice, destination, source, count int, r *I64Slice) {
		s.BlockCopy(destination, source, count)
		if !r.Equal(s) {
			t.Fatalf("BlockCopy(%v, %v, %v) should be %v but is %v", destination, source, count, r, s)
		}
	}

	ConfirmBlockCopy(I64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, 0, 4, I64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockCopy(I64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 9, 9, 4, I64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockCopy(I64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 2, 4, I64List(0, 1, 2, 3, 4, 2, 3, 4, 5, 9))
	ConfirmBlockCopy(I64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 2, 5, 4, I64List(0, 1, 5, 6, 7, 8, 6, 7, 8, 9))
}

func TestI64SliceBlockClear(t *testing.T) {
	ConfirmBlockClear := func(s *I64Slice, start, count int, r *I64Slice) {
		s.BlockClear(start, count)
		if !r.Equal(s) {
			t.Fatalf("BlockClear(%v, %v) should be %v but is %v", start, count, r, s)
		}
	}

	ConfirmBlockClear(I64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, 4, I64List(0, 0, 0, 0, 4, 5, 6, 7, 8, 9))
	ConfirmBlockClear(I64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, 4, I64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockClear(I64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 4, I64List(0, 1, 2, 3, 4, 0, 0, 0, 0, 9))
}

func TestI64SliceOverwrite(t *testing.T) {
	ConfirmOverwrite := func(s *I64Slice, offset int, v, r *I64Slice) {
		s.Overwrite(offset, *v)
		if !r.Equal(s) {
			t.Fatalf("Overwrite(%v, %v) should be %v but is %v", offset, v, r, s)
		}
	}

	ConfirmOverwrite(I64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, I64List(10, 9, 8, 7), I64List(10, 9, 8, 7, 4, 5, 6, 7, 8, 9))
	ConfirmOverwrite(I64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, I64List(10, 9, 8, 7), I64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmOverwrite(I64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, I64List(11, 12, 13, 14), I64List(0, 1, 2, 3, 4, 11, 12, 13, 14, 9))
}

func TestI64SliceReallocate(t *testing.T) {
	ConfirmReallocate := func(s *I64Slice, l, c int, r *I64Slice) {
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

	ConfirmReallocate(I64List(), 0, 10, I64List())
	ConfirmReallocate(I64List(0, 1, 2, 3, 4), 3, 10, I64List(0, 1, 2))
	ConfirmReallocate(I64List(0, 1, 2, 3, 4), 5, 10, I64List(0, 1, 2, 3, 4))
	ConfirmReallocate(I64List(0, 1, 2, 3, 4), 10, 10, I64List(0, 1, 2, 3, 4, 0, 0, 0, 0, 0))
	ConfirmReallocate(I64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 1, 5, I64List(0))
	ConfirmReallocate(I64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 5, I64List(0, 1, 2, 3, 4))
	ConfirmReallocate(I64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, 5, I64List(0, 1, 2, 3, 4))
}

func TestI64SliceExtend(t *testing.T) {
	ConfirmExtend := func(s *I64Slice, n int, r *I64Slice) {
		c := s.Cap()
		s.Extend(n)
		switch {
		case s.Len() != r.Len():	t.Fatalf("Extend(%v) len should be %v but is %v", n, r.Len(), s.Len())
		case s.Cap() != c + n:		t.Fatalf("Extend(%v) cap should be %v but is %v", n, c + n, s.Cap())
		case !r.Equal(s):			t.Fatalf("Extend(%v) should be %v but is %v", n, r, s)
		}
	}

	ConfirmExtend(I64List(), 1, I64List(0))
	ConfirmExtend(I64List(), 2, I64List(0, 0))
}

func TestI64SliceExpand(t *testing.T) {
	ConfirmExpand := func(s *I64Slice, i, n int, r *I64Slice) {
		c := s.Cap()
		s.Expand(i, n)
		switch {
		case s.Len() != r.Len():	t.Fatalf("Expand(%v, %v) len should be %v but is %v", i, n, r.Len(), s.Len())
		case s.Cap() != c + n:		t.Fatalf("Expand(%v, %v) cap should be %v but is %v", i, n, c + n, s.Cap())
		case !r.Equal(s):			t.Fatalf("Expand(%v, %v) should be %v but is %v", i, n, r, s)
		}
	}

	ConfirmExpand(I64List(), -1, 1, I64List(0))
	ConfirmExpand(I64List(), 0, 1, I64List(0))
	ConfirmExpand(I64List(), 1, 1, I64List(0))
	ConfirmExpand(I64List(), 0, 2, I64List(0, 0))

	ConfirmExpand(I64List(0, 1, 2), -1, 2, I64List(0, 0, 0, 1, 2))
	ConfirmExpand(I64List(0, 1, 2), 0, 2, I64List(0, 0, 0, 1, 2))
	ConfirmExpand(I64List(0, 1, 2), 1, 2, I64List(0, 0, 0, 1, 2))
	ConfirmExpand(I64List(0, 1, 2), 2, 2, I64List(0, 1, 0, 0, 2))
	ConfirmExpand(I64List(0, 1, 2), 3, 2, I64List(0, 1, 2, 0, 0))
	ConfirmExpand(I64List(0, 1, 2), 4, 2, I64List(0, 1, 2, 0, 0))
}

func TestI64SliceDepth(t *testing.T) {
	ConfirmDepth := func(s *I64Slice, i int) {
		if x := s.Depth(); x != i {
			t.Fatalf("%v.Depth() should be %v but is %v", s, i, x)
		}
	}
	ConfirmDepth(I64List(0, 1), 0)
}

func TestI64SliceReverse(t *testing.T) {
	sxp := I64List(1, 2, 3, 4, 5)
	rxp := I64List(5, 4, 3, 2, 1)
	sxp.Reverse()
	if !rxp.Equal(sxp) {
		t.Fatalf("Reversal failed: %v", sxp)
	}
}

func TestI64SliceAppend(t *testing.T) {
	ConfirmAppend := func(s *I64Slice, v interface{}, r *I64Slice) {
		s.Append(v)
		if !r.Equal(s) {
			t.Fatalf("Append(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmAppend(I64List(), int64(0), I64List(0))
}

func TestI64SliceAppendSlice(t *testing.T) {
	ConfirmAppendSlice := func(s, v, r *I64Slice) {
		s.AppendSlice(*v)
		if !r.Equal(s) {
			t.Fatalf("AppendSlice(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmAppendSlice(I64List(), I64List(0), I64List(0))
	ConfirmAppendSlice(I64List(), I64List(0, 1), I64List(0, 1))
	ConfirmAppendSlice(I64List(0, 1, 2), I64List(3, 4), I64List(0, 1, 2, 3, 4))
}

func TestI64SlicePrepend(t *testing.T) {
	ConfirmPrepend := func(s *I64Slice, v interface{}, r *I64Slice) {
		if s.Prepend(v); !r.Equal(s) {
			t.Fatalf("Prepend(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmPrepend(I64List(), int64(0), I64List(0))
	ConfirmPrepend(I64List(0), int64(1), I64List(1, 0))
}

func TestI64SlicePrependSlice(t *testing.T) {
	ConfirmPrependSlice := func(s, v, r *I64Slice) {
		if s.PrependSlice(*v); !r.Equal(s) {
			t.Fatalf("PrependSlice(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmPrependSlice(I64List(), I64List(0), I64List(0))
	ConfirmPrependSlice(I64List(), I64List(0, 1), I64List(0, 1))
	ConfirmPrependSlice(I64List(0, 1, 2), I64List(3, 4), I64List(3, 4, 0, 1, 2))
}

func TestI64SliceRepeat(t *testing.T) {
	ConfirmRepeat := func(s *I64Slice, count int, r *I64Slice) {
		if x := s.Repeat(count); !x.Equal(r) {
			t.Fatalf("%v.Repeat(%v) should be %v but is %v", s, count, r, x)
		}
	}

	ConfirmRepeat(I64List(), 5, I64List())
	ConfirmRepeat(I64List(0), 1, I64List(0))
	ConfirmRepeat(I64List(0), 2, I64List(0, 0))
	ConfirmRepeat(I64List(0), 3, I64List(0, 0, 0))
	ConfirmRepeat(I64List(0), 4, I64List(0, 0, 0, 0))
	ConfirmRepeat(I64List(0), 5, I64List(0, 0, 0, 0, 0))
}

func TestI64SliceCar(t *testing.T) {
	ConfirmCar := func(s *I64Slice, r int64) {
		n := s.Car().(int64)
		if ok := n == r; !ok {
			t.Fatalf("head should be '%v' but is '%v'", r, n)
		}
	}
	ConfirmCar(I64List(1, 2, 3), 1)
}

func TestI64SliceCdr(t *testing.T) {
	ConfirmCdr := func(s, r *I64Slice) {
		if n := s.Cdr(); !n.Equal(r) {
			t.Fatalf("tail should be '%v' but is '%v'", r, n)
		}
	}
	ConfirmCdr(I64List(1, 2, 3), I64List(2, 3))
}

func TestI64SliceRplaca(t *testing.T) {
	ConfirmRplaca := func(s *I64Slice, v interface{}, r *I64Slice) {
		if s.Rplaca(v); !s.Equal(r) {
			t.Fatalf("slice should be '%v' but is '%v'", r, s)
		}
	}
	ConfirmRplaca(I64List(1, 2, 3, 4, 5), int64(0), I64List(0, 2, 3, 4, 5))
}

func TestI64SliceRplacd(t *testing.T) {
	ConfirmRplacd := func(s *I64Slice, v interface{}, r *I64Slice) {
		if s.Rplacd(v); !s.Equal(r) {
			t.Fatalf("slice should be '%v' but is '%v'", r, s)
		}
	}
	ConfirmRplacd(I64List(1, 2, 3, 4, 5), nil, I64List(1))
	ConfirmRplacd(I64List(1, 2, 3, 4, 5), int64(10), I64List(1, 10))
	ConfirmRplacd(I64List(1, 2, 3, 4, 5), I64List(5, 4, 3, 2), I64List(1, 5, 4, 3, 2))
	ConfirmRplacd(I64List(1, 2, 3, 4, 5, 6), I64List(2, 4, 8, 16), I64List(1, 2, 4, 8, 16))
}

func TestI64SliceSetIntersection(t *testing.T) {
	ConfirmSetIntersection := func(s, o, r *I64Slice) {
		x := s.SetIntersection(*o)
		x.Sort()
		if !r.Equal(x) {
			t.Fatalf("%v.SetIntersection(%v) should be %v but is %v", s, o, r, x)
		}
	}

	ConfirmSetIntersection(I64List(1, 2, 3), I64List(), I64List())
	ConfirmSetIntersection(I64List(1, 2, 3), I64List(1), I64List(1))
	ConfirmSetIntersection(I64List(1, 2, 3), I64List(1, 1), I64List(1))
	ConfirmSetIntersection(I64List(1, 2, 3), I64List(1, 2, 1), I64List(1, 2))
}

func TestI64SliceSetUnion(t *testing.T) {
	ConfirmSetUnion := func(s, o, r *I64Slice) {
		x := s.SetUnion(*o)
		x.Sort()
		if !r.Equal(x) {
			t.Fatalf("%v.SetUnion(%v) should be %v but is %v", s, o, r, x)
		}
	}

	ConfirmSetUnion(I64List(1, 2, 3), I64List(), I64List(1, 2, 3))
	ConfirmSetUnion(I64List(1, 2, 3), I64List(1), I64List(1, 2, 3))
	ConfirmSetUnion(I64List(1, 2, 3), I64List(1, 1), I64List(1, 2, 3))
	ConfirmSetUnion(I64List(1, 2, 3), I64List(1, 2, 1), I64List(1, 2, 3))
}

func TestI64SliceSetDifference(t *testing.T) {
	ConfirmSetUnion := func(s, o, r *I64Slice) {
		x := s.SetDifference(*o)
		x.Sort()
		if !r.Equal(x) {
			t.Fatalf("%v.SetUnion(%v) should be %v but is %v", s, o, r, x)
		}
	}

	ConfirmSetUnion(I64List(1, 2, 3), I64List(), I64List(1, 2, 3))
	ConfirmSetUnion(I64List(1, 2, 3), I64List(1), I64List(2, 3))
	ConfirmSetUnion(I64List(1, 2, 3), I64List(1, 1), I64List(2, 3))
	ConfirmSetUnion(I64List(1, 2, 3), I64List(1, 2, 1), I64List(3))
}