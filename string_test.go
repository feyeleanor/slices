package slices

import "testing"

func TestSSlice(t *testing.T) {
	sxp := SList("A")
	switch {
	case sxp.Len() != 1:			t.Fatalf("SList(A) should allocate 1 cells, not %v cells", sxp.Len())
	case sxp.At(0) != "A":			t.Fatalf("SList(A) element 0 should be A and not %v", sxp.At(0))
	}

	sxp = SList("A", "B")
	switch {
	case sxp.Len() != 2:			t.Fatalf("SList(A B) should allocate 2 cells, not %v cells", sxp.Len())
	case sxp.At(0) != "A":			t.Fatalf("SList(A B) element 0 should be A and not %v", sxp.At(0))
	case sxp.At(1) != "B":			t.Fatalf("SList(A B) element 1 should be B and not %v", sxp.At(1))
	}

	sxp = SList("A", "B", "C")
	switch {
	case sxp.Len() != 3:			t.Fatalf("SList(A B C) should allocate 3 cells, not %v cells", sxp.Len())
	case sxp.At(0) != "A":			t.Fatalf("SList(A B C) element 0 should be A and not %v", sxp.At(0))
	case sxp.At(1) != "B":			t.Fatalf("SList(A B C) element 1 should be B and not %v", sxp.At(1))
	case sxp.At(2) != "C":			t.Fatalf("SList(A B C) element 2 should be C and not %v", sxp.At(2))
	}
}

func TestSSliceString(t *testing.T) {
	ConfirmString := func(s *SSlice, r string) {
		if x := s.String(); x != r {
			t.Fatalf("%v erroneously serialised as '%v'", r, x)
		}
	}

	ConfirmString(SList(), "()")
	ConfirmString(SList("A"), "(A)")
	ConfirmString(SList("A", "B"), "(A B)")
}

func TestSSliceLen(t *testing.T) {
	ConfirmLength := func(s *SSlice, i int) {
		if x := s.Len(); x != i {
			t.Fatalf("%v.Len() should be %v but is %v", s, i, x)
		}
	}
	
	ConfirmLength(SList("A"), 1)
	ConfirmLength(SList("A", "B"), 2)
}

func TestSSliceSwap(t *testing.T) {
	ConfirmSwap := func(s *SSlice, i, j int, r *SSlice) {
		if s.Swap(i, j); !r.Equal(s) {
			t.Fatalf("Swap(%v, %v) should be %v but is %v", i, j, r, s)
		}
	}
	ConfirmSwap(SList("A", "B", "C"), 0, 1, SList("B", "A", "C"))
	ConfirmSwap(SList("A", "B", "C"), 0, 2, SList("C", "B", "A"))
}

func TestSSliceSort(t *testing.T) {
	ConfirmSort := func(s, r *SSlice) {
		if s.Sort(); !r.Equal(s) {
			t.Fatalf("Sort() should be %v but is %v", r, s)
		}
	}

	ConfirmSort(SList("D", "C", "B", "E", "F", "A"), SList("A", "B", "C", "D", "E", "F"))
}

func TestSSliceCompare(t *testing.T) {
	ConfirmCompare := func(s *SSlice, i, j, r int) {
		if x := s.Compare(i, j); x != r {
			t.Fatalf("Compare(%v, %v) should be %v but is %v", i, j, r, x)
		}
	}

	ConfirmCompare(SList("A", "B"), 0, 0, IS_SAME_AS)
	ConfirmCompare(SList("A", "B"), 0, 1, IS_LESS_THAN)
	ConfirmCompare(SList("A", "B"), 1, 0, IS_GREATER_THAN)
}

