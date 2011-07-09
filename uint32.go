package slices

import "fmt"
import "sort"

func U32List(n... uint32) *U32Slice {
	return (*U32Slice)(&n)
}

type U32Slice	[]uint32

func (s U32Slice) Len() int							{ return len(s) }
func (s U32Slice) Cap() int							{ return cap(s) }

func (s U32Slice) At(i int) interface{}				{ return s[i] }
func (s U32Slice) U32At(i int) uint32				{ return s[i] }
func (s U32Slice) Set(i int, v interface{})			{ s[i] = v.(uint32) }
func (s U32Slice) U32Set(i int, v uint32)			{ s[i] = v }
func (s U32Slice) Clear(i int)						{ s[i] = 0 }
func (s U32Slice) Swap(i, j int)					{ s[i], s[j] = s[j], s[i] }

func (s U32Slice) Negate(i int)						{ s[i] = -s[i] }
func (s U32Slice) Increment(i int)					{ s[i]++ }
func (s U32Slice) Decrement(i int)					{ s[i]++ }

func (s U32Slice) Add(i, j int)						{ s[i] += s[j] }
func (s U32Slice) Subtract(i, j int)				{ s[i] -= s[j] }
func (s U32Slice) Multiply(i, j int)				{ s[i] *= s[j] }
func (s U32Slice) Divide(i, j int)					{ s[i] /= s[j] }
func (s U32Slice) Remainder(i, j int)				{ s[i] %= s[j] }

func (s U32Slice) And(i, j int)						{ s[i] &= s[j] }
func (s U32Slice) Or(i, j int)						{ s[i] |= s[j] }
func (s U32Slice) Xor(i, j int)						{ s[i] ^= s[j] }
func (s U32Slice) Invert(i int)						{ s[i] = ^s[i] }
func (s U32Slice) ShiftLeft(i, j int)				{ s[i] <<= s[j] }
func (s U32Slice) ShiftRight(i, j int)				{ s[i] >>= s[j] }

func (s U32Slice) Less(i, j int) bool				{ return s[i] < s[j] }
func (s U32Slice) AtLeast(i, j int) bool			{ return s[i] <= s[j] }
func (s U32Slice) Same(i, j int) bool				{ return s[i] == s[j] }
func (s U32Slice) AtMost(i, j int) bool				{ return s[i] >= s[j] }
func (s U32Slice) More(i, j int) bool				{ return s[i] > s[j] }
func (s U32Slice) ZeroLess(i int) bool				{ return 0 < s[i] }
func (s U32Slice) ZeroAtLeast(i, j int) bool		{ return true }
func (s U32Slice) ZeroSame(i int) bool				{ return 0 == s[i] }
func (s U32Slice) ZeroAtMost(i, j int) bool			{ return 0 == s[j] }
func (s U32Slice) ZeroMore(i int) bool				{ return false }

func (s U32Slice) Sort()							{ sort.Sort(s) }

func (s *U32Slice) RestrictTo(i, j int)				{ *s = (*s)[i:j] }

func (s U32Slice) Compare(i, j int) (r int) {
	switch {
	case s[i] < s[j]:		r = IS_LESS_THAN
	case s[i] > s[j]:		r = IS_GREATER_THAN
	default:				r = IS_SAME_AS
	}
	return
}

func (s U32Slice) ZeroCompare(i int) (r int) {
	switch {
	case 0 < s[i]:			r = IS_LESS_THAN
	default:				r = IS_SAME_AS
	}
	return
}

func (s *U32Slice) Cut(i, j int) {
	a := *s
	l := len(a)
	if i < 0 {
		i = 0
	}
	if j > l {
		j = l
	}
	if j > i {
		if m := l - (j - i); m > 0 && l > m {
			copy(a[i:m], a[j:l])
			*s = a[:m]
		}
	}
}

func (s *U32Slice) Trim(i, j int) {
	a := *s
	n := len(a)
	if i < 0 {
		i = 0
	}
	if j > n {
		j = n
	}
	if j > i {
		copy(a, a[i:j])
		*s = a[:j - i]
	}
}

