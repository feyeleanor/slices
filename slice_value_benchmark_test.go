package slices

import "testing"

//func BenchmarkVSlice(b *testing.B) {}

//func BenchmarkSliceValueBlockCopy(b *testing.B) {}

//func BenchmarkSliceValueOverwrite(b *testing.B) {}

func BenchmarkSliceValueAppend1x1(b *testing.B) {
	v := VSlice([]int{ 0: 0 })
	for i := 0; i < b.N; i++ {
		VSlice([]int{ 0: 0 }).Append(v)
	}
}

func BenchmarkSliceValueAppend1x10(b *testing.B) {
	b.StopTimer()
		v := VSlice(make([]int, 0, 10))
		for i := 0; i < 10; i++ {
			v.Append(i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		VSlice([]int{ 0: 0 }).Append(v)
	}
}

func BenchmarkSliceValueAppend1x100(b *testing.B) {
	b.StopTimer()
		v := VSlice(make([]int, 0, 100))
		for i := 0; i < 100; i++ {
			v.Append(i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		VSlice([]int{ 0: 0 }).Append(v)
	}
}

func BenchmarkSliceValueAppend1x1000(b *testing.B) {
	b.StopTimer()
		v := VSlice(make([]int, 0, 1000))
		for i := 0; i < 1000; i++ {
			v.Append(i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		VSlice([]int{ 0: 0 }).Append(v)
	}
}

//func BenchmarkSliceValuePrepend(b *testing.B) {}

func BenchmarkSliceValueAt(b *testing.B) {
	v := VSlice([]int{ 0 })
	for i := 0; i < b.N; i++ {
		_ = v.At(0)
	}
}

func BenchmarkSliceValueSet(b *testing.B) {
	v := VSlice([]int{ 0: 0 })
	for i := 0; i < b.N; i++ {
		v.Set(0, 0)
	}
}

func BenchmarkSliceValueRepeat1x1(b *testing.B) {
	v := VSlice([]int{ 0: 0 })
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(1)
	}
}

func BenchmarkSliceValueRepeat1x10(b *testing.B) {
	v := VSlice([]int{ 0: 0 })
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(10)
	}
}

func BenchmarkSliceValueRepeat1x100(b *testing.B) {
	v := VSlice([]int{ 0: 0 })
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(100)
	}
}

func BenchmarkSliceValueRepeat1x1000(b *testing.B) {
	v := VSlice([]int{ 0: 0 })
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(1000)
	}
}

func BenchmarkSliceValueRepeat10x1(b *testing.B) {
	b.StopTimer()
		v := VSlice(make([]int, 0, 10))
		for i := 0; i < 10; i++ {
			v.Append(i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(1)
	}
}

func BenchmarkSliceValueRepeat10x10(b *testing.B) {
	b.StopTimer()
		v := VSlice(make([]int, 0, 10))
		for i := 0; i < 10; i++ {
			v.Append(i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(10)
	}
}

func BenchmarkSliceValueRepeat10x100(b *testing.B) {
	b.StopTimer()
		v := VSlice(make([]int, 0, 10))
		for i := 0; i < 10; i++ {
			v.Append(i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(100)
	}
}

func BenchmarkSliceValueRepeat10x1000(b *testing.B) {
	b.StopTimer()
		v := VSlice(make([]int, 0, 10))
		for i := 0; i < 10; i++ {
			v.Append(i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(1000)
	}
}

func BenchmarkSliceValueRepeat100x1(b *testing.B) {
	b.StopTimer()
		v := VSlice(make([]int, 0, 100))
		for i := 0; i < 100; i++ {
			v.Append(i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(1)
	}
}

func BenchmarkSliceValueRepeat100x10(b *testing.B) {
	b.StopTimer()
		v := VSlice(make([]int, 0, 100))
		for i := 0; i < 100; i++ {
			v.Append(i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(10)
	}
}

func BenchmarkSliceValueRepeat100x100(b *testing.B) {
	b.StopTimer()
		v := VSlice(make([]int, 0, 100))
		for i := 0; i < 100; i++ {
			v.Append(i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(100)
	}
}

func BenchmarkSliceValueRepeat100x1000(b *testing.B) {
	b.StopTimer()
		v := VSlice(make([]int, 0, 100))
		for i := 0; i < 100; i++ {
			v.Append(i)
		}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Repeat(1000)
	}
}

//func BenchmarkSliceValueSection(b *testing.B) {}

//func BenchmarkSliceValueReallocate(b *testing.B) {}