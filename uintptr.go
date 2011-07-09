package slices

import "fmt"
import "sort"

func AList(n... uintptr) *ASlice {
	return (*ASlice)(&n)
}

type ASlice	[]uintptr

func (s ASlice) Len() int							{ return len(s) }
func (s ASlice) Cap() int							{ return cap(s) }

func (s ASlice) At(i int) interface{}				{ return s[i] }
func (s ASlice) AAt(i int) uintptr					{ return s[i] }
func (s ASlice) Set(i int, v interface{})			{ s[i] = v.(uintptr) }
func (s ASlice) ASet(i int, v uintptr)				{ s[i] = v }
func (s ASlice) Clear(i int)						{ s[i] = 0 }
func (s ASlice) Swap(i, j int)						{ s[i], s[j] = s[j], s[i] }

func (s ASlice) Negate(i int)						{ s[i] = -s[i] }
func (s ASlice) Increment(i int)					{ s[i]++ }
func (s ASlice) Decrement(i int)					{ s[i]-- }

func (s ASlice) Add(i, j int)						{ s[i] += s[j] }
func (s ASlice) Subtract(i, j int)					{ s[i] -= s[j] }

func (s ASlice) And(i, j int)						{ s[i] &= s[j] }
func (s ASlice) Or(i, j int)						{ s[i] |= s[j] }
func (s ASlice) Xor(i, j int)						{ s[i] ^= s[j] }
func (s ASlice) Invert(i int)						{ s[i] = ^s[i] }
func (s ASlice) ShiftLeft(i, j int)					{ s[i] <<= s[j] }
func (s ASlice) ShiftRight(i, j int)				{ s[i] >>= s[j] }

func (s ASlice) Less(i, j int) bool					{ return s[i] < s[j] }
func (s ASlice) AtLeast(i, j int) bool				{ return s[i] <= s[j] }
func (s ASlice) Same(i, j int) bool					{ return s[i] == s[j] }
func (s ASlice) AtMost(i, j int) bool				{ return s[i] >= s[j] }
func (s ASlice) More(i, j int) bool					{ return s[i] > s[j] }
func (s ASlice) ZeroLess(i int) bool				{ return 0 < s[i] }
func (s ASlice) ZeroAtLeast(i, j int) bool			{ return 0 <= s[j] }
func (s ASlice) ZeroSame(i int) bool				{ return 0 == s[i] }
func (s ASlice) ZeroAtMost(i, j int) bool			{ return 0 >= s[j] }
func (s ASlice) ZeroMore(i int) bool				{ return 0 > s[i] }

func (s ASlice) Sort()								{ sort.Sort(s) }

func (s *ASlice) RestrictTo(i, j int)				{ *s = (*s)[i:j] }

func (s ASlice) Compare(i, j int) (r int) {
	switch {
	case s[i] < s[j]:		r = IS_LESS_THAN
	case s[i] > s[j]:		r = IS_GREATER_THAN
	default:				r = IS_SAME_AS
	}
	return
}

func (s ASlice) ZeroCompare(i int) (r int) {
	switch {
	case 0 < s[i]:			r = IS_LESS_THAN
	case 0 > s[i]:			r = IS_GREATER_THAN
	default:				r = IS_SAME_AS
	}
	return
}

func (s *ASlice) Cut(i, j int) {
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
			*s = a[:m]
		}
	}
}

func (s *ASlice) Trim(i, j int) {
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
		*s = a[:j - i]
	}
}

func (s *ASlice) Delete(i int) {
	a := *s
	n := len(a)
	if i > -1 && i < n {
		copy(a[i:n - 1], a[i + 1:n])
		*s = a[:n - 1]
	}
}

func (s *ASlice) DeleteIf(f interface{}) {
	a := *s
	p := 0
	switch f := f.(type) {
	case uintptr:					for i, v := range a {
										if i != p {
											a[p] = v
										}
										if v != f {
											p++
										}
									}

	case func(uintptr) bool:		for i, v := range a {
										if i != p {
											a[p] = v
										}
										if !f(v) {
											p++
										}
									}

	case func(interface{}) bool:	for i, v := range a {
										if i != p {
											a[p] = v
										}
										if !f(v) {
											p++
										}
									}

	default:						p = len(a)
	}
	*s = a[:p]
}

func (s ASlice) Each(f interface{}) {
	switch f := f.(type) {
	case func(uintptr):						for _, v := range s { f(v) }
	case func(int, uintptr):				for i, v := range s { f(i, v) }
	case func(interface{}, uintptr):		for i, v := range s { f(i, v) }
	case func(interface{}):					for _, v := range s { f(v) }
	case func(int, interface{}):			for i, v := range s { f(i, v) }
	case func(interface{}, interface{}):	for i, v := range s { f(i, v) }
	}
}

func (s ASlice) String() (t string) {
	for _, v := range s {
		if len(t) > 0 {
			t += " "
		}
		t += fmt.Sprintf("%v", v)
	}
	return fmt.Sprintf("(%v)", t)
}

func (s ASlice) BlockCopy(destination, source, count int) {
	end := source + count
	if end > len(s) {
		end = len(s)
	}
	copy(s[destination:], s[source:end])
}

func (s ASlice) BlockClear(start, count int) {
	copy(s[start:], make(ASlice, count, count))
}

func (s ASlice) Overwrite(offset int, container interface{}) {
	switch container := container.(type) {
	case ASlice:			copy(s[offset:], container)
	case []uintptr:			copy(s[offset:], container)
	}
}