func (s *U32Slice) Delete(i int) {
	a := *s
	n := len(a)
	if i > -1 && i < n {
		copy(a[i:n - 1], a[i + 1:n])
		*s = a[:n - 1]
	}
}

func (s *U32Slice) DeleteIf(f interface{}) {
	a := *s
	p := 0
	switch f := f.(type) {
	case uint32:					for i, v := range a {
										if i != p {
											a[p] = v
										}
										if v != f {
											p++
										}
									}

	case func(uint32) bool:			for i, v := range a {
										if i != p {
											a[p] = v
										}
										if !f(v) {
											p++
										}
									}

	case func(interface{}) bool:	for i, v := range a {
										if i != p {
											a[p] = v
										}
										if !f(v) {
											p++
										}
									}

	default:						p = len(a)
	}
	*s = a[:p]
}

func (s U32Slice) Each(f interface{}) {
	switch f := f.(type) {
	case func(uint32):						for _, v := range s { f(v) }
	case func(int, uint32):					for i, v := range s { f(i, v) }
	case func(interface{}, uint32):			for i, v := range s { f(i, v) }
	case func(interface{}):					for _, v := range s { f(v) }
	case func(int, interface{}):			for i, v := range s { f(i, v) }
	case func(interface{}, interface{}):	for i, v := range s { f(i, v) }
	}
}

func (s U32Slice) String() (t string) {
	for _, v := range s {
		if len(t) > 0 {
			t += " "
		}
		t += fmt.Sprintf("%v", v)
	}
	return fmt.Sprintf("(%v)", t)
}

func (s U32Slice) BlockCopy(destination, source, count int) {
	end := source + count
	if end > len(s) {
		end = len(s)
	}
	copy(s[destination:], s[source:end])
}

func (s U32Slice) BlockClear(start, count int) {
	copy(s[start:], make(U32Slice, count, count))
}

func (s U32Slice) Overwrite(offset int, container interface{}) {
	switch container := container.(type) {
	case U32Slice:			copy(s[offset:], container)
	case []uint32:			copy(s[offset:], container)
	}
}

func (s *U32Slice) Reallocate(length, capacity int) {
	switch {
	case length > capacity:		s.Reallocate(capacity, capacity)
	case capacity != cap(*s):	x := make(U32Slice, length, capacity)
								copy(x, *s)
								*s = x
	default:					*s = (*s)[:length]
	}
}

func (s *U32Slice) Extend(n int) {
	c := cap(*s)
	l := len(*s) + n
	if l > c {
		c = l
	}
	s.Reallocate(l, c)
}

func (s *U32Slice) Expand(i, n int) {
	if i < 0 {
		i = 0
	}

	l := s.Len()
	if l < i {
		i = l
	}

	l += n
	c := s.Cap()
	if c < l {
		c = l
	}

	if c != s.Cap() {
		x := make(U32Slice, l, c)
		copy(x, (*s)[:i])
		copy(x[i + n:], (*s)[i:])
		*s = x
	} else {
		a := (*s)[:l]
		for j := l - 1; j >= i; j-- {
			a[j] = a[j - n]
		}
		*s = a
	}
}

func (s U32Slice) Reverse() {
	end := s.Len() - 1
	for i := 0; i < end; i++ {
		s[i], s[end] = s[end], s[i]
		end--
	}
}

func (s U32Slice) Depth() int {
	return 0
}

func (s *U32Slice) Append(v interface{}) {
	switch v := v.(type) {
	case uint32:			*s = append(*s, v)
	case U32Slice:			*s = append(*s, v...)
	case *U32Slice:			*s = append(*s, (*v)...)
	case []uint32:			s.Append(U32Slice(v))
	case *[]uint32:			s.Append(U32Slice(*v))
	default:				panic(v)
	}
}

func (s *U32Slice) Prepend(v interface{}) {
	switch v := v.(type) {
	case uint32:			l := s.Len() + 1
							n := make(U32Slice, l, l)
							n[0] = v
							copy(n[1:], *s)
							*s = n

	case U32Slice:			l := s.Len() + len(v)
							n := make(U32Slice, l, l)
							copy(n, v)
							copy(n[len(v):], *s)
							*s = n

	case *U32Slice:			s.Prepend(*v)
	case []uint32:			s.Prepend(U32Slice(v))
	case *[]uint32:			s.Prepend(U32Slice(*v))
	default:				panic(v)
	}
}

