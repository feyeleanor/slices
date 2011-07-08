package slices

import "testing"

func TestU16Slice(t *testing.T) {
	sxp := U16List(1)
	switch {
	case sxp.Len() != 1:			t.Fatalf("U16List(1) should allocate 1 cells, not %v cells", sxp.Len())
	case sxp.U16At(0) != 1:			t.Fatalf("U16List(1) element 0 should be 1 and not %v", sxp.U16At(0))
	}

	sxp = U16List(1, 2)
	switch {
	case sxp.Len() != 2:			t.Fatalf("U16List(1 2) should allocate 2 cells, not %v cells", sxp.Len())
	case sxp.U16At(0) != 1:			t.Fatalf("U16List(1 2) element 0 should be 1 and not %v", sxp.U16At(0))
	case sxp.U16At(1) != 2:			t.Fatalf("U16List(1 2) element 1 should be 2 and not %v", sxp.U16At(1))
	}

	sxp = U16List(1, 2, 3)
	switch {
	case sxp.Len() != 3:			t.Fatalf("U16List(1 2 3) should allocate 3 cells, not %v cells", sxp.Len())
	case sxp.U16At(0) != 1:			t.Fatalf("U16List(1 2 3) element 0 should be 1 and not %v", sxp.U16At(0))
	case sxp.U16At(1) != 2:			t.Fatalf("U16List(1 2 3) element 1 should be 2 and not %v", sxp.U16At(1))
	case sxp.U16At(2) != 3:			t.Fatalf("U16List(1 2 3) element 2 should be 3 and not %v", sxp.U16At(2))
	}
}

func TestU16SliceString(t *testing.T) {
	ConfirmString := func(s *U16Slice, r string) {
		if x := s.String(); x != r {
			t.Fatalf("%v erroneously serialised as '%v'", r, x)
		}
	}

	ConfirmString(U16List(), "()")
	ConfirmString(U16List(0), "(0)")
	ConfirmString(U16List(0, 1), "(0 1)")
}

func TestU16SliceLen(t *testing.T) {
	ConfirmLength := func(s *U16Slice, i int) {
		if x := s.Len(); x != i {
			t.Fatalf("%v.Len() should be %v but is %v", s, i, x)
		}
	}
	
	ConfirmLength(U16List(0), 1)
	ConfirmLength(U16List(0, 1), 2)
}

func TestU16SliceSwap(t *testing.T) {
	ConfirmSwap := func(s *U16Slice, i, j int, r *U16Slice) {
		if s.Swap(i, j); !r.Equal(s) {
			t.Fatalf("Swap(%v, %v) should be %v but is %v", i, j, r, s)
		}
	}
	ConfirmSwap(U16List(0, 1, 2), 0, 1, U16List(1, 0, 2))
	ConfirmSwap(U16List(0, 1, 2), 0, 2, U16List(2, 1, 0))
}

func TestU16SliceSort(t *testing.T) {
	ConfirmSort := func(s, r *U16Slice) {
		if s.Sort(); !r.Equal(s) {
			t.Fatalf("Sort() should be %v but is %v", r, s)
		}
	}

	ConfirmSort(U16List(3, 2, 1, 4, 5, 0), U16List(0, 1, 2, 3, 4, 5))
}

func TestU16SliceCompare(t *testing.T) {
	ConfirmCompare := func(s *U16Slice, i, j, r int) {
		if x := s.Compare(i, j); x != r {
			t.Fatalf("Compare(%v, %v) should be %v but is %v", i, j, r, x)
		}
	}

	ConfirmCompare(U16List(0, 1), 0, 0, IS_SAME_AS)
	ConfirmCompare(U16List(0, 1), 0, 1, IS_LESS_THAN)
	ConfirmCompare(U16List(0, 1), 1, 0, IS_GREATER_THAN)
}

func TestU16SliceZeroCompare(t *testing.T) {
	ConfirmCompare := func(s *U16Slice, i, r int) {
		if x := s.ZeroCompare(i); x != r {
			t.Fatalf("ZeroCompare(%v) should be %v but is %v", i, r, x)
		}
	}

	ConfirmCompare(U16List(0, 1, 2), 0, IS_SAME_AS)
	ConfirmCompare(U16List(0, 1, 2), 1, IS_LESS_THAN)
	ConfirmCompare(U16List(0, 1, 2), 2, IS_LESS_THAN)
}

