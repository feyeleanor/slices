package slices

import "fmt"
import "sort"

func SList(n... string) *SSlice {
	return (*SSlice)(&n)
}

type SSlice	[]string

func (s SSlice) Len() int							{ return len(s) }
func (s SSlice) Cap() int							{ return cap(s) }

func (s SSlice) At(i int) interface{}				{ return s[i] }
func (s SSlice) SAt(i int) string					{ return s[i] }
func (s SSlice) Set(i int, v interface{})			{ s[i] = v.(string) }
func (s SSlice) SSet(i int, v string)				{ s[i] = v }
func (s SSlice) Clear(i int)						{ s[i] = "" }
func (s SSlice) Swap(i, j int)						{ s[i], s[j] = s[j], s[i] }

func (s SSlice) Add(i, j int)						{ s[i] += s[j] }

func (s SSlice) Less(i, j int) bool					{ return s[i] < s[j] }
func (s SSlice) AtLeast(i, j int) bool				{ return s[i] <= s[j] }
func (s SSlice) Same(i, j int) bool					{ return s[i] == s[j] }
func (s SSlice) AtMost(i, j int) bool				{ return s[i] >= s[j] }
func (s SSlice) More(i, j int) bool					{ return s[i] > s[j] }

func (s SSlice) Sort()								{ sort.Sort(s) }

func (s *SSlice) RestrictTo(i, j int)				{ *s = (*s)[i:j] }

func (s SSlice) Compare(i, j int) (r int) {
	switch {
	case s[i] < s[j]:		r = IS_LESS_THAN
	case s[i] > s[j]:		r = IS_GREATER_THAN
	default:				r = IS_SAME_AS
	}
	return
}

func (s *SSlice) Cut(i, j int) {
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

func (s *SSlice) Trim(i, j int) {
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

func (s *SSlice) Delete(i int) {
	a := *s
	n := len(a)
	if i > -1 && i < n {
		copy(a[i:n - 1], a[i + 1:n])
		*s = a[0 : n - 1]
	}
}

func (s SSlice) Each(f func(interface{})) {
	for _, v := range s {
		f(v)
	}
}

func (s SSlice) EachWithIndex(f func(int, interface{})) {
	for i, v := range s {
		f(i, v)
	}
}

func (s SSlice) EachWithKey(f func(key, value interface{})) {
	for i, v := range s {
		f(i, v)
	}
}

func (s SSlice) SEach(f func(string)) {
	for _, v := range s {
		f(v)
	}
}

func (s SSlice) SEachWithIndex(f func(int, string)) {
	for i, v := range s {
		f(i, v)
	}
}

func (s SSlice) SEachWithKey(f func(interface{}, string)) {
	for i, v := range s {
		f(i, v)
	}
}

func (s SSlice) String() (t string) {
	for _, v := range s {
		if len(t) > 0 {
			t += " "
		}
		t += fmt.Sprintf("%v", v)
	}
	return fmt.Sprintf("(%v)", t)
}

func (s SSlice) BlockCopy(destination, source, count int) {
	end := source + count
	if end > len(s) {
		end = len(s)
	}
	copy(s[destination:], s[source:end])
}

func (s SSlice) BlockClear(start, count int) {
	copy(s[start:], make(SSlice, count, count))
}

func (s SSlice) Overwrite(offset int, container interface{}) {
	switch container := container.(type) {
	case SSlice:			copy(s[offset:], container)
	case []string:			copy(s[offset:], container)
	}
}

func (s *SSlice) Reallocate(length, capacity int) {
	switch {
	case length > capacity:		s.Reallocate(capacity, capacity)
	case capacity != cap(*s):	x := make(SSlice, length, capacity)
								copy(x, *s)
								*s = x
	default:					*s = (*s)[:length]
	}
}

func (s *SSlice) Extend(n int) {
	c := cap(*s)
	l := len(*s) + n
	if l > c {
		c = l
	}
	s.Reallocate(l, c)
}

func (s *SSlice) Expand(i, n int) {
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
		x := make(SSlice, l, c)
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

func (s SSlice) Reverse() {
	end := s.Len() - 1
	for i := 0; i < end; i++ {
		s[i], s[end] = s[end], s[i]
		end--
	}
}

func (s SSlice) Depth() int {
	return 0
}

func (s *SSlice) Append(v interface{}) {
	s.SAppend(v.(string))
}

func (s *SSlice) SAppend(v string) {
	*s = append(*s, v)
}

func (s *SSlice) AppendSlice(o SSlice) {
	*s = append(*s, o...)
}

func (s *SSlice) Prepend(v interface{}) {
	s.SPrepend(v.(string))
}

func (s *SSlice) SPrepend(v string) {
	l := s.Len() + 1
	n := make(SSlice, l, l)
	n[0] = v
	copy(n[1:], *s)
	*s = n
}

func (s *SSlice) PrependSlice(o SSlice) {
	l := s.Len() + o.Len()
	n := make(SSlice, l, l)
	copy(n, o)
	copy(n[o.Len():], *s)
	*s = n
}

func (s SSlice) Repeat(count int) SSlice {
	length := len(s) * count
	capacity := cap(s)
	if capacity < length {
		capacity = length
	}
	destination := make(SSlice, length, capacity)
	for start, end := 0, len(s); count > 0; count-- {
		copy(destination[start:end], s)
		start = end
		end += len(s)
	}
	return destination
}

func (s *SSlice) Flatten() {
	//	Flatten is a non-op for the SSlice as they cannot contain nested elements
}

func (s SSlice) equal(o SSlice) (r bool) {
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

func (s SSlice) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case *SSlice:			r = o != nil && s.equal(*o)
	case SSlice:			r = s.equal(o)
	case *[]string:			r = o != nil && s.equal(*o)
	case []string:			r = s.equal(o)
	}
	return
}

func (s SSlice) Car() (h interface{}) {
	if s.Len() > 0 {
		h = s[0]
	}
	return
}

func (s SSlice) Cdr() (t SSlice) {
	if s.Len() > 1 {
		t = s[1:]
	}
	return
}

func (s *SSlice) Rplaca(v interface{}) {
	switch {
	case s == nil:			*s = *SList(v.(string))
	case s.Len() == 0:		*s = append(*s, v.(string))
	default:				(*s)[0] = v.(string)
	}
}

func (s *SSlice) Rplacd(v interface{}) {
	if s == nil {
		*s = *SList(v.(string))
	} else {
		ReplaceSlice := func(v SSlice) {
			if l := len(v); l < cap(*s) {
				copy((*s)[1:], v)
				*s = (*s)[:l + 1]
			} else {
				l++
				n := make(SSlice, l, l)
				copy(n, (*s)[:1])
				copy(n[1:], v)
				*s = n
			}
		}

		switch v := v.(type) {
		case *SSlice:		ReplaceSlice(*v)
		case SSlice:		ReplaceSlice(v)
		case *[]string:		ReplaceSlice(SSlice(*v))
		case []string:		ReplaceSlice(SSlice(v))
		case nil:			*s = (*s)[:1]
		default:			(*s)[1] = v.(string)
							*s = (*s)[:2]
		}
	}
}

func (s SSlice) SetIntersection(o SSlice) (r SSlice) {
	cache := make(map[string]bool)
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

func (s SSlice) SetUnion(o SSlice) (r SSlice) {
	cache := make(map[string]bool)
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