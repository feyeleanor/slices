package slices

import "fmt"

func U64List(n... uint64) *U64Slice {
	return (*U64Slice)(&n)
}

type U64Slice	[]uint64

func (s U64Slice) Len() int							{ return len(s) }
func (s U64Slice) Cap() int							{ return cap(s) }

func (s U64Slice) At(i int) interface{}				{ return s[i] }
func (s U64Slice) U64At(i int) uint64				{ return s[i] }
func (s U64Slice) Set(i int, v interface{})			{ s[i] = v.(uint64) }
func (s U64Slice) U64Set(i int, v uint64)			{ s[i] = v }
func (s U64Slice) Clear(i int)						{ s[i] = 0 }
func (s U64Slice) Swap(i, j int)					{ s[i], s[j] = s[j], s[i] }

func (s U64Slice) Negate(i int)						{ s[i] = -s[i] }
func (s U64Slice) Increment(i int)					{ s[i]++ }
func (s U64Slice) Decrement(i int)					{ s[i]-- }

func (s U64Slice) Add(i, j int)						{ s[i] += s[j] }
func (s U64Slice) Subtract(i, j int)				{ s[i] -= s[j] }
func (s U64Slice) Multiply(i, j int)				{ s[i] *= s[j] }
func (s U64Slice) Divide(i, j int)					{ s[i] /= s[j] }
func (s U64Slice) Remainder(i, j int)				{ s[i] %= s[j] }

func (s U64Slice) And(i, j int)						{ s[i] &= s[j] }
func (s U64Slice) Or(i, j int)						{ s[i] |= s[j] }
func (s U64Slice) Xor(i, j int)						{ s[i] ^= s[j] }
func (s U64Slice) Invert(i int)						{ s[i] = ^s[i] }
func (s U64Slice) ShiftLeft(i, j int)				{ s[i] <<= s[j] }
func (s U64Slice) ShiftRight(i, j int)				{ s[i] >>= s[j] }

func (s U64Slice) Less(i, j int) bool				{ return s[i] < s[j] }
func (s U64Slice) AtLeast(i, j int) bool			{ return s[i] <= s[j] }
func (s U64Slice) Same(i, j int) bool				{ return s[i] == s[j] }
func (s U64Slice) AtMost(i, j int) bool				{ return s[i] >= s[j] }
func (s U64Slice) More(i, j int) bool				{ return s[i] > s[j] }
func (s U64Slice) ZeroLess(i int) bool				{ return 0 < s[i] }
func (s U64Slice) ZeroAtLeast(i, j int) bool		{ return true }
func (s U64Slice) ZeroSame(i int) bool				{ return 0 == s[i] }
func (s U64Slice) ZeroAtMost(i, j int) bool			{ return 0 == s[j] }
func (s U64Slice) ZeroMore(i int) bool				{ return false }

func (s *U64Slice) RestrictTo(i, j int)				{ *s = (*s)[i:j] }

func (s U64Slice) Compare(i, j int) (r int) {
	switch {
	case s[i] < s[j]:		r = IS_LESS_THAN
	case s[i] > s[j]:		r = IS_GREATER_THAN
	default:				r = IS_SAME_AS
	}
	return
}

func (s U64Slice) ZeroCompare(i int) (r int) {
	switch {
	case 0 < s[i]:			r = IS_LESS_THAN
	default:				r = IS_SAME_AS
	}
	return
}

func (s *U64Slice) Cut(i, j int) {
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
			*s = a[0:m]
		}
	}
}

func (s *U64Slice) Trim(i, j int) {
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
		*s = a[0:j - i]
	}
}

func (s *U64Slice) Delete(i int) {
	a := *s
	n := len(a)
	if i > -1 && i < n {
		copy(a[i:n - 1], a[i + 1:n])
		*s = a[0 : n - 1]
	}
}

func (s U64Slice) Each(f func(interface{})) {
	for _, v := range s {
		f(v)
	}
}

func (s U64Slice) EachWithIndex(f func(int, interface{})) {
	for i, v := range s {
		f(i, v)
	}
}

func (s U64Slice) EachWithKey(f func(key, value interface{})) {
	for i, v := range s {
		f(i, v)
	}
}

func (s U64Slice) U64Each(f func(uint64)) {
	for _, v := range s {
		f(v)
	}
}

func (s U64Slice) U64EachWithIndex(f func(int, uint64)) {
	for i, v := range s {
		f(i, v)
	}
}

func (s U64Slice) U64EachWithKey(f func(interface{}, uint64)) {
	for i, v := range s {
		f(i, v)
	}
}

