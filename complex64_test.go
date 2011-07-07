package slices

import "testing"

func TestC64Slice(t *testing.T) {
	sxp := C64List(1)
	switch {
	case sxp.Len() != 1:			t.Fatalf("C64List(1) should allocate 1 cells, not %v cells", sxp.Len())
	case sxp.C64At(0) != 1:			t.Fatalf("C64List(1) element 0 should be 1 and not %v", sxp.C64At(0))
	}

	sxp = C64List(1, 2)
	switch {
	case sxp.Len() != 2:			t.Fatalf("C64List(1 2) should allocate 2 cells, not %v cells", sxp.Len())
	case sxp.C64At(0) != 1:			t.Fatalf("C64List(1 2) element 0 should be 1 and not %v", sxp.C64At(0))
	case sxp.C64At(1) != 2:			t.Fatalf("C64List(1 2) element 1 should be 2 and not %v", sxp.C64At(1))
	}

	sxp = C64List(1, 2, 3)
	switch {
	case sxp.Len() != 3:			t.Fatalf("C64List(1 2 3) should allocate 3 cells, not %v cells", sxp.Len())
	case sxp.C64At(0) != 1:			t.Fatalf("C64List(1 2 3) element 0 should be 1 and not %v", sxp.C64At(0))
	case sxp.C64At(1) != 2:			t.Fatalf("C64List(1 2 3) element 1 should be 2 and not %v", sxp.C64At(1))
	case sxp.C64At(2) != 3:			t.Fatalf("C64List(1 2 3) element 2 should be 3 and not %v", sxp.C64At(2))
	}
}

func TestC64SliceString(t *testing.T) {
	ConfirmString := func(s *C64Slice, r string) {
		if x := s.String(); x != r {
			t.Fatalf("%v erroneously serialised as '%v'", r, x)
		}
	}

	ConfirmString(C64List(), "()")
	ConfirmString(C64List(0), "((0+0i))")
	ConfirmString(C64List(0, 1), "((0+0i) (1+0i))")
	ConfirmString(C64List(0, 1i), "((0+0i) (0+1i))")
}

func TestC64SliceLen(t *testing.T) {
	ConfirmLength := func(s *C64Slice, i int) {
		if x := s.Len(); x != i {
			t.Fatalf("%v.Len() should be %v but is %v", s, i, x)
		}
	}
	
	ConfirmLength(C64List(0), 1)
	ConfirmLength(C64List(0, 1), 2)
}

func TestC64SliceSwap(t *testing.T) {
	ConfirmSwap := func(s *C64Slice, i, j int, r *C64Slice) {
		if s.Swap(i, j); !r.Equal(s) {
			t.Fatalf("Swap(%v, %v) should be %v but is %v", i, j, r, s)
		}
	}
	ConfirmSwap(C64List(0, 1, 2), 0, 1, C64List(1, 0, 2))
	ConfirmSwap(C64List(0, 1, 2), 0, 2, C64List(2, 1, 0))
}

func TestC64SliceCut(t *testing.T) {
	ConfirmCut := func(s *C64Slice, start, end int, r *C64Slice) {
		if s.Cut(start, end); !r.Equal(s) {
			t.Fatalf("Cut(%v, %v) should be %v but is %v", start, end, r, s)
		}
	}

	ConfirmCut(C64List(0, 1, 2, 3, 4, 5), 0, 1, C64List(1, 2, 3, 4, 5))
	ConfirmCut(C64List(0, 1, 2, 3, 4, 5), 1, 2, C64List(0, 2, 3, 4, 5))
	ConfirmCut(C64List(0, 1, 2, 3, 4, 5), 2, 3, C64List(0, 1, 3, 4, 5))
	ConfirmCut(C64List(0, 1, 2, 3, 4, 5), 3, 4, C64List(0, 1, 2, 4, 5))
	ConfirmCut(C64List(0, 1, 2, 3, 4, 5), 4, 5, C64List(0, 1, 2, 3, 5))
	ConfirmCut(C64List(0, 1, 2, 3, 4, 5), 5, 6, C64List(0, 1, 2, 3, 4))

	ConfirmCut(C64List(0, 1, 2, 3, 4, 5), -1, 1, C64List(1, 2, 3, 4, 5))
	ConfirmCut(C64List(0, 1, 2, 3, 4, 5), 0, 2, C64List(2, 3, 4, 5))
	ConfirmCut(C64List(0, 1, 2, 3, 4, 5), 1, 3, C64List(0, 3, 4, 5))
	ConfirmCut(C64List(0, 1, 2, 3, 4, 5), 2, 4, C64List(0, 1, 4, 5))
	ConfirmCut(C64List(0, 1, 2, 3, 4, 5), 3, 5, C64List(0, 1, 2, 5))
	ConfirmCut(C64List(0, 1, 2, 3, 4, 5), 4, 6, C64List(0, 1, 2, 3))
	ConfirmCut(C64List(0, 1, 2, 3, 4, 5), 5, 7, C64List(0, 1, 2, 3, 4))
}

