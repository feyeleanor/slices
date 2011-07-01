package slices

import "fmt"

func C64List(n... complex64) *C64Slice {
	return (*C64Slice)(&n)
}

type C64Slice	[]complex64

func (s C64Slice) Len() int							{ return len(s) }
func (s C64Slice) Cap() int							{ return cap(s) }

func (s C64Slice) At(i int) interface{}				{ return s[i] }
func (s C64Slice) C64At(i int) complex64			{ return s[i] }
func (s C64Slice) Set(i int, v interface{})			{ s[i] = v.(complex64) }
func (s C64Slice) C64Set(i int, v complex64)		{ s[i] = v }
func (s C64Slice) Clear(i int)						{ s[i] = 0 }
func (s C64Slice) Swap(i, j int)					{ s[i], s[j] = s[j], s[i] }

func (s C64Slice) Negate(i int)						{ s[i] = -s[i] }
func (s C64Slice) Increment(i int)					{ s[i]++ }
func (s C64Slice) Decrement(i int)					{ s[i]-- }

func (s C64Slice) Add(i, j int)						{ s[i] += s[j] }
func (s C64Slice) Subtract(i, j int)				{ s[i] -= s[j] }
func (s C64Slice) Multiply(i, j int)				{ s[i] *= s[j] }
func (s C64Slice) Divide(i, j int)					{ s[i] /= s[j] }

func (s C64Slice) Same(i, j int) bool				{ return s[i] == s[j] }

func (s *C64Slice) Cut(i, j int) {
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

func (s *C64Slice) Delete(i int) {
	a := *s
	n := len(a)
	if i > -1 && i < n {
		copy(a[i:n - 1], a[i + 1:n])
		*s = a[0 : n - 1]
	}
}

func (s C64Slice) Each(f func(interface{})) {
	for _, v := range s {
		f(v)
	}
}

func (s C64Slice) EachWithIndex(f func(int, interface{})) {
	for i, v := range s {
		f(i, v)
	}
}

func (s C64Slice) EachWithKey(f func(key, value interface{})) {
	for i, v := range s {
		f(i, v)
	}
}

func (s C64Slice) C64Each(f func(complex64)) {
	for _, v := range s {
		f(v)
	}
}

func (s C64Slice) C64EachWithIndex(f func(int, complex64)) {
	for i, v := range s {
		f(i, v)
	}
}

func (s C64Slice) C64EachWithKey(f func(interface{}, complex64)) {
	for i, v := range s {
		f(i, v)
	}
}

func (s C64Slice) String() (t string) {
	for _, v := range s {
		if len(t) > 0 {
			t += " "
		}
		t += fmt.Sprintf("%v", v)
	}
	return fmt.Sprintf("(%v)", t)
}

func (s C64Slice) BlockCopy(destination, source, count int) {
	end := source + count
	if end > len(s) {
		end = len(s)
	}
	copy(s[destination:], s[source:end])
}

func (s C64Slice) BlockClear(start, count int) {
	copy(s[start:], make(C64Slice, count, count))
}

func (s C64Slice) Overwrite(offset int, container interface{}) {
	switch container := container.(type) {
	case C64Slice:			copy(s[offset:], container)
	case []complex64:			copy(s[offset:], container)
	}
}

func (s *C64Slice) Reallocate(length, capacity int) {
	switch {
	case length > capacity:		s.Reallocate(capacity, capacity)
	case capacity != cap(*s):	x := make(C64Slice, length, capacity)
								copy(x, *s)
								*s = x
	default:					*s = (*s)[:length]
	}
}

func (s *C64Slice) Extend(n int) {
	c := cap(*s)
	l := len(*s) + n
	if l > c {
		c = l
	}
	s.Reallocate(l, c)
}

func (s *C64Slice) Expand(i, n int) {
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
		x := make(C64Slice, l, c)
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

func (s C64Slice) Reverse() {
	end := s.Len() - 1
	for i := 0; i < end; i++ {
		s[i], s[end] = s[end], s[i]
		end--
	}
}

func (s C64Slice) Depth() int {
	return 0
}

func (s *C64Slice) Append(v interface{}) {
	s.C64Append(v.(complex64))
}

func (s *C64Slice) C64Append(v complex64) {
	*s = append(*s, v)
}

func (s *C64Slice) AppendSlice(o C64Slice) {
	*s = append(*s, o...)
}

func (s *C64Slice) Prepend(v interface{}) {
	s.C64Prepend(v.(complex64))
}

func (s *C64Slice) C64Prepend(v complex64) {
	l := s.Len() + 1
	n := make(C64Slice, l, l)
	n[0] = v
	copy(n[1:], *s)
	*s = n
}

func (s *C64Slice) PrependSlice(o C64Slice) {
	l := s.Len() + o.Len()
	n := make(C64Slice, l, l)
	copy(n, o)
	copy(n[o.Len():], *s)
	*s = n
}

func (s C64Slice) Repeat(count int) C64Slice {
	length := len(s) * count
	capacity := cap(s)
	if capacity < length {
		capacity = length
	}
	destination := make(C64Slice, length, capacity)
	for start, end := 0, len(s); count > 0; count-- {
		copy(destination[start:end], s)
		start = end
		end += len(s)
	}
	return destination
}

func (s *C64Slice) Flatten() {
	//	Flatten is a non-op for the C64Slice as they cannot contain nested elements
}

func (s C64Slice) equal(o C64Slice) (r bool) {
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

func (s C64Slice) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case *C64Slice:			r = o != nil && s.equal(*o)
	case C64Slice:			r = s.equal(o)
	case *[]complex64:			r = o != nil && s.equal(*o)
	case []complex64:			r = s.equal(o)
	}
	return
}

func (s C64Slice) Car() (h interface{}) {
	if s.Len() > 0 {
		h = s[0]
	}
	return
}

func (s C64Slice) Cdr() (t C64Slice) {
	if s.Len() > 1 {
		t = s[1:]
	}
	return
}

func (s *C64Slice) Rplaca(v interface{}) {
	switch {
	case s == nil:			*s = *C64List(v.(complex64))
	case s.Len() == 0:		*s = append(*s, v.(complex64))
	default:				(*s)[0] = v.(complex64)
	}
}

func (s *C64Slice) Rplacd(v interface{}) {
	if s == nil {
		*s = *C64List(v.(complex64))
	} else {
		ReplaceSlice := func(v C64Slice) {
			if l := len(v); l < cap(*s) {
				copy((*s)[1:], v)
			} else {
				l++
				n := make(C64Slice, l, l)
				copy(n, (*s)[:1])
				copy(n[1:], v)
				*s = n
			}
		}

		switch v := v.(type) {
		case *C64Slice:		ReplaceSlice(*v)
		case C64Slice:		ReplaceSlice(v)
		case *[]complex64:		ReplaceSlice(C64Slice(*v))
		case []complex64:		ReplaceSlice(C64Slice(v))
		case nil:			*s = (*s)[:1]
		default:			(*s)[1] = v.(complex64)
							*s = (*s)[:2]
		}
	}
}