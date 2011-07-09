package slices

import "fmt"
import "sort"

func I64List(n... int64) *I64Slice {
	return (*I64Slice)(&n)
}

type I64Slice	[]int64

func (s I64Slice) Len() int							{ return len(s) }
func (s I64Slice) Cap() int							{ return cap(s) }

func (s I64Slice) At(i int) interface{}				{ return s[i] }
func (s I64Slice) Set(i int, v interface{})			{ s[i] = v.(int64) }
func (s I64Slice) Clear(i int)						{ s[i] = 0 }
func (s I64Slice) Swap(i, j int)					{ s[i], s[j] = s[j], s[i] }

func (s I64Slice) Negate(i int)						{ s[i] = -s[i] }
func (s I64Slice) Increment(i int)					{ s[i]++ }
func (s I64Slice) Decrement(i int)					{ s[i]-- }

func (s I64Slice) Add(i, j int)						{ s[i] += s[j] }
func (s I64Slice) Subtract(i, j int)				{ s[i] -= s[j] }
func (s I64Slice) Multiply(i, j int)				{ s[i] *= s[j] }
func (s I64Slice) Divide(i, j int)					{ s[i] /= s[j] }
func (s I64Slice) Remainder(i, j int)				{ s[i] %= s[j] }

func (s I64Slice) And(i, j int)						{ s[i] &= s[j] }
func (s I64Slice) Or(i, j int)						{ s[i] |= s[j] }
func (s I64Slice) Xor(i, j int)						{ s[i] ^= s[j] }
func (s I64Slice) Invert(i int)						{ s[i] = ^s[i] }
func (s I64Slice) ShiftLeft(i, j int)				{ s[i] <<= uint(s[j]) }
func (s I64Slice) ShiftRight(i, j int)				{ s[i] >>= uint(s[j]) }

func (s I64Slice) Less(i, j int) bool				{ return s[i] < s[j] }
func (s I64Slice) AtLeast(i, j int) bool			{ return s[i] <= s[j] }
func (s I64Slice) Same(i, j int) bool				{ return s[i] == s[j] }
func (s I64Slice) AtMost(i, j int) bool				{ return s[i] >= s[j] }
func (s I64Slice) More(i, j int) bool				{ return s[i] > s[j] }
func (s I64Slice) ZeroLess(i int) bool				{ return 0 < s[i] }
func (s I64Slice) ZeroAtLeast(i, j int) bool		{ return 0 <= s[j] }
func (s I64Slice) ZeroSame(i int) bool				{ return 0 == s[i] }
func (s I64Slice) ZeroAtMost(i, j int) bool			{ return 0 >= s[j] }
func (s I64Slice) ZeroMore(i int) bool				{ return 0 > s[i] }

func (s I64Slice) Sort()							{ sort.Sort(s) }

func (s *I64Slice) RestrictTo(i, j int)				{ *s = (*s)[i:j] }

func (s I64Slice) Compare(i, j int) (r int) {
	switch {
	case s[i] < s[j]:		r = IS_LESS_THAN
	case s[i] > s[j]:		r = IS_GREATER_THAN
	default:				r = IS_SAME_AS
	}
	return
}

func (s I64Slice) ZeroCompare(i int) (r int) {
	switch {
	case 0 < s[i]:			r = IS_LESS_THAN
	case 0 > s[i]:			r = IS_GREATER_THAN
	default:				r = IS_SAME_AS
	}
	return
}