func TestSSliceCut(t *testing.T) {
	ConfirmCut := func(s *SSlice, start, end int, r *SSlice) {
		if s.Cut(start, end); !r.Equal(s) {
			t.Fatalf("Cut(%v, %v) should be %v but is %v", start, end, r, s)
		}
	}

	ConfirmCut(SList("A", "B", "C", "D", "E", "F"), 0, 1, SList("B", "C", "D", "E", "F"))
	ConfirmCut(SList("A", "B", "C", "D", "E", "F"), 1, 2, SList("A", "C", "D", "E", "F"))
	ConfirmCut(SList("A", "B", "C", "D", "E", "F"), 2, 3, SList("A", "B", "D", "E", "F"))
	ConfirmCut(SList("A", "B", "C", "D", "E", "F"), 3, 4, SList("A", "B", "C", "E", "F"))
	ConfirmCut(SList("A", "B", "C", "D", "E", "F"), 4, 5, SList("A", "B", "C", "D", "F"))
	ConfirmCut(SList("A", "B", "C", "D", "E", "F"), 5, 6, SList("A", "B", "C", "D", "E"))

	ConfirmCut(SList("A", "B", "C", "D", "E", "F"), -1, 1, SList("B", "C", "D", "E", "F"))
	ConfirmCut(SList("A", "B", "C", "D", "E", "F"), 0, 2, SList("C", "D", "E", "F"))
	ConfirmCut(SList("A", "B", "C", "D", "E", "F"), 1, 3, SList("A", "D", "E", "F"))
	ConfirmCut(SList("A", "B", "C", "D", "E", "F"), 2, 4, SList("A", "B", "E", "F"))
	ConfirmCut(SList("A", "B", "C", "D", "E", "F"), 3, 5, SList("A", "B", "C", "F"))
	ConfirmCut(SList("A", "B", "C", "D", "E", "F"), 4, 6, SList("A", "B", "C", "D"))
	ConfirmCut(SList("A", "B", "C", "D", "E", "F"), 5, 7, SList("A", "B", "C", "D", "E"))
}

func TestSSliceTrim(t *testing.T) {
	ConfirmTrim := func(s *SSlice, start, end int, r *SSlice) {
		if s.Trim(start, end); !r.Equal(s) {
			t.Fatalf("Trim(%v, %v) should be %v but is %v", start, end, r, s)
		}
	}

	ConfirmTrim(SList("A", "B", "C", "D", "E", "F"), 0, 1, SList("A"))
	ConfirmTrim(SList("A", "B", "C", "D", "E", "F"), 1, 2, SList("B"))
	ConfirmTrim(SList("A", "B", "C", "D", "E", "F"), 2, 3, SList("C"))
	ConfirmTrim(SList("A", "B", "C", "D", "E", "F"), 3, 4, SList("D"))
	ConfirmTrim(SList("A", "B", "C", "D", "E", "F"), 4, 5, SList("E"))
	ConfirmTrim(SList("A", "B", "C", "D", "E", "F"), 5, 6, SList("F"))

	ConfirmTrim(SList("A", "B", "C", "D", "E", "F"), -1, 1, SList("A"))
	ConfirmTrim(SList("A", "B", "C", "D", "E", "F"), 0, 2, SList("A", "B"))
	ConfirmTrim(SList("A", "B", "C", "D", "E", "F"), 1, 3, SList("B", "C"))
	ConfirmTrim(SList("A", "B", "C", "D", "E", "F"), 2, 4, SList("C", "D"))
	ConfirmTrim(SList("A", "B", "C", "D", "E", "F"), 3, 5, SList("D", "E"))
	ConfirmTrim(SList("A", "B", "C", "D", "E", "F"), 4, 6, SList("E", "F"))
	ConfirmTrim(SList("A", "B", "C", "D", "E", "F"), 5, 7, SList("F"))
}

func TestSSliceDelete(t *testing.T) {
	ConfirmCut := func(s *SSlice, index int, r *SSlice) {
		if s.Delete(index); !r.Equal(s) {
			t.Fatalf("Delete(%v) should be %v but is %v", index, r, s)
		}
	}

	ConfirmCut(SList("A", "B", "C", "D", "E", "F"), -1, SList("A", "B", "C", "D", "E", "F"))
	ConfirmCut(SList("A", "B", "C", "D", "E", "F"), 0, SList("B", "C", "D", "E", "F"))
	ConfirmCut(SList("A", "B", "C", "D", "E", "F"), 1, SList("A", "C", "D", "E", "F"))
	ConfirmCut(SList("A", "B", "C", "D", "E", "F"), 2, SList("A", "B", "D", "E", "F"))
	ConfirmCut(SList("A", "B", "C", "D", "E", "F"), 3, SList("A", "B", "C", "E", "F"))
	ConfirmCut(SList("A", "B", "C", "D", "E", "F"), 4, SList("A", "B", "C", "D", "F"))
	ConfirmCut(SList("A", "B", "C", "D", "E", "F"), 5, SList("A", "B", "C", "D", "E"))
	ConfirmCut(SList("A", "B", "C", "D", "E", "F"), 6, SList("A", "B", "C", "D", "E", "F"))
}

