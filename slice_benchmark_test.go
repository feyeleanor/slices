package slices

import "testing"

func BenchmarkList2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = List(0, 1)
	}
}

func BenchmarkList10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	}
}

func BenchmarkList2x2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = List(0, List(0, 1))
	}
}

func BenchmarkList2x10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = List(0, List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	}
}

func BenchmarkList10x2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = List(List(0, 1), List(1, 2), List(2, 3), List(3, 4), List(4, 5), List(5, 6), List(6, 7), List(7, 8), List(8, 9), List(9, 0))
	}
}

func BenchmarkList10x10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = List(	List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)	)
	}
}

func BenchmarkSliceLen1(b *testing.B) {
	v := List(0)
	for i := 0; i < b.N; i++ {
		_ = v.Len()
	}
}

func BenchmarkSliceLen1x1(b *testing.B) {
	v := List(0, List(0))
	for i := 0; i < b.N; i++ {
		_ = v.Len()
	}
}

func BenchmarkSliceLen1x10(b *testing.B) {
	v := List(List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	for i := 0; i < b.N; i++ {
		_ = v.Len()
	}
}

func BenchmarkSliceLen10(b *testing.B) {
	v := List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	for i := 0; i < b.N; i++ {
		_ = v.Len()
	}
}

func BenchmarkSliceLen10x2(b *testing.B) {
	v := List(List(0, 1), List(1, 2), List(2, 3), List(3, 4), List(4, 5), List(5, 6), List(6, 7), List(7, 8), List(8, 9), List(9, 0))
	for i := 0; i < b.N; i++ {
		_ = v.Len()
	}
}

func BenchmarkSliceLen10x10(b *testing.B) {
	v := List(	List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)	)
	for i := 0; i < b.N; i++ {
		_ = v.Len()
	}
}

func BenchmarkSliceDepth1(b *testing.B) {
	v := List(0)
	for i := 0; i < b.N; i++ {
		_ = v.Depth()
	}
}

func BenchmarkSliceDepth1x1(b *testing.B) {
	v := List(0, List(0))
	for i := 0; i < b.N; i++ {
		_ = v.Depth()
	}
}

func BenchmarkSliceDepth1x10(b *testing.B) {
	v := List(List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	for i := 0; i < b.N; i++ {
		_ = v.Depth()
	}
}

func BenchmarkSliceDepth10(b *testing.B) {
	v := List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	for i := 0; i < b.N; i++ {
		_ = v.Depth()
	}
}

func BenchmarkSliceDepth10x2(b *testing.B) {
	v := List(List(0, 1), List(1, 2), List(2, 3), List(3, 4), List(4, 5), List(5, 6), List(6, 7), List(7, 8), List(8, 9), List(9, 0))
	for i := 0; i < b.N; i++ {
		_ = v.Depth()
	}
}

func BenchmarkSliceDepth10x10(b *testing.B) {
	v := List(	List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)	)
	for i := 0; i < b.N; i++ {
		_ = v.Depth()
	}
}

func BenchmarkSliceReverse10(b *testing.B) {
	b.StopTimer()
		v := List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		v.Reverse()
	}
}

func BenchmarkSliceReverse10x10(b *testing.B) {
	b.StopTimer()
		v := List(	List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)	)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		v.Reverse()
	}
}
func BenchmarkSliceFlatten1(b *testing.B) {
	v := List(0)
	for i := 0; i < b.N; i++ {
		v.Flatten()
	}
}

func BenchmarkSliceFlatten1x1(b *testing.B) {
	v := List(0, List(0))
	for i := 0; i < b.N; i++ {
		v.Flatten()
	}
}

func BenchmarkSliceFlatten1x10(b *testing.B) {
	v := List(List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	for i := 0; i < b.N; i++ {
		v.Flatten()
	}
}

func BenchmarkSliceFlatten10(b *testing.B) {
	v := List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	for i := 0; i < b.N; i++ {
		v.Flatten()
	}
}

func BenchmarkSliceFlatten10x2(b *testing.B) {
	v := List(List(0, 1), List(1, 2), List(2, 3), List(3, 4), List(4, 5), List(5, 6), List(6, 7), List(7, 8), List(8, 9), List(9, 0))
	for i := 0; i < b.N; i++ {
		v.Flatten()
	}
}

func BenchmarkSliceFlatten10x10(b *testing.B) {
	v := List(	List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)	)
	for i := 0; i < b.N; i++ {
		v.Flatten()
	}
}

func BenchmarkSliceCar(b *testing.B) {
	v := List(0, 1)
	for i := 0; i < b.N; i++ {
		_ = v.Car()
	}
}

func BenchmarkSliceCdr(b *testing.B) {
	v := List(0, 1)
	for i := 0; i < b.N; i++ {
		_ = v.Cdr()
	}
}

func BenchmarkSliceRplaca(b *testing.B) {}

func BenchmarkSliceRplacd(b *testing.B) {}