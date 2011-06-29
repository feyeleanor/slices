package slices

import "testing"

func initSliceValueTest() (b []int, g *SliceValue) {
	b = []int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	g = VSlice(b)
	return
}

func TestSliceValueMakeSlice(t *testing.T) {
//	t.Fatal()
}

func TestSliceValueVSlice(t *testing.T) {
	g := VSlice([]int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 })
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

func TestSliceValueAt(t *testing.T) {
//	t.Fatal()
}

func TestSliceValueSet(t *testing.T) {
//	t.Fatal()
}

func TestSliceValueEach(t *testing.T) {
//	t.Fatal()
}

func TestSliceValueString(t *testing.T) {
//	t.Fatal()
}

func TestSliceValueLen(t *testing.T) {
//	t.Fatal()
}

func TestSliceValueCap(t *testing.T) {
//	t.Fatal()
}


func TestSliceValuenew(t *testing.T) {
//	t.Fatal()
}

func TestSliceValueBlockCopy(t *testing.T) {
//	t.Fatal()
/*	SHOULD_MATCH := "Slice elements g[%v] and c[%v] should match but are %v and %v"

	_, g := initSliceValueTest()
	c := Copy(g)
	c.BlockCopy(0, 5, 5)
	switch {
	case c.Len() != g.Len():	t.Fatalf("Slice length should be %v not %v", g.Len(), c.Len())
	case c.Cap() != g.Cap():	t.Fatalf("Slice capacity should be %v not %v", g.Cap(), c.Cap())
	case c.At(0) != g.At(5):	t.Fatalf(SHOULD_MATCH, 0, 0, g.At(5), c.At(0))
	case c.At(1) != g.At(6):	t.Fatalf(SHOULD_MATCH, 1, 1, g.At(6), c.At(1))
	case c.At(2) != g.At(7):	t.Fatalf(SHOULD_MATCH, 2, 2, g.At(7), c.At(2))
	case c.At(3) != g.At(8):	t.Fatalf(SHOULD_MATCH, 3, 3, g.At(8), c.At(3))
	case c.At(4) != g.At(9):	t.Fatalf(SHOULD_MATCH, 4, 4, g.At(9), c.At(4))
	case c.At(5) != g.At(5):	t.Fatalf(SHOULD_MATCH, 5, 5, g.At(5), c.At(5))
	case c.At(6) != g.At(6):	t.Fatalf(SHOULD_MATCH, 6, 6, g.At(6), c.At(6))
	case c.At(7) != g.At(7):	t.Fatalf(SHOULD_MATCH, 7, 7, g.At(7), c.At(7))
	case c.At(8) != g.At(8):	t.Fatalf(SHOULD_MATCH, 8, 8, g.At(8), c.At(8))
	case c.At(9) != g.At(9):	t.Fatalf(SHOULD_MATCH, 9, 9, g.At(9), c.At(9))
	}
*/}

func TestSliceValueOverwrite(t *testing.T) {
	g := VSlice([]int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 })
	c := VSlice(make([]int, g.Len(), g.Cap()))
	c.Overwrite(0, g)
	for i := 0; i < g.Len(); i++ {
		if c.At(i) != g.At(i) {
			t.Fatalf("Slice elements g[%v] and c[%v] should match but are %v and %v", i, i, g.At(0), c.At(0))
		}
	}
}

func TestSliceValueReallocate(t *testing.T) {
	b, g := initSliceValueTest()
	switch g.Reallocate(10, 20); {
	case b == nil:				t.Fatal("Reallocate() created a nil value for original slice")
	case g == nil:				t.Fatal("Reallocate() created a nil value for Slice")
	case cap(b) != 10:			t.Fatalf("original slice capacity should be 10 but is %v", cap(b))
	case len(b) != 10:			t.Fatalf("original slice length should be 10 but is %v", len(b))
	case g.Cap() != 20:			t.Fatalf("Slice capacity should be 20 but is %v", g.Cap())
	case g.Len() != 10:			t.Fatalf("Slice length should be 10 but is %v", g.Len())
	}

	switch g.Reallocate(10, 5); {
	case b == nil:				t.Fatal("Reallocate() created a nil value for original slice")
	case g == nil:				t.Fatal("Reallocate() created a nil value for Slice")
	case cap(b) != 10:			t.Fatalf("original slice capacity should be 10 but is %v", cap(b))
	case len(b) != 10:			t.Fatalf("original slice length should be 10 but is %v", len(b))
	case g.Cap() != 5:			t.Fatalf("Slice capacity should be 5 but is %v", g.Cap())
	case g.Len() != 5:			t.Fatalf("Slice length should be 5 but is %v", g.Len())
	}
}

