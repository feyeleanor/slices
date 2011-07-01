package slices

import "fmt"

func IList(n... int) *ISlice {
	return (*ISlice)(&n)
}

type ISlice	[]int

func (s ISlice) Len() int							{ return len(s) }
func (s ISlice) Cap() int							{ return cap(s) }

func (s ISlice) At(i int) interface{}				{ return s[i] }
func (s ISlice) AtInt(i int) int					{ return s[i] }
func (s ISlice) Set(i int, v interface{})			{ s[i] = v.(int) }
func (s ISlice) SetInt(i, v int)					{ s[i] = v }
func (s ISlice) Clear(i int)						{ s[i] = 0 }
func (s ISlice) Swap(i, j int)						{ s[i], s[j] = s[j], s[i] }

func (s ISlice) Negate(i int)						{ s[i] = -s[i] }
func (s ISlice) Increment(i int)					{ s[i] += 1 }
func (s ISlice) Decrement(i int)					{ s[i] -= 1 }

func (s ISlice) Add(i, j int)						{ s[i] += s[j] }
func (s ISlice) Subtract(i, j int)					{ s[i] -= s[j] }
func (s ISlice) Multiply(i, j int)					{ s[i] *= s[j] }
func (s ISlice) Divide(i, j int)					{ s[i] /= s[j] }
func (s ISlice) Remainder(i, j int)					{ s[i] %= s[j] }

func (s ISlice) And(i, j int)						{ s[i] &= s[j] }
func (s ISlice) Or(i, j int)						{ s[i] |= s[j] }
func (s ISlice) Xor(i, j int)						{ s[i] ^= s[j] }
func (s ISlice) Invert(i int)						{ s[i] = ^s[i] }
func (s ISlice) ShiftLeft(i, j int)					{ s[i] <<= uint(s[j]) }
func (s ISlice) ShiftRight(i, j int)				{ s[i] >>= uint(s[j]) }

func (s ISlice) Less(i, j int) bool					{ return s[i] < s[j] }
func (s ISlice) AtLeast(i, j int) bool				{ return s[i] <= s[j] }
func (s ISlice) Same(i, j int) bool					{ return s[i] == s[j] }
func (s ISlice) AtMost(i, j int) bool				{ return s[i] >= s[j] }
func (s ISlice) More(i, j int) bool					{ return s[i] > s[j] }
func (s ISlice) ZeroLess(i int) bool				{ return 0 < s[i] }
func (s ISlice) ZeroAtLeast(i, j int) bool			{ return 0 <= s[j] }
func (s ISlice) ZeroSame(i int) bool				{ return 0 == s[i] }
func (s ISlice) ZeroAtMost(i, j int) bool			{ return 0 >= s[j] }
func (s ISlice) ZeroMore(i int) bool				{ return 0 > s[i] }

func (s ISlice) Compare(i, j int) (r int) {
	switch {
	case s[i] < s[j]:		r = IS_LESS_THAN
	case s[i] > s[j]:		r = IS_GREATER_THAN
	default:				r = IS_SAME_AS
	}
	return
}

func (s ISlice) ZeroCompare(i int) (r int) {
	switch {
	case 0 < s[i]:			r = IS_LESS_THAN
	case 0 > s[i]:			r = IS_GREATER_THAN
	default:				r = IS_SAME_AS
	}
	return
}

