package slices

import "fmt"

func U16List(n... uint16) *U16Slice {
	return (*U16Slice)(&n)
}

type U16Slice	[]uint16

func (s U16Slice) Len() int							{ return len(s) }
func (s U16Slice) Cap() int							{ return cap(s) }

func (s U16Slice) At(i int) interface{}				{ return s[i] }
func (s U16Slice) U16At(i int) uint16				{ return s[i] }
func (s U16Slice) Set(i int, v interface{})			{ s[i] = v.(uint16) }
func (s U16Slice) U16Set(i int, v uint16)			{ s[i] = v }
func (s U16Slice) Clear(i int)						{ s[i] = 0 }
func (s U16Slice) Swap(i, j int)					{ s[i], s[j] = s[j], s[i] }

func (s U16Slice) Negate(i int)						{ s[i] = -s[i] }
func (s U16Slice) Increment(i int)					{ s[i]++ }
func (s U16Slice) Decrement(i int)					{ s[i]-- }

func (s U16Slice) Add(i, j int)						{ s[i] += s[j] }
func (s U16Slice) Subtract(i, j int)				{ s[i] -= s[j] }
func (s U16Slice) Multiply(i, j int)				{ s[i] *= s[j] }
func (s U16Slice) Divide(i, j int)					{ s[i] /= s[j] }
func (s U16Slice) Remainder(i, j int)				{ s[i] %= s[j] }

func (s U16Slice) And(i, j int)						{ s[i] &= s[j] }
func (s U16Slice) Or(i, j int)						{ s[i] |= s[j] }
func (s U16Slice) Xor(i, j int)						{ s[i] ^= s[j] }
func (s U16Slice) Invert(i int)						{ s[i] = ^s[i] }
func (s U16Slice) ShiftLeft(i, j int)				{ s[i] <<= s[j] }
func (s U16Slice) ShiftRight(i, j int)				{ s[i] >>= s[j] }

func (s U16Slice) Less(i, j int) bool				{ return s[i] < s[j] }
func (s U16Slice) AtLeast(i, j int) bool			{ return s[i] <= s[j] }
func (s U16Slice) Same(i, j int) bool				{ return s[i] == s[j] }
func (s U16Slice) AtMost(i, j int) bool				{ return s[i] >= s[j] }
func (s U16Slice) More(i, j int) bool				{ return s[i] > s[j] }
func (s U16Slice) ZeroLess(i int) bool				{ return 0 < s[i] }
func (s U16Slice) ZeroAtLeast(i, j int) bool		{ return true }
func (s U16Slice) ZeroSame(i int) bool				{ return 0 == s[i] }
func (s U16Slice) ZeroAtMost(i, j int) bool			{ return 0 == s[j] }
func (s U16Slice) ZeroMore(i int) bool				{ return false }

func (s U16Slice) Compare(i, j int) (r int) {
	switch {
	case s[i] < s[j]:		r = IS_LESS_THAN
	case s[i] > s[j]:		r = IS_GREATER_THAN
	default:				r = IS_SAME_AS
	}
	return
}

func (s U16Slice) ZeroCompare(i int) (r int) {
	switch {
	case 0 < s[i]:			r = IS_LESS_THAN
	default:				r = IS_SAME_AS
	}
	return
}