func TestU16SliceCut(t *testing.T) {
	ConfirmCut := func(s *U16Slice, start, end int, r *U16Slice) {
		if s.Cut(start, end); !r.Equal(s) {
			t.Fatalf("Cut(%v, %v) should be %v but is %v", start, end, r, s)
		}
	}

	ConfirmCut(U16List(0, 1, 2, 3, 4, 5), 0, 1, U16List(1, 2, 3, 4, 5))
	ConfirmCut(U16List(0, 1, 2, 3, 4, 5), 1, 2, U16List(0, 2, 3, 4, 5))
	ConfirmCut(U16List(0, 1, 2, 3, 4, 5), 2, 3, U16List(0, 1, 3, 4, 5))
	ConfirmCut(U16List(0, 1, 2, 3, 4, 5), 3, 4, U16List(0, 1, 2, 4, 5))
	ConfirmCut(U16List(0, 1, 2, 3, 4, 5), 4, 5, U16List(0, 1, 2, 3, 5))
	ConfirmCut(U16List(0, 1, 2, 3, 4, 5), 5, 6, U16List(0, 1, 2, 3, 4))

	ConfirmCut(U16List(0, 1, 2, 3, 4, 5), -1, 1, U16List(1, 2, 3, 4, 5))
	ConfirmCut(U16List(0, 1, 2, 3, 4, 5), 0, 2, U16List(2, 3, 4, 5))
	ConfirmCut(U16List(0, 1, 2, 3, 4, 5), 1, 3, U16List(0, 3, 4, 5))
	ConfirmCut(U16List(0, 1, 2, 3, 4, 5), 2, 4, U16List(0, 1, 4, 5))
	ConfirmCut(U16List(0, 1, 2, 3, 4, 5), 3, 5, U16List(0, 1, 2, 5))
	ConfirmCut(U16List(0, 1, 2, 3, 4, 5), 4, 6, U16List(0, 1, 2, 3))
	ConfirmCut(U16List(0, 1, 2, 3, 4, 5), 5, 7, U16List(0, 1, 2, 3, 4))
}

func TestU16SliceTrim(t *testing.T) {
	ConfirmTrim := func(s *U16Slice, start, end int, r *U16Slice) {
		if s.Trim(start, end); !r.Equal(s) {
			t.Fatalf("Trim(%v, %v) should be %v but is %v", start, end, r, s)
		}
	}

	ConfirmTrim(U16List(0, 1, 2, 3, 4, 5), 0, 1, U16List(0))
	ConfirmTrim(U16List(0, 1, 2, 3, 4, 5), 1, 2, U16List(1))
	ConfirmTrim(U16List(0, 1, 2, 3, 4, 5), 2, 3, U16List(2))
	ConfirmTrim(U16List(0, 1, 2, 3, 4, 5), 3, 4, U16List(3))
	ConfirmTrim(U16List(0, 1, 2, 3, 4, 5), 4, 5, U16List(4))
	ConfirmTrim(U16List(0, 1, 2, 3, 4, 5), 5, 6, U16List(5))

	ConfirmTrim(U16List(0, 1, 2, 3, 4, 5), -1, 1, U16List(0))
	ConfirmTrim(U16List(0, 1, 2, 3, 4, 5), 0, 2, U16List(0, 1))
	ConfirmTrim(U16List(0, 1, 2, 3, 4, 5), 1, 3, U16List(1, 2))
	ConfirmTrim(U16List(0, 1, 2, 3, 4, 5), 2, 4, U16List(2, 3))
	ConfirmTrim(U16List(0, 1, 2, 3, 4, 5), 3, 5, U16List(3, 4))
	ConfirmTrim(U16List(0, 1, 2, 3, 4, 5), 4, 6, U16List(4, 5))
	ConfirmTrim(U16List(0, 1, 2, 3, 4, 5), 5, 7, U16List(5))
}

func TestU16SliceDelete(t *testing.T) {
	ConfirmDelete := func(s *U16Slice, index int, r *U16Slice) {
		if s.Delete(index); !r.Equal(s) {
			t.Fatalf("Delete(%v) should be %v but is %v", index, r, s)
		}
	}

	ConfirmDelete(U16List(0, 1, 2, 3, 4, 5), -1, U16List(0, 1, 2, 3, 4, 5))
	ConfirmDelete(U16List(0, 1, 2, 3, 4, 5), 0, U16List(1, 2, 3, 4, 5))
	ConfirmDelete(U16List(0, 1, 2, 3, 4, 5), 1, U16List(0, 2, 3, 4, 5))
	ConfirmDelete(U16List(0, 1, 2, 3, 4, 5), 2, U16List(0, 1, 3, 4, 5))
	ConfirmDelete(U16List(0, 1, 2, 3, 4, 5), 3, U16List(0, 1, 2, 4, 5))
	ConfirmDelete(U16List(0, 1, 2, 3, 4, 5), 4, U16List(0, 1, 2, 3, 5))
	ConfirmDelete(U16List(0, 1, 2, 3, 4, 5), 5, U16List(0, 1, 2, 3, 4))
	ConfirmDelete(U16List(0, 1, 2, 3, 4, 5), 6, U16List(0, 1, 2, 3, 4, 5))
}