func (s *ISlice) Cut(i, j int) {
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

func (s *ISlice) Delete(i int) {
	a := *s
	n := len(a)
	if i > -1 && i < n {
		copy(a[i:n - 1], a[i + 1:n])
		*s = a[0 : n - 1]
	}
}

func (s ISlice) Each(f func(interface{})) {
	for _, v := range s {
		f(v)
	}
}

func (s ISlice) EachWithIndex(f func(int, interface{})) {
	for i, v := range s {
		f(i, v)
	}
}

func (s ISlice) EachWithKey(f func(key, value interface{})) {
	for i, v := range s {
		f(i, v)
	}
}

func (s ISlice) EachInt(f func(int)) {
	for _, v := range s {
		f(v)
	}
}

func (s ISlice) EachIntWithIndex(f func(int, int)) {
	for i, v := range s {
		f(i, v)
	}
}

func (s ISlice) EachIntWithKey(f func(interface{}, int)) {
	for i, v := range s {
		f(i, v)
	}
}

func (s ISlice) String() (t string) {
	for _, v := range s {
		if len(t) > 0 {
			t += " "
		}
		t += fmt.Sprintf("%v", v)
	}
	return fmt.Sprintf("(%v)", t)
}

func (s ISlice) BlockCopy(destination, source, count int) {
	end := source + count
	if end > len(s) {
		end = len(s)
	}
	copy(s[destination:], s[source:end])
}

func (s ISlice) BlockClear(start, count int) {
	copy(s[start:], make(ISlice, count, count))
}

func (s ISlice) Overwrite(offset int, container interface{}) {
	switch container := container.(type) {
	case ISlice:			copy(s[offset:], container)
	case []int:				copy(s[offset:], container)
	}
}

func (s *ISlice) Reallocate(length, capacity int) {
	switch {
	case length > capacity:		s.Reallocate(capacity, capacity)
	case capacity != cap(*s):	x := make(ISlice, length, capacity)
								copy(x, *s)
								*s = x
	default:					*s = (*s)[:length]
	}
}

func (s *ISlice) Extend(n int) {
	c := cap(*s)
	l := len(*s) + n
	if l > c {
		c = l
	}
	s.Reallocate(l, c)
}

func (s *ISlice) Expand(i, n int) {
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
		x := make(ISlice, l, c)
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

func (s ISlice) Reverse() {
	end := s.Len() - 1
	for i := 0; i < end; i++ {
		s[i], s[end] = s[end], s[i]
		end--
	}
}

func (s ISlice) Depth() int {
	return 0
}

func (s *ISlice) Append(v interface{}) {
	s.AppendInt(v.(int))
}

func (s *ISlice) AppendInt(v int) {
	*s = append(*s, v)
}

func (s *ISlice) AppendSlice(o ISlice) {
	*s = append(*s, o...)
}

func (s *ISlice) Prepend(v interface{}) {
	s.PrependInt(v.(int))
}

func (s *ISlice) PrependInt(v int) {
	l := s.Len() + 1
	n := make(ISlice, l, l)
	n[0] = v
	copy(n[1:], *s)
	*s = n
}

func (s *ISlice) PrependSlice(o ISlice) {
	l := s.Len() + o.Len()
	n := make(ISlice, l, l)
	copy(n, o)
	copy(n[o.Len():], *s)
	*s = n
}

func (s ISlice) Subslice(start, end int) interface{} {
	return s[start:end]
}

func (s ISlice) Repeat(count int) ISlice {
	length := len(s) * count
	capacity := cap(s)
	if capacity < length {
		capacity = length
	}
	destination := make(ISlice, length, capacity)
	for start, end := 0, len(s); count > 0; count-- {
		copy(destination[start:end], s)
		start = end
		end += len(s)
	}
	return destination
}

func (s *ISlice) Flatten() {
	//	Flatten is a non-op for the ISlice as they cannot contain nested elements
}

func (s ISlice) equal(o ISlice) (r bool) {
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

func (s ISlice) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case *ISlice:			r = o != nil && s.equal(*o)
	case ISlice:			r = s.equal(o)
	case *[]int:			r = o != nil && s.equal(*o)
	case []int:				r = s.equal(o)
	}
	return
}

func (s ISlice) Car() (h interface{}) {
	if s.Len() > 0 {
		h = s[0]
	}
	return
}

func (s ISlice) Cdr() (t ISlice) {
	if s.Len() > 1 {
		t = s[1:]
	}
	return
}

func (s *ISlice) Rplaca(v interface{}) {
	switch {
	case s == nil:			*s = *IList(v.(int))
	case s.Len() == 0:		*s = append(*s, v.(int))
	default:				(*s)[0] = v.(int)
	}
}

func (s *ISlice) Rplacd(v interface{}) {
	if s == nil {
		*s = *IList(v.(int))
	} else {
		ReplaceSlice := func(v ISlice) {
			if l := len(v); l < cap(*s) {
				copy((*s)[1:], v)
			} else {
				l++
				n := make(ISlice, l, l)
				copy(n, (*s)[:1])
				copy(n[1:], v)
				*s = n
			}
		}

		switch v := v.(type) {
		case *ISlice:		ReplaceSlice(*v)
		case ISlice:		ReplaceSlice(v)
		case *[]int:		ReplaceSlice(ISlice(*v))
		case []int:			ReplaceSlice(ISlice(v))
		case nil:			*s = (*s)[:1]
		default:			(*s)[1] = v.(int)
							*s = (*s)[:2]
		}
	}
}