func (s *ASlice) Reallocate(length, capacity int) {
	switch {
	case length > capacity:		s.Reallocate(capacity, capacity)
	case capacity != cap(*s):	x := make(ASlice, length, capacity)
								copy(x, *s)
								*s = x
	default:					*s = (*s)[:length]
	}
}

func (s *ASlice) Extend(n int) {
	c := cap(*s)
	l := len(*s) + n
	if l > c {
		c = l
	}
	s.Reallocate(l, c)
}

func (s *ASlice) Expand(i, n int) {
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
		x := make(ASlice, l, c)
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

func (s ASlice) Reverse() {
	end := s.Len() - 1
	for i := 0; i < end; i++ {
		s[i], s[end] = s[end], s[i]
		end--
	}
}

func (s ASlice) Depth() int {
	return 0
}

func (s *ASlice) Append(v interface{}) {
	switch v := v.(type) {
	case uintptr:			*s = append(*s, v)
	case ASlice:			*s = append(*s, v...)
	case *ASlice:			*s = append(*s, (*v)...)
	case []uintptr:			s.Append(ASlice(v))
	case *[]uintptr:		s.Append(ASlice(*v))
	default:				panic(v)
	}
}

func (s *ASlice) Prepend(v interface{}) {
	switch v := v.(type) {
	case uintptr:			l := s.Len() + 1
							n := make(ASlice, l, l)
							n[0] = v
							copy(n[1:], *s)
							*s = n

	case ASlice:			l := s.Len() + len(v)
							n := make(ASlice, l, l)
							copy(n, v)
							copy(n[len(v):], *s)
							*s = n

	case *ASlice:			s.Prepend(*v)
	case []uintptr:			s.Prepend(ASlice(v))
	case *[]uintptr:		s.Prepend(ASlice(*v))
	default:				panic(v)
	}
}

func (s ASlice) Repeat(count int) ASlice {
	length := len(s) * count
	capacity := cap(s)
	if capacity < length {
		capacity = length
	}
	destination := make(ASlice, length, capacity)
	for start, end := 0, len(s); count > 0; count-- {
		copy(destination[start:end], s)
		start = end
		end += len(s)
	}
	return destination
}

func (s *ASlice) Flatten() {
	//	Flatten is a non-op for the ASlice as they cannot contain nested elements
}

func (s ASlice) equal(o ASlice) (r bool) {
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

func (s ASlice) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case *ASlice:			r = o != nil && s.equal(*o)
	case ASlice:			r = s.equal(o)
	case *[]uintptr:			r = o != nil && s.equal(*o)
	case []uintptr:			r = s.equal(o)
	}
	return
}

func (s ASlice) Car() (h interface{}) {
	if s.Len() > 0 {
		h = s[0]
	}
	return
}

func (s ASlice) Cdr() (t ASlice) {
	if s.Len() > 1 {
		t = s[1:]
	}
	return
}

func (s *ASlice) Rplaca(v interface{}) {
	switch {
	case s == nil:			*s = *AList(v.(uintptr))
	case s.Len() == 0:		*s = append(*s, v.(uintptr))
	default:				(*s)[0] = v.(uintptr)
	}
}

func (s *ASlice) Rplacd(v interface{}) {
	if s == nil {
		*s = *AList(v.(uintptr))
	} else {
		ReplaceSlice := func(v ASlice) {
			if l := len(v); l < cap(*s) {
				copy((*s)[1:], v)
				*s = (*s)[:l + 1]
			} else {
				l++
				n := make(ASlice, l, l)
				copy(n, (*s)[:1])
				copy(n[1:], v)
				*s = n
			}
		}

		switch v := v.(type) {
		case *ASlice:		ReplaceSlice(*v)
		case ASlice:		ReplaceSlice(v)
		case *[]uintptr:	ReplaceSlice(ASlice(*v))
		case []uintptr:		ReplaceSlice(ASlice(v))
		case nil:			*s = (*s)[:1]
		default:			(*s)[1] = v.(uintptr)
							*s = (*s)[:2]
		}
	}
}

func (s ASlice) SetIntersection(o ASlice) (r ASlice) {
	cache := make(map[uintptr]bool)
	for _, v := range s {
		if ok := cache[v]; !ok {
			cache[v] = true
		}
	}
	for _, v := range o {
		if _, ok := cache[v]; ok {
			cache[v] = false, false
			r = append(r, v)
		}
	}
	return
}

func (s ASlice) SetUnion(o ASlice) (r ASlice) {
	cache := make(map[uintptr]bool)
	for _, v := range s {
		if ok := cache[v]; !ok {
			cache[v] = true
		}
	}
	for _, v := range o {
		if ok := cache[v]; !ok {
			cache[v] = true
		}
	}
	for k, _ := range cache {
		r = append(r, k)
	}
	return
}

func (s ASlice) SetDifference(o ASlice) (r ASlice) {
	left := make(map[uintptr]bool)
	right := make(map[uintptr]bool)
	for _, v := range s {
		if ok := left[v]; !ok {
			left[v] = true
		}
	}
	for _, v := range o {
		if ok := right[v]; !ok {
			right[v] = true
		}
	}
	for k, _ := range left {
		if ok := right[k]; ok {
			right[k] = false, false
		} else {
			r = append(r, k)
		}
	}
	for k, _ := range right {
		if ok := left[k]; !ok {
			r = append(r, k)
		}
	}
	return
}