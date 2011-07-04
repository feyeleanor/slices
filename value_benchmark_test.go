package slices

import "testing"

//func BenchmarkVSlice(b *testing.B) {}

//func BenchmarkVSliceBlockCopy(b *testing.B) {}

//func BenchmarkVSliceOverwrite(b *testing.B) {}

func BenchmarkVSliceAppend1x1(b *testing.B) {
	v := VWrap([]int{ 0: 0 })
	for i := 0; i < b.N; i++ {
		VWrap([]int{ 0: 0 }).Append(v)
	}
}

func BenchmarkVSliceAppend1x10(b *testing.B) {
	b.StopTimer()
		v := VWrap(make([]int, 0, 10))
		for i := 0; i < 10; i++ {
			v.Append(i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		VWrap([]int{ 0: 0 }).Append(v)
	}
}

func BenchmarkVSliceAppend1x100(b *testing.B) {
	b.StopTimer()
		v := VWrap(make([]int, 0, 100))
		for i := 0; i < 100; i++ {
			v.Append(i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		VWrap([]int{ 0: 0 }).Append(v)
	}
}

func BenchmarkVSliceAppend1x1000(b *testing.B) {
	b.StopTimer()
		v := VWrap(make([]int, 0, 1000))
		for i := 0; i < 1000; i++ {
			v.Append(i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		VWrap([]int{ 0: 0 }).Append(v)
	}
}

//func BenchmarkVSlicePrepend(b *testing.B) {}

func BenchmarkVSliceAt(b *testing.B) {
	v := VWrap([]int{ 0 })
	for i := 0; i < b.N; i++ {
		_ = v.At(0)
	}
}

func BenchmarkVSliceSet(b *testing.B) {
	v := VWrap([]int{ 0: 0 })
	for i := 0; i < b.N; i++ {
		v.Set(0, 0)
	}
}

func BenchmarkVSliceRepeat1x1(b *testing.B) {
	v := VWrap([]int{ 0: 0 })
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(1)
	}
}

func BenchmarkVSliceRepeat1x10(b *testing.B) {
	v := VWrap([]int{ 0: 0 })
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(10)
	}
}

func BenchmarkVSliceRepeat1x100(b *testing.B) {
	v := VWrap([]int{ 0: 0 })
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(100)
	}
}

func BenchmarkVSliceRepeat1x1000(b *testing.B) {
	v := VWrap([]int{ 0: 0 })
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(1000)
	}
}

func BenchmarkVSliceRepeat10x1(b *testing.B) {
	b.StopTimer()
		v := VWrap(make([]int, 0, 10))
		for i := 0; i < 10; i++ {
			v.Append(i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(1)
	}
}

func BenchmarkVSliceRepeat10x10(b *testing.B) {
	b.StopTimer()
		v := VWrap(make([]int, 0, 10))
		for i := 0; i < 10; i++ {
			v.Append(i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(10)
	}
}

func BenchmarkVSliceRepeat10x100(b *testing.B) {
	b.StopTimer()
		v := VWrap(make([]int, 0, 10))
		for i := 0; i < 10; i++ {
			v.Append(i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(100)
	}
}

func BenchmarkVSliceRepeat10x1000(b *testing.B) {
	b.StopTimer()
		v := VWrap(make([]int, 0, 10))
		for i := 0; i < 10; i++ {
			v.Append(i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(1000)
	}
}

func BenchmarkVSliceRepeat100x1(b *testing.B) {
	b.StopTimer()
		v := VWrap(make([]int, 0, 100))
		for i := 0; i < 100; i++ {
			v.Append(i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(1)
	}
}

func BenchmarkVSliceRepeat100x10(b *testing.B) {
	b.StopTimer()
		v := VWrap(make([]int, 0, 100))
		for i := 0; i < 100; i++ {
			v.Append(i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(10)
	}
}

func BenchmarkVSliceRepeat100x100(b *testing.B) {
	b.StopTimer()
		v := VWrap(make([]int, 0, 100))
		for i := 0; i < 100; i++ {
			v.Append(i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(100)
	}
}

func BenchmarkVSliceRepeat100x1000(b *testing.B) {
	b.StopTimer()
		v := VWrap(make([]int, 0, 100))
		for i := 0; i < 100; i++ {
			v.Append(i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(1000)
	}
}

//func BenchmarkVSliceSection(b *testing.B) {}

//func BenchmarkVSliceReallocate(b *testing.B) {}