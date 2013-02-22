package slices

import "testing"

var sC64 = C64Slice{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func BenchmarkC64SliceString(b *testing.B) {
    for i := 0; i < b.N; i++ {
		sC64.String()
	}
}

func BenchmarkC64Slice_len(b *testing.B) {
    for i := 0; i < b.N; i++ {
		_ = len(sC64)
	}
}

func BenchmarkC64SliceLen(b *testing.B) {
    for i := 0; i < b.N; i++ {
		sC64.Len()
	}
}

func BenchmarkC64Slice_cap(b *testing.B) {
    for i := 0; i < b.N; i++ {
		_ = cap(sC64)
	}
}

func BenchmarkC64SliceCap(b *testing.B) {
    for i := 0; i < b.N; i++ {
		sC64.Cap()
	}
}

func BenchmarkC64Slice_at(b *testing.B) {
    for i := 0; i < b.N; i++ {
		_ = sC64[0]
	}
}

func BenchmarkC64SliceAt(b *testing.B) {
    for i := 0; i < b.N; i++ {
		sC64.At(0)
	}
}

func BenchmarkC64Slice_set(b *testing.B) {
    for i := 0; i < b.N; i++ {
		sC64[0] = 0
	}
}

func BenchmarkC64SliceSet(b *testing.B) {
    for i := 0; i < b.N; i++ {
		sC64.Set(0, complex64(0))
	}
}

func BenchmarkC64SliceExpand10(b *testing.B) {
    for i := 0; i < b.N; i++ {
		s := C64Slice{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		s.Expand(5, 10)
	}
}

func BenchmarkC64SliceExpand100(b *testing.B) {
    for i := 0; i < b.N; i++ {
		s := C64Slice{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		s.Expand(5, 100)
	}
}