func TestSliceValueAppend(t *testing.T) {
	ConfirmAppend := func(b, v interface{}, r *SliceValue) {
		g := VSlice(b)
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

	ConfirmAppend([]int{0, 1, 2}, 3, 							VSlice([]int{0, 1, 2, 3}))
	ConfirmAppend([]int{0, 1, 2}, []int{3, 4, 5}, 				VSlice([]int{0, 1, 2, 3, 4, 5}))
	ConfirmAppend([]int{0, 1, 2}, VSlice([]int{3, 4, 5}),		VSlice([]int{0, 1, 2, 3, 4, 5}))
	ConfirmAppend([]int{0, 1, 2}, *VSlice([]int{3, 4, 5}),	VSlice([]int{0, 1, 2, 3, 4, 5}))
}

func TestSliceValuePrepend(t *testing.T) {
	ConfirmPrepend := func(b, v interface{}, r *SliceValue) {
		g := VSlice(b)
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

	ConfirmPrepend([]int{3, 4, 5}, 2, 							VSlice([]int{2, 3, 4, 5}))
	ConfirmPrepend([]int{3, 4, 5}, []int{0, 1, 2},				VSlice([]int{0, 1, 2, 3, 4, 5}))
	ConfirmPrepend([]int{3, 4, 5}, VSlice([]int{0, 1, 2}),	VSlice([]int{0, 1, 2, 3, 4, 5}))
	ConfirmPrepend([]int{3, 4, 5}, *VSlice([]int{0, 1, 2}),	VSlice([]int{0, 1, 2, 3, 4, 5}))
}

func TestSliceValueRepeat(t *testing.T) {
	SHOULD_MATCH := "Slice elements g[%v] and g[%v] should match but are %v and %v"

	b, g := initSliceValueTest()
	c := 3
	g = g.Repeat(c)
	switch {
	case g.Len() != len(b) * c:	t.Fatalf("Slice length should be %v not %v", len(b) * c, g.Len())
	case g.Cap() != cap(b) * c:	t.Fatalf("Slice capacity should be %v not %v", cap(b) * c, g.Cap())
	case g.At(0) != g.At(10):	t.Fatalf(SHOULD_MATCH, 0, 10, g.At(0), g.At(10))
	case g.At(1) != g.At(11):	t.Fatalf(SHOULD_MATCH, 1, 11, g.At(1), g.At(11))
	case g.At(9) != g.At(19):	t.Fatalf(SHOULD_MATCH, 9, 19, g.At(9), g.At(19))
	case g.At(0) != g.At(20):	t.Fatalf(SHOULD_MATCH, 0, 20, g.At(0), g.At(20))
	case g.At(1) != g.At(21):	t.Fatalf(SHOULD_MATCH, 1, 21, g.At(1), g.At(21))
	case g.At(9) != g.At(29):	t.Fatalf(SHOULD_MATCH, 9, 19, g.At(9), g.At(29))
	}
}

func TestSliceValueFlatten(t *testing.T) {
	
}

func TestSliceValueEqual(t *testing.T) {
	ConfirmEqual := func(s *SliceValue, o interface{}) {
		if !s.Equal(o) {
			t.Fatalf("%v should equal %v", s, o)
		}
	}
	RefuteEqual := func(s *SliceValue, o interface{}) {
		if s.Equal(o) {
			t.Fatalf("%v should not equal %v", s, o)
		}
	}

	ConfirmEqual(VSlice([]int{ 0 }), VSlice([]int{ 0 }))
	RefuteEqual(VSlice([]int{ 0 }), VSlice([]int{ 1 }))
}

//	func TestSliceValueFeed(t *testing.T) { t.Fatal() }
//	func TestSliceValuePipe(t *testing.T) { t.Fatal() }