func (s *I64Slice) Cut(i, j int) {
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

func (s *I64Slice) Trim(i, j int) {
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

func (s *I64Slice) Delete(i int) {
	a := *s
	n := len(a)
	if i > -1 && i < n {
		copy(a[i:n - 1], a[i + 1:n])
		*s = a[:n - 1]
	}
}

func (s *I64Slice) DeleteIf(f interface{}) {
	a := *s
	p := 0
	switch f := f.(type) {
	case int64:						for i, v := range a {
										if i != p {
											a[p] = v
										}
										if v != f {
											p++
										}
									}

	case func(int64) bool:			for i, v := range a {
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

func (s I64Slice) Each(f interface{}) {
	switch f := f.(type) {
	case func(int64):						for _, v := range s { f(v) }
	case func(int, int64):					for i, v := range s { f(i, v) }
	case func(interface{}, int64):			for i, v := range s { f(i, v) }
	case func(interface{}):					for _, v := range s { f(v) }
	case func(int, interface{}):			for i, v := range s { f(i, v) }
	case func(interface{}, interface{}):	for i, v := range s { f(i, v) }
	}
}

func (s I64Slice) String() (t string) {
	for _, v := range s {
		if len(t) > 0 {
			t += " "
		}
		t += fmt.Sprintf("%v", v)
	}
	return fmt.Sprintf("(%v)", t)
}

func (s I64Slice) BlockCopy(destination, source, count int) {
	end := source + count
	if end > len(s) {
		end = len(s)
	}
	copy(s[destination:], s[source:end])
}

func (s I64Slice) BlockClear(start, count int) {
	copy(s[start:], make(I64Slice, count, count))
}

func (s I64Slice) Overwrite(offset int, container interface{}) {
	switch container := container.(type) {
	case I64Slice:			copy(s[offset:], container)
	case []int64:			copy(s[offset:], container)
	}
}

func (s *I64Slice) Reallocate(length, capacity int) {
	switch {
	case length > capacity:		s.Reallocate(capacity, capacity)
	case capacity != cap(*s):	x := make(I64Slice, length, capacity)
								copy(x, *s)
								*s = x
	default:					*s = (*s)[:length]
	}
}

func (s *I64Slice) Extend(n int) {
	c := cap(*s)
	l := len(*s) + n
	if l > c {
		c = l
	}
	s.Reallocate(l, c)
}

func (s *I64Slice) Expand(i, n int) {
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
		x := make(I64Slice, l, c)
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

func (s I64Slice) Reverse() {
	end := s.Len() - 1
	for i := 0; i < end; i++ {
		s[i], s[end] = s[end], s[i]
		end--
	}
}

func (s I64Slice) Depth() int {
	return 0
}

func (s *I64Slice) Append(v interface{}) {
	switch v := v.(type) {
	case int64:				*s = append(*s, v)
	case I64Slice:			*s = append(*s, v...)
	case *I64Slice:			*s = append(*s, (*v)...)
	case []int64:			s.Append(I64Slice(v))
	case *[]int64:			s.Append(I64Slice(*v))
	default:				panic(v)
	}
}

func (s *I64Slice) Prepend(v interface{}) {
	switch v := v.(type) {
	case int64:				l := s.Len() + 1
							n := make(I64Slice, l, l)
							n[0] = v
							copy(n[1:], *s)
							*s = n

	case I64Slice:			l := s.Len() + len(v)
							n := make(I64Slice, l, l)
							copy(n, v)
							copy(n[len(v):], *s)
							*s = n

	case *I64Slice:			s.Prepend(*v)
	case []int64:			s.Prepend(I64Slice(v))
	case *[]int64:			s.Prepend(I64Slice(*v))
	default:				panic(v)
	}
}

func (s I64Slice) Repeat(count int) I64Slice {
	length := len(s) * count
	capacity := cap(s)
	if capacity < length {
		capacity = length
	}
	destination := make(I64Slice, length, capacity)
	for start, end := 0, len(s); count > 0; count-- {
		copy(destination[start:end], s)
		start = end
		end += len(s)
	}
	return destination
}

func (s *I64Slice) Flatten() {
	//	Flatten is a non-op for the I64Slice as they cannot contain nested elements
}

func (s I64Slice) equal(o I64Slice) (r bool) {
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

func (s I64Slice) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case *I64Slice:			r = o != nil && s.equal(*o)
	case I64Slice:			r = s.equal(o)
	case *[]int64:			r = o != nil && s.equal(*o)
	case []int64:			r = s.equal(o)
	}
	return
}

func (s I64Slice) Car() (h interface{}) {
	if s.Len() > 0 {
		h = s[0]
	}
	return
}

func (s I64Slice) Cdr() (t I64Slice) {
	if s.Len() > 1 {
		t = s[1:]
	}
	return
}

func (s *I64Slice) Rplaca(v interface{}) {
	switch {
	case s == nil:			*s = *I64List(v.(int64))
	case s.Len() == 0:		*s = append(*s, v.(int64))
	default:				(*s)[0] = v.(int64)
	}
}

func (s *I64Slice) Rplacd(v interface{}) {
	if s == nil {
		*s = *I64List(v.(int64))
	} else {
		ReplaceSlice := func(v I64Slice) {
			if l := len(v); l < cap(*s) {
				copy((*s)[1:], v)
				*s = (*s)[:l + 1]
			} else {
				l++
				n := make(I64Slice, l, l)
				copy(n, (*s)[:1])
				copy(n[1:], v)
				*s = n
			}
		}

		switch v := v.(type) {
		case *I64Slice:		ReplaceSlice(*v)
		case I64Slice:		ReplaceSlice(v)
		case *[]int64:		ReplaceSlice(I64Slice(*v))
		case []int64:		ReplaceSlice(I64Slice(v))
		case nil:			*s = (*s)[:1]
		default:			(*s)[1] = v.(int64)
							*s = (*s)[:2]
		}
	}
}

func (s I64Slice) SetIntersection(o I64Slice) (r I64Slice) {
	cache := make(map[int64]bool)
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

func (s I64Slice) SetUnion(o I64Slice) (r I64Slice) {
	cache := make(map[int64]bool)
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

func (s I64Slice) SetDifference(o I64Slice) (r I64Slice) {
	left := make(map[int64]bool)
	right := make(map[int64]bool)
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