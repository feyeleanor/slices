package slices

import "fmt"

func I16List(n... int16) *I16Slice {
	return (*I16Slice)(&n)
}

type I16Slice	[]int16

func (s I16Slice) Len() int							{ return len(s) }
func (s I16Slice) Cap() int							{ return cap(s) }

func (s I16Slice) At(i int) interface{}				{ return s[i] }
func (s I16Slice) I16At(i int) int16				{ return s[i] }
func (s I16Slice) Set(i int, v interface{})			{ s[i] = v.(int16) }
func (s I16Slice) I16Set(i int, v int16)			{ s[i] = v }
func (s I16Slice) Clear(i int)						{ s[i] = 0 }
func (s I16Slice) Swap(i, j int)					{ s[i], s[j] = s[j], s[i] }

func (s I16Slice) Negate(i int)						{ s[i] = -s[i] }
func (s I16Slice) Increment(i int)					{ s[i]++ }
func (s I16Slice) Decrement(i int)					{ s[i]-- }

func (s I16Slice) Add(i, j int)						{ s[i] += s[j] }
func (s I16Slice) Subtract(i, j int)				{ s[i] -= s[j] }
func (s I16Slice) Multiply(i, j int)				{ s[i] *= s[j] }
func (s I16Slice) Divide(i, j int)					{ s[i] /= s[j] }
func (s I16Slice) Remainder(i, j int)				{ s[i] %= s[j] }

func (s I16Slice) And(i, j int)						{ s[i] &= s[j] }
func (s I16Slice) Or(i, j int)						{ s[i] |= s[j] }
func (s I16Slice) Xor(i, j int)						{ s[i] ^= s[j] }
func (s I16Slice) Invert(i int)						{ s[i] = ^s[i] }
func (s I16Slice) ShiftLeft(i, j int)				{ s[i] <<= uint(s[j]) }
func (s I16Slice) ShiftRight(i, j int)				{ s[i] >>= uint(s[j]) }

func (s I16Slice) Less(i, j int) bool				{ return s[i] < s[j] }
func (s I16Slice) AtLeast(i, j int) bool			{ return s[i] <= s[j] }
func (s I16Slice) Same(i, j int) bool				{ return s[i] == s[j] }
func (s I16Slice) AtMost(i, j int) bool				{ return s[i] >= s[j] }
func (s I16Slice) More(i, j int) bool				{ return s[i] > s[j] }
func (s I16Slice) ZeroLess(i int) bool				{ return 0 < s[i] }
func (s I16Slice) ZeroAtLeast(i, j int) bool		{ return 0 <= s[j] }
func (s I16Slice) ZeroSame(i int) bool				{ return 0 == s[i] }
func (s I16Slice) ZeroAtMost(i, j int) bool			{ return 0 >= s[j] }
func (s I16Slice) ZeroMore(i int) bool				{ return 0 > s[i] }

func (s *I16Slice) RestrictTo(i, j int)				{ *s = (*s)[i:j] }

func (s I16Slice) Compare(i, j int) (r int) {
	switch {
	case s[i] < s[j]:		r = IS_LESS_THAN
	case s[i] > s[j]:		r = IS_GREATER_THAN
	default:				r = IS_SAME_AS
	}
	return
}

func (s I16Slice) ZeroCompare(i int) (r int) {
	switch {
	case 0 < s[i]:			r = IS_LESS_THAN
	case 0 > s[i]:			r = IS_GREATER_THAN
	default:				r = IS_SAME_AS
	}
	return
}

