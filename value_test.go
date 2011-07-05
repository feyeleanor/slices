package slices

import "github.com/feyeleanor/lists"
import "reflect"
import "testing"

func TestVSliceMakeSlice(t *testing.T) {}

func TestVSliceVSlice(t *testing.T) {
	g := VWrap([]int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 })
	if g == nil {
		t.Fatal("Make slice returned a nil value")
	}

	if g.Len() != 10 {
		t.Fatalf("Slice length should be %v not %v", 10, g.Len())
	}

	for i := 0; i < g.Len(); i++ {
		if g.At(i) != i {
			t.Fatalf("g[%v] should contain %v but contains %v", 0, g.At(0))
		}
	}
}

func TestVSliceString(t *testing.T) {
	ConfirmString := func(s *VSlice, r string) {
		if x := s.String(); x != r {
			t.Fatalf("%v erroneously serialised as '%v'", r, x)
		}
	}

	ConfirmString(VList(), "()")
	ConfirmString(VList(0), "(0)")
	ConfirmString(VList(0, 1), "(0 1)")
}

func TestVSliceLen(t *testing.T) {
	ConfirmLength := func(s *VSlice, i int) {
		if x := s.Len(); x != i {
			t.Fatalf("%v.Len() should be %v but is %v", s, i, x)
		}
	}
	
	ConfirmLength(VList(0), 1)
	ConfirmLength(VList(0, 1), 2)
}

func TestVSliceClear(t *testing.T) {
	ConfirmClear := func(s *VSlice, i int, r *VSlice) {
		if s.Clear(i); !r.Equal(s) {
			t.Fatalf("Clear(%v) should be %v but is %v", i, r, s)
		}
	}

	ConfirmClear(VList(0, 1, 2, 3, 4), 0, VList(nil, 1, 2, 3, 4))
	ConfirmClear(VList(0, 1, 2, 3, 4), 1, VList(0, nil, 2, 3, 4))
	ConfirmClear(VList(0, 1, 2, 3, 4), 2, VList(0, 1, nil, 3, 4))
	ConfirmClear(VList(0, 1, 2, 3, 4), 3, VList(0, 1, 2, nil, 4))
	ConfirmClear(VList(0, 1, 2, 3, 4), 4, VList(0, 1, 2, 3, nil))

	ConfirmClear(VWrap([]int{0, 1, 2, 3, 4}), 0, VList(0, 1, 2, 3, 4))
	ConfirmClear(VWrap([]int{0, 1, 2, 3, 4}), 1, VList(0, 0, 2, 3, 4))
	ConfirmClear(VWrap([]int{0, 1, 2, 3, 4}), 2, VList(0, 1, 0, 3, 4))
	ConfirmClear(VWrap([]int{0, 1, 2, 3, 4}), 3, VList(0, 1, 2, 0, 4))
	ConfirmClear(VWrap([]int{0, 1, 2, 3, 4}), 4, VList(0, 1, 2, 3, 0))
}

func TestVSliceSwap(t *testing.T) {
	ConfirmSwap := func(s *VSlice, i, j int, r *VSlice) {
		if s.Swap(i, j); !r.Equal(s) {
			t.Fatalf("Swap(%v, %v) should be %v but is %v", i, j, r, s)
		}
	}
	ConfirmSwap(VList(0, 1, 2), 0, 1, VList(1, 0, 2))
	ConfirmSwap(VList(0, 1, 2), 0, 2, VList(2, 1, 0))
}

func TestVSliceCut(t *testing.T) {
	ConfirmCut := func(s *VSlice, start, end int, r *VSlice) {
		if s.Cut(start, end); !r.Equal(s) {
			t.Fatalf("Cut(%v, %v) should be %v but is %v", start, end, r, s)
		}
	}

	ConfirmCut(VList(0, 1, 2, 3, 4, 5), 0, 1, VList(1, 2, 3, 4, 5))
	ConfirmCut(VList(0, 1, 2, 3, 4, 5), 1, 2, VList(0, 2, 3, 4, 5))
	ConfirmCut(VList(0, 1, 2, 3, 4, 5), 2, 3, VList(0, 1, 3, 4, 5))
	ConfirmCut(VList(0, 1, 2, 3, 4, 5), 3, 4, VList(0, 1, 2, 4, 5))
	ConfirmCut(VList(0, 1, 2, 3, 4, 5), 4, 5, VList(0, 1, 2, 3, 5))
	ConfirmCut(VList(0, 1, 2, 3, 4, 5), 5, 6, VList(0, 1, 2, 3, 4))

	ConfirmCut(VList(0, 1, 2, 3, 4, 5), -1, 1, VList(1, 2, 3, 4, 5))
	ConfirmCut(VList(0, 1, 2, 3, 4, 5), 0, 2, VList(2, 3, 4, 5))
	ConfirmCut(VList(0, 1, 2, 3, 4, 5), 1, 3, VList(0, 3, 4, 5))
	ConfirmCut(VList(0, 1, 2, 3, 4, 5), 2, 4, VList(0, 1, 4, 5))
	ConfirmCut(VList(0, 1, 2, 3, 4, 5), 3, 5, VList(0, 1, 2, 5))
	ConfirmCut(VList(0, 1, 2, 3, 4, 5), 4, 6, VList(0, 1, 2, 3))
	ConfirmCut(VList(0, 1, 2, 3, 4, 5), 5, 7, VList(0, 1, 2, 3, 4))
}

