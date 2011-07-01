package slices

import "fmt"

func U8List(n... uint8) *U8Slice {
	return (*U8Slice)(&n)
}

type U8Slice	[]uint8

func (s U8Slice) Len() int							{ return len(s) }
func (s U8Slice) Cap() int							{ return cap(s) }

func (s U8Slice) At(i int) interface{}				{ return s[i] }
func (s U8Slice) U8At(i int) uint8					{ return s[i] }
func (s U8Slice) Set(i int, v interface{})			{ s[i] = v.(uint8) }
func (s U8Slice) U8Set(i int, v uint8)				{ s[i] = v }
func (s U8Slice) Clear(i int)						{ s[i] = 0 }
func (s U8Slice) Swap(i, j int)						{ s[i], s[j] = s[j], s[i] }

func (s U8Slice) Negate(i int)						{ s[i] = -s[i] }
func (s U8Slice) Increment(i int)					{ s[i]++ }
func (s U8Slice) Decrement(i int)					{ s[i]-- }

func (s U8Slice) Add(i, j int)						{ s[i] += s[j] }
func (s U8Slice) Subtract(i, j int)					{ s[i] -= s[j] }
func (s U8Slice) Multiply(i, j int)					{ s[i] *= s[j] }
func (s U8Slice) Divide(i, j int)					{ s[i] /= s[j] }
func (s U8Slice) Remainder(i, j int)				{ s[i] %= s[j] }

func (s U8Slice) And(i, j int)						{ s[i] &= s[j] }
func (s U8Slice) Or(i, j int)						{ s[i] |= s[j] }
func (s U8Slice) Xor(i, j int)						{ s[i] ^= s[j] }
func (s U8Slice) Invert(i int)						{ s[i] = ^s[i] }
func (s U8Slice) ShiftLeft(i, j int)				{ s[i] <<= s[j] }
func (s U8Slice) ShiftRight(i, j int)				{ s[i] >>= s[j] }

func (s U8Slice) Less(i, j int) bool				{ return s[i] < s[j] }
func (s U8Slice) AtLeast(i, j int) bool				{ return s[i] <= s[j] }
func (s U8Slice) Same(i, j int) bool				{ return s[i] == s[j] }
func (s U8Slice) AtMost(i, j int) bool				{ return s[i] >= s[j] }
func (s U8Slice) More(i, j int) bool				{ return s[i] > s[j] }
func (s U8Slice) ZeroLess(i int) bool				{ return 0 < s[i] }
func (s U8Slice) ZeroAtLeast(i, j int) bool			{ return true }
func (s U8Slice) ZeroSame(i int) bool				{ return 0 == s[i] }
func (s U8Slice) ZeroAtMost(i, j int) bool			{ return 0 == s[j] }
func (s U8Slice) ZeroMore(i int) bool				{ return false }

func (s U8Slice) Compare(i, j int) (r int) {
	switch {
	case s[i] < s[j]:		r = IS_LESS_THAN
	case s[i] > s[j]:		r = IS_GREATER_THAN
	default:				r = IS_SAME_AS
	}
	return
}

func (s U8Slice) ZeroCompare(i int) (r int) {
	switch {
	case 0 < s[i]:			r = IS_LESS_THAN
	default:				r = IS_SAME_AS
	}
	return
}

