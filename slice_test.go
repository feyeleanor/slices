package slices

import "github.com/feyeleanor/lists"
import "testing"

func TestList(t *testing.T) {
	sxp := List(nil, nil)
	switch {
	case sxp.Len() != 2:			t.Fatalf("List(nil nil) should allocate 2 cells, not %v cells", sxp.Len())
	case sxp.At(0) != nil:			t.Fatalf("List(nil nil) element 0 should be nil and not %v", sxp.At(0))
	case sxp.At(1) != nil:			t.Fatalf("List(nil nil) element 1 should be nil and not %v", sxp.At(1))
	}

	sxp = List(1, nil)
	switch {
	case sxp.Len() != 2:			t.Fatalf("List(1 nil) should allocate 2 cells, not %v cells", sxp.Len())
	case sxp.At(0) != 1:			t.Fatalf("List(1 nil) element 0 should be 1 and not %v", sxp.At(0))
	case sxp.At(1) != nil:			t.Fatalf("List(1 nil) element 1 should be nil and not %v", sxp.At(1))
	}

	sxp = List(1, 2)
	switch {
	case sxp.Len() != 2:			t.Fatalf("List(1 2) should allocate 2 cells, not %v cells", sxp.Len())
	case sxp.At(0) != 1:			t.Fatalf("List(1 2) element 0 should be 1 and not %v", sxp.At(0))
	case sxp.At(1) != 2:			t.Fatalf("List(1 2) element 1 should be 2 and not %v", sxp.At(1))
	}

	sxp = List(1, 2, 3)
	switch {
	case sxp.Len() != 3:			t.Fatalf("List(1 2 3) should allocate 3 cells, not %v cells", sxp.Len())
	case sxp.At(0) != 1:			t.Fatalf("List(1 2 3) element 0 should be 1 and not %v", sxp.At(0))
	case sxp.At(1) != 2:			t.Fatalf("List(1 2 3) element 1 should be 2 and not %v", sxp.At(1))
	case sxp.At(2) != 3:			t.Fatalf("List(1 2 3) element 2 should be 3 and not %v", sxp.At(2))
	}

	sxp = List(1, List(10, 20), 3)
	rxp := List(10, 20)
	switch {
	case sxp.Len() != 3:			t.Fatalf("List(1 (10 20) 3) should allocate 3 cells, not %v cells", sxp.Len())
	case sxp.At(0) != 1:			t.Fatalf("List(1 (10 20) 3) element 0 should be 1 and not %v", sxp.At(0))
	case !rxp.Equal(sxp.At(1)):		t.Fatalf("List(1 (10 20) 3) element 1 should be (10 20) and not %v", sxp.At(1))
	case sxp.At(2) != 3:			t.Fatalf("List(1 (10 20) 3) element 2 should be 3 and not %v", sxp.At(2))
	}


	sxp = List(1, List(10, List(-10, -30)), 3)
	rxp = List(10, List(-10, -30))
	switch {
	case sxp.Len() != 3:			t.Fatalf("List(1 (10 20) 3) should allocate 3 cells, not %v cells", sxp.Len())
	case sxp.At(0) != 1:			t.Fatalf("List(1 (10 20) 3) element 0 should be 1 and not %v", sxp.At(0))
	case !rxp.Equal(sxp.At(1)):		t.Fatalf("List(1 (10 20) 3) element 1 should be (10 20) and not %v", sxp.At(1))
	case sxp.At(2) != 3:			t.Fatalf("List(1 (10 20) 3) element 2 should be 3 and not %v", sxp.At(2))
	}
}

func TestSliceString(t *testing.T) {
	ConfirmString := func(s *Slice, r string) {
		if x := s.String(); x != r {
			t.Fatalf("%v erroneously serialised as '%v'", r, x)
		}
	}

	ConfirmString(List(), "()")
	ConfirmString(List(0), "(0)")
	ConfirmString(List(0, 1), "(0 1)")
	ConfirmString(List(List(0, 1), 1), "((0 1) 1)")
	ConfirmString(List(List(0, 1), List(0, 1)), "((0 1) (0 1))")
}

