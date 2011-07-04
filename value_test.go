package slices

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

func TestVSliceAt(t *testing.T) {}
func TestVSliceSet(t *testing.T) {}
func TestVSliceEach(t *testing.T) {}
func TestVSliceString(t *testing.T) {}
func TestVSliceLen(t *testing.T) {}
func TestVSliceCap(t *testing.T) {}
func TestVSlicenew(t *testing.T) {}
func TestVSliceBlockCopy(t *testing.T) {}

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

func TestVSliceAppend(t *testing.T) {
	ConfirmAppend := func(b, v interface{}, r *VSlice) {
		g := VWrap(b)
		g.Append(v)
		if g.Len() != r.Len() {
			t.Fatalf("Slice length should be %v but is %v", r.Len(), g.Len())
		}
		for i := 0; i < r.Len(); i++ {
			if g.At(i) != r.At(i) {
				t.Fatalf("Slice elements b[%v] and r[%v] should match but are %v and %v", i, i, g.At(i), r.At(i))
			}
		}
	}

	ConfirmAppend([]int{0, 1, 2}, 3, 						VWrap([]int{0, 1, 2, 3}))
}

func TestVSliceAppendSlice(t *testing.T) {
	ConfirmAppendSlice := func(b, v interface{}, r *VSlice) {
		g := VWrap(b)
		g.AppendSlice(VWrap(v))
		if g.Len() != r.Len() {
			t.Fatalf("Slice length should be %v but is %v", r.Len(), g.Len())
		}
		for i := 0; i < r.Len(); i++ {
			if g.At(i) != r.At(i) {
				t.Fatalf("Slice elements b[%v] and r[%v] should match but are %v and %v", i, i, g.At(i), r.At(i))
			}
		}
	}

	ConfirmAppendSlice([]int{0, 1, 2}, []int{3, 4, 5}, 			VWrap([]int{0, 1, 2, 3, 4, 5}))
	ConfirmAppendSlice([]int{0, 1, 2}, VWrap([]int{3, 4, 5}),	VWrap([]int{0, 1, 2, 3, 4, 5}))
	ConfirmAppendSlice([]int{0, 1, 2}, *VWrap([]int{3, 4, 5}),	VWrap([]int{0, 1, 2, 3, 4, 5}))
}

func TestVSlicePrepend(t *testing.T) {
	ConfirmPrepend := func(b, v interface{}, r *VSlice) {
		g := VWrap(b)
		g.Prepend(v)
		if g.Len() != r.Len() {
			t.Fatalf("Slice length should be %v but is %v", r.Len(), g.Len())
		}
		for i := 0; i < r.Len(); i++ {
			if g.At(i) != r.At(i) {
				t.Fatalf("Slice elements b[%v] and r[%v] should match but are %v and %v", i, i, g.At(i), r.At(i))
			}
		}
	}

	ConfirmPrepend([]int{3, 4, 5}, 2, VList(2, 3, 4, 5))
}

func TestVSlicePrependSlice(t *testing.T) {
	ConfirmPrependSlice := func(b, v interface{}, r *VSlice) {
		g := VWrap(b)
		g.PrependSlice(VWrap(v))
		if g.Len() != r.Len() {
			t.Fatalf("Slice length should be %v but is %v", r.Len(), g.Len())
		}
		for i := 0; i < r.Len(); i++ {
			if g.At(i) != r.At(i) {
				t.Fatalf("Slice elements b[%v] and r[%v] should match but are %v and %v", i, i, g.At(i), r.At(i))
			}
		}
	}

	ConfirmPrependSlice([]int{3, 4, 5}, []int{0, 1, 2},			VWrap([]int{0, 1, 2, 3, 4, 5}))
	ConfirmPrependSlice([]int{3, 4, 5}, VWrap([]int{0, 1, 2}),	VWrap([]int{0, 1, 2, 3, 4, 5}))
	ConfirmPrependSlice([]int{3, 4, 5}, *VWrap([]int{0, 1, 2}),	VWrap([]int{0, 1, 2, 3, 4, 5}))
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
	
}

func TestVSliceEqual(t *testing.T) {
	ConfirmEqual := func(s *VSlice, o interface{}) {
		if !s.Equal(o) {
			t.Fatalf("%v should equal %v", s, o)
		}
	}
	RefuteEqual := func(s *VSlice, o interface{}) {
		if s.Equal(o) {
			t.Fatalf("%v should not equal %v", s, o)
		}
	}

	ConfirmEqual(VWrap([]int{ 0 }), VWrap([]int{ 0 }))
	RefuteEqual(VWrap([]int{ 0 }), VWrap([]int{ 1 }))
}