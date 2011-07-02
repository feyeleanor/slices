package slices

import "fmt"

func UList(n... uint) *USlice {
	return (*USlice)(&n)
}

type USlice	[]uint

func (s USlice) Len() int							{ return len(s) }
func (s USlice) Cap() int							{ return cap(s) }

func (s USlice) At(i int) interface{}				{ return s[i] }
func (s USlice) UAt(i int) uint						{ return s[i] }
func (s USlice) Set(i int, v interface{})			{ s[i] = v.(uint) }
func (s USlice) USet(i int, v uint)					{ s[i] = v }
func (s USlice) Clear(i int)						{ s[i] = 0 }
func (s USlice) Swap(i, j int)						{ s[i], s[j] = s[j], s[i] }

func (s USlice) Negate(i int)						{ s[i] = -s[i] }
func (s USlice) Increment(i int)					{ s[i]++ }
func (s USlice) Decrement(i int)					{ s[i]-- }

func (s USlice) Add(i, j int)						{ s[i] += s[j] }
func (s USlice) Subtract(i, j int)					{ s[i] -= s[j] }
func (s USlice) Multiply(i, j int)					{ s[i] *= s[j] }
func (s USlice) Divide(i, j int)					{ s[i] /= s[j] }
func (s USlice) Remainder(i, j int)					{ s[i] %= s[j] }

func (s USlice) And(i, j int)						{ s[i] &= s[j] }
func (s USlice) Or(i, j int)						{ s[i] |= s[j] }
func (s USlice) Xor(i, j int)						{ s[i] ^= s[j] }
func (s USlice) Invert(i int)						{ s[i] = ^s[i] }
func (s USlice) ShiftLeft(i, j int)					{ s[i] <<= s[j] }
func (s USlice) ShiftRight(i, j int)				{ s[i] >>= s[j] }

func (s USlice) Less(i, j int) bool					{ return s[i] < s[j] }
func (s USlice) AtLeast(i, j int) bool				{ return s[i] <= s[j] }
func (s USlice) Same(i, j int) bool					{ return s[i] == s[j] }
func (s USlice) AtMost(i, j int) bool				{ return s[i] >= s[j] }
func (s USlice) More(i, j int) bool					{ return s[i] > s[j] }
func (s USlice) ZeroLess(i int) bool				{ return 0 < s[i] }
func (s USlice) ZeroAtLeast(i, j int) bool			{ return true }
func (s USlice) ZeroSame(i int) bool				{ return 0 == s[i] }
func (s USlice) ZeroAtMost(i, j int) bool			{ return 0 == s[j] }
func (s USlice) ZeroMore(i int) bool				{ return false }

func (s USlice) Compare(i, j int) (r int) {
	switch {
	case s[i] < s[j]:		r = IS_LESS_THAN
	case s[i] > s[j]:		r = IS_GREATER_THAN
	default:				r = IS_SAME_AS
	}
	return
}

func (s USlice) ZeroCompare(i int) (r int) {
	switch {
	case 0 < s[i]:			r = IS_LESS_THAN
	default:				r = IS_SAME_AS
	}
	return
}