func TestSliceLen(t *testing.T) {
	ConfirmLength := func(s *Slice, i int) {
		if x := s.Len(); x != i {
			t.Fatalf("%v.Len() should be %v but is %v", s, i, x)
		}
	}
	
	ConfirmLength(List(0), 1)
	ConfirmLength(List(0, 1), 2)
	ConfirmLength(List(List(0, 1), 2), 2)
	ConfirmLength(List(0, 1), 2)
	ConfirmLength(List(List(0, 1), 2), 2)

	sxp := List(0, 1, List(2, List(3, 4, 5)), List(6, 7, 8, 9))
	ConfirmLength(sxp, 4)
	ConfirmLength(List(0, 1, List(2, List(3, 4, 5)), sxp, List(6, 7, 8, 9)), 5)
}

func TestSliceSwap(t *testing.T) {
	ConfirmSwap := func(s *Slice, i, j int, r *Slice) {
		if s.Swap(i, j); !r.Equal(s) {
			t.Fatalf("Swap(%v, %v) should be %v but is %v", i, j, r, s)
		}
	}
	ConfirmSwap(List(0, 1, 2), 0, 1, List(1, 0, 2))
	ConfirmSwap(List(0, 1, 2), 0, 2, List(2, 1, 0))
}

func TestSliceRestrictTo(t *testing.T) {
	ConfirmRestrictTo := func(s *Slice, start, end int, r *Slice) {
		if s.RestrictTo(start, end); !r.Equal(s) {
			t.Fatalf("Trim(%v, %v) should be %v but is %v", start, end, r, s)
		}
	}

	ConfirmRestrictTo(List(0, 1, 2, 3, 4, 5), 0, 1, List(0))
	ConfirmRestrictTo(List(0, 1, 2, 3, 4, 5), 1, 2, List(1))
	ConfirmRestrictTo(List(0, 1, 2, 3, 4, 5), 2, 3, List(2))
	ConfirmRestrictTo(List(0, 1, 2, 3, 4, 5), 3, 4, List(3))
	ConfirmRestrictTo(List(0, 1, 2, 3, 4, 5), 4, 5, List(4))
	ConfirmRestrictTo(List(0, 1, 2, 3, 4, 5), 5, 6, List(5))

	ConfirmRestrictTo(List(0, 1, 2, 3, 4, 5), 0, 2, List(0, 1))
	ConfirmRestrictTo(List(0, 1, 2, 3, 4, 5), 1, 3, List(1, 2))
	ConfirmRestrictTo(List(0, 1, 2, 3, 4, 5), 2, 4, List(2, 3))
	ConfirmRestrictTo(List(0, 1, 2, 3, 4, 5), 3, 5, List(3, 4))
	ConfirmRestrictTo(List(0, 1, 2, 3, 4, 5), 4, 6, List(4, 5))
}

func TestSliceCut(t *testing.T) {
	ConfirmCut := func(s *Slice, start, end int, r *Slice) {
		if s.Cut(start, end); !r.Equal(s) {
			t.Fatalf("Cut(%v, %v) should be %v but is %v", start, end, r, s)
		}
	}

	ConfirmCut(List(0, 1, 2, 3, 4, 5), 0, 1, List(1, 2, 3, 4, 5))
	ConfirmCut(List(0, 1, 2, 3, 4, 5), 1, 2, List(0, 2, 3, 4, 5))
	ConfirmCut(List(0, 1, 2, 3, 4, 5), 2, 3, List(0, 1, 3, 4, 5))
	ConfirmCut(List(0, 1, 2, 3, 4, 5), 3, 4, List(0, 1, 2, 4, 5))
	ConfirmCut(List(0, 1, 2, 3, 4, 5), 4, 5, List(0, 1, 2, 3, 5))
	ConfirmCut(List(0, 1, 2, 3, 4, 5), 5, 6, List(0, 1, 2, 3, 4))

	ConfirmCut(List(0, 1, 2, 3, 4, 5), -1, 1, List(1, 2, 3, 4, 5))
	ConfirmCut(List(0, 1, 2, 3, 4, 5), 0, 2, List(2, 3, 4, 5))
	ConfirmCut(List(0, 1, 2, 3, 4, 5), 1, 3, List(0, 3, 4, 5))
	ConfirmCut(List(0, 1, 2, 3, 4, 5), 2, 4, List(0, 1, 4, 5))
	ConfirmCut(List(0, 1, 2, 3, 4, 5), 3, 5, List(0, 1, 2, 5))
	ConfirmCut(List(0, 1, 2, 3, 4, 5), 4, 6, List(0, 1, 2, 3))
	ConfirmCut(List(0, 1, 2, 3, 4, 5), 5, 7, List(0, 1, 2, 3, 4))
}