func TestC64SliceTrim(t *testing.T) {
	ConfirmTrim := func(s *C64Slice, start, end int, r *C64Slice) {
		if s.Trim(start, end); !r.Equal(s) {
			t.Fatalf("Trim(%v, %v) should be %v but is %v", start, end, r, s)
		}
	}

	ConfirmTrim(C64List(0, 1, 2, 3, 4, 5), 0, 1, C64List(0))
	ConfirmTrim(C64List(0, 1, 2, 3, 4, 5), 1, 2, C64List(1))
	ConfirmTrim(C64List(0, 1, 2, 3, 4, 5), 2, 3, C64List(2))
	ConfirmTrim(C64List(0, 1, 2, 3, 4, 5), 3, 4, C64List(3))
	ConfirmTrim(C64List(0, 1, 2, 3, 4, 5), 4, 5, C64List(4))
	ConfirmTrim(C64List(0, 1, 2, 3, 4, 5), 5, 6, C64List(5))

	ConfirmTrim(C64List(0, 1, 2, 3, 4, 5), -1, 1, C64List(0))
	ConfirmTrim(C64List(0, 1, 2, 3, 4, 5), 0, 2, C64List(0, 1))
	ConfirmTrim(C64List(0, 1, 2, 3, 4, 5), 1, 3, C64List(1, 2))
	ConfirmTrim(C64List(0, 1, 2, 3, 4, 5), 2, 4, C64List(2, 3))
	ConfirmTrim(C64List(0, 1, 2, 3, 4, 5), 3, 5, C64List(3, 4))
	ConfirmTrim(C64List(0, 1, 2, 3, 4, 5), 4, 6, C64List(4, 5))
	ConfirmTrim(C64List(0, 1, 2, 3, 4, 5), 5, 7, C64List(5))
}

func TestC64SliceDelete(t *testing.T) {
	ConfirmCut := func(s *C64Slice, index int, r *C64Slice) {
		if s.Delete(index); !r.Equal(s) {
			t.Fatalf("Delete(%v) should be %v but is %v", index, r, s)
		}
	}

	ConfirmCut(C64List(0, 1, 2, 3, 4, 5), -1, C64List(0, 1, 2, 3, 4, 5))
	ConfirmCut(C64List(0, 1, 2, 3, 4, 5), 0, C64List(1, 2, 3, 4, 5))
	ConfirmCut(C64List(0, 1, 2, 3, 4, 5), 1, C64List(0, 2, 3, 4, 5))
	ConfirmCut(C64List(0, 1, 2, 3, 4, 5), 2, C64List(0, 1, 3, 4, 5))
	ConfirmCut(C64List(0, 1, 2, 3, 4, 5), 3, C64List(0, 1, 2, 4, 5))
	ConfirmCut(C64List(0, 1, 2, 3, 4, 5), 4, C64List(0, 1, 2, 3, 5))
	ConfirmCut(C64List(0, 1, 2, 3, 4, 5), 5, C64List(0, 1, 2, 3, 4))
	ConfirmCut(C64List(0, 1, 2, 3, 4, 5), 6, C64List(0, 1, 2, 3, 4, 5))
}

func TestC64SliceEach(t *testing.T) {
	var count	complex64
	C64List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).Each(func(i interface{}) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})
}

func TestC64SliceEachWithIndex(t *testing.T) {
	C64List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).EachWithIndex(func(index int, i interface{}) {
		if index != int(real(i.(complex64))) {
			t.Fatalf("element %v erroneously reported as %v", index, i)
		}
	})
}