func (s U64Slice) String() (t string) {
	for _, v := range s {
		if len(t) > 0 {
			t += " "
		}
		t += fmt.Sprintf("%v", v)
	}
	return fmt.Sprintf("(%v)", t)
}

func (s U64Slice) BlockCopy(destination, source, count int) {
	end := source + count
	if end > len(s) {
		end = len(s)
	}
	copy(s[destination:], s[source:end])
}

func (s U64Slice) BlockClear(start, count int) {
	copy(s[start:], make(U64Slice, count, count))
}

func (s U64Slice) Overwrite(offset int, container interface{}) {
	switch container := container.(type) {
	case U64Slice:			copy(s[offset:], container)
	case []uint64:			copy(s[offset:], container)
	}
}

func (s *U64Slice) Reallocate(length, capacity int) {
	switch {
	case length > capacity:		s.Reallocate(capacity, capacity)
	case capacity != cap(*s):	x := make(U64Slice, length, capacity)
								copy(x, *s)
								*s = x
	default:					*s = (*s)[:length]
	}
}

func (s *U64Slice) Extend(n int) {
	c := cap(*s)
	l := len(*s) + n
	if l > c {
		c = l
	}
	s.Reallocate(l, c)
}

func (s *U64Slice) Expand(i, n int) {
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
		x := make(U64Slice, l, c)
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

func (s U64Slice) Reverse() {
	end := s.Len() - 1
	for i := 0; i < end; i++ {
		s[i], s[end] = s[end], s[i]
		end--
	}
}

func (s U64Slice) Depth() int {
	return 0
}

func (s *U64Slice) Append(v interface{}) {
	s.U64Append(v.(uint64))
}

func (s *U64Slice) U64Append(v uint64) {
	*s = append(*s, v)
}

func (s *U64Slice) AppendSlice(o U64Slice) {
	*s = append(*s, o...)
}

func (s *U64Slice) Prepend(v interface{}) {
	s.U64Prepend(v.(uint64))
}

func (s *U64Slice) U64Prepend(v uint64) {
	l := s.Len() + 1
	n := make(U64Slice, l, l)
	n[0] = v
	copy(n[1:], *s)
	*s = n
}

func (s *U64Slice) PrependSlice(o U64Slice) {
	l := s.Len() + o.Len()
	n := make(U64Slice, l, l)
	copy(n, o)
	copy(n[o.Len():], *s)
	*s = n
}

func (s U64Slice) Repeat(count int) U64Slice {
	length := len(s) * count
	capacity := cap(s)
	if capacity < length {
		capacity = length
	}
	destination := make(U64Slice, length, capacity)
	for start, end := 0, len(s); count > 0; count-- {
		copy(destination[start:end], s)
		start = end
		end += len(s)
	}
	return destination
}

func (s *U64Slice) Flatten() {
	//	Flatten is a non-op for the U64Slice as they cannot contain nested elements
}

func (s U64Slice) equal(o U64Slice) (r bool) {
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

func (s U64Slice) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case *U64Slice:			r = o != nil && s.equal(*o)
	case U64Slice:			r = s.equal(o)
	case *[]uint64:			r = o != nil && s.equal(*o)
	case []uint64:			r = s.equal(o)
	}
	return
}

func (s U64Slice) Car() (h interface{}) {
	if s.Len() > 0 {
		h = s[0]
	}
	return
}

func (s U64Slice) Cdr() (t U64Slice) {
	if s.Len() > 1 {
		t = s[1:]
	}
	return
}

func (s *U64Slice) Rplaca(v interface{}) {
	switch {
	case s == nil:			*s = *U64List(v.(uint64))
	case s.Len() == 0:		*s = append(*s, v.(uint64))
	default:				(*s)[0] = v.(uint64)
	}
}

func (s *U64Slice) Rplacd(v interface{}) {
	if s == nil {
		*s = *U64List(v.(uint64))
	} else {
		ReplaceSlice := func(v U64Slice) {
			if l := len(v); l < cap(*s) {
				copy((*s)[1:], v)
				*s = (*s)[0:l + 1]
			} else {
				l++
				n := make(U64Slice, l, l)
				copy(n, (*s)[:1])
				copy(n[1:], v)
				*s = n
			}
		}

		switch v := v.(type) {
		case *U64Slice:		ReplaceSlice(*v)
		case U64Slice:		ReplaceSlice(v)
		case *[]uint64:		ReplaceSlice(U64Slice(*v))
		case []uint64:		ReplaceSlice(U64Slice(v))
		case nil:			*s = (*s)[:1]
		default:			(*s)[1] = v.(uint64)
							*s = (*s)[:2]
		}
	}
}