func TestVSliceTrim(t *testing.T) {
	ConfirmTrim := func(s *VSlice, start, end int, r *VSlice) {
		if s.Trim(start, end); !r.Equal(s) {
			t.Fatalf("Trim(%v, %v) should be %v but is %v", start, end, r, s)
		}
	}

	ConfirmTrim(VList(0, 1, 2, 3, 4, 5), 0, 1, VList(0))
	ConfirmTrim(VList(0, 1, 2, 3, 4, 5), 1, 2, VList(1))
	ConfirmTrim(VList(0, 1, 2, 3, 4, 5), 2, 3, VList(2))
	ConfirmTrim(VList(0, 1, 2, 3, 4, 5), 3, 4, VList(3))
	ConfirmTrim(VList(0, 1, 2, 3, 4, 5), 4, 5, VList(4))
	ConfirmTrim(VList(0, 1, 2, 3, 4, 5), 5, 6, VList(5))

	ConfirmTrim(VList(0, 1, 2, 3, 4, 5), -1, 1, VList(0))
	ConfirmTrim(VList(0, 1, 2, 3, 4, 5), 0, 2, VList(0, 1))
	ConfirmTrim(VList(0, 1, 2, 3, 4, 5), 1, 3, VList(1, 2))
	ConfirmTrim(VList(0, 1, 2, 3, 4, 5), 2, 4, VList(2, 3))
	ConfirmTrim(VList(0, 1, 2, 3, 4, 5), 3, 5, VList(3, 4))
	ConfirmTrim(VList(0, 1, 2, 3, 4, 5), 4, 6, VList(4, 5))
	ConfirmTrim(VList(0, 1, 2, 3, 4, 5), 5, 7, VList(5))
}

func TestVSliceDelete(t *testing.T) {
	ConfirmCut := func(s *VSlice, index int, r *VSlice) {
		if s.Delete(index); !r.Equal(s) {
			t.Fatalf("Delete(%v) should be %v but is %v", index, r, s)
		}
	}

	ConfirmCut(VList(0, 1, 2, 3, 4, 5), -1, VList(0, 1, 2, 3, 4, 5))
	ConfirmCut(VList(0, 1, 2, 3, 4, 5), 0, VList(1, 2, 3, 4, 5))
	ConfirmCut(VList(0, 1, 2, 3, 4, 5), 1, VList(0, 2, 3, 4, 5))
	ConfirmCut(VList(0, 1, 2, 3, 4, 5), 2, VList(0, 1, 3, 4, 5))
	ConfirmCut(VList(0, 1, 2, 3, 4, 5), 3, VList(0, 1, 2, 4, 5))
	ConfirmCut(VList(0, 1, 2, 3, 4, 5), 4, VList(0, 1, 2, 3, 5))
	ConfirmCut(VList(0, 1, 2, 3, 4, 5), 5, VList(0, 1, 2, 3, 4))
	ConfirmCut(VList(0, 1, 2, 3, 4, 5), 6, VList(0, 1, 2, 3, 4, 5))
}

func TestVSliceEach(t *testing.T) {
	c := VList(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9)
	count := 0
	c.Each(func(i interface{}) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})
}

func TestVSliceEachWithIndex(t *testing.T) {
	c := VList(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9)
	c.EachWithIndex(func(index int, i interface{}) {
		if i != index {
			t.Fatalf("element %v erroneously reported as %v", index, i)
		}
	})
}

func TestVSliceEachWithKey(t *testing.T) {
	c := VList(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9)
	c.EachWithKey(func(key, i interface{}) {
		if i != key {
			t.Fatalf("element %v erroneously reported as %v", key, i)
		}
	})
}

