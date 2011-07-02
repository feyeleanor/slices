package slices

import "fmt"

func I32List(n... int32) *I32Slice {
	return (*I32Slice)(&n)
}

type I32Slice	[]int32

func (s I32Slice) Len() int							{ return len(s) }
func (s I32Slice) Cap() int							{ return cap(s) }

func (s I32Slice) At(i int) interface{}				{ return s[i] }
func (s I32Slice) I32At(i int) int32				{ return s[i] }
func (s I32Slice) Set(i int, v interface{})			{ s[i] = v.(int32) }
func (s I32Slice) I32Set(i int, v int32)			{ s[i] = v }
func (s I32Slice) Clear(i int)						{ s[i] = 0 }
func (s I32Slice) Swap(i, j int)					{ s[i], s[j] = s[j], s[i] }

func (s I32Slice) Negate(i int)						{ s[i] = -s[i] }
func (s I32Slice) Increment(i int)					{ s[i]++ }
func (s I32Slice) Decrement(i int)					{ s[i]-- }

func (s I32Slice) Add(i, j int)						{ s[i] += s[j] }
func (s I32Slice) Subtract(i, j int)				{ s[i] -= s[j] }
func (s I32Slice) Multiply(i, j int)				{ s[i] *= s[j] }
func (s I32Slice) Divide(i, j int)					{ s[i] /= s[j] }
func (s I32Slice) Remainder(i, j int)				{ s[i] %= s[j] }

func (s I32Slice) And(i, j int)						{ s[i] &= s[j] }
func (s I32Slice) Or(i, j int)						{ s[i] |= s[j] }
func (s I32Slice) Xor(i, j int)						{ s[i] ^= s[j] }
func (s I32Slice) Invert(i int)						{ s[i] = ^s[i] }
func (s I32Slice) ShiftLeft(i, j int)				{ s[i] <<= uint(s[j]) }
func (s I32Slice) ShiftRight(i, j int)				{ s[i] >>= uint(s[j]) }

func (s I32Slice) Less(i, j int) bool				{ return s[i] < s[j] }
func (s I32Slice) AtLeast(i, j int) bool			{ return s[i] <= s[j] }
func (s I32Slice) Same(i, j int) bool				{ return s[i] == s[j] }
func (s I32Slice) AtMost(i, j int) bool				{ return s[i] >= s[j] }
func (s I32Slice) More(i, j int) bool				{ return s[i] > s[j] }
func (s I32Slice) ZeroLess(i int) bool				{ return 0 < s[i] }
func (s I32Slice) ZeroAtLeast(i, j int) bool		{ return 0 <= s[j] }
func (s I32Slice) ZeroSame(i int) bool				{ return 0 == s[i] }
func (s I32Slice) ZeroAtMost(i, j int) bool			{ return 0 >= s[j] }
func (s I32Slice) ZeroMore(i int) bool				{ return 0 > s[i] }

func (s *I32Slice) RestrictTo(i, j int)				{ *s = (*s)[i:j] }

func (s I32Slice) Compare(i, j int) (r int) {
	switch {
	case s[i] < s[j]:		r = IS_LESS_THAN
	case s[i] > s[j]:		r = IS_GREATER_THAN
	default:				r = IS_SAME_AS
	}
	return
}

func (s I32Slice) ZeroCompare(i int) (r int) {
	switch {
	case 0 < s[i]:			r = IS_LESS_THAN
	case 0 > s[i]:			r = IS_GREATER_THAN
	default:				r = IS_SAME_AS
	}
	return
}

