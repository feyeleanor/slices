package slices

import "fmt"

func C128List(n... complex128) *C128Slice {
	return (*C128Slice)(&n)
}

type C128Slice	[]complex128

func (s C128Slice) Len() int						{ return len(s) }
func (s C128Slice) Cap() int						{ return cap(s) }

func (s C128Slice) At(i int) interface{}			{ return s[i] }
func (s C128Slice) C128At(i int) complex128			{ return s[i] }
func (s C128Slice) Set(i int, v interface{})		{ s[i] = v.(complex128) }
func (s C128Slice) C128Set(i int, v complex128)		{ s[i] = v }
func (s C128Slice) Clear(i int)						{ s[i] = 0 }
func (s C128Slice) Swap(i, j int)					{ s[i], s[j] = s[j], s[i] }

func (s C128Slice) Negate(i int)					{ s[i] = -s[i] }
func (s C128Slice) Increment(i int)					{ s[i]++ }
func (s C128Slice) Decrement(i int)					{ s[i]-- }

func (s C128Slice) Add(i, j int)					{ s[i] += s[j] }
func (s C128Slice) Subtract(i, j int)				{ s[i] -= s[j] }
func (s C128Slice) Multiply(i, j int)				{ s[i] *= s[j] }
func (s C128Slice) Divide(i, j int)					{ s[i] /= s[j] }

func (s C128Slice) Same(i, j int) bool				{ return s[i] == s[j] }

func (s *C128Slice) RestrictTo(i, j int)			{ *s = (*s)[i:j] }

func (s *C128Slice) Cut(i, j int) {
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

func (s *C128Slice) Trim(i, j int) {
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

func (s *C128Slice) Delete(i int) {
	a := *s
	n := len(a)
	if i > -1 && i < n {
		copy(a[i:n - 1], a[i + 1:n])
		*s = a[0 : n - 1]
	}
}

func (s C128Slice) Each(f func(interface{})) {
	for _, v := range s {
		f(v)
	}
}

func (s C128Slice) EachWithIndex(f func(int, interface{})) {
	for i, v := range s {
		f(i, v)
	}
}

func (s C128Slice) EachWithKey(f func(key, value interface{})) {
	for i, v := range s {
		f(i, v)
	}
}

func (s C128Slice) C128Each(f func(complex128)) {
	for _, v := range s {
		f(v)
	}
}

func (s C128Slice) C128EachWithIndex(f func(int, complex128)) {
	for i, v := range s {
		f(i, v)
	}
}

func (s C128Slice) C128EachWithKey(f func(interface{}, complex128)) {
	for i, v := range s {
		f(i, v)
	}
}

func (s C128Slice) String() (t string) {
	for _, v := range s {
		if len(t) > 0 {
			t += " "
		}
		t += fmt.Sprintf("%v", v)
	}
	return fmt.Sprintf("(%v)", t)
}

func (s C128Slice) BlockCopy(destination, source, count int) {
	end := source + count
	if end > len(s) {
		end = len(s)
	}
	copy(s[destination:], s[source:end])
}

func (s C128Slice) BlockClear(start, count int) {
	copy(s[start:], make(C128Slice, count, count))
}

func (s C128Slice) Overwrite(offset int, container interface{}) {
	switch container := container.(type) {
	case C128Slice:				copy(s[offset:], container)
	case []complex128:			copy(s[offset:], container)
	}
}

func (s *C128Slice) Reallocate(length, capacity int) {
	switch {
	case length > capacity:		s.Reallocate(capacity, capacity)
	case capacity != cap(*s):	x := make(C128Slice, length, capacity)
								copy(x, *s)
								*s = x
	default:					*s = (*s)[:length]
	}
}

func (s *C128Slice) Extend(n int) {
	c := cap(*s)
	l := len(*s) + n
	if l > c {
		c = l
	}
	s.Reallocate(l, c)
}

func (s *C128Slice) Expand(i, n int) {
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
		x := make(C128Slice, l, c)
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

func (s C128Slice) Reverse() {
	end := s.Len() - 1
	for i := 0; i < end; i++ {
		s[i], s[end] = s[end], s[i]
		end--
	}
}

func (s C128Slice) Depth() int {
	return 0
}

func (s *C128Slice) Append(v interface{}) {
	s.C128Append(v.(complex128))
}

func (s *C128Slice) C128Append(v complex128) {
	*s = append(*s, v)
}

func (s *C128Slice) AppendSlice(o C128Slice) {
	*s = append(*s, o...)
}

func (s *C128Slice) Prepend(v interface{}) {
	s.C128Prepend(v.(complex128))
}

func (s *C128Slice) C128Prepend(v complex128) {
	l := s.Len() + 1
	n := make(C128Slice, l, l)
	n[0] = v
	copy(n[1:], *s)
	*s = n
}

func (s *C128Slice) PrependSlice(o C128Slice) {
	l := s.Len() + o.Len()
	n := make(C128Slice, l, l)
	copy(n, o)
	copy(n[o.Len():], *s)
	*s = n
}

func (s C128Slice) Repeat(count int) C128Slice {
	length := len(s) * count
	capacity := cap(s)
	if capacity < length {
		capacity = length
	}
	destination := make(C128Slice, length, capacity)
	for start, end := 0, len(s); count > 0; count-- {
		copy(destination[start:end], s)
		start = end
		end += len(s)
	}
	return destination
}

func (s *C128Slice) Flatten() {
	//	Flatten is a non-op for the C128Slice as they cannot contain nested elements
}

func (s C128Slice) equal(o C128Slice) (r bool) {
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

func (s C128Slice) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case *C128Slice:			r = o != nil && s.equal(*o)
	case C128Slice:				r = s.equal(o)
	case *[]complex128:			r = o != nil && s.equal(*o)
	case []complex128:			r = s.equal(o)
	}
	return
}

func (s C128Slice) Car() (h interface{}) {
	if s.Len() > 0 {
		h = s[0]
	}
	return
}

func (s C128Slice) Cdr() (t C128Slice) {
	if s.Len() > 1 {
		t = s[1:]
	}
	return
}

func (s *C128Slice) Rplaca(v interface{}) {
	switch {
	case s == nil:			*s = *C128List(v.(complex128))
	case s.Len() == 0:		*s = append(*s, v.(complex128))
	default:				(*s)[0] = v.(complex128)
	}
}

func (s *C128Slice) Rplacd(v interface{}) {
	if s == nil {
		*s = *C128List(v.(complex128))
	} else {
		ReplaceSlice := func(v C128Slice) {
			if l := len(v); l < cap(*s) {
				copy((*s)[1:], v)
				*s = (*s)[0:l + 1]
			} else {
				l++
				n := make(C128Slice, l, l)
				copy(n, (*s)[:1])
				copy(n[1:], v)
				*s = n
			}
		}

		switch v := v.(type) {
		case *C128Slice:		ReplaceSlice(*v)
		case C128Slice:			ReplaceSlice(v)
		case *[]complex128:		ReplaceSlice(C128Slice(*v))
		case []complex128:		ReplaceSlice(C128Slice(v))
		case nil:				*s = (*s)[:1]
		default:				(*s)[1] = v.(complex128)
								*s = (*s)[:2]
		}
	}
}