package slices

import "fmt"
import "sort"

func I8List(n... int8) *I8Slice {
	return (*I8Slice)(&n)
}

type I8Slice	[]int8

func (s I8Slice) Len() int							{ return len(s) }
func (s I8Slice) Cap() int							{ return cap(s) }

func (s I8Slice) At(i int) interface{}				{ return s[i] }
func (s I8Slice) I8At(i int) int8					{ return s[i] }
func (s I8Slice) Set(i int, v interface{})			{ s[i] = v.(int8) }
func (s I8Slice) I8Set(i int, v int8)				{ s[i] = v }
func (s I8Slice) Clear(i int)						{ s[i] = 0 }
func (s I8Slice) Swap(i, j int)						{ s[i], s[j] = s[j], s[i] }

func (s I8Slice) Negate(i int)						{ s[i] = -s[i] }
func (s I8Slice) Increment(i int)					{ s[i]++ }
func (s I8Slice) Decrement(i int)					{ s[i]-- }

func (s I8Slice) Add(i, j int)						{ s[i] += s[j] }
func (s I8Slice) Subtract(i, j int)					{ s[i] -= s[j] }
func (s I8Slice) Multiply(i, j int)					{ s[i] *= s[j] }
func (s I8Slice) Divide(i, j int)					{ s[i] /= s[j] }
func (s I8Slice) Remainder(i, j int)				{ s[i] %= s[j] }

func (s I8Slice) And(i, j int)						{ s[i] &= s[j] }
func (s I8Slice) Or(i, j int)						{ s[i] |= s[j] }
func (s I8Slice) Xor(i, j int)						{ s[i] ^= s[j] }
func (s I8Slice) Invert(i int)						{ s[i] = ^s[i] }
func (s I8Slice) ShiftLeft(i, j int)				{ s[i] <<= uint(s[j]) }
func (s I8Slice) ShiftRight(i, j int)				{ s[i] >>= uint(s[j]) }

func (s I8Slice) Less(i, j int) bool				{ return s[i] < s[j] }
func (s I8Slice) AtLeast(i, j int) bool				{ return s[i] <= s[j] }
func (s I8Slice) Same(i, j int) bool				{ return s[i] == s[j] }
func (s I8Slice) AtMost(i, j int) bool				{ return s[i] >= s[j] }
func (s I8Slice) More(i, j int) bool				{ return s[i] > s[j] }
func (s I8Slice) ZeroLess(i int) bool				{ return 0 < s[i] }
func (s I8Slice) ZeroAtLeast(i, j int) bool			{ return 0 <= s[j] }
func (s I8Slice) ZeroSame(i int) bool				{ return 0 == s[i] }
func (s I8Slice) ZeroAtMost(i, j int) bool			{ return 0 >= s[j] }
func (s I8Slice) ZeroMore(i int) bool				{ return 0 > s[i] }

func (s I8Slice) Sort()								{ sort.Sort(s) }

func (s *I8Slice) RestrictTo(i, j int)				{ *s = (*s)[i:j] }

func (s I8Slice) Compare(i, j int) (r int) {
	switch {
	case s[i] < s[j]:		r = IS_LESS_THAN
	case s[i] > s[j]:		r = IS_GREATER_THAN
	default:				r = IS_SAME_AS
	}
	return
}

func (s I8Slice) ZeroCompare(i int) (r int) {
	switch {
	case 0 < s[i]:			r = IS_LESS_THAN
	case 0 > s[i]:			r = IS_GREATER_THAN
	default:				r = IS_SAME_AS
	}
	return
}