func TestSliceTrim(t *testing.T) {
	ConfirmTrim := func(s *Slice, start, end int, r *Slice) {
		if s.Trim(start, end); !r.Equal(s) {
			t.Fatalf("Trim(%v, %v) should be %v but is %v", start, end, r, s)
		}
	}

	ConfirmTrim(List(0, 1, 2, 3, 4, 5), 0, 1, List(0))
	ConfirmTrim(List(0, 1, 2, 3, 4, 5), 1, 2, List(1))
	ConfirmTrim(List(0, 1, 2, 3, 4, 5), 2, 3, List(2))
	ConfirmTrim(List(0, 1, 2, 3, 4, 5), 3, 4, List(3))
	ConfirmTrim(List(0, 1, 2, 3, 4, 5), 4, 5, List(4))
	ConfirmTrim(List(0, 1, 2, 3, 4, 5), 5, 6, List(5))

	ConfirmTrim(List(0, 1, 2, 3, 4, 5), -1, 1, List(0))
	ConfirmTrim(List(0, 1, 2, 3, 4, 5), 0, 2, List(0, 1))
	ConfirmTrim(List(0, 1, 2, 3, 4, 5), 1, 3, List(1, 2))
	ConfirmTrim(List(0, 1, 2, 3, 4, 5), 2, 4, List(2, 3))
	ConfirmTrim(List(0, 1, 2, 3, 4, 5), 3, 5, List(3, 4))
	ConfirmTrim(List(0, 1, 2, 3, 4, 5), 4, 6, List(4, 5))
	ConfirmTrim(List(0, 1, 2, 3, 4, 5), 5, 7, List(5))
}

func TestSliceDelete(t *testing.T) {
	ConfirmCut := func(s *Slice, index int, r *Slice) {
		if s.Delete(index); !r.Equal(s) {
			t.Fatalf("Delete(%v) should be %v but is %v", index, r, s)
		}
	}

	ConfirmCut(List(0, 1, 2, 3, 4, 5), -1, List(0, 1, 2, 3, 4, 5))
	ConfirmCut(List(0, 1, 2, 3, 4, 5), 0, List(1, 2, 3, 4, 5))
	ConfirmCut(List(0, 1, 2, 3, 4, 5), 1, List(0, 2, 3, 4, 5))
	ConfirmCut(List(0, 1, 2, 3, 4, 5), 2, List(0, 1, 3, 4, 5))
	ConfirmCut(List(0, 1, 2, 3, 4, 5), 3, List(0, 1, 2, 4, 5))
	ConfirmCut(List(0, 1, 2, 3, 4, 5), 4, List(0, 1, 2, 3, 5))
	ConfirmCut(List(0, 1, 2, 3, 4, 5), 5, List(0, 1, 2, 3, 4))
	ConfirmCut(List(0, 1, 2, 3, 4, 5), 6, List(0, 1, 2, 3, 4, 5))
}

func TestSliceEach(t *testing.T) {
	c := List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9)
	count := 0
	c.Each(func(i interface{}) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})
}