func (s *I32Slice) Cut(i, j int) {
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

func (s *I32Slice) Trim(i, j int) {
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

func (s *I32Slice) Delete(i int) {
	a := *s
	n := len(a)
	if i > -1 && i < n {
		copy(a[i:n - 1], a[i + 1:n])
		*s = a[0 : n - 1]
	}
}

func (s I32Slice) Each(f func(interface{})) {
	for _, v := range s {
		f(v)
	}
}

func (s I32Slice) EachWithIndex(f func(int, interface{})) {
	for i, v := range s {
		f(i, v)
	}
}

func (s I32Slice) EachWithKey(f func(key, value interface{})) {
	for i, v := range s {
		f(i, v)
	}
}

func (s I32Slice) I32Each(f func(int32)) {
	for _, v := range s {
		f(v)
	}
}

func (s I32Slice) I32EachWithIndex(f func(int, int32)) {
	for i, v := range s {
		f(i, v)
	}
}

func (s I32Slice) I32EachWithKey(f func(interface{}, int32)) {
	for i, v := range s {
		f(i, v)
	}
}

func (s I32Slice) String() (t string) {
	for _, v := range s {
		if len(t) > 0 {
			t += " "
		}
		t += fmt.Sprintf("%v", v)
	}
	return fmt.Sprintf("(%v)", t)
}

func (s I32Slice) BlockCopy(destination, source, count int) {
	end := source + count
	if end > len(s) {
		end = len(s)
	}
	copy(s[destination:], s[source:end])
}

func (s I32Slice) BlockClear(start, count int) {
	copy(s[start:], make(I32Slice, count, count))
}

func (s I32Slice) Overwrite(offset int, container interface{}) {
	switch container := container.(type) {
	case I32Slice:			copy(s[offset:], container)
	case []int32:			copy(s[offset:], container)
	}
}

func (s *I32Slice) Reallocate(length, capacity int) {
	switch {
	case length > capacity:		s.Reallocate(capacity, capacity)
	case capacity != cap(*s):	x := make(I32Slice, length, capacity)
								copy(x, *s)
								*s = x
	default:					*s = (*s)[:length]
	}
}

func (s *I32Slice) Extend(n int) {
	c := cap(*s)
	l := len(*s) + n
	if l > c {
		c = l
	}
	s.Reallocate(l, c)
}

func (s *I32Slice) Expand(i, n int) {
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
		x := make(I32Slice, l, c)
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

func (s I32Slice) Reverse() {
	end := s.Len() - 1
	for i := 0; i < end; i++ {
		s[i], s[end] = s[end], s[i]
		end--
	}
}

func (s I32Slice) Depth() int {
	return 0
}

func (s *I32Slice) Append(v interface{}) {
	s.I32Append(v.(int32))
}

func (s *I32Slice) I32Append(v int32) {
	*s = append(*s, v)
}

func (s *I32Slice) AppendSlice(o I32Slice) {
	*s = append(*s, o...)
}

func (s *I32Slice) Prepend(v interface{}) {
	s.I32Prepend(v.(int32))
}

func (s *I32Slice) I32Prepend(v int32) {
	l := s.Len() + 1
	n := make(I32Slice, l, l)
	n[0] = v
	copy(n[1:], *s)
	*s = n
}

func (s *I32Slice) PrependSlice(o I32Slice) {
	l := s.Len() + o.Len()
	n := make(I32Slice, l, l)
	copy(n, o)
	copy(n[o.Len():], *s)
	*s = n
}

func (s I32Slice) Repeat(count int) I32Slice {
	length := len(s) * count
	capacity := cap(s)
	if capacity < length {
		capacity = length
	}
	destination := make(I32Slice, length, capacity)
	for start, end := 0, len(s); count > 0; count-- {
		copy(destination[start:end], s)
		start = end
		end += len(s)
	}
	return destination
}

func (s *I32Slice) Flatten() {
	//	Flatten is a non-op for the I32Slice as they cannot contain nested elements
}

func (s I32Slice) equal(o I32Slice) (r bool) {
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

func (s I32Slice) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case *I32Slice:			r = o != nil && s.equal(*o)
	case I32Slice:			r = s.equal(o)
	case *[]int32:			r = o != nil && s.equal(*o)
	case []int32:			r = s.equal(o)
	}
	return
}

func (s I32Slice) Car() (h interface{}) {
	if s.Len() > 0 {
		h = s[0]
	}
	return
}

func (s I32Slice) Cdr() (t I32Slice) {
	if s.Len() > 1 {
		t = s[1:]
	}
	return
}

func (s *I32Slice) Rplaca(v interface{}) {
	switch {
	case s == nil:			*s = *I32List(v.(int32))
	case s.Len() == 0:		*s = append(*s, v.(int32))
	default:				(*s)[0] = v.(int32)
	}
}

func (s *I32Slice) Rplacd(v interface{}) {
	if s == nil {
		*s = *I32List(v.(int32))
	} else {
		ReplaceSlice := func(v I32Slice) {
			if l := len(v); l < cap(*s) {
				copy((*s)[1:], v)
			} else {
				l++
				n := make(I32Slice, l, l)
				copy(n, (*s)[:1])
				copy(n[1:], v)
				*s = n
			}
		}

		switch v := v.(type) {
		case *I32Slice:		ReplaceSlice(*v)
		case I32Slice:		ReplaceSlice(v)
		case *[]int32:		ReplaceSlice(I32Slice(*v))
		case []int32:		ReplaceSlice(I32Slice(v))
		case nil:			*s = (*s)[:1]
		default:			(*s)[1] = v.(int32)
							*s = (*s)[:2]
		}
	}
}