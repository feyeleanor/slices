package slices

import "testing"

func BenchmarkSList2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = SList(0, 1)
	}
}

func BenchmarkSList10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	}
}

func BenchmarkSList2x2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = SList(0, SList(0, 1))
	}
}

func BenchmarkSList2x10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = SList(0, SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	}
}

func BenchmarkSList10x2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = SList(SList(0, 1), SList(1, 2), SList(2, 3), SList(3, 4), SList(4, 5), SList(5, 6), SList(6, 7), SList(7, 8), SList(8, 9), SList(9, 0))
	}
}

func BenchmarkSList10x10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = SList(	SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)	)
	}
}

func BenchmarkSliceLen1(b *testing.B) {
	v := SList(0)
	for i := 0; i < b.N; i++ {
		_ = v.Len()
	}
}

func BenchmarkSliceLen1x1(b *testing.B) {
	v := SList(0, SList(0))
	for i := 0; i < b.N; i++ {
		_ = v.Len()
	}
}

func BenchmarkSliceLen1x10(b *testing.B) {
	v := SList(SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	for i := 0; i < b.N; i++ {
		_ = v.Len()
	}
}

func BenchmarkSliceLen10(b *testing.B) {
	v := SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	for i := 0; i < b.N; i++ {
		_ = v.Len()
	}
}

func BenchmarkSliceLen10x2(b *testing.B) {
	v := SList(SList(0, 1), SList(1, 2), SList(2, 3), SList(3, 4), SList(4, 5), SList(5, 6), SList(6, 7), SList(7, 8), SList(8, 9), SList(9, 0))
	for i := 0; i < b.N; i++ {
		_ = v.Len()
	}
}

func BenchmarkSliceLen10x10(b *testing.B) {
	v := SList(	SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)	)
	for i := 0; i < b.N; i++ {
		_ = v.Len()
	}
}

func BenchmarkSliceDepth1(b *testing.B) {
	v := SList(0)
	for i := 0; i < b.N; i++ {
		_ = v.Depth()
	}
}

func BenchmarkSliceDepth1x1(b *testing.B) {
	v := SList(0, SList(0))
	for i := 0; i < b.N; i++ {
		_ = v.Depth()
	}
}

func BenchmarkSliceDepth1x10(b *testing.B) {
	v := SList(SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	for i := 0; i < b.N; i++ {
		_ = v.Depth()
	}
}

func BenchmarkSliceDepth10(b *testing.B) {
	v := SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	for i := 0; i < b.N; i++ {
		_ = v.Depth()
	}
}

func BenchmarkSliceDepth10x2(b *testing.B) {
	v := SList(SList(0, 1), SList(1, 2), SList(2, 3), SList(3, 4), SList(4, 5), SList(5, 6), SList(6, 7), SList(7, 8), SList(8, 9), SList(9, 0))
	for i := 0; i < b.N; i++ {
		_ = v.Depth()
	}
}

func BenchmarkSliceDepth10x10(b *testing.B) {
	v := SList(	SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)	)
	for i := 0; i < b.N; i++ {
		_ = v.Depth()
	}
}

func BenchmarkSliceReverse10(b *testing.B) {
	b.StopTimer()
		v := SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		v.Reverse()
	}
}

func BenchmarkSliceReverse10x10(b *testing.B) {
	b.StopTimer()
		v := SList(	SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)	)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		v.Reverse()
	}
}
func BenchmarkSliceFlatten1(b *testing.B) {
	v := SList(0)
	for i := 0; i < b.N; i++ {
		v.Flatten()
	}
}

func BenchmarkSliceFlatten1x1(b *testing.B) {
	v := SList(0, SList(0))
	for i := 0; i < b.N; i++ {
		v.Flatten()
	}
}

func BenchmarkSliceFlatten1x10(b *testing.B) {
	v := SList(SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	for i := 0; i < b.N; i++ {
		v.Flatten()
	}
}

func BenchmarkSliceFlatten10(b *testing.B) {
	v := SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	for i := 0; i < b.N; i++ {
		v.Flatten()
	}
}

func BenchmarkSliceFlatten10x2(b *testing.B) {
	v := SList(SList(0, 1), SList(1, 2), SList(2, 3), SList(3, 4), SList(4, 5), SList(5, 6), SList(6, 7), SList(7, 8), SList(8, 9), SList(9, 0))
	for i := 0; i < b.N; i++ {
		v.Flatten()
	}
}

func BenchmarkSliceFlatten10x10(b *testing.B) {
	v := SList(	SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)	)
	for i := 0; i < b.N; i++ {
		v.Flatten()
	}
}

func BenchmarkSliceCar(b *testing.B) {
	v := SList(0, 1)
	for i := 0; i < b.N; i++ {
		_ = v.Car()
	}
}

func BenchmarkSliceCaar(b *testing.B) {
	v := SList(SList(0, 1), 2)
	for i := 0; i < b.N; i++ {
		_ = v.Caar()
	}
}

func BenchmarkSliceCdr(b *testing.B) {
	v := SList(0, 1)
	for i := 0; i < b.N; i++ {
		_ = v.Cdr()
	}
}

func BenchmarkSliceCddr(b *testing.B) {
	v := SList(0, SList(1, 2))
	for i := 0; i < b.N; i++ {
		_ = v.Cddr()
	}
}

func BenchmarkSliceRplaca(b *testing.B) {}

func BenchmarkSliceRplacd(b *testing.B) {}