func TestSliceEachWithIndex(t *testing.T) {
	c := List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9)
	c.EachWithIndex(func(index int, i interface{}) {
		if i != index {
			t.Fatalf("element %v erroneously reported as %v", index, i)
		}
	})
}

func TestSliceEachWithKey(t *testing.T) {
	c := List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9)
	c.EachWithKey(func(key, i interface{}) {
		if i != key {
			t.Fatalf("element %v erroneously reported as %v", key, i)
		}
	})
}

func TestSliceBlockCopy(t *testing.T) {
	ConfirmBlockCopy := func(s *Slice, destination, source, count int, r *Slice) {
		s.BlockCopy(destination, source, count)
		if !r.Equal(s) {
			t.Fatalf("BlockCopy(%v, %v, %v) should be %v but is %v", destination, source, count, r, s)
		}
	}

	ConfirmBlockCopy(List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, 0, 4, List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockCopy(List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 9, 9, 4, List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockCopy(List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 2, 4, List(0, 1, 2, 3, 4, 2, 3, 4, 5, 9))
	ConfirmBlockCopy(List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 2, 5, 4, List(0, 1, 5, 6, 7, 8, 6, 7, 8, 9))
}

func TestSliceBlockClear(t *testing.T) {
	ConfirmBlockClear := func(s *Slice, start, count int, r *Slice) {
		s.BlockClear(start, count)
		if !r.Equal(s) {
			t.Fatalf("BlockClear(%v, %v) should be %v but is %v", start, count, r, s)
		}
	}

	ConfirmBlockClear(List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, 4, List(nil, nil, nil, nil, 4, 5, 6, 7, 8, 9))
	ConfirmBlockClear(List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, 4, List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockClear(List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 4, List(0, 1, 2, 3, 4, nil, nil, nil, nil, 9))
}

func TestSliceOverwrite(t *testing.T) {
	ConfirmOverwrite := func(s *Slice, offset int, v, r *Slice) {
		s.Overwrite(offset, *v)
		if !r.Equal(s) {
			t.Fatalf("Overwrite(%v, %v) should be %v but is %v", offset, v, r, s)
		}
	}

	ConfirmOverwrite(List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, List(10, 9, 8, 7), List(10, 9, 8, 7, 4, 5, 6, 7, 8, 9))
	ConfirmOverwrite(List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, List(10, 9, 8, 7), List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmOverwrite(List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, List(11, 12, 13, 14), List(0, 1, 2, 3, 4, 11, 12, 13, 14, 9))
}

func TestSliceReallocate(t *testing.T) {
	ConfirmReallocate := func(s *Slice, l, c int, r *Slice) {
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

	ConfirmReallocate(List(), 0, 10, List())
	ConfirmReallocate(List(0, 1, 2, 3, 4), 3, 10, List(0, 1, 2))
	ConfirmReallocate(List(0, 1, 2, 3, 4), 5, 10, List(0, 1, 2, 3, 4))
	ConfirmReallocate(List(0, 1, 2, 3, 4), 10, 10, List(0, 1, 2, 3, 4, nil, nil, nil, nil, nil))
	ConfirmReallocate(List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 1, 5, List(0))
	ConfirmReallocate(List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 5, List(0, 1, 2, 3, 4))
	ConfirmReallocate(List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, 5, List(0, 1, 2, 3, 4))
}

func TestSliceExtend(t *testing.T) {
	ConfirmExtend := func(s *Slice, n int, r *Slice) {
		c := s.Cap()
		s.Extend(n)
		switch {
		case s.Len() != r.Len():	t.Fatalf("Extend(%v) len should be %v but is %v", n, r.Len(), s.Len())
		case s.Cap() != c + n:		t.Fatalf("Extend(%v) cap should be %v but is %v", n, c + n, s.Cap())
		case !r.Equal(s):			t.Fatalf("Extend(%v) should be %v but is %v", n, r, s)
		}
	}

	ConfirmExtend(List(), 1, List(nil))
	ConfirmExtend(List(), 2, List(nil, nil))
}

func TestSliceExpand(t *testing.T) {
	ConfirmExpand := func(s *Slice, i, n int, r *Slice) {
		c := s.Cap()
		s.Expand(i, n)
		switch {
		case s.Len() != r.Len():	t.Fatalf("Expand(%v, %v) len should be %v but is %v", i, n, r.Len(), s.Len())
		case s.Cap() != c + n:		t.Fatalf("Expand(%v, %v) cap should be %v but is %v", i, n, c + n, s.Cap())
		case !r.Equal(s):			t.Fatalf("Expand(%v, %v) should be %v but is %v", i, n, r, s)
		}
	}

	ConfirmExpand(List(), -1, 1, List(nil))
	ConfirmExpand(List(), 0, 1, List(nil))
	ConfirmExpand(List(), 1, 1, List(nil))
	ConfirmExpand(List(), 0, 2, List(nil, nil))

	ConfirmExpand(List(0, 1, 2), -1, 2, List(nil, nil, 0, 1, 2))
	ConfirmExpand(List(0, 1, 2), 0, 2, List(nil, nil, 0, 1, 2))
	ConfirmExpand(List(0, 1, 2), 1, 2, List(0, nil, nil, 1, 2))
	ConfirmExpand(List(0, 1, 2), 2, 2, List(0, 1, nil, nil, 2))
	ConfirmExpand(List(0, 1, 2), 3, 2, List(0, 1, 2, nil, nil))
	ConfirmExpand(List(0, 1, 2), 4, 2, List(0, 1, 2, nil, nil))
}

func TestSliceDepth(t *testing.T) {
	ConfirmDepth := func(s *Slice, i int) {
		if x := s.Depth(); x != i {
			t.Fatalf("%v.Depth() should be %v but is %v", s, i, x)
		}
	}
	ConfirmDepth(List(0, 1), 0)
	ConfirmDepth(List(List(0, 1), 2), 1)
	ConfirmDepth(List(0, List(1, 2)), 1)
	ConfirmDepth(List(0, 1, List(2, List(3, 4, 5))), 2)

	sxp := List(0, 1,
				List(2, List(3, 4, 5)),
				List(6, List(7, List(8, List(9, 0)))),
				List(2, List(3, 4, 5)))
	ConfirmDepth(sxp, 4)

	rxp := List(0, sxp, sxp)
	ConfirmDepth(rxp, 5)
	ConfirmDepth(List(rxp, sxp), 6)
	t.Log("Need tests for circular recursive Slice?")
}

func TestSliceReverse(t *testing.T) {
	sxp := List(1, 2, 3, 4, 5)
	rxp := List(5, 4, 3, 2, 1)
	if sxp.Reverse(); !rxp.Equal(sxp) {
		t.Fatalf("Reversal failed: %v", sxp)
	}
}

func TestSliceAppend(t *testing.T) {
	ConfirmAppend := func(s *Slice, v interface{}, r *Slice) {
		if s.Append(v); !r.Equal(s) {
			t.Fatalf("Append(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmAppend(List(), 0, List(0))
	ConfirmAppend(List(), List(0, 1), List(List(0, 1)))
	ConfirmAppend(List(0, 1, 2), List(3, 4), List(0, 1, 2, List(3, 4)))
}

func TestSliceAppendSlice(t *testing.T) {
	ConfirmAppendSlice := func(s, v, r *Slice) {
		if s.AppendSlice(*v); !r.Equal(s) {
			t.Fatalf("AppendSlice(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmAppendSlice(List(), List(0), List(0))
	ConfirmAppendSlice(List(), List(0, 1), List(0, 1))
	ConfirmAppendSlice(List(0, 1, 2), List(3, 4), List(0, 1, 2, 3, 4))
}

func TestSlicePrepend(t *testing.T) {
	ConfirmPrepend := func(s *Slice, v interface{}, r *Slice) {
		if s.Prepend(v); !r.Equal(s) {
			t.Fatalf("Prepend(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmPrepend(List(), 0, List(0))
	ConfirmPrepend(List(0), 1, List(1, 0))
	ConfirmPrepend(List(0, 1, 2), List(3, 4), List(List(3, 4), 0, 1, 2))
}

func TestSlicePrependSlice(t *testing.T) {
	ConfirmPrependSlice := func(s, v, r *Slice) {
		if s.PrependSlice(*v); !r.Equal(s) {
			t.Fatalf("PrependSlice(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmPrependSlice(List(), List(0), List(0))
	ConfirmPrependSlice(List(), List(0, 1), List(0, 1))
	ConfirmPrependSlice(List(0, 1, 2), List(3, 4), List(3, 4, 0, 1, 2))
}

func TestSliceRepeat(t *testing.T) {
	ConfirmRepeat := func(s *Slice, count int, r *Slice) {
		if x := s.Repeat(count); !x.Equal(r) {
			t.Fatalf("%v.Repeat(%v) should be %v but is %v", s, count, r, x)
		}
	}

	ConfirmRepeat(List(), 5, List())
	ConfirmRepeat(List(0), 1, List(0))
	ConfirmRepeat(List(0), 2, List(0, 0))
	ConfirmRepeat(List(0), 3, List(0, 0, 0))
	ConfirmRepeat(List(0), 4, List(0, 0, 0, 0))
	ConfirmRepeat(List(0), 5, List(0, 0, 0, 0, 0))
}

func TestSliceFlatten(t *testing.T) {
	ConfirmFlatten := func(s, r *Slice) {
		if s.Flatten(); !s.Equal(r) {
			t.Fatalf("%v should be %v", s, r)
		}
	}
	ConfirmFlatten(List(), List())
	ConfirmFlatten(List(1), List(1))
	ConfirmFlatten(List(1, List(2)), List(1, 2))
	ConfirmFlatten(List(1, List(2, List(3))), List(1, 2, 3))
	ConfirmFlatten(List(1, 2, List(3, List(4, 5), List(6, List(7, 8, 9), List(10, 11)))), List(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11))

	ConfirmFlatten(List(0, lists.List(1, 2, List(3, 4))), List(0, lists.List(1, 2, List(3, 4))))
	ConfirmFlatten(List(0, lists.List(1, 2, lists.List(3, 4))), List(0, lists.List(1, 2, 3, 4)))

	ConfirmFlatten(List(0, lists.Loop(1, 2)), List(0, lists.Loop(1, 2)))
	ConfirmFlatten(List(0, lists.List(1, lists.Loop(2, 3))), List(0, lists.List(1, 2, 3)))

	ConfirmFlatten(List(0, lists.List(1, 2, lists.Loop(3, 4))), List(0, lists.List(1, 2, 3, 4)))
	ConfirmFlatten(List(3, 4, List(5, 6, 7)), List(3, 4, 5, 6, 7))
	ConfirmFlatten(List(0, lists.Loop(1, 2, List(3, 4, List(5, 6, 7)))), List(0, lists.Loop(1, 2, List(3, 4, 5, 6, 7))))

	sxp := List(1, 2, List(3, List(4, 5), List(6, List(7, 8, 9), List(10, 11))))
	rxp := List(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	sxp.Flatten()
	if !rxp.Equal(sxp) {
		t.Fatalf("Flatten failed: %v", sxp)
	}

	rxp = List(1, 2, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 3, 4, 5, 6, 7, 8, 9, 10, 11, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	sxp = List(1, 2, sxp, List(3, List(4, 5), List(6, List(7, 8, 9), List(10, 11), sxp)))
	sxp.Flatten()
	if !rxp.Equal(sxp) {
		t.Fatalf("Flatten failed with explicit expansions: %v", sxp)
	}
}

func TestSliceCar(t *testing.T) {
	ConfirmCar := func(s *Slice, r interface{}) {
		var ok bool
		n := s.Car()
		switch n := n.(type) {
		case Equatable:		ok = n.Equal(r)
		default:			ok = n == r
		}
		if !ok {
			t.Fatalf("head should be '%v' but is '%v'", r, n)
		}
	}
	ConfirmCar(List(1, 2, 3), 1)
	ConfirmCar(List(List(10, 20), 2, 3), List(10, 20))
}

func TestSliceCdr(t *testing.T) {
	ConfirmCdr := func(s, r *Slice) {
		if n := s.Cdr(); !n.Equal(r) {
			t.Fatalf("tail should be '%v' but is '%v'", r, n)
		}
	}
	ConfirmCdr(List(1, 2, 3), List(2, 3))
}

func TestSliceRplaca(t *testing.T) {
	ConfirmRplaca := func(s *Slice, v interface{}, r *Slice) {
		if s.Rplaca(v); !s.Equal(r) {
			t.Fatalf("slice should be '%v' but is '%v'", r, s)
		}
	}
	ConfirmRplaca(List(1, 2, 3, 4, 5), 0, List(0, 2, 3, 4, 5))
	ConfirmRplaca(List(1, 2, 3, 4, 5), List(1, 2, 3), List(List(1, 2, 3), 2, 3, 4, 5))
}

func TestSliceRplacd(t *testing.T) {
	ConfirmRplacd := func(s *Slice, v interface{}, r *Slice) {
		if s.Rplacd(v); !s.Equal(r) {
			t.Fatalf("slice should be '%v' but is '%v'", r, s)
		}
	}
	ConfirmRplacd(List(1, 2, 3, 4, 5), nil, List(1))
	ConfirmRplacd(List(1, 2, 3, 4, 5), 10, List(1, 10))
	ConfirmRplacd(List(1, 2, 3, 4, 5), List(5, 4, 3, 2), List(1, 5, 4, 3, 2))
	ConfirmRplacd(List(1, 2, 3, 4, 5, 6), List(2, 4, 8, 16), List(1, 2, 4, 8, 16))
}

func TestSliceSetIntersection(t *testing.T) {
	ConfirmSetIntersection := func(s, o, r *Slice) {
		if x := s.SetIntersection(*o); !r.Equal(x) {
			t.Fatalf("%v.SetIntersection(%v) should be %v but is %v", s, o, r, x)
		}
	}

	ConfirmSetIntersection(List(1, 2, 3), List(), List())
	ConfirmSetIntersection(List(1, 2, 3), List(1), List(1))
	ConfirmSetIntersection(List(1, 2, 3), List(1, 1), List(1))
	ConfirmSetIntersection(List(1, 2, 3), List(1, 2, 1), List(1, 2))
}

func TestSliceSetUnion(t *testing.T) {
	ConfirmSetUnion := func(s, o, r *Slice) {
		if x := s.SetUnion(*o); !r.Equal(x) {
			t.Fatalf("%v.SetUnion(%v) should be %v but is %v", s, o, r, x)
		}
	}

	ConfirmSetUnion(List(1, 2, 3), List(), List(1, 2, 3))
	ConfirmSetUnion(List(1, 2, 3), List(1), List(1, 2, 3))
	ConfirmSetUnion(List(1, 2, 3), List(1, 1), List(1, 2, 3))
	ConfirmSetUnion(List(1, 2, 3), List(1, 2, 1), List(1, 2, 3))
}

func TestSliceSetDifference(t *testing.T) {
	ConfirmSetUnion := func(s, o, r *Slice) {
		if x := s.SetDifference(*o); !r.Equal(x) {
			t.Fatalf("%v.SetUnion(%v) should be %v but is %v", s, o, r, x)
		}
	}

	ConfirmSetUnion(List(1, 2, 3), List(), List(1, 2, 3))
	ConfirmSetUnion(List(1, 2, 3), List(1), List(2, 3))
	ConfirmSetUnion(List(1, 2, 3), List(1, 1), List(2, 3))
	ConfirmSetUnion(List(1, 2, 3), List(1, 2, 1), List(3))
}