func TestU16SliceDeleteAll(t *testing.T) {
	ConfirmDeleteAll := func(s *U16Slice, v interface{}, r *U16Slice) {
		if s.DeleteAll(v); !r.Equal(s) {
			t.Fatalf("DeleteAll(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmDeleteAll(U16List(0, 1, 0, 3, 0, 5), uint16(0), U16List(1, 3, 5))
	ConfirmDeleteAll(U16List(0, 1, 0, 3, 0, 5), uint16(1), U16List(0, 0, 3, 0, 5))
	ConfirmDeleteAll(U16List(0, 1, 0, 3, 0, 5), uint16(6), U16List(0, 1, 0, 3, 0, 5))
}

func TestU16SliceU16DeleteAll(t *testing.T) {
	ConfirmDeleteAll := func(s *U16Slice, v uint16, r *U16Slice) {
		if s.U16DeleteAll(v); !r.Equal(s) {
			t.Fatalf("DeleteAll(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmDeleteAll(U16List(0, 1, 0, 3, 0, 5), uint16(0), U16List(1, 3, 5))
	ConfirmDeleteAll(U16List(0, 1, 0, 3, 0, 5), uint16(1), U16List(0, 0, 3, 0, 5))
	ConfirmDeleteAll(U16List(0, 1, 0, 3, 0, 5), uint16(6), U16List(0, 1, 0, 3, 0, 5))
}

func TestU16SliceDeleteIf(t *testing.T) {
	ConfirmDeleteIf := func(s, r *U16Slice, f func(interface{}) bool) {
		if s.DeleteIf(f); !r.Equal(s) {
			t.Fatalf("DeleteIf(%v) should be %v but is %v", f, r, s)
		}
	}

	ConfirmDeleteIf(U16List(0, 1, 0, 3, 0, 5), U16List(1, 3, 5), func(x interface{}) bool { return x == uint16(0) })
	ConfirmDeleteIf(U16List(0, 1, 0, 3, 0, 5), U16List(0, 0, 3, 0, 5), func(x interface{}) bool { return x == uint16(1) })
	ConfirmDeleteIf(U16List(0, 1, 0, 3, 0, 5), U16List(0, 1, 0, 3, 0, 5), func(x interface{}) bool { return x == uint16(6) })
}

func TestU16SliceU16DeleteIf(t *testing.T) {
	ConfirmDeleteIf := func(s, r *U16Slice, f func(uint16) bool) {
		if s.U16DeleteIf(f); !r.Equal(s) {
			t.Fatalf("DeleteIf(%v) should be %v but is %v", f, r, s)
		}
	}

	ConfirmDeleteIf(U16List(0, 1, 0, 3, 0, 5), U16List(1, 3, 5), func(x uint16) bool { return x == uint16(0) })
	ConfirmDeleteIf(U16List(0, 1, 0, 3, 0, 5), U16List(0, 0, 3, 0, 5), func(x uint16) bool { return x == uint16(1) })
	ConfirmDeleteIf(U16List(0, 1, 0, 3, 0, 5), U16List(0, 1, 0, 3, 0, 5), func(x uint16) bool { return x == uint16(6) })
}

func TestU16SliceEach(t *testing.T) {
	var count	uint16
	U16List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).Each(func(i interface{}) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})
}

func TestU16SliceEachWithIndex(t *testing.T) {
	U16List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).EachWithIndex(func(index int, i interface{}) {
		if i != uint16(index) {
			t.Fatalf("element %v erroneously reported as %v", index, i)
		}
	})
}

func TestU16SliceEachWithKey(t *testing.T) {
	U16List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).EachWithKey(func(key, i interface{}) {
		if i != uint16(key.(int)) {
			t.Fatalf("element %v erroneously reported as %v", key, i)
		}
	})
}

func TestU16SliceU16Each(t *testing.T) {
	var count	uint16
	U16List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).U16Each(func(i uint16) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})
}

