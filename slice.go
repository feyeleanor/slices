package slices

import (
	"fmt"
	"math/rand"
	"reflect"
)

type Slice	[]interface{}

func (s Slice) release_reference(i int) {
	var zero interface{}
	s[i] = zero
}

func (s Slice) Len() int					{ return len(s) }
func (s Slice) Cap() int					{ return cap(s) }
func (s Slice) At(i int) interface{}		{ return s[i] }
func (s Slice) Set(i int, v interface{})	{ s[i] = v }
func (s Slice) Clear(i int)					{ s[i] = nil }
func (s Slice) Swap(i, j int)				{ s[i], s[j] = s[j], s[i] }
func (s *Slice) RestrictTo(i, j int)		{ *s = (*s)[i:j] }

func (s *Slice) Cut(i, j int) {
	a := *s
	n := len(a)
	if i < 0 {
		i = 0
	}
	if j > n {
		j = n
	}
	if j > i {
		if m := n - (j - i); m > 0 && m <= n {
			copy(a[i:m], a[j:n])
//			var zero interface{}
			for k := m; k < n; k++ {
//				a[k] = zero
				a.release_reference(k)
			}
			*s = a[:m]
		}
	}
}

func (s *Slice) Trim(i, j int) {
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
//		var zero interface{}
		for k, base := n - 1, i + 1; k > base; k-- {
//			a[k] = zero
			a.release_reference(k)
		}
		*s = a[:j - i]
	}
}

func (s *Slice) Delete(i int) {
	a := *s
	n := len(a)
	if i > -1 && i < n {
		copy(a[i:n - 1], a[i + 1:n])
//		var zero interface{}
//		a[n - 1] = zero
		a.release_reference(n - 1)
		*s = a[:n - 1]
	}
}

func (s *Slice) DeleteIf(f interface{}) {
	a := *s
	p := 0
	switch f := f.(type) {
	case func(interface{}) bool:	for i, v := range a {
										if i != p {
											a[p] = v
										}
										if !f(v) {
											p++
										}
									}

	default:						for i, v := range a {
										if i != p {
											a[p] = v
										}
										if v != f {
											p++
										}
									}
	}
	*s = a[:p]
}

func (s Slice) Each(f interface{}) {
	switch f := f.(type) {
	case func(interface{}):						for _, v := range s { f(v) }
	case func(int, interface{}):				for i, v := range s { f(i, v) }
	case func(interface{}, interface{}):		for i, v := range s { f(i, v) }
	}
}

func (s Slice) While(f interface{}) int {
	switch f := f.(type) {
	case func(interface{}) bool:				for i, v := range s {
													if !f(v) {
														return i
													}
												}
	case func(int, interface{}) bool:			for i, v := range s {
													if !f(i, v) {
														return i
													}
												}
	case func(interface{}, interface{}) bool:	for i, v := range s {
													if !f(i, v) {
														return i
													}
												}
	}
	return len(s)
}

func (s Slice) Until(f interface{}) int {
	switch f := f.(type) {
	case func(interface{}) bool:				for i, v := range s {
													if f(v) {
														return i
													}
												}
	case func(int, interface{}) bool:			for i, v := range s {
													if f(i, v) {
														return i
													}
												}
	case func(interface{}, interface{}) bool:	for i, v := range s {
													if f(i, v) {
														return i
													}
												}
	}
	return len(s)
}

func (s Slice) String() (t string) {
	for _, v := range s {
		if len(t) > 0 {
			t += " "
		}
		t += fmt.Sprintf("%v", v)
	}
	return fmt.Sprintf("(%v)", t)
}

func (s Slice) BlockCopy(destination, source, count int) {
	if destination < len(s) {
		if end := destination + count; end >= len(s) {
			copy(s[destination:], s[source:])
		} else {
			copy(s[destination : end], s[source:])
		}
	}
}

func (s Slice) BlockClear(start, count int) {
	copy(s[start:], make(Slice, count, count))
}

func (s Slice) Overwrite(offset int, container interface{}) {
	switch container := container.(type) {
	case Slice:					copy(s[offset:], container)
	case []interface{}:			copy(s[offset:], container)
	}
}

func (s *Slice) Reallocate(length, capacity int) {
	switch {
	case length > capacity:		s.Reallocate(capacity, capacity)
	case capacity != cap(*s):	x := make(Slice, length, capacity)
								copy(x, *s)
								*s = x
	default:					*s = (*s)[:length]
	}
}

