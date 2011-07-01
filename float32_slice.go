package slices

import "fmt"

func F32List(n... float32) *F32Slice {
	return (*F32Slice)(&n)
}

type F32Slice	[]float32

func (s F32Slice) Len() int							{ return len(s) }
func (s F32Slice) Cap() int							{ return cap(s) }

func (s F32Slice) At(i int) interface{}				{ return s[i] }
func (s F32Slice) F32At(i int) float32				{ return s[i] }
func (s F32Slice) F64At(i int) float64				{ return float64(s[i]) }
func (s F32Slice) Set(i int, v interface{})			{ s[i] = v.(float32) }
func (s F32Slice) F32Set(i int, v float32)			{ s[i] = v }
func (s F32Slice) F64Set(i int, v float64)			{ s[i] = float32(v) }
func (s F32Slice) Clear(i int)						{ s[i] = 0 }
func (s F32Slice) Swap(i, j int)					{ s[i], s[j] = s[j], s[i] }

func (s F32Slice) Negate(i int)						{ s[i] = -s[i] }
func (s F32Slice) Increment(i int)					{ s[i] += 1 }
func (s F32Slice) Decrement(i int)					{ s[i] -= 1 }

func (s F32Slice) Add(i, j int)						{ s[i] += s[j] }
func (s F32Slice) Subtract(i, j int)				{ s[i] -= s[j] }
func (s F32Slice) Multiply(i, j int)				{ s[i] *= s[j] }
func (s F32Slice) Divide(i, j int)					{ s[i] /= s[j] }

func (s F32Slice) Less(i, j int) bool				{ return s[i] < s[j] }
func (s F32Slice) AtLeast(i, j int) bool			{ return s[i] <= s[j] }
func (s F32Slice) Same(i, j int) bool				{ return s[i] == s[j] }
func (s F32Slice) AtMost(i, j int) bool				{ return s[i] >= s[j] }
func (s F32Slice) More(i, j int) bool				{ return s[i] > s[j] }

func (s F32Slice) ZeroLess(i int) bool				{ return 0 < s[i] }
func (s F32Slice) ZeroAtLeast(i, j int) bool		{ return 0 <= s[j] }
func (s F32Slice) ZeroSame(i int) bool				{ return 0 == s[i] }
func (s F32Slice) ZeroAtMost(i, j int) bool			{ return 0 >= s[j] }
func (s F32Slice) ZeroMore(i int) bool				{ return 0 > s[i] }

func (s F32Slice) Compare(i, j int) (r int) {
	switch {
	case s[i] < s[j]:		r = IS_LESS_THAN
	case s[i] > s[j]:		r = IS_GREATER_THAN
	default:				r = IS_SAME_AS
	}
	return
}

func (s F32Slice) ZeroCompare(i int) (r int) {
	switch {
	case 0 < s[i]:			r = IS_LESS_THAN
	case 0 > s[i]:			r = IS_GREATER_THAN
	default:				r = IS_SAME_AS
	}
	return
}