func (s *USlice) Cut(i, j int) {
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

func (s *USlice) Trim(i, j int) {
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

func (s *USlice) Delete(i int) {
	a := *s
	n := len(a)
	if i > -1 && i < n {
		copy(a[i:n - 1], a[i + 1:n])
		*s = a[0 : n - 1]
	}
}

func (s USlice) Each(f func(interface{})) {
	for _, v := range s {
		f(v)
	}
}

func (s USlice) EachWithIndex(f func(int, interface{})) {
	for i, v := range s {
		f(i, v)
	}
}

func (s USlice) EachWithKey(f func(key, value interface{})) {
	for i, v := range s {
		f(i, v)
	}
}

func (s USlice) UEach(f func(uint)) {
	for _, v := range s {
		f(v)
	}
}

func (s USlice) UEachWithIndex(f func(int, uint)) {
	for i, v := range s {
		f(i, v)
	}
}

func (s USlice) UEachWithKey(f func(interface{}, uint)) {
	for i, v := range s {
		f(i, v)
	}
}

func (s USlice) String() (t string) {
	for _, v := range s {
		if len(t) > 0 {
			t += " "
		}
		t += fmt.Sprintf("%v", v)
	}
	return fmt.Sprintf("(%v)", t)
}

func (s USlice) BlockCopy(destination, source, count int) {
	end := source + count
	if end > len(s) {
		end = len(s)
	}
	copy(s[destination:], s[source:end])
}

func (s USlice) BlockClear(start, count int) {
	copy(s[start:], make(USlice, count, count))
}

func (s USlice) Overwrite(offset int, container interface{}) {
	switch container := container.(type) {
	case USlice:			copy(s[offset:], container)
	case []uint:			copy(s[offset:], container)
	}
}

func (s *USlice) Reallocate(length, capacity int) {
	switch {
	case length > capacity:		s.Reallocate(capacity, capacity)
	case capacity != cap(*s):	x := make(USlice, length, capacity)
								copy(x, *s)
								*s = x
	default:					*s = (*s)[:length]
	}
}

func (s *USlice) Extend(n int) {
	c := cap(*s)
	l := len(*s) + n
	if l > c {
		c = l
	}
	s.Reallocate(l, c)
}

func (s *USlice) Expand(i, n int) {
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
		x := make(USlice, l, c)
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

func (s USlice) Reverse() {
	end := s.Len() - 1
	for i := 0; i < end; i++ {
		s[i], s[end] = s[end], s[i]
		end--
	}
}

func (s USlice) Depth() int {
	return 0
}

func (s *USlice) Append(v interface{}) {
	s.UAppend(v.(uint))
}

func (s *USlice) UAppend(v uint) {
	*s = append(*s, v)
}

func (s *USlice) AppendSlice(o USlice) {
	*s = append(*s, o...)
}

func (s *USlice) Prepend(v interface{}) {
	s.UPrepend(v.(uint))
}

func (s *USlice) UPrepend(v uint) {
	l := s.Len() + 1
	n := make(USlice, l, l)
	n[0] = v
	copy(n[1:], *s)
	*s = n
}

func (s *USlice) PrependSlice(o USlice) {
	l := s.Len() + o.Len()
	n := make(USlice, l, l)
	copy(n, o)
	copy(n[o.Len():], *s)
	*s = n
}

func (s USlice) Repeat(count int) USlice {
	length := len(s) * count
	capacity := cap(s)
	if capacity < length {
		capacity = length
	}
	destination := make(USlice, length, capacity)
	for start, end := 0, len(s); count > 0; count-- {
		copy(destination[start:end], s)
		start = end
		end += len(s)
	}
	return destination
}

func (s *USlice) Flatten() {
	//	Flatten is a non-op for the USlice as they cannot contain nested elements
}

func (s USlice) equal(o USlice) (r bool) {
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

func (s USlice) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case *USlice:			r = o != nil && s.equal(*o)
	case USlice:			r = s.equal(o)
	case *[]uint:			r = o != nil && s.equal(*o)
	case []uint:			r = s.equal(o)
	}
	return
}

func (s USlice) Car() (h interface{}) {
	if s.Len() > 0 {
		h = s[0]
	}
	return
}

func (s USlice) Cdr() (t USlice) {
	if s.Len() > 1 {
		t = s[1:]
	}
	return
}

func (s *USlice) Rplaca(v interface{}) {
	switch {
	case s == nil:			*s = *UList(v.(uint))
	case s.Len() == 0:		*s = append(*s, v.(uint))
	default:				(*s)[0] = v.(uint)
	}
}

func (s *USlice) Rplacd(v interface{}) {
	if s == nil {
		*s = *UList(v.(uint))
	} else {
		ReplaceSlice := func(v USlice) {
			if l := len(v); l < cap(*s) {
				copy((*s)[1:], v)
			} else {
				l++
				n := make(USlice, l, l)
				copy(n, (*s)[:1])
				copy(n[1:], v)
				*s = n
			}
		}

		switch v := v.(type) {
		case *USlice:		ReplaceSlice(*v)
		case USlice:		ReplaceSlice(v)
		case *[]uint:		ReplaceSlice(USlice(*v))
		case []uint:		ReplaceSlice(USlice(v))
		case nil:			*s = (*s)[:1]
		default:			(*s)[1] = v.(uint)
							*s = (*s)[:2]
		}
	}
}