func TestVSliceVEach(t *testing.T) {
	var count	int
	VList(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).VEach(func(i reflect.Value) {
		if i.Interface() != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})
}

func TestVSliceVEachWithIndex(t *testing.T) {
	VList(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9).VEachWithIndex(func(index int, i reflect.Value) {
		if i.Interface() != index {
			t.Fatalf("element %v erroneously reported as %v", index, i)
		}
	})
}

func TestVSliceVEachWithKey(t *testing.T) {
	c := VList(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9)
	c.VEachWithKey(func(key interface{}, i reflect.Value) {
		if i.Interface() != key {
			t.Fatalf("element %v erroneously reported as %v", key, i)
		}
	})
}

func TestVSliceBlockCopy(t *testing.T) {
	ConfirmBlockCopy := func(s *VSlice, destination, source, count int, r *VSlice) {
		s.BlockCopy(destination, source, count)
		if !r.Equal(s) {
			t.Fatalf("BlockCopy(%v, %v, %v) should be %v but is %v", destination, source, count, r, s)
		}
	}

	ConfirmBlockCopy(VList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, 0, 4, VList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockCopy(VList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 9, 9, 4, VList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockCopy(VList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 2, 4, VList(0, 1, 2, 3, 4, 2, 3, 4, 5, 9))
	ConfirmBlockCopy(VList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 2, 5, 4, VList(0, 1, 5, 6, 7, 8, 6, 7, 8, 9))
}

func TestVSliceBlockClear(t *testing.T) {
	ConfirmBlockClear := func(s *VSlice, start, count int, r *VSlice) {
		s.BlockClear(start, count)
		if !r.Equal(s) {
			t.Fatalf("BlockClear(%v, %v) should be %v but is %v", start, count, r, s)
		}
	}

	ConfirmBlockClear(VList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, 4, VList(nil, nil, nil, nil, 4, 5, 6, 7, 8, 9))
	ConfirmBlockClear(VList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, 4, VList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockClear(VList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 4, VList(0, 1, 2, 3, 4, nil, nil, nil, nil, 9))

	ConfirmBlockClear(VWrap([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}), 0, 4, VList(0, 0, 0, 0, 4, 5, 6, 7, 8, 9))
	ConfirmBlockClear(VWrap([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}), 10, 4, VList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockClear(VWrap([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}), 5, 4, VList(0, 1, 2, 3, 4, 0, 0, 0, 0, 9))
}

func TestVSliceOverwrite(t *testing.T) {
	g := VWrap([]int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 })
	c := VWrap(make([]int, g.Len(), g.Cap()))
	c.Overwrite(0, g)
	for i := 0; i < g.Len(); i++ {
		if c.At(i) != g.At(i) {
			t.Fatalf("Slice elements g[%v] and c[%v] should match but are %v and %v", i, i, g.At(0), c.At(0))
		}
	}
}

func TestVSliceReallocate(t *testing.T) {
	ConfirmReallocate := func(s *VSlice, l, c int, r *VSlice) {
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

	ConfirmReallocate(VList(), 0, 10, VList())
	ConfirmReallocate(VList(0, 1, 2, 3, 4), 3, 10, VList(0, 1, 2))
	ConfirmReallocate(VList(0, 1, 2, 3, 4), 5, 10, VList(0, 1, 2, 3, 4))
	ConfirmReallocate(VList(0, 1, 2, 3, 4), 10, 10, VList(0, 1, 2, 3, 4, nil, nil, nil, nil, nil))
	ConfirmReallocate(VList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 1, 5, VList(0))
	ConfirmReallocate(VList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 5, VList(0, 1, 2, 3, 4))
	ConfirmReallocate(VList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, 5, VList(0, 1, 2, 3, 4))
}

func TestVSliceExtend(t *testing.T) {
	ConfirmExtend := func(s *VSlice, n int, r *VSlice) {
		c := s.Cap()
		s.Extend(n)
		switch {
		case s.Len() != r.Len():	t.Fatalf("Extend(%v) len should be %v but is %v", n, r.Len(), s.Len())
		case s.Cap() != c + n:		t.Fatalf("Extend(%v) cap should be %v but is %v", n, c + n, s.Cap())
		case !r.Equal(s):			t.Fatalf("Extend(%v) should be %v but is %v", n, r, s)
		}
	}

	ConfirmExtend(VList(), 1, VList(nil))
	ConfirmExtend(VList(), 2, VList(nil, nil))
}

func TestVSliceExpand(t *testing.T) {
	ConfirmExpand := func(s *VSlice, i, n int, r *VSlice) {
		c := s.Cap()
		s.Expand(i, n)
		switch {
		case s.Len() != r.Len():	t.Fatalf("Expand(%v, %v) len should be %v but is %v", i, n, r.Len(), s.Len())
		case s.Cap() != c + n:		t.Fatalf("Expand(%v, %v) cap should be %v but is %v", i, n, c + n, s.Cap())
		case !r.Equal(s):			t.Fatalf("Expand(%v, %v) should be %v but is %v", i, n, r, s)
		}
	}

	ConfirmExpand(VList(), -1, 1, VList(nil))
	ConfirmExpand(VList(), 0, 1, VList(nil))
	ConfirmExpand(VList(), 1, 1, VList(nil))
	ConfirmExpand(VList(), 0, 2, VList(nil, nil))

	ConfirmExpand(VList(0, 1, 2), -1, 2, VList(nil, nil, 0, 1, 2))
	ConfirmExpand(VList(0, 1, 2), 0, 2, VList(nil, nil, 0, 1, 2))
	ConfirmExpand(VList(0, 1, 2), 1, 2, VList(0, nil, nil, 1, 2))
	ConfirmExpand(VList(0, 1, 2), 2, 2, VList(0, 1, nil, nil, 2))
	ConfirmExpand(VList(0, 1, 2), 3, 2, VList(0, 1, 2, nil, nil))
	ConfirmExpand(VList(0, 1, 2), 4, 2, VList(0, 1, 2, nil, nil))

	ConfirmExpand(VWrap([]int{0, 1, 2}), -1, 2, VList(0, 0, 0, 1, 2))
	ConfirmExpand(VWrap([]int{0, 1, 2}), 0, 2, VList(0, 0, 0, 1, 2))
	ConfirmExpand(VWrap([]int{0, 1, 2}), 1, 2, VList(0, 0, 0, 1, 2))
	ConfirmExpand(VWrap([]int{0, 1, 2}), 2, 2, VList(0, 1, 0, 0, 2))
	ConfirmExpand(VWrap([]int{0, 1, 2}), 3, 2, VList(0, 1, 2, 0, 0))
	ConfirmExpand(VWrap([]int{0, 1, 2}), 4, 2, VList(0, 1, 2, 0, 0))
}

func TestVSliceDepth(t *testing.T) {
	ConfirmDepth := func(s *VSlice, i int) {
		if x := s.Depth(); x != i {
			t.Fatalf("%v.Depth() should be %v but is %v", s, i, x)
		}
	}
	ConfirmDepth(VList(0, 1), 0)
	ConfirmDepth(VList(0, 1), 0)
	ConfirmDepth(VList(VList(0, 1), 2), 1)
	ConfirmDepth(VList(0, VList(1, 2)), 1)
	ConfirmDepth(VList(0, 1, VList(2, VList(3, 4, 5))), 2)

	sxp := VList(0, 1,
				VList(2, VList(3, 4, 5)),
				VList(6, VList(7, VList(8, VList(9, 0)))),
				VList(2, VList(3, 4, 5)))
	ConfirmDepth(sxp, 4)

	rxp := VList(0, sxp, sxp)
	ConfirmDepth(rxp, 5)
	ConfirmDepth(VList(rxp, sxp), 6)

	ConfirmDepth(VList(0, 1), 0)
	ConfirmDepth(VList(List(0, 1), 2), 1)
	ConfirmDepth(VList(0, List(1, 2)), 1)
	ConfirmDepth(VList(0, 1, List(2, List(3, 4, 5))), 2)

	sxp = VList(0, 1,
				List(2, List(3, 4, 5)),
				List(6, List(7, List(8, List(9, 0)))),
				List(2, List(3, 4, 5)))
	ConfirmDepth(sxp, 4)

	rxp = VList(0, sxp, sxp)
	ConfirmDepth(rxp, 5)
	ConfirmDepth(VList(rxp, sxp), 6)
}

func TestVSliceReverse(t *testing.T) {
	ConfirmReverse := func(s, r *VSlice) {
		if s.Reverse(); !r.Equal(s) {
			t.Fatalf("Reverse() should be %v but is %v", r, s)
		}
	}

	ConfirmReverse(VList(1, 2, 3, 4, 5), VList(5, 4, 3, 2, 1))
	ConfirmReverse(VList(1, 3, 2, 5, 4), VList(4, 5, 2, 3, 1))
}

func TestVSliceAppend(t *testing.T) {
	ConfirmAppend := func(s *VSlice, v interface{}, r *VSlice) {
		if s.Append(v); !r.Equal(s) {
			t.Fatalf("Append(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmAppend(VList(0, 1, 2), 3, VList(0, 1, 2, 3))
	ConfirmAppend(VWrap([]int{0, 1, 2}), 3, VList(0, 1, 2, 3))
	ConfirmAppend(VList(0, 1, 2), 3, VWrap([]int{0, 1, 2, 3}))
}

func TestVSliceAppendSlice(t *testing.T) {
	ConfirmAppendSlice := func(s *VSlice, v interface{}, r *VSlice) {
		if s.AppendSlice(VWrap(v)); !r.Equal(s) {
			t.Fatalf("AppendSlice(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmAppendSlice(VList(0, 1, 2), VList(3, 4, 5), VList(0, 1, 2, 3, 4, 5))
	ConfirmAppendSlice(VWrap([]int{0, 1, 2}), []int{3, 4, 5}, 			VList(0, 1, 2, 3, 4, 5))
	ConfirmAppendSlice(VWrap([]int{0, 1, 2}), VWrap([]int{3, 4, 5}),	VList(0, 1, 2, 3, 4, 5))
	ConfirmAppendSlice(VWrap([]int{0, 1, 2}), *VWrap([]int{3, 4, 5}),	VList(0, 1, 2, 3, 4, 5))
}

func TestVSlicePrepend(t *testing.T) {
	ConfirmPrepend := func(s *VSlice, v interface{}, r *VSlice) {
		if s.Prepend(v); !r.Equal(s) {
			t.Fatalf("Prepend(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmPrepend(VList(0, 1, 2), 3, VList(3, 0, 1, 2))
	ConfirmPrepend(VWrap([]int{0, 1, 2}), 3, VList(3, 0, 1, 2))
	ConfirmPrepend(VList(0, 1, 2), 3, VWrap([]int{3, 0, 1, 2}))
}

func TestVSlicePrependSlice(t *testing.T) {
	ConfirmPrependSlice := func(s *VSlice, v interface{}, r *VSlice) {
		if s.PrependSlice(VWrap(v)); !r.Equal(s) {
			t.Fatalf("AppendSlice(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmPrependSlice(VList(0, 1, 2), VList(3, 4, 5),			VList(3, 4, 5, 0, 1, 2))
	ConfirmPrependSlice(VWrap([]int{0, 1, 2}), []int{3, 4, 5}, 			VList(3, 4, 5, 0, 1, 2))
	ConfirmPrependSlice(VWrap([]int{0, 1, 2}), VWrap([]int{3, 4, 5}),	VList(3, 4, 5, 0, 1, 2))
	ConfirmPrependSlice(VWrap([]int{0, 1, 2}), *VWrap([]int{3, 4, 5}),	VList(3, 4, 5, 0, 1, 2))
}

func TestVSliceRepeat(t *testing.T) {
	ConfirmRepeat := func(s *VSlice, count int, r *VSlice) {
		if x := s.Repeat(count); !x.Equal(r) {
			t.Fatalf("%v.Repeat(%v) should be %v but is %v", s, count, r, x)
		}
	}

	ConfirmRepeat(VList(), 5, VList())
	ConfirmRepeat(VList(0), 1,VList(0))
	ConfirmRepeat(VList(0), 2, VList(0, 0))
	ConfirmRepeat(VList(0), 3, VList(0, 0, 0))
	ConfirmRepeat(VList(0), 4, VList(0, 0, 0, 0))
	ConfirmRepeat(VList(0), 5, VList(0, 0, 0, 0, 0))
}

func TestVSliceFlatten(t *testing.T) {
	ConfirmFlatten := func(s, r *VSlice) {
		o := s.String()
		if s.Flatten(); !r.Equal(s) {
			t.Fatalf("Flatten(%v) should be %v but is %v", o, r, s)
		}
	}
	ConfirmFlatten(VList(), VList())
	ConfirmFlatten(VList(1), VList(1))
	ConfirmFlatten(VList(1, VList(2)), VList(1, 2))
	ConfirmFlatten(VList(1, VList(2, VList(3))), VList(1, 2, 3))
	ConfirmFlatten(VList(1, 2, VList(3, VList(4, 5), VList(6, VList(7, 8, 9), VList(10, 11)))), VList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11))

	ConfirmFlatten(VList(0, lists.List(1, 2, VList(3, 4))), VList(0, lists.List(1, 2, VList(3, 4))))
	ConfirmFlatten(VList(0, lists.List(1, 2, lists.List(3, 4))), VList(0, lists.List(1, 2, 3, 4)))

	ConfirmFlatten(VList(0, lists.Loop(1, 2)), VList(0, lists.Loop(1, 2)))
	ConfirmFlatten(VList(0, lists.List(1, lists.Loop(2, 3))), VList(0, lists.List(1, 2, 3)))

	ConfirmFlatten(VList(0, lists.List(1, 2, lists.Loop(3, 4))), VList(0, lists.List(1, 2, 3, 4)))
	ConfirmFlatten(VList(3, 4, VList(5, 6, 7)), VList(3, 4, 5, 6, 7))
	ConfirmFlatten(VList(0, lists.Loop(1, 2, VList(3, 4, VList(5, 6, 7)))), VList(0, lists.Loop(1, 2, VList(3, 4, 5, 6, 7))))

	sxp := VList(1, 2, VList(3, VList(4, 5), VList(6, VList(7, 8, 9), VList(10, 11))))
	rxp := VList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	ConfirmFlatten(sxp, rxp)

	rxp = VList(1, 2, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 3, 4, 5, 6, 7, 8, 9, 10, 11, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	sxp = VList(1, 2, sxp, VList(3, VList(4, 5), VList(6, VList(7, 8, 9), VList(10, 11), sxp)))
	ConfirmFlatten(sxp, rxp)
}

func TestVSliceEqual(t *testing.T) {
	ConfirmEqual := func(s *VSlice, o interface{}) {
		if !s.Equal(o) {
			t.Fatalf("%v.Equal(%v) should be equal", s, o)
		}
	}
	RefuteEqual := func(s *VSlice, o interface{}) {
		if s.Equal(o) {
			t.Fatalf("%v.Equal(%v) should not be equal", s, o)
		}
	}

	ConfirmEqual(VList(0), VWrap([]int{ 0 }))
	RefuteEqual(VList(0), VWrap([]uint{ 0 }))
	RefuteEqual(VList(0), VWrap([]int{ 1 }))
}


func TestVSliceCar(t *testing.T) {
	ConfirmCar := func(s *VSlice, r interface{}) {
		var ok bool
		n := s.Car()
		switch n := n.(type) {
		case Equatable:		ok = n.Equal(r)
		default:			ok = n == r
		}
		if !ok {
			t.Fatalf("%s.Car() should be %v but is %v", s, r, n)
		}
	}
	ConfirmCar(VList(1, 2, 3), 1)
	ConfirmCar(VList(VList(10, 20), 2, 3), VList(10, 20))
}

func TestVSliceCdr(t *testing.T) {
	ConfirmCdr := func(s, r *VSlice) {
		if n := s.Cdr(); !r.Equal(n) {
			t.Fatalf("%v.Cdr() should be %v but is %v", s, r, n)
		}
	}
//	ConfirmCdr(VList(), VList())
	ConfirmCdr(VList(1), VList())
	ConfirmCdr(VList(1, 2, 3), VList(2, 3))
}

func TestVSliceRplaca(t *testing.T) {
	ConfirmRplaca := func(s *VSlice, v interface{}, r *VSlice) {
		if s.Rplaca(v); !s.Equal(r) {
			t.Fatalf("Rplaca() should be %v but is %v", r, s)
		}
	}
	ConfirmRplaca(VList(1, 2, 3, 4, 5), 0, VList(0, 2, 3, 4, 5))
	ConfirmRplaca(VList(1, 2, 3, 4, 5), VList(1, 2, 3), VList(VList(1, 2, 3), 2, 3, 4, 5))
}

func TestVSliceRplacd(t *testing.T) {
	ConfirmRplacd := func(s *VSlice, v interface{}, r *VSlice) {
		if s.Rplacd(v); !s.Equal(r) {
			t.Fatalf("Rplacd() should be %v but is %v", r, s)
		}
	}
	ConfirmRplacd(VList(1, 2, 3, 4, 5), nil, VList(1))
	ConfirmRplacd(VList(1, 2, 3, 4, 5), 10, VList(1, 10))
	ConfirmRplacd(VList(1, 2, 3, 4, 5), VList(5, 4, 3, 2), VList(1, 5, 4, 3, 2))
	ConfirmRplacd(VList(1, 2, 3, 4, 5, 6), VList(2, 4, 8, 16), VList(1, 2, 4, 8, 16))
}