func (s *Slice) Extend(n int) {
	c := cap(*s)
	l := len(*s) + n
	if l > c {
		c = l
	}
	s.Reallocate(l, c)
}

func (s *Slice) Expand(i, n int) {
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
		x := make(Slice, l, c)
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

func (s Slice) Depth() (c int) {
	for _, v := range s {
		if v, ok := v.(Nested); ok {
			if r := v.Depth() + 1; r > c {
				c = r
			}
		}
	}
	return
}

func (s Slice) Reverse() {
	end := s.Len() - 1
	for i := 0; i < end; i++ {
		s[i], s[end] = s[end], s[i]
		end--
	}
}

func (s *Slice) Append(v interface{}) {
	switch v := v.(type) {
	case Slice:				*s = append(*s, v...)
	case []interface{}:		*s = append(*s, v...)
	default:				*s = append(*s, v)
	}
}

func (s *Slice) Prepend(v interface{}) {
	switch v := v.(type) {
	case Slice:				l := s.Len() + len(v)
							n := make(Slice, l, l)
							copy(n, v)
							copy(n[len(v):], *s)
							*s = n

	case []interface{}:		s.Prepend(Slice(v))
	default:				l := s.Len() + 1
							n := make(Slice, l, l)
							n[0] = v
							copy(n[1:], *s)
							*s = n
	}
}

func (s *Slice) AppendSlice(v interface{}) {
	*s = append(*s, v)
}

func (s *Slice) PrependSlice(v interface{}) {
	l := s.Len() + 1
	n := make(Slice, l, l)
	n[0] = v
	copy(n[1:], *s)
	*s = n
}

func (s Slice) Repeat(count int) Slice {
	length := len(s) * count
	capacity := cap(s)
	if capacity < length {
		capacity = length
	}
	destination := make(Slice, length, capacity)
	for start, end := 0, len(s); count > 0; count-- {
		copy(destination[start:end], s)
		start = end
		end += len(s)
	}
	return destination
}

func (s *Slice) Flatten() {
	if s != nil {
		n := make(Slice, 0, len(*s))
		for _, v := range *s {
			switch v := v.(type) {
			case Slice:				(&v).Flatten()
									n = append(n, v...)
			case []interface{}:		n = append(n, v...)
			case Flattenable:		v.Flatten()
									n = append(n, v)
			default:				n = append(n, v)
			}
		}
		*s = n
	}
}

func (s Slice) equal(o Slice) (r bool) {
	if len(s) == len(o) {
		r = true
		for i, v := range s {
			switch v := v.(type) {
			case Equatable:		r = v.Equal(o[i])
			default:			r = v == o[i]
			}
			if !r {
				return
			}
		}
	}
	return
}

func (s Slice) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case Slice:				r = s.equal(o)
	case []interface{}:		r = s.equal(o)
	}
	return
}

func (s Slice) Car() (h interface{}) {
	if s.Len() > 0 {
		h = s[0]
	}
	return
}

func (s Slice) Cdr() (t Slice) {
	if s.Len() > 1 {
		t = s[1:]
	}
	return
}

func (s *Slice) Rplaca(v interface{}) {
	switch {
	case s == nil:			*s = Slice{v}
	case s.Len() == 0:		*s = append(*s, v)
	default:				(*s)[0] = v
	}
}

func (s *Slice) Rplacd(v interface{}) {
	if s == nil {
		*s = Slice{v}
	} else {
		ReplaceSlice := func(v Slice) {
			if l := len(v); l < cap(*s) {
				copy((*s)[1:], v)
				*s = (*s)[:l + 1]
			} else {
				l++
				n := make(Slice, l, l)
				copy(n, (*s)[:1])
				copy(n[1:], v)
				*s = n
			}
		}

		switch v := v.(type) {
		case *Slice:			ReplaceSlice(*v)
		case Slice:				ReplaceSlice(v)
		case *[]interface{}:	ReplaceSlice(Slice(*v))
		case []interface{}:		ReplaceSlice(Slice(v))
		case nil:				*s = (*s)[:1]
		default:				(*s)[1] = v
								*s = (*s)[:2]
		}
	}
}

func (s Slice) SetIntersection(o Slice) (r Slice) {
	cache := make(map[interface{}]bool)
	for _, v := range s {
		if ok := cache[v]; !ok {
			cache[v] = true
		}
	}
	for _, v := range o {
		if _, ok := cache[v]; ok {
			delete(cache, v)
			r = append(r, v)
		}
	}
	return
}