func TestC64SliceEachWithKey(t *testing.T) {
	C64List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).EachWithKey(func(key, i interface{}) {
		if complex(float32(key.(int)), 0) != i {
			t.Fatalf("element %v erroneously reported as %v", key, i)
		}
	})
}

func TestC64SliceC64Each(t *testing.T) {
	var count	complex64
	C64List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).C64Each(func(i complex64) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})
}

func TestC64SliceC64EachWithIndex(t *testing.T) {
	C64List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).C64EachWithIndex(func(index int, i complex64) {
		if int(real(i)) != index {
			t.Fatalf("element %v erroneously reported as %v", index, i)
		}
	})
}

func TestC64SliceC64EachWithKey(t *testing.T) {
	c := C64List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9)
	c.C64EachWithKey(func(key interface{}, i complex64) {
		if key.(int) != int(real(i)) {
			t.Fatalf("element %v erroneously reported as %v", key, i)
		}
	})
}

func TestC64SliceBlockCopy(t *testing.T) {
	ConfirmBlockCopy := func(s *C64Slice, destination, source, count int, r *C64Slice) {
		s.BlockCopy(destination, source, count)
		if !r.Equal(s) {
			t.Fatalf("BlockCopy(%v, %v, %v) should be %v but is %v", destination, source, count, r, s)
		}
	}

	ConfirmBlockCopy(C64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, 0, 4, C64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockCopy(C64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 9, 9, 4, C64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockCopy(C64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 2, 4, C64List(0, 1, 2, 3, 4, 2, 3, 4, 5, 9))
	ConfirmBlockCopy(C64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 2, 5, 4, C64List(0, 1, 5, 6, 7, 8, 6, 7, 8, 9))
}

func TestC64SliceBlockClear(t *testing.T) {
	ConfirmBlockClear := func(s *C64Slice, start, count int, r *C64Slice) {
		s.BlockClear(start, count)
		if !r.Equal(s) {
			t.Fatalf("BlockClear(%v, %v) should be %v but is %v", start, count, r, s)
		}
	}

	ConfirmBlockClear(C64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, 4, C64List(0, 0, 0, 0, 4, 5, 6, 7, 8, 9))
	ConfirmBlockClear(C64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, 4, C64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockClear(C64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 4, C64List(0, 1, 2, 3, 4, 0, 0, 0, 0, 9))
}

func TestC64SliceOverwrite(t *testing.T) {
	ConfirmOverwrite := func(s *C64Slice, offset int, v, r *C64Slice) {
		s.Overwrite(offset, *v)
		if !r.Equal(s) {
			t.Fatalf("Overwrite(%v, %v) should be %v but is %v", offset, v, r, s)
		}
	}

	ConfirmOverwrite(C64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, C64List(10, 9, 8, 7), C64List(10, 9, 8, 7, 4, 5, 6, 7, 8, 9))
	ConfirmOverwrite(C64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, C64List(10, 9, 8, 7), C64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmOverwrite(C64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, C64List(11, 12, 13, 14), C64List(0, 1, 2, 3, 4, 11, 12, 13, 14, 9))
}

func TestC64SliceReallocate(t *testing.T) {
	ConfirmReallocate := func(s *C64Slice, l, c int, r *C64Slice) {
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

	ConfirmReallocate(C64List(), 0, 10, C64List())
	ConfirmReallocate(C64List(0, 1, 2, 3, 4), 3, 10, C64List(0, 1, 2))
	ConfirmReallocate(C64List(0, 1, 2, 3, 4), 5, 10, C64List(0, 1, 2, 3, 4))
	ConfirmReallocate(C64List(0, 1, 2, 3, 4), 10, 10, C64List(0, 1, 2, 3, 4, 0, 0, 0, 0, 0))
	ConfirmReallocate(C64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 1, 5, C64List(0))
	ConfirmReallocate(C64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 5, C64List(0, 1, 2, 3, 4))
	ConfirmReallocate(C64List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, 5, C64List(0, 1, 2, 3, 4))
}

func TestC64SliceExtend(t *testing.T) {
	ConfirmExtend := func(s *C64Slice, n int, r *C64Slice) {
		c := s.Cap()
		s.Extend(n)
		switch {
		case s.Len() != r.Len():	t.Fatalf("Extend(%v) len should be %v but is %v", n, r.Len(), s.Len())
		case s.Cap() != c + n:		t.Fatalf("Extend(%v) cap should be %v but is %v", n, c + n, s.Cap())
		case !r.Equal(s):			t.Fatalf("Extend(%v) should be %v but is %v", n, r, s)
		}
	}

	ConfirmExtend(C64List(), 1, C64List(0))
	ConfirmExtend(C64List(), 2, C64List(0, 0))
}

func TestC64SliceExpand(t *testing.T) {
	ConfirmExpand := func(s *C64Slice, i, n int, r *C64Slice) {
		c := s.Cap()
		s.Expand(i, n)
		switch {
		case s.Len() != r.Len():	t.Fatalf("Expand(%v, %v) len should be %v but is %v", i, n, r.Len(), s.Len())
		case s.Cap() != c + n:		t.Fatalf("Expand(%v, %v) cap should be %v but is %v", i, n, c + n, s.Cap())
		case !r.Equal(s):			t.Fatalf("Expand(%v, %v) should be %v but is %v", i, n, r, s)
		}
	}

	ConfirmExpand(C64List(), -1, 1, C64List(0))
	ConfirmExpand(C64List(), 0, 1, C64List(0))
	ConfirmExpand(C64List(), 1, 1, C64List(0))
	ConfirmExpand(C64List(), 0, 2, C64List(0, 0))

	ConfirmExpand(C64List(0, 1, 2), -1, 2, C64List(0, 0, 0, 1, 2))
	ConfirmExpand(C64List(0, 1, 2), 0, 2, C64List(0, 0, 0, 1, 2))
	ConfirmExpand(C64List(0, 1, 2), 1, 2, C64List(0, 0, 0, 1, 2))
	ConfirmExpand(C64List(0, 1, 2), 2, 2, C64List(0, 1, 0, 0, 2))
	ConfirmExpand(C64List(0, 1, 2), 3, 2, C64List(0, 1, 2, 0, 0))
	ConfirmExpand(C64List(0, 1, 2), 4, 2, C64List(0, 1, 2, 0, 0))
}

func TestC64SliceDepth(t *testing.T) {
	ConfirmDepth := func(s *C64Slice, i int) {
		if x := s.Depth(); x != i {
			t.Fatalf("%v.Depth() should be %v but is %v", s, i, x)
		}
	}
	ConfirmDepth(C64List(0, 1), 0)
}

func TestC64SliceReverse(t *testing.T) {
	sxp := C64List(1, 2, 3, 4, 5)
	rxp := C64List(5, 4, 3, 2, 1)
	sxp.Reverse()
	if !rxp.Equal(sxp) {
		t.Fatalf("Reversal failed: %v", sxp)
	}
}

func TestC64SliceAppend(t *testing.T) {
	ConfirmAppend := func(s *C64Slice, v interface{}, r *C64Slice) {
		s.Append(v)
		if !r.Equal(s) {
			t.Fatalf("Append(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmAppend(C64List(), complex64(0), C64List(0))
}

func TestC64SliceAppendSlice(t *testing.T) {
	ConfirmAppendSlice := func(s, v, r *C64Slice) {
		s.AppendSlice(*v)
		if !r.Equal(s) {
			t.Fatalf("AppendSlice(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmAppendSlice(C64List(), C64List(0), C64List(0))
	ConfirmAppendSlice(C64List(), C64List(0, 1), C64List(0, 1))
	ConfirmAppendSlice(C64List(0, 1, 2), C64List(3, 4), C64List(0, 1, 2, 3, 4))
}

func TestC64SlicePrepend(t *testing.T) {
	ConfirmPrepend := func(s *C64Slice, v interface{}, r *C64Slice) {
		if s.Prepend(v); !r.Equal(s) {
			t.Fatalf("Prepend(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmPrepend(C64List(), complex64(0), C64List(0))
	ConfirmPrepend(C64List(0), complex64(1), C64List(1, 0))
}

func TestC64SlicePrependSlice(t *testing.T) {
	ConfirmPrependSlice := func(s, v, r *C64Slice) {
		if s.PrependSlice(*v); !r.Equal(s) {
			t.Fatalf("PrependSlice(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmPrependSlice(C64List(), C64List(0), C64List(0))
	ConfirmPrependSlice(C64List(), C64List(0, 1), C64List(0, 1))
	ConfirmPrependSlice(C64List(0, 1, 2), C64List(3, 4), C64List(3, 4, 0, 1, 2))
}

func TestC64SliceRepeat(t *testing.T) {
	ConfirmRepeat := func(s *C64Slice, count int, r *C64Slice) {
		if x := s.Repeat(count); !x.Equal(r) {
			t.Fatalf("%v.Repeat(%v) should be %v but is %v", s, count, r, x)
		}
	}

	ConfirmRepeat(C64List(), 5, C64List())
	ConfirmRepeat(C64List(0), 1, C64List(0))
	ConfirmRepeat(C64List(0), 2, C64List(0, 0))
	ConfirmRepeat(C64List(0), 3, C64List(0, 0, 0))
	ConfirmRepeat(C64List(0), 4, C64List(0, 0, 0, 0))
	ConfirmRepeat(C64List(0), 5, C64List(0, 0, 0, 0, 0))
}

func TestC64SliceCar(t *testing.T) {
	ConfirmCar := func(s *C64Slice, r complex64) {
		n := s.Car().(complex64)
		if ok := n == r; !ok {
			t.Fatalf("head should be '%v' but is '%v'", r, n)
		}
	}
	ConfirmCar(C64List(1, 2, 3), 1)
}

func TestC64SliceCdr(t *testing.T) {
	ConfirmCdr := func(s, r *C64Slice) {
		if n := s.Cdr(); !n.Equal(r) {
			t.Fatalf("tail should be '%v' but is '%v'", r, n)
		}
	}
	ConfirmCdr(C64List(1, 2, 3), C64List(2, 3))
}

func TestC64SliceRplaca(t *testing.T) {
	ConfirmRplaca := func(s *C64Slice, v interface{}, r *C64Slice) {
		if s.Rplaca(v); !s.Equal(r) {
			t.Fatalf("slice should be '%v' but is '%v'", r, s)
		}
	}
	ConfirmRplaca(C64List(1, 2, 3, 4, 5), complex64(0), C64List(0, 2, 3, 4, 5))
}

func TestC64SliceRplacd(t *testing.T) {
	ConfirmRplacd := func(s *C64Slice, v interface{}, r *C64Slice) {
		if s.Rplacd(v); !s.Equal(r) {
			t.Fatalf("slice should be '%v' but is '%v'", r, s)
		}
	}
	ConfirmRplacd(C64List(1, 2, 3, 4, 5), nil, C64List(1))
	ConfirmRplacd(C64List(1, 2, 3, 4, 5), complex64(10), C64List(1, 10))
	ConfirmRplacd(C64List(1, 2, 3, 4, 5), C64List(5, 4, 3, 2), C64List(1, 5, 4, 3, 2))
	ConfirmRplacd(C64List(1, 2, 3, 4, 5, 6), C64List(2, 4, 8, 16), C64List(1, 2, 4, 8, 16))
}

func TestC64SliceSetIntersection(t *testing.T) {
	ConfirmSetIntersection := func(s, o, r *C64Slice) {
		if x := s.SetIntersection(*o); !r.Equal(x) {
			t.Fatalf("%v.SetIntersection(%v) should be %v but is %v", s, o, r, x)
		}
	}

	ConfirmSetIntersection(C64List(1, 2, 3), C64List(), C64List())
	ConfirmSetIntersection(C64List(1, 2, 3), C64List(1), C64List(1))
	ConfirmSetIntersection(C64List(1, 2, 3), C64List(1, 1), C64List(1))
	ConfirmSetIntersection(C64List(1, 2, 3), C64List(1, 2, 1), C64List(1, 2))
}

func TestC64SliceSetUnion(t *testing.T) {
	ConfirmSetUnion := func(s, o, r *C64Slice) {
		if x := s.SetUnion(*o); !r.Equal(x) {
			t.Fatalf("%v.SetUnion(%v) should be %v but is %v", s, o, r, x)
		}
	}

	ConfirmSetUnion(C64List(1, 2, 3), C64List(), C64List(1, 2, 3))
	ConfirmSetUnion(C64List(1, 2, 3), C64List(1), C64List(1, 2, 3))
	ConfirmSetUnion(C64List(1, 2, 3), C64List(1, 1), C64List(1, 2, 3))
	ConfirmSetUnion(C64List(1, 2, 3), C64List(1, 2, 1), C64List(1, 2, 3))
}