func (s *I8Slice) Cut(i, j int) {
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

func (s *I8Slice) Trim(i, j int) {
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

func (s *I8Slice) Delete(i int) {
	a := *s
	n := len(a)
	if i > -1 && i < n {
		copy(a[i:n - 1], a[i + 1:n])
		*s = a[0 : n - 1]
	}
}

func (s I8Slice) Each(f func(interface{})) {
	for _, v := range s {
		f(v)
	}
}

func (s I8Slice) EachWithIndex(f func(int, interface{})) {
	for i, v := range s {
		f(i, v)
	}
}

func (s I8Slice) EachWithKey(f func(key, value interface{})) {
	for i, v := range s {
		f(i, v)
	}
}

func (s I8Slice) I8Each(f func(int8)) {
	for _, v := range s {
		f(v)
	}
}

func (s I8Slice) I8EachWithIndex(f func(int, int8)) {
	for i, v := range s {
		f(i, v)
	}
}

func (s I8Slice) I8EachWithKey(f func(interface{}, int8)) {
	for i, v := range s {
		f(i, v)
	}
}

func (s I8Slice) String() (t string) {
	for _, v := range s {
		if len(t) > 0 {
			t += " "
		}
		t += fmt.Sprintf("%v", v)
	}
	return fmt.Sprintf("(%v)", t)
}

func (s I8Slice) BlockCopy(destination, source, count int) {
	end := source + count
	if end > len(s) {
		end = len(s)
	}
	copy(s[destination:], s[source:end])
}

func (s I8Slice) BlockClear(start, count int) {
	copy(s[start:], make(I8Slice, count, count))
}

func (s I8Slice) Overwrite(offset int, container interface{}) {
	switch container := container.(type) {
	case I8Slice:			copy(s[offset:], container)
	case []int8:			copy(s[offset:], container)
	}
}

func (s *I8Slice) Reallocate(length, capacity int) {
	switch {
	case length > capacity:		s.Reallocate(capacity, capacity)
	case capacity != cap(*s):	x := make(I8Slice, length, capacity)
								copy(x, *s)
								*s = x
	default:					*s = (*s)[:length]
	}
}

func (s *I8Slice) Extend(n int) {
	c := cap(*s)
	l := len(*s) + n
	if l > c {
		c = l
	}
	s.Reallocate(l, c)
}

func (s *I8Slice) Expand(i, n int) {
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
		x := make(I8Slice, l, c)
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

func (s I8Slice) Reverse() {
	end := s.Len() - 1
	for i := 0; i < end; i++ {
		s[i], s[end] = s[end], s[i]
		end--
	}
}

func (s I8Slice) Depth() int {
	return 0
}

func (s *I8Slice) Append(v interface{}) {
	s.I8Append(v.(int8))
}

func (s *I8Slice) I8Append(v int8) {
	*s = append(*s, v)
}

func (s *I8Slice) AppendSlice(o I8Slice) {
	*s = append(*s, o...)
}

func (s *I8Slice) Prepend(v interface{}) {
	s.I8Prepend(v.(int8))
}

func (s *I8Slice) I8Prepend(v int8) {
	l := s.Len() + 1
	n := make(I8Slice, l, l)
	n[0] = v
	copy(n[1:], *s)
	*s = n
}

func (s *I8Slice) PrependSlice(o I8Slice) {
	l := s.Len() + o.Len()
	n := make(I8Slice, l, l)
	copy(n, o)
	copy(n[o.Len():], *s)
	*s = n
}

func (s I8Slice) Repeat(count int) I8Slice {
	length := len(s) * count
	capacity := cap(s)
	if capacity < length {
		capacity = length
	}
	destination := make(I8Slice, length, capacity)
	for start, end := 0, len(s); count > 0; count-- {
		copy(destination[start:end], s)
		start = end
		end += len(s)
	}
	return destination
}

func (s *I8Slice) Flatten() {
	//	Flatten is a non-op for the I8Slice as they cannot contain nested elements
}

func (s I8Slice) equal(o I8Slice) (r bool) {
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

func (s I8Slice) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case *I8Slice:			r = o != nil && s.equal(*o)
	case I8Slice:			r = s.equal(o)
	case *[]int8:			r = o != nil && s.equal(*o)
	case []int8:			r = s.equal(o)
	}
	return
}

func (s I8Slice) Car() (h interface{}) {
	if s.Len() > 0 {
		h = s[0]
	}
	return
}

func (s I8Slice) Cdr() (t I8Slice) {
	if s.Len() > 1 {
		t = s[1:]
	}
	return
}

func (s *I8Slice) Rplaca(v interface{}) {
	switch {
	case s == nil:			*s = *I8List(v.(int8))
	case s.Len() == 0:		*s = append(*s, v.(int8))
	default:				(*s)[0] = v.(int8)
	}
}

func (s *I8Slice) Rplacd(v interface{}) {
	if s == nil {
		*s = *I8List(v.(int8))
	} else {
		ReplaceSlice := func(v I8Slice) {
			if l := len(v); l < cap(*s) {
				copy((*s)[1:], v)
				*s = (*s)[0:l + 1]
			} else {
				l++
				n := make(I8Slice, l, l)
				copy(n, (*s)[:1])
				copy(n[1:], v)
				*s = n
			}
		}

		switch v := v.(type) {
		case *I8Slice:		ReplaceSlice(*v)
		case I8Slice:		ReplaceSlice(v)
		case *[]int8:		ReplaceSlice(I8Slice(*v))
		case []int8:		ReplaceSlice(I8Slice(v))
		case nil:			*s = (*s)[:1]
		default:			(*s)[1] = v.(int8)
							*s = (*s)[:2]
		}
	}
}