func (s Slice) SetUnion(o Slice) (r Slice) {
	cache := make(map[interface{}]bool)
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

func (s Slice) SetDifference(o Slice) (r Slice) {
	left := make(map[interface{}]bool)
	right := make(map[interface{}]bool)
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
			delete(right, k)
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

func (s Slice) Find(v interface{}) (i int, found bool) {
	for j, x := range s {
		if x == v {
			i = j
			found = true
			break
		}
	}
	return
}

func (s Slice) FindN(v interface{}, n int) (i ISlice) {
	i = make(ISlice, 0, 0)
	for j, x := range s {
		if x == v {
			i = append(i, j)
			if len(i) == n {
				break
			}
		}
	}
	return
}

func (s *Slice) KeepIf(f interface{}) {
	a := *s
	p := 0
	switch f := f.(type) {
	case func(interface{}) bool:	for i, v := range a {
										if i != p {
											a[p] = v
										}
										if f(v) {
											p++
										}
									}

	default:						for i, v := range a {
										if i != p {
											a[p] = v
										}	
										if v == f {
											p++
										}
									}
	}
	*s = a[:p]
}

func (s Slice) ReverseEach(f interface{}) {
	switch f := f.(type) {
	case func(interface{}):					for i := len(s) - 1; i > -1; i-- { f(s[i]) }
	case func(int, interface{}):			for i := len(s) - 1; i > -1; i-- { f(i, s[i]) }
	case func(interface{}, interface{}):	for i := len(s) - 1; i > -1; i-- { f(i, s[i]) }
	}
}

func (s Slice) ReplaceIf(f interface{}, r interface{}) {
	replacement := r.(interface{})
	switch f := f.(type) {
	case func(interface{}) bool:	for i, v := range s {
										if f(v) {
											s[i] = replacement
										}
									}

	default:						for i, v := range s {
										if v == f {
											s[i] = replacement
										}
									}
	}
}

func (s *Slice) Replace(o interface{}) {
	switch o := o.(type) {
	case Slice:				*s = o
	case []interface{}:		*s = Slice(o)

	case []reflect.Value:	n := make(Slice, len(o), len(o))
							for i, v := range o {
								n[i] = v.Interface()
							}
							*s = n

	default:				if v := reflect.ValueOf(o); v.Kind() == reflect.Slice {
								vl := v.Len()
								n := make(Slice, vl, vl)
								for i := 0; i < vl; i++ {
									n[i] = v.Index(i).Interface()
								}
								*s = n
							} else {
								*s= Slice{ v.Interface() }
							}
	}
}

func (s Slice) Select(f interface{}) interface{} {
	r := make(Slice, 0, len(s) / 4)
	switch f := f.(type) {
	case func(interface{}) bool:	for _, v := range s {
										if f(v) {
											r = append(r, v)
										}
									}

	default:						for _, v := range s {
										if v == f {
											r = append(r, v)
										}
									}
	}
	return r
}

func (s *Slice) Uniq() {
	a := *s
	if len(a) > 0 {
		p := 0
		m := make(map[interface{}] bool)
		for _, v := range a {
			if ok := m[v]; !ok {
				m[v] = true
				a[p] = v
				p++
			}
		}
		*s = a[:p]
	}
}

func (s Slice) Shuffle() {
	l := len(s) - 1
	for i, _ := range s {
		r := i + rand.Intn(l - i)
		s.Swap(i, r)
	}
}

func (s Slice) ValuesAt(n ...int) interface{} {
	r := make(Slice, 0, len(n))
	for _, v := range n {
		r = append(r, s[v])
	}
	return r
}

func (s *Slice) Insert(i int, v interface{}) {
	switch v := v.(type) {
	case Slice:				l := s.Len() + len(v)
							n := make(Slice, l, l)
							copy(n, (*s)[:i])
							copy(n[i:], v)
							copy(n[i + len(v):], (*s)[i:])
							*s = n

	case []interface{}:		s.Insert(i, Slice(v))

	default:				l := s.Len() + 1
							n := make(Slice, l, l)
							copy(n, (*s)[:i])
							n[i] = v
							copy(n[i + 1:], (*s)[i:])
							*s = n
	}
}

func (s *Slice) Pop() (r interface{}, ok bool) {
	if end := s.Len() - 1; end > -1 {
		r = (*s)[end]
		*s = (*s)[:end]
		ok = true
	}
	return
}