func TestU16SliceU16EachWithIndex(t *testing.T) {
	U16List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).U16EachWithIndex(func(index int, i uint16) {
		if i != uint16(index) {
			t.Fatalf("element %v erroneously reported as %v", index, i)
		}
	})
}

func TestU16SliceU16EachWithKey(t *testing.T) {
	c := U16List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9)
	c.U16EachWithKey(func(key interface{}, i uint16) {
		if i != uint16(key.(int)) {
			t.Fatalf("element %v erroneously reported as %v", key, i)
		}
	})
}

func TestU16SliceBlockCopy(t *testing.T) {
	ConfirmBlockCopy := func(s *U16Slice, destination, source, count int, r *U16Slice) {
		s.BlockCopy(destination, source, count)
		if !r.Equal(s) {
			t.Fatalf("BlockCopy(%v, %v, %v) should be %v but is %v", destination, source, count, r, s)
		}
	}

	ConfirmBlockCopy(U16List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, 0, 4, U16List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockCopy(U16List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 9, 9, 4, U16List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockCopy(U16List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 2, 4, U16List(0, 1, 2, 3, 4, 2, 3, 4, 5, 9))
	ConfirmBlockCopy(U16List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 2, 5, 4, U16List(0, 1, 5, 6, 7, 8, 6, 7, 8, 9))
}

func TestU16SliceBlockClear(t *testing.T) {
	ConfirmBlockClear := func(s *U16Slice, start, count int, r *U16Slice) {
		s.BlockClear(start, count)
		if !r.Equal(s) {
			t.Fatalf("BlockClear(%v, %v) should be %v but is %v", start, count, r, s)
		}
	}

	ConfirmBlockClear(U16List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, 4, U16List(0, 0, 0, 0, 4, 5, 6, 7, 8, 9))
	ConfirmBlockClear(U16List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, 4, U16List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockClear(U16List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 4, U16List(0, 1, 2, 3, 4, 0, 0, 0, 0, 9))
}

func TestU16SliceOverwrite(t *testing.T) {
	ConfirmOverwrite := func(s *U16Slice, offset int, v, r *U16Slice) {
		s.Overwrite(offset, *v)
		if !r.Equal(s) {
			t.Fatalf("Overwrite(%v, %v) should be %v but is %v", offset, v, r, s)
		}
	}

	ConfirmOverwrite(U16List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, U16List(10, 9, 8, 7), U16List(10, 9, 8, 7, 4, 5, 6, 7, 8, 9))
	ConfirmOverwrite(U16List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, U16List(10, 9, 8, 7), U16List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmOverwrite(U16List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, U16List(11, 12, 13, 14), U16List(0, 1, 2, 3, 4, 11, 12, 13, 14, 9))
}

func TestU16SliceReallocate(t *testing.T) {
	ConfirmReallocate := func(s *U16Slice, l, c int, r *U16Slice) {
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

	ConfirmReallocate(U16List(), 0, 10, U16List())
	ConfirmReallocate(U16List(0, 1, 2, 3, 4), 3, 10, U16List(0, 1, 2))
	ConfirmReallocate(U16List(0, 1, 2, 3, 4), 5, 10, U16List(0, 1, 2, 3, 4))
	ConfirmReallocate(U16List(0, 1, 2, 3, 4), 10, 10, U16List(0, 1, 2, 3, 4, 0, 0, 0, 0, 0))
	ConfirmReallocate(U16List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 1, 5, U16List(0))
	ConfirmReallocate(U16List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 5, U16List(0, 1, 2, 3, 4))
	ConfirmReallocate(U16List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, 5, U16List(0, 1, 2, 3, 4))
}

func TestU16SliceExtend(t *testing.T) {
	ConfirmExtend := func(s *U16Slice, n int, r *U16Slice) {
		c := s.Cap()
		s.Extend(n)
		switch {
		case s.Len() != r.Len():	t.Fatalf("Extend(%v) len should be %v but is %v", n, r.Len(), s.Len())
		case s.Cap() != c + n:		t.Fatalf("Extend(%v) cap should be %v but is %v", n, c + n, s.Cap())
		case !r.Equal(s):			t.Fatalf("Extend(%v) should be %v but is %v", n, r, s)
		}
	}

	ConfirmExtend(U16List(), 1, U16List(0))
	ConfirmExtend(U16List(), 2, U16List(0, 0))
}

func TestU16SliceExpand(t *testing.T) {
	ConfirmExpand := func(s *U16Slice, i, n int, r *U16Slice) {
		c := s.Cap()
		s.Expand(i, n)
		switch {
		case s.Len() != r.Len():	t.Fatalf("Expand(%v, %v) len should be %v but is %v", i, n, r.Len(), s.Len())
		case s.Cap() != c + n:		t.Fatalf("Expand(%v, %v) cap should be %v but is %v", i, n, c + n, s.Cap())
		case !r.Equal(s):			t.Fatalf("Expand(%v, %v) should be %v but is %v", i, n, r, s)
		}
	}

	ConfirmExpand(U16List(), -1, 1, U16List(0))
	ConfirmExpand(U16List(), 0, 1, U16List(0))
	ConfirmExpand(U16List(), 1, 1, U16List(0))
	ConfirmExpand(U16List(), 0, 2, U16List(0, 0))

	ConfirmExpand(U16List(0, 1, 2), -1, 2, U16List(0, 0, 0, 1, 2))
	ConfirmExpand(U16List(0, 1, 2), 0, 2, U16List(0, 0, 0, 1, 2))
	ConfirmExpand(U16List(0, 1, 2), 1, 2, U16List(0, 0, 0, 1, 2))
	ConfirmExpand(U16List(0, 1, 2), 2, 2, U16List(0, 1, 0, 0, 2))
	ConfirmExpand(U16List(0, 1, 2), 3, 2, U16List(0, 1, 2, 0, 0))
	ConfirmExpand(U16List(0, 1, 2), 4, 2, U16List(0, 1, 2, 0, 0))
}

func TestU16SliceDepth(t *testing.T) {
	ConfirmDepth := func(s *U16Slice, i int) {
		if x := s.Depth(); x != i {
			t.Fatalf("%v.Depth() should be %v but is %v", s, i, x)
		}
	}
	ConfirmDepth(U16List(0, 1), 0)
}

func TestU16SliceReverse(t *testing.T) {
	sxp := U16List(1, 2, 3, 4, 5)
	rxp := U16List(5, 4, 3, 2, 1)
	sxp.Reverse()
	if !rxp.Equal(sxp) {
		t.Fatalf("Reversal failed: %v", sxp)
	}
}

func TestU16SliceAppend(t *testing.T) {
	ConfirmAppend := func(s *U16Slice, v interface{}, r *U16Slice) {
		s.Append(v)
		if !r.Equal(s) {
			t.Fatalf("Append(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmAppend(U16List(), uint16(0), U16List(0))
}

func TestU16SliceAppendSlice(t *testing.T) {
	ConfirmAppendSlice := func(s, v, r *U16Slice) {
		s.AppendSlice(*v)
		if !r.Equal(s) {
			t.Fatalf("AppendSlice(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmAppendSlice(U16List(), U16List(0), U16List(0))
	ConfirmAppendSlice(U16List(), U16List(0, 1), U16List(0, 1))
	ConfirmAppendSlice(U16List(0, 1, 2), U16List(3, 4), U16List(0, 1, 2, 3, 4))
}

func TestU16SlicePrepend(t *testing.T) {
	ConfirmPrepend := func(s *U16Slice, v interface{}, r *U16Slice) {
		if s.Prepend(v); !r.Equal(s) {
			t.Fatalf("Prepend(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmPrepend(U16List(), uint16(0), U16List(0))
	ConfirmPrepend(U16List(0), uint16(1), U16List(1, 0))
}

func TestU16SlicePrependSlice(t *testing.T) {
	ConfirmPrependSlice := func(s, v, r *U16Slice) {
		if s.PrependSlice(*v); !r.Equal(s) {
			t.Fatalf("PrependSlice(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmPrependSlice(U16List(), U16List(0), U16List(0))
	ConfirmPrependSlice(U16List(), U16List(0, 1), U16List(0, 1))
	ConfirmPrependSlice(U16List(0, 1, 2), U16List(3, 4), U16List(3, 4, 0, 1, 2))
}

func TestU16SliceRepeat(t *testing.T) {
	ConfirmRepeat := func(s *U16Slice, count int, r *U16Slice) {
		if x := s.Repeat(count); !x.Equal(r) {
			t.Fatalf("%v.Repeat(%v) should be %v but is %v", s, count, r, x)
		}
	}

	ConfirmRepeat(U16List(), 5, U16List())
	ConfirmRepeat(U16List(0), 1, U16List(0))
	ConfirmRepeat(U16List(0), 2, U16List(0, 0))
	ConfirmRepeat(U16List(0), 3, U16List(0, 0, 0))
	ConfirmRepeat(U16List(0), 4, U16List(0, 0, 0, 0))
	ConfirmRepeat(U16List(0), 5, U16List(0, 0, 0, 0, 0))
}

func TestU16SliceCar(t *testing.T) {
	ConfirmCar := func(s *U16Slice, r uint16) {
		n := s.Car().(uint16)
		if ok := n == r; !ok {
			t.Fatalf("head should be '%v' but is '%v'", r, n)
		}
	}
	ConfirmCar(U16List(1, 2, 3), 1)
}

func TestU16SliceCdr(t *testing.T) {
	ConfirmCdr := func(s, r *U16Slice) {
		if n := s.Cdr(); !n.Equal(r) {
			t.Fatalf("tail should be '%v' but is '%v'", r, n)
		}
	}
	ConfirmCdr(U16List(1, 2, 3), U16List(2, 3))
}

func TestU16SliceRplaca(t *testing.T) {
	ConfirmRplaca := func(s *U16Slice, v interface{}, r *U16Slice) {
		if s.Rplaca(v); !s.Equal(r) {
			t.Fatalf("slice should be '%v' but is '%v'", r, s)
		}
	}
	ConfirmRplaca(U16List(1, 2, 3, 4, 5), uint16(0), U16List(0, 2, 3, 4, 5))
}

func TestU16SliceRplacd(t *testing.T) {
	ConfirmRplacd := func(s *U16Slice, v interface{}, r *U16Slice) {
		if s.Rplacd(v); !s.Equal(r) {
			t.Fatalf("slice should be '%v' but is '%v'", r, s)
		}
	}
	ConfirmRplacd(U16List(1, 2, 3, 4, 5), nil, U16List(1))
	ConfirmRplacd(U16List(1, 2, 3, 4, 5), uint16(10), U16List(1, 10))
	ConfirmRplacd(U16List(1, 2, 3, 4, 5), U16List(5, 4, 3, 2), U16List(1, 5, 4, 3, 2))
	ConfirmRplacd(U16List(1, 2, 3, 4, 5, 6), U16List(2, 4, 8, 16), U16List(1, 2, 4, 8, 16))
}

func TestU16SliceSetIntersection(t *testing.T) {
	ConfirmSetIntersection := func(s, o, r *U16Slice) {
		x := s.SetIntersection(*o)
		x.Sort()
		if !r.Equal(x) {
			t.Fatalf("%v.SetIntersection(%v) should be %v but is %v", s, o, r, x)
		}
	}

	ConfirmSetIntersection(U16List(1, 2, 3), U16List(), U16List())
	ConfirmSetIntersection(U16List(1, 2, 3), U16List(1), U16List(1))
	ConfirmSetIntersection(U16List(1, 2, 3), U16List(1, 1), U16List(1))
	ConfirmSetIntersection(U16List(1, 2, 3), U16List(1, 2, 1), U16List(1, 2))
}

func TestU16SliceSetUnion(t *testing.T) {
	ConfirmSetUnion := func(s, o, r *U16Slice) {
		x := s.SetUnion(*o)
		x.Sort()
		if !r.Equal(x) {
			t.Fatalf("%v.SetUnion(%v) should be %v but is %v", s, o, r, x)
		}
	}

	ConfirmSetUnion(U16List(1, 2, 3), U16List(), U16List(1, 2, 3))
	ConfirmSetUnion(U16List(1, 2, 3), U16List(1), U16List(1, 2, 3))
	ConfirmSetUnion(U16List(1, 2, 3), U16List(1, 1), U16List(1, 2, 3))
	ConfirmSetUnion(U16List(1, 2, 3), U16List(1, 2, 1), U16List(1, 2, 3))
}

func TestU16SliceSetDifference(t *testing.T) {
	ConfirmSetUnion := func(s, o, r *U16Slice) {
		x := s.SetDifference(*o)
		x.Sort()
		if !r.Equal(x) {
			t.Fatalf("%v.SetUnion(%v) should be %v but is %v", s, o, r, x)
		}
	}

	ConfirmSetUnion(U16List(1, 2, 3), U16List(), U16List(1, 2, 3))
	ConfirmSetUnion(U16List(1, 2, 3), U16List(1), U16List(2, 3))
	ConfirmSetUnion(U16List(1, 2, 3), U16List(1, 1), U16List(2, 3))
	ConfirmSetUnion(U16List(1, 2, 3), U16List(1, 2, 1), U16List(3))
}