func (s *U16Slice) Cut(i, j int) {
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

func (s *U16Slice) Delete(i int) {
	a := *s
	n := len(a)
	if i > -1 && i < n {
		copy(a[i:n - 1], a[i + 1:n])
		*s = a[0 : n - 1]
	}
}

func (s U16Slice) Each(f func(interface{})) {
	for _, v := range s {
		f(v)
	}
}

func (s U16Slice) EachWithIndex(f func(int, interface{})) {
	for i, v := range s {
		f(i, v)
	}
}

func (s U16Slice) EachWithKey(f func(key, value interface{})) {
	for i, v := range s {
		f(i, v)
	}
}

func (s U16Slice) U16Each(f func(uint16)) {
	for _, v := range s {
		f(v)
	}
}

func (s U16Slice) U16EachWithIndex(f func(int, uint16)) {
	for i, v := range s {
		f(i, v)
	}
}

func (s U16Slice) U16EachWithKey(f func(interface{}, uint16)) {
	for i, v := range s {
		f(i, v)
	}
}

func (s U16Slice) String() (t string) {
	for _, v := range s {
		if len(t) > 0 {
			t += " "
		}
		t += fmt.Sprintf("%v", v)
	}
	return fmt.Sprintf("(%v)", t)
}

func (s U16Slice) BlockCopy(destination, source, count int) {
	end := source + count
	if end > len(s) {
		end = len(s)
	}
	copy(s[destination:], s[source:end])
}

func (s U16Slice) BlockClear(start, count int) {
	copy(s[start:], make(U16Slice, count, count))
}

func (s U16Slice) Overwrite(offset int, container interface{}) {
	switch container := container.(type) {
	case U16Slice:			copy(s[offset:], container)
	case []uint16:			copy(s[offset:], container)
	}
}

func (s *U16Slice) Reallocate(length, capacity int) {
	switch {
	case length > capacity:		s.Reallocate(capacity, capacity)
	case capacity != cap(*s):	x := make(U16Slice, length, capacity)
								copy(x, *s)
								*s = x
	default:					*s = (*s)[:length]
	}
}

func (s *U16Slice) Extend(n int) {
	c := cap(*s)
	l := len(*s) + n
	if l > c {
		c = l
	}
	s.Reallocate(l, c)
}

func (s *U16Slice) Expand(i, n int) {
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
		x := make(U16Slice, l, c)
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

func (s U16Slice) Reverse() {
	end := s.Len() - 1
	for i := 0; i < end; i++ {
		s[i], s[end] = s[end], s[i]
		end--
	}
}

func (s U16Slice) Depth() int {
	return 0
}

func (s *U16Slice) Append(v interface{}) {
	s.U16Append(v.(uint16))
}

func (s *U16Slice) U16Append(v uint16) {
	*s = append(*s, v)
}

func (s *U16Slice) AppendSlice(o U16Slice) {
	*s = append(*s, o...)
}

func (s *U16Slice) Prepend(v interface{}) {
	s.U16Prepend(v.(uint16))
}

func (s *U16Slice) U16Prepend(v uint16) {
	l := s.Len() + 1
	n := make(U16Slice, l, l)
	n[0] = v
	copy(n[1:], *s)
	*s = n
}

func (s *U16Slice) PrependSlice(o U16Slice) {
	l := s.Len() + o.Len()
	n := make(U16Slice, l, l)
	copy(n, o)
	copy(n[o.Len():], *s)
	*s = n
}

func (s U16Slice) Repeat(count int) U16Slice {
	length := len(s) * count
	capacity := cap(s)
	if capacity < length {
		capacity = length
	}
	destination := make(U16Slice, length, capacity)
	for start, end := 0, len(s); count > 0; count-- {
		copy(destination[start:end], s)
		start = end
		end += len(s)
	}
	return destination
}

func (s *U16Slice) Flatten() {
	//	Flatten is a non-op for the U16Slice as they cannot contain nested elements
}

func (s U16Slice) equal(o U16Slice) (r bool) {
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

func (s U16Slice) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case *U16Slice:			r = o != nil && s.equal(*o)
	case U16Slice:			r = s.equal(o)
	case *[]uint16:			r = o != nil && s.equal(*o)
	case []uint16:			r = s.equal(o)
	}
	return
}

func (s U16Slice) Car() (h interface{}) {
	if s.Len() > 0 {
		h = s[0]
	}
	return
}

func (s U16Slice) Cdr() (t U16Slice) {
	if s.Len() > 1 {
		t = s[1:]
	}
	return
}

func (s *U16Slice) Rplaca(v interface{}) {
	switch {
	case s == nil:			*s = *U16List(v.(uint16))
	case s.Len() == 0:		*s = append(*s, v.(uint16))
	default:				(*s)[0] = v.(uint16)
	}
}

func (s *U16Slice) Rplacd(v interface{}) {
	if s == nil {
		*s = *U16List(v.(uint16))
	} else {
		ReplaceSlice := func(v U16Slice) {
			if l := len(v); l < cap(*s) {
				copy((*s)[1:], v)
			} else {
				l++
				n := make(U16Slice, l, l)
				copy(n, (*s)[:1])
				copy(n[1:], v)
				*s = n
			}
		}

		switch v := v.(type) {
		case *U16Slice:		ReplaceSlice(*v)
		case U16Slice:		ReplaceSlice(v)
		case *[]uint16:		ReplaceSlice(U16Slice(*v))
		case []uint16:		ReplaceSlice(U16Slice(v))
		case nil:			*s = (*s)[:1]
		default:			(*s)[1] = v.(uint16)
							*s = (*s)[:2]
		}
	}
}