func (s *F32Slice) Cut(i, j int) {
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

func (s *F32Slice) Delete(i int) {
	a := *s
	n := len(a)
	if i > -1 && i < n {
		copy(a[i:n - 1], a[i + 1:n])
		*s = a[0 : n - 1]
	}
}

func (s F32Slice) Each(f func(interface{})) {
	for _, v := range s {
		f(v)
	}
}

func (s F32Slice) EachWithIndex(f func(int, interface{})) {
	for i, v := range s {
		f(i, v)
	}
}

func (s F32Slice) EachWithKey(f func(key, value interface{})) {
	for i, v := range s {
		f(i, v)
	}
}

func (s F32Slice) F32Each(f func(float32)) {
	for _, v := range s {
		f(v)
	}
}

func (s F32Slice) F32EachWithIndex(f func(int, float32)) {
	for i, v := range s {
		f(i, v)
	}
}

func (s F32Slice) F32EachWithKey(f func(interface{}, float32)) {
	for i, v := range s {
		f(i, v)
	}
}

func (s F32Slice) String() (t string) {
	for _, v := range s {
		if len(t) > 0 {
			t += " "
		}
		t += fmt.Sprintf("%v", v)
	}
	return fmt.Sprintf("(%v)", t)
}

func (s F32Slice) BlockCopy(destination, source, count int) {
	end := source + count
	if end > len(s) {
		end = len(s)
	}
	copy(s[destination:], s[source:end])
}

func (s F32Slice) BlockClear(start, count int) {
	copy(s[start:], make(F32Slice, count, count))
}

func (s F32Slice) Overwrite(offset int, container interface{}) {
	switch container := container.(type) {
	case F32Slice:			copy(s[offset:], container)
	case []float32:			copy(s[offset:], container)
	}
}

func (s *F32Slice) Reallocate(length, capacity int) {
	switch {
	case length > capacity:		s.Reallocate(capacity, capacity)
	case capacity != cap(*s):	x := make(F32Slice, length, capacity)
								copy(x, *s)
								*s = x
	default:					*s = (*s)[:length]
	}
}

func (s *F32Slice) Extend(n int) {
	c := cap(*s)
	l := len(*s) + n
	if l > c {
		c = l
	}
	s.Reallocate(l, c)
}

func (s *F32Slice) Expand(i, n int) {
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
		x := make(F32Slice, l, c)
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

func (s F32Slice) Reverse() {
	end := s.Len() - 1
	for i := 0; i < end; i++ {
		s[i], s[end] = s[end], s[i]
		end--
	}
}

func (s F32Slice) Depth() int {
	return 0
}

func (s *F32Slice) Append(v interface{}) {
	s.F32Append(v.(float32))
}

func (s *F32Slice) F32Append(v float32) {
	*s = append(*s, v)
}

func (s *F32Slice) AppendSlice(o F32Slice) {
	*s = append(*s, o...)
}

func (s *F32Slice) Prepend(v interface{}) {
	s.F32Prepend(v.(float32))
}

func (s *F32Slice) F32Prepend(v float32) {
	l := s.Len() + 1
	n := make(F32Slice, l, l)
	n[0] = v
	copy(n[1:], *s)
	*s = n
}

func (s *F32Slice) PrependSlice(o F32Slice) {
	l := s.Len() + o.Len()
	n := make(F32Slice, l, l)
	copy(n, o)
	copy(n[o.Len():], *s)
	*s = n
}

func (s F32Slice) Subslice(start, end int) interface{} {
	return s[start:end]
}

func (s F32Slice) Repeat(count int) F32Slice {
	length := len(s) * count
	capacity := cap(s)
	if capacity < length {
		capacity = length
	}
	destination := make(F32Slice, length, capacity)
	for start, end := 0, len(s); count > 0; count-- {
		copy(destination[start:end], s)
		start = end
		end += len(s)
	}
	return destination
}

func (s *F32Slice) Flatten() {
	//	Flatten is a non-op for the F32Slice as they cannot contain nested elements
}

func (s F32Slice) equal(o F32Slice) (r bool) {
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

func (s F32Slice) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case *F32Slice:			r = o != nil && s.equal(*o)
	case F32Slice:			r = s.equal(o)
	case *[]float32:		r = o != nil && s.equal(*o)
	case []float32:			r = s.equal(o)
	}
	return
}

func (s F32Slice) Car() (h interface{}) {
	if s.Len() > 0 {
		h = s[0]
	}
	return
}

func (s F32Slice) Cdr() (t F32Slice) {
	if s.Len() > 1 {
		t = s[1:]
	}
	return
}

func (s *F32Slice) Rplaca(v interface{}) {
	switch {
	case s == nil:			*s = *F32List(v.(float32))
	case s.Len() == 0:		*s = append(*s, v.(float32))
	default:				(*s)[0] = v.(float32)
	}
}

func (s *F32Slice) Rplacd(v interface{}) {
	if s == nil {
		*s = *F32List(v.(float32))
	} else {
		ReplaceSlice := func(v F32Slice) {
			if l := len(v); l < cap(*s) {
				copy((*s)[1:], v)
			} else {
				l++
				n := make(F32Slice, l, l)
				copy(n, (*s)[:1])
				copy(n[1:], v)
				*s = n
			}
		}

		switch v := v.(type) {
		case *F32Slice:		ReplaceSlice(*v)
		case F32Slice:		ReplaceSlice(v)
		case *[]float32:	ReplaceSlice(F32Slice(*v))
		case []float32:		ReplaceSlice(F32Slice(v))
		case nil:			*s = (*s)[:1]
		default:			(*s)[1] = v.(float32)
							*s = (*s)[:2]
		}
	}
}