func (s *I16Slice) Cut(i, j int) {
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

func (s *I16Slice) Trim(i, j int) {
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

func (s *I16Slice) Delete(i int) {
	a := *s
	n := len(a)
	if i > -1 && i < n {
		copy(a[i:n - 1], a[i + 1:n])
		*s = a[0 : n - 1]
	}
}

func (s I16Slice) Each(f func(interface{})) {
	for _, v := range s {
		f(v)
	}
}

func (s I16Slice) EachWithIndex(f func(int, interface{})) {
	for i, v := range s {
		f(i, v)
	}
}

func (s I16Slice) EachWithKey(f func(key, value interface{})) {
	for i, v := range s {
		f(i, v)
	}
}

func (s I16Slice) I16Each(f func(int16)) {
	for _, v := range s {
		f(v)
	}
}

func (s I16Slice) I16EachWithIndex(f func(int, int16)) {
	for i, v := range s {
		f(i, v)
	}
}

func (s I16Slice) I16EachWithKey(f func(interface{}, int16)) {
	for i, v := range s {
		f(i, v)
	}
}

func (s I16Slice) String() (t string) {
	for _, v := range s {
		if len(t) > 0 {
			t += " "
		}
		t += fmt.Sprintf("%v", v)
	}
	return fmt.Sprintf("(%v)", t)
}

func (s I16Slice) BlockCopy(destination, source, count int) {
	end := source + count
	if end > len(s) {
		end = len(s)
	}
	copy(s[destination:], s[source:end])
}

func (s I16Slice) BlockClear(start, count int) {
	copy(s[start:], make(I16Slice, count, count))
}

func (s I16Slice) Overwrite(offset int, container interface{}) {
	switch container := container.(type) {
	case I16Slice:			copy(s[offset:], container)
	case []int16:			copy(s[offset:], container)
	}
}

func (s *I16Slice) Reallocate(length, capacity int) {
	switch {
	case length > capacity:		s.Reallocate(capacity, capacity)
	case capacity != cap(*s):	x := make(I16Slice, length, capacity)
								copy(x, *s)
								*s = x
	default:					*s = (*s)[:length]
	}
}

func (s *I16Slice) Extend(n int) {
	c := cap(*s)
	l := len(*s) + n
	if l > c {
		c = l
	}
	s.Reallocate(l, c)
}

func (s *I16Slice) Expand(i, n int) {
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
		x := make(I16Slice, l, c)
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

func (s I16Slice) Reverse() {
	end := s.Len() - 1
	for i := 0; i < end; i++ {
		s[i], s[end] = s[end], s[i]
		end--
	}
}

func (s I16Slice) Depth() int {
	return 0
}

func (s *I16Slice) Append(v interface{}) {
	s.I16Append(v.(int16))
}

func (s *I16Slice) I16Append(v int16) {
	*s = append(*s, v)
}

func (s *I16Slice) AppendSlice(o I16Slice) {
	*s = append(*s, o...)
}

func (s *I16Slice) Prepend(v interface{}) {
	s.I16Prepend(v.(int16))
}

func (s *I16Slice) I16Prepend(v int16) {
	l := s.Len() + 1
	n := make(I16Slice, l, l)
	n[0] = v
	copy(n[1:], *s)
	*s = n
}

func (s *I16Slice) PrependSlice(o I16Slice) {
	l := s.Len() + o.Len()
	n := make(I16Slice, l, l)
	copy(n, o)
	copy(n[o.Len():], *s)
	*s = n
}

func (s I16Slice) Repeat(count int) I16Slice {
	length := len(s) * count
	capacity := cap(s)
	if capacity < length {
		capacity = length
	}
	destination := make(I16Slice, length, capacity)
	for start, end := 0, len(s); count > 0; count-- {
		copy(destination[start:end], s)
		start = end
		end += len(s)
	}
	return destination
}

func (s *I16Slice) Flatten() {
	//	Flatten is a non-op for the I16Slice as they cannot contain nested elements
}

func (s I16Slice) equal(o I16Slice) (r bool) {
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

func (s I16Slice) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case *I16Slice:			r = o != nil && s.equal(*o)
	case I16Slice:			r = s.equal(o)
	case *[]int16:			r = o != nil && s.equal(*o)
	case []int16:			r = s.equal(o)
	}
	return
}

func (s I16Slice) Car() (h interface{}) {
	if s.Len() > 0 {
		h = s[0]
	}
	return
}

func (s I16Slice) Cdr() (t I16Slice) {
	if s.Len() > 1 {
		t = s[1:]
	}
	return
}

func (s *I16Slice) Rplaca(v interface{}) {
	switch {
	case s == nil:			*s = *I16List(v.(int16))
	case s.Len() == 0:		*s = append(*s, v.(int16))
	default:				(*s)[0] = v.(int16)
	}
}

func (s *I16Slice) Rplacd(v interface{}) {
	if s == nil {
		*s = *I16List(v.(int16))
	} else {
		ReplaceSlice := func(v I16Slice) {
			if l := len(v); l < cap(*s) {
				copy((*s)[1:], v)
				*s = (*s)[0:l + 1]
			} else {
				l++
				n := make(I16Slice, l, l)
				copy(n, (*s)[:1])
				copy(n[1:], v)
				*s = n
			}
		}

		switch v := v.(type) {
		case *I16Slice:		ReplaceSlice(*v)
		case I16Slice:		ReplaceSlice(v)
		case *[]int16:		ReplaceSlice(I16Slice(*v))
		case []int16:		ReplaceSlice(I16Slice(v))
		case nil:			*s = (*s)[:1]
		default:			(*s)[1] = v.(int16)
							*s = (*s)[:2]
		}
	}
}