func (s *U8Slice) Cut(i, j int) {
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

func (s *U8Slice) Delete(i int) {
	a := *s
	n := len(a)
	if i > -1 && i < n {
		copy(a[i:n - 1], a[i + 1:n])
		*s = a[0 : n - 1]
	}
}

func (s U8Slice) Each(f func(interface{})) {
	for _, v := range s {
		f(v)
	}
}

func (s U8Slice) EachWithIndex(f func(int, interface{})) {
	for i, v := range s {
		f(i, v)
	}
}

func (s U8Slice) EachWithKey(f func(key, value interface{})) {
	for i, v := range s {
		f(i, v)
	}
}

func (s U8Slice) U8Each(f func(uint8)) {
	for _, v := range s {
		f(v)
	}
}

func (s U8Slice) U8EachWithIndex(f func(int, uint8)) {
	for i, v := range s {
		f(i, v)
	}
}

func (s U8Slice) U8EachWithKey(f func(interface{}, uint8)) {
	for i, v := range s {
		f(i, v)
	}
}

func (s U8Slice) String() (t string) {
	for _, v := range s {
		if len(t) > 0 {
			t += " "
		}
		t += fmt.Sprintf("%v", v)
	}
	return fmt.Sprintf("(%v)", t)
}

func (s U8Slice) BlockCopy(destination, source, count int) {
	end := source + count
	if end > len(s) {
		end = len(s)
	}
	copy(s[destination:], s[source:end])
}

func (s U8Slice) BlockClear(start, count int) {
	copy(s[start:], make(U8Slice, count, count))
}

func (s U8Slice) Overwrite(offset int, container interface{}) {
	switch container := container.(type) {
	case U8Slice:			copy(s[offset:], container)
	case []uint8:			copy(s[offset:], container)
	}
}

func (s *U8Slice) Reallocate(length, capacity int) {
	switch {
	case length > capacity:		s.Reallocate(capacity, capacity)
	case capacity != cap(*s):	x := make(U8Slice, length, capacity)
								copy(x, *s)
								*s = x
	default:					*s = (*s)[:length]
	}
}

func (s *U8Slice) Extend(n int) {
	c := cap(*s)
	l := len(*s) + n
	if l > c {
		c = l
	}
	s.Reallocate(l, c)
}

func (s *U8Slice) Expand(i, n int) {
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
		x := make(U8Slice, l, c)
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

func (s U8Slice) Reverse() {
	end := s.Len() - 1
	for i := 0; i < end; i++ {
		s[i], s[end] = s[end], s[i]
		end--
	}
}

func (s U8Slice) Depth() int {
	return 0
}

func (s *U8Slice) Append(v interface{}) {
	s.U8Append(v.(uint8))
}

func (s *U8Slice) U8Append(v uint8) {
	*s = append(*s, v)
}

func (s *U8Slice) AppendSlice(o U8Slice) {
	*s = append(*s, o...)
}

func (s *U8Slice) Prepend(v interface{}) {
	s.U8Prepend(v.(uint8))
}

func (s *U8Slice) U8Prepend(v uint8) {
	l := s.Len() + 1
	n := make(U8Slice, l, l)
	n[0] = v
	copy(n[1:], *s)
	*s = n
}

func (s *U8Slice) PrependSlice(o U8Slice) {
	l := s.Len() + o.Len()
	n := make(U8Slice, l, l)
	copy(n, o)
	copy(n[o.Len():], *s)
	*s = n
}

func (s U8Slice) Repeat(count int) U8Slice {
	length := len(s) * count
	capacity := cap(s)
	if capacity < length {
		capacity = length
	}
	destination := make(U8Slice, length, capacity)
	for start, end := 0, len(s); count > 0; count-- {
		copy(destination[start:end], s)
		start = end
		end += len(s)
	}
	return destination
}

func (s *U8Slice) Flatten() {
	//	Flatten is a non-op for the U8Slice as they cannot contain nested elements
}

func (s U8Slice) equal(o U8Slice) (r bool) {
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

func (s U8Slice) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case *U8Slice:			r = o != nil && s.equal(*o)
	case U8Slice:			r = s.equal(o)
	case *[]uint8:			r = o != nil && s.equal(*o)
	case []uint8:			r = s.equal(o)
	}
	return
}

func (s U8Slice) Car() (h interface{}) {
	if s.Len() > 0 {
		h = s[0]
	}
	return
}

func (s U8Slice) Cdr() (t U8Slice) {
	if s.Len() > 1 {
		t = s[1:]
	}
	return
}

func (s *U8Slice) Rplaca(v interface{}) {
	switch {
	case s == nil:			*s = *U8List(v.(uint8))
	case s.Len() == 0:		*s = append(*s, v.(uint8))
	default:				(*s)[0] = v.(uint8)
	}
}

func (s *U8Slice) Rplacd(v interface{}) {
	if s == nil {
		*s = *U8List(v.(uint8))
	} else {
		ReplaceSlice := func(v U8Slice) {
			if l := len(v); l < cap(*s) {
				copy((*s)[1:], v)
			} else {
				l++
				n := make(U8Slice, l, l)
				copy(n, (*s)[:1])
				copy(n[1:], v)
				*s = n
			}
		}

		switch v := v.(type) {
		case *U8Slice:		ReplaceSlice(*v)
		case U8Slice:		ReplaceSlice(v)
		case *[]uint8:		ReplaceSlice(U8Slice(*v))
		case []uint8:		ReplaceSlice(U8Slice(v))
		case nil:			*s = (*s)[:1]
		default:			(*s)[1] = v.(uint8)
							*s = (*s)[:2]
		}
	}
}