func (s U32Slice) Repeat(count int) U32Slice {
	length := len(s) * count
	capacity := cap(s)
	if capacity < length {
		capacity = length
	}
	destination := make(U32Slice, length, capacity)
	for start, end := 0, len(s); count > 0; count-- {
		copy(destination[start:end], s)
		start = end
		end += len(s)
	}
	return destination
}

func (s *U32Slice) Flatten() {
	//	Flatten is a non-op for the U32Slice as they cannot contain nested elements
}

func (s U32Slice) equal(o U32Slice) (r bool) {
	switch {
	case s == nil:				r = o == nil
	case s.Len() == o.Len():	r = true
								for i, v := range s {
									if r = v == o[i]; !r {
										return
									}
								}
	}
	return
}

func (s U32Slice) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case *U32Slice:			r = o != nil && s.equal(*o)
	case U32Slice:			r = s.equal(o)
	case *[]uint32:			r = o != nil && s.equal(*o)
	case []uint32:			r = s.equal(o)
	}
	return
}

func (s U32Slice) Car() (h interface{}) {
	if s.Len() > 0 {
		h = s[0]
	}
	return
}

func (s U32Slice) Cdr() (t U32Slice) {
	if s.Len() > 1 {
		t = s[1:]
	}
	return
}

func (s *U32Slice) Rplaca(v interface{}) {
	switch {
	case s == nil:			*s = *U32List(v.(uint32))
	case s.Len() == 0:		*s = append(*s, v.(uint32))
	default:				(*s)[0] = v.(uint32)
	}
}

func (s *U32Slice) Rplacd(v interface{}) {
	if s == nil {
		*s = *U32List(v.(uint32))
	} else {
		ReplaceSlice := func(v U32Slice) {
			if l := len(v); l < cap(*s) {
				copy((*s)[1:], v)
				*s = (*s)[:l + 1]
			} else {
				l++
				n := make(U32Slice, l, l)
				copy(n, (*s)[:1])
				copy(n[1:], v)
				*s = n
			}
		}

		switch v := v.(type) {
		case *U32Slice:		ReplaceSlice(*v)
		case U32Slice:		ReplaceSlice(v)
		case *[]uint32:		ReplaceSlice(U32Slice(*v))
		case []uint32:		ReplaceSlice(U32Slice(v))
		case nil:			*s = (*s)[:1]
		default:			(*s)[1] = v.(uint32)
							*s = (*s)[:2]
		}
	}
}

func (s U32Slice) SetIntersection(o U32Slice) (r U32Slice) {
	cache := make(map[uint32]bool)
	for _, v := range s {
		if ok := cache[v]; !ok {
			cache[v] = true
		}
	}
	for _, v := range o {
		if _, ok := cache[v]; ok {
			cache[v] = false, false
			r = append(r, v)
		}
	}
	return
}

func (s U32Slice) SetUnion(o U32Slice) (r U32Slice) {
	cache := make(map[uint32]bool)
	for _, v := range s {
		if ok := cache[v]; !ok {
			cache[v] = true
		}
	}
	for _, v := range o {
		if ok := cache[v]; !ok {
			cache[v] = true
		}
	}
	for k, _ := range cache {
		r = append(r, k)
	}
	return
}

func (s U32Slice) SetDifference(o U32Slice) (r U32Slice) {
	left := make(map[uint32]bool)
	right := make(map[uint32]bool)
	for _, v := range s {
		if ok := left[v]; !ok {
			left[v] = true
		}
	}
	for _, v := range o {
		if ok := right[v]; !ok {
			right[v] = true
		}
	}
	for k, _ := range left {
		if ok := right[k]; ok {
			right[k] = false, false
		} else {
			r = append(r, k)
		}
	}
	for k, _ := range right {
		if ok := left[k]; !ok {
			r = append(r, k)
		}
	}
	return
}