func TestSSliceEach(t *testing.T) {
	c := SList("A", "B", "C", "D", "E", "F")
	count := 0
	c.Each(func(i interface{}) {
		if i != string([]byte{ byte(count) + "A"[0] }) {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})
}

func TestSSliceEachWithIndex(t *testing.T) {
	c := SList("A", "B", "C", "D", "E", "F")
	c.EachWithIndex(func(index int, i interface{}) {
		if i != string([]byte{ byte(index) + "A"[0] }) {
			t.Fatalf("element %v erroneously reported as %v", index, i)
		}
	})
}

func TestSSliceEachWithKey(t *testing.T) {
	c := SList("A", "B", "C", "D", "E", "F")
	c.EachWithKey(func(key, i interface{}) {
		if i != string([]byte{ byte(key.(int)) + "A"[0] }) {
			t.Fatalf("element %v erroneously reported as %v", key, i)
		}
	})
}

func TestSSliceIEach(t *testing.T) {
	c := SList("A", "B", "C", "D", "E", "F")
	count := 0
	c.SEach(func(i string) {
		if i != string([]byte{ byte(count) + "A"[0] }) {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})
}

func TestSSliceIEachWithIndex(t *testing.T) {
	c := SList("A", "B", "C", "D", "E", "F")
	c.SEachWithIndex(func(index int, i string) {
		if i != string([]byte{ byte(index) + "A"[0] }) {
			t.Fatalf("element %v erroneously reported as %v", index, i)
		}
	})
}

func TestSSliceIEachWithKey(t *testing.T) {
	c := SList("A", "B", "C", "D", "E", "F")
	c.SEachWithKey(func(key interface{}, i string) {
		if i != string([]byte{ byte(key.(int)) + "A"[0] }) {
			t.Fatalf("element %v erroneously reported as %v", key, i)
		}
	})
}

func TestSSliceBlockCopy(t *testing.T) {
	ConfirmBlockCopy := func(s *SSlice, destination, source, count int, r *SSlice) {
		s.BlockCopy(destination, source, count)
		if !r.Equal(s) {
			t.Fatalf("BlockCopy(%v, %v, %v) should be %v but is %v", destination, source, count, r, s)
		}
	}

	ConfirmBlockCopy(SList("A", "B", "C", "D", "E", "F"), 0, 0, 4, SList("A", "B", "C", "D", "E", "F"))
	ConfirmBlockCopy(SList("A", "B", "C", "D", "E", "F"), 6, 6, 4, SList("A", "B", "C", "D", "E", "F"))
	ConfirmBlockCopy(SList("A", "B", "C", "D", "E", "F"), 4, 2, 2, SList("A", "B", "C", "D", "C", "D"))
	ConfirmBlockCopy(SList("A", "B", "C", "D", "E", "F"), 2, 4, 4, SList("A", "B", "E", "F", "E", "F"))
}

func TestSSliceBlockClear(t *testing.T) {
	ConfirmBlockClear := func(s *SSlice, start, count int, r *SSlice) {
		s.BlockClear(start, count)
		if !r.Equal(s) {
			t.Fatalf("BlockClear(%v, %v) should be %v but is %v", start, count, r, s)
		}
	}

	ConfirmBlockClear(SList("A", "B", "C", "D", "E", "F"), 0, 4, SList("", "", "", "", "E", "F"))
	ConfirmBlockClear(SList("A", "B", "C", "D", "E", "F"), 1, 4, SList("A", "", "", "", "", "F"))
	ConfirmBlockClear(SList("A", "B", "C", "D", "E", "F"), 2, 4, SList("A", "B", "", "", "", ""))
}

func TestSSliceOverwrite(t *testing.T) {
	ConfirmOverwrite := func(s *SSlice, offset int, v, r *SSlice) {
		s.Overwrite(offset, *v)
		if !r.Equal(s) {
			t.Fatalf("Overwrite(%v, %v) should be %v but is %v", offset, v, r, s)
		}
	}

	ConfirmOverwrite(SList("A", "B", "C", "D", "E", "F"), 0, SList("Z", "Y", "X"), SList("Z", "Y", "X", "D", "E", "F"))
	ConfirmOverwrite(SList("A", "B", "C", "D", "E", "F"), 6, SList("Z", "Y", "X"), SList("A", "B", "C", "D", "E", "F"))
	ConfirmOverwrite(SList("A", "B", "C", "D", "E", "F"), 2, SList("Z", "Y", "X"), SList("A", "B", "Z", "Y", "X", "F"))
}

func TestSSliceReallocate(t *testing.T) {
	ConfirmReallocate := func(s *SSlice, l, c int, r *SSlice) {
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

	ConfirmReallocate(SList(), 0, 10, SList())
	ConfirmReallocate(SList("A", "B", "C", "D", "E", "F"), 3, 10, SList("A", "B", "C"))
	ConfirmReallocate(SList("A", "B", "C", "D", "E", "F"), 6, 10, SList("A", "B", "C", "D", "E", "F"))
	ConfirmReallocate(SList("A", "B", "C", "D", "E", "F"), 10, 10, SList("A", "B", "C", "D", "E", "F", "", "", "", ""))
	ConfirmReallocate(SList("A", "B", "C", "D", "E", "F"), 1, 3, SList("A"))
	ConfirmReallocate(SList("A", "B", "C", "D", "E", "F"), 3, 3, SList("A", "B", "C"))
	ConfirmReallocate(SList("A", "B", "C", "D", "E", "F"), 6, 3, SList("A", "B", "C"))
}

func TestSSliceExtend(t *testing.T) {
	ConfirmExtend := func(s *SSlice, n int, r *SSlice) {
		c := s.Cap()
		s.Extend(n)
		switch {
		case s.Len() != r.Len():	t.Fatalf("Extend(%v) len should be %v but is %v", n, r.Len(), s.Len())
		case s.Cap() != c + n:		t.Fatalf("Extend(%v) cap should be %v but is %v", n, c + n, s.Cap())
		case !r.Equal(s):			t.Fatalf("Extend(%v) should be %v but is %v", n, r, s)
		}
	}

	ConfirmExtend(SList(), 1, SList(""))
	ConfirmExtend(SList(), 2, SList("", ""))
}

func TestSSliceExpand(t *testing.T) {
	ConfirmExpand := func(s *SSlice, i, n int, r *SSlice) {
		c := s.Cap()
		s.Expand(i, n)
		switch {
		case s.Len() != r.Len():	t.Fatalf("Expand(%v, %v) len should be %v but is %v", i, n, r.Len(), s.Len())
		case s.Cap() != c + n:		t.Fatalf("Expand(%v, %v) cap should be %v but is %v", i, n, c + n, s.Cap())
		case !r.Equal(s):			t.Fatalf("Expand(%v, %v) should be %v but is %v", i, n, r, s)
		}
	}

	ConfirmExpand(SList(), -1, 1, SList(""))
	ConfirmExpand(SList(), 0, 1, SList(""))
	ConfirmExpand(SList(), 1, 1, SList(""))
	ConfirmExpand(SList(), 0, 2, SList("", ""))

	ConfirmExpand(SList("A", "B", "C"), -1, 2, SList("", "", "A", "B", "C"))
	ConfirmExpand(SList("A", "B", "C"), 0, 2, SList("", "", "A", "B", "C"))
	ConfirmExpand(SList("A", "B", "C"), 1, 2, SList("A", "", "", "B", "C"))
	ConfirmExpand(SList("A", "B", "C"), 2, 2, SList("A", "B", "", "", "C"))
	ConfirmExpand(SList("A", "B", "C"), 3, 2, SList("A", "B", "C", "", ""))
	ConfirmExpand(SList("A", "B", "C"), 4, 2, SList("A", "B", "C", "", ""))
}

func TestSSliceDepth(t *testing.T) {
	ConfirmDepth := func(s *SSlice, i int) {
		if x := s.Depth(); x != i {
			t.Fatalf("%v.Depth() should be %v but is %v", s, i, x)
		}
	}
	ConfirmDepth(SList("A", "B"), 0)
}

func TestSSliceReverse(t *testing.T) {
	sxp := SList("A", "B", "C", "D", "E", "F")
	rxp := SList("F", "E", "D", "C", "B", "A")
	sxp.Reverse()
	if !rxp.Equal(sxp) {
		t.Fatalf("Reversal failed: %v", sxp)
	}
}

func TestSSliceAppend(t *testing.T) {
	ConfirmAppend := func(s *SSlice, v interface{}, r *SSlice) {
		s.Append(v)
		if !r.Equal(s) {
			t.Fatalf("Append(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmAppend(SList(), "A", SList("A"))
}

func TestSSliceAppendSlice(t *testing.T) {
	ConfirmAppendSlice := func(s, v, r *SSlice) {
		s.AppendSlice(*v)
		if !r.Equal(s) {
			t.Fatalf("AppendSlice(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmAppendSlice(SList(), SList("A"), SList("A"))
	ConfirmAppendSlice(SList(), SList("A", "B"), SList("A", "B"))
	ConfirmAppendSlice(SList("A", "B", "C"), SList("D", "E", "F"), SList("A", "B", "C", "D", "E", "F"))
}

func TestSSlicePrepend(t *testing.T) {
	ConfirmPrepend := func(s *SSlice, v interface{}, r *SSlice) {
		if s.Prepend(v); !r.Equal(s) {
			t.Fatalf("Prepend(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmPrepend(SList(), "A", SList("A"))
	ConfirmPrepend(SList("A"), "B", SList("B", "A"))
}

func TestSSlicePrependSlice(t *testing.T) {
	ConfirmPrependSlice := func(s, v, r *SSlice) {
		if s.PrependSlice(*v); !r.Equal(s) {
			t.Fatalf("PrependSlice(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmPrependSlice(SList(), SList("A"), SList("A"))
	ConfirmPrependSlice(SList(), SList("A", "B"), SList("A", "B"))
	ConfirmPrependSlice(SList("A", "B", "C"), SList("D", "E", "F"), SList("D", "E", "F", "A", "B", "C"))
}

func TestSSliceRepeat(t *testing.T) {
	ConfirmRepeat := func(s *SSlice, count int, r *SSlice) {
		if x := s.Repeat(count); !x.Equal(r) {
			t.Fatalf("%v.Repeat(%v) should be %v but is %v", s, count, r, x)
		}
	}

	ConfirmRepeat(SList(), 5, SList())
	ConfirmRepeat(SList("A"), 1, SList("A"))
	ConfirmRepeat(SList("A"), 2, SList("A", "A"))
	ConfirmRepeat(SList("A"), 3, SList("A", "A", "A"))
	ConfirmRepeat(SList("A"), 4, SList("A", "A", "A", "A"))
	ConfirmRepeat(SList("A"), 5, SList("A", "A", "A", "A", "A"))
}

func TestSSliceCar(t *testing.T) {
	ConfirmCar := func(s *SSlice, r string) {
		n := s.Car()
		if ok := n == r; !ok {
			t.Fatalf("head should be '%v' but is '%v'", r, n)
		}
	}
	ConfirmCar(SList("A", "B", "C", "D", "E", "F"), "A")
}

func TestSSliceCdr(t *testing.T) {
	ConfirmCdr := func(s, r *SSlice) {
		if n := s.Cdr(); !n.Equal(r) {
			t.Fatalf("tail should be '%v' but is '%v'", r, n)
		}
	}
	ConfirmCdr(SList("A", "B", "C", "D", "E", "F"), SList("B", "C", "D", "E", "F"))
}

func TestSSliceRplaca(t *testing.T) {
	ConfirmRplaca := func(s *SSlice, v interface{}, r *SSlice) {
		if s.Rplaca(v); !s.Equal(r) {
			t.Fatalf("slice should be '%v' but is '%v'", r, s)
		}
	}
	ConfirmRplaca(SList("A", "B", "C", "D", "E", "F"), "B", SList("B", "B", "C", "D", "E", "F"))
}

func TestSSliceRplacd(t *testing.T) {
	ConfirmRplacd := func(s *SSlice, v interface{}, r *SSlice) {
		if s.Rplacd(v); !s.Equal(r) {
			t.Fatalf("slice should be '%v' but is '%v'", r, s)
		}
	}
	ConfirmRplacd(SList("A", "B", "C", "D", "E", "F"), nil, SList("A"))
	ConfirmRplacd(SList("A", "B", "C", "D", "E", "F"), "B", SList("A", "B"))
	ConfirmRplacd(SList("A", "B", "C", "D", "E", "F"), SList("F", "E", "D", "C"), SList("A", "F", "E", "D", "C"))
}

func TestSSliceSetIntersection(t *testing.T) {
	ConfirmSetIntersection := func(s, o, r *SSlice) {
		x := s.SetIntersection(*o)
		x.Sort()
		if !r.Equal(x) {
			t.Fatalf("%v.SetIntersection(%v) should be %v but is %v", s, o, r, x)
		}
	}

	ConfirmSetIntersection(SList("A", "B", "C"), SList(), SList())
	ConfirmSetIntersection(SList("A", "B", "C"), SList("A"), SList("A"))
	ConfirmSetIntersection(SList("A", "B", "C"), SList("A", "A"), SList("A"))
	ConfirmSetIntersection(SList("A", "B", "C"), SList("A", "B", "A"), SList("A", "B"))
}

func TestSSliceSetUnion(t *testing.T) {
	ConfirmSetUnion := func(s, o, r *SSlice) {
		x := s.SetUnion(*o)
		x.Sort()
		if !r.Equal(x) {
			t.Fatalf("%v.SetUnion(%v) should be %v but is %v", s, o, r, x)
		}
	}

	ConfirmSetUnion(SList("A", "B", "C"), SList(), SList("A", "B", "C"))
	ConfirmSetUnion(SList("A", "B", "C"), SList("A"), SList("A", "B", "C"))
	ConfirmSetUnion(SList("A", "B", "C"), SList("A", "A"), SList("A", "B", "C"))
	ConfirmSetUnion(SList("A", "B", "C"), SList("A", "B", "A"), SList("A", "B", "C"))
}