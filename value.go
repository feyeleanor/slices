package slices

import "fmt"
import "reflect"

func MakeSlice(t interface{}, length, capacity int) (s *VSlice) {
	switch t := t.(type) {
	case reflect.Type:	s = &VSlice{ reflect.MakeSlice(t, length, capacity) }
	case Typed:			s = MakeSlice(t.Type(), length, capacity)
	default:			s = MakeSlice(reflect.ValueOf(t), length, capacity)
	}
	return
}

func VWrap(i interface{}) (s *VSlice) {
	switch v := i.(type) {
	case *VSlice:		s = v
	case VSlice:		s = &v
	default:			switch v := reflect.ValueOf(i); v.Kind() {
						case reflect.Slice:						if !v.CanAddr() {
																	x := reflect.New(v.Type()).Elem()
																	x.Set(v)
																	v = x
																}
																s = &VSlice{ v }

						case reflect.Ptr, reflect.Interface:	s = VWrap(v.Elem().Interface())
						default:								panic(v.Kind())
						}
	}
	return
}

func VList(n... interface{}) *VSlice {
	s := VSlice{ reflect.ValueOf(n) }
	return &s
}

type VSlice struct {
	reflect.Value
}

func (s *VSlice) setValue(v reflect.Value) {
	if !s.CanAddr() {
		x := reflect.New(s.Type()).Elem()
		x.Set(s.Value)
		s.Value = x
	}
	s.Value = v
}

func (s *VSlice) SetValue(i interface{})			{ s.setValue(reflect.ValueOf(i)) }
func (s *VSlice) At(i int) interface{}				{ return s.Index(i).Interface() }
func (s *VSlice) VAt(i int) reflect.Value			{ return s.Index(i) }
func (s *VSlice) Set(i int, value interface{})		{ s.Index(i).Set(reflect.ValueOf(value)) }
func (s *VSlice) VSet(i int, value reflect.Value)	{ s.Index(i).Set(value) }
func (s *VSlice) Clear(i int)						{ s.Index(i).Set(reflect.Zero(s.Type().Elem())) }

func (s VSlice) Swap(i, j int) {
	temp := s.Index(i).Interface()
	s.Index(i).Set(s.Index(j))
	s.Index(j).Set(reflect.ValueOf(temp))
}

func (s VSlice) RestrictTo(i, j int) {
	s.setValue(s.Slice(i, j))
}

func (s VSlice) Cut(i, j int) {
	n := s.Len()
	if i < 0 {
		i = 0
	}
	if j > n {
		j = n
	}
	if j > i {
		if m := n - (j - i); m > 0 && m <= n {
			reflect.Copy(s.Slice(i, m), s.Slice(j, n))
			for k := m; k < n; k++ {
				s.Clear(i)
			}
			s.SetLen(m)
		}
	}
}

func (s VSlice) Trim(i, j int) {
	n := s.Len()
	if i < 0 {
		i = 0
	}
	if j > n {
		j = n
	}
	if j > i {
		reflect.Copy(s.Value, s.Slice(i, j))
		for k, base := n - 1, i + 1; k > base; k-- {
			s.Clear(i)
		}
		s.SetLen(j - i)
	}
}

func (s VSlice) Delete(i int) {
	n := s.Len()
	if i > -1 && i < n {
		l := n - 1
		reflect.Copy(s.Slice(i, l), s.Slice(i + 1, n))
		s.Clear(l)
		s.SetLen(l)
	}
}

func (s *VSlice) Each(f func(interface{})) {
	for i := 0; i < s.Len(); i++ {
		f(s.At(i))
	}
}

func (s *VSlice) EachWithIndex(f func(int, interface{})) {
	for i := 0; i < s.Len(); i++ {
		f(i, s.At(i))
	}
}

func (s *VSlice) EachWithKey(f func(key, value interface{})) {
	for i := 0; i < s.Len(); i++ {
		f(i, s.At(i))
	}
}

func (s *VSlice) VEach(f func(reflect.Value)) {
	for i := 0; i < s.Len(); i++ {
		f(s.Index(i))
	}
}

func (s *VSlice) VEachWithIndex(f func(int, reflect.Value)) {
	for i := 0; i < s.Len(); i++ {
		f(i, s.Index(i))
	}
}

func (s *VSlice) VEachWithKey(f func(key interface{}, v reflect.Value)) {
	for i := 0; i < s.Len(); i++ {
		f(i, s.Index(i))
	}
}

func (s *VSlice) String() (t string) {
	s.Each(func( v interface{}) {
		if len(t) > 0 {
			t += " "
		}
		t += fmt.Sprintf("%v", v)
	})
	return fmt.Sprintf("(%v)", t)
}

func (s *VSlice) BlockCopy(destination, source, count int) {
	reflect.Copy(s.Slice(destination, destination + count), s.Slice(source, source + count))
}

func (s *VSlice) BlockClear(start, count int) {
	for i := start + count; i > start; i-- {
		s.Clear(i)
	} 
}

func (s *VSlice) Overwrite(offset int, source interface{}) {
	switch source := source.(type) {
	case *VSlice:		s.Overwrite(offset, *source)
	case VSlice:		if offset == 0 {
							reflect.Copy(s.Value, source.Value)
						} else {
							reflect.Copy(s.Slice(offset, s.Len()), source.Value)
						}
	default:			switch v := reflect.ValueOf(source); v.Kind() {
						case reflect.Slice:		s.Overwrite(offset, VWrap(source))
						default:				s.Set(offset, v.Interface())
						}
	}
}

func (s *VSlice) Reallocate(length, capacity int) {
	switch {
	case length > capacity:		s.Reallocate(capacity, capacity)

	case capacity != s.Cap():	x := reflect.MakeSlice(s.Type(), length, capacity)
								reflect.Copy(x, s.Value)
								s.setValue(x)

	default:					s.setValue(s.Slice(0, length))
	}
}

func (s *VSlice) Extend(n int) {
	c := s.Cap()
	l := s.Len() + n
	if l > c {
		c = l
	}
	s.Reallocate(l, c)
}

func (s *VSlice) Expand(i, n int) {
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
		x := reflect.MakeSlice(s.Type(), l, c)
		reflect.Copy(x, s.Slice(0, i))
		reflect.Copy(x.Slice(i + n, l - 1), s.Slice(i, s.Len() - 1))
		s.setValue(x)
	} else {
		for j := l - 1; j >= i; j-- {
			s.Index(j).Set(s.Index(j - n))
		}
		s.SetLen(l)
	}
}

func (s VSlice) Depth() (c int) {
	for i := s.Len(); i > 0; i-- {
		if v, ok := s.At(i).(Nested); ok {
			if r := v.Depth() + 1; r > c {
				c = r
			}
		}
	}
	return
}

func (s *VSlice) Reverse() {
	end := s.Len() - 1
	for i := 0; i < end; i++ {
		s.Swap(i, end)
		end--
	}
}

func (s *VSlice) Append(v interface{}) {
	s.setValue(reflect.Append(s.Value, reflect.ValueOf(v)))
}

func (s *VSlice) VAppend(v reflect.Value) {
	s.setValue(reflect.Append(s.Value, v))
}

func (s *VSlice) AppendSlice(o *VSlice) {
	s.setValue(reflect.AppendSlice(s.Value, o.Value))
}

func (s *VSlice) Prepend(v interface{}) {
	s.VPrepend(reflect.ValueOf(v))
}

func (s *VSlice) VPrepend(v reflect.Value) {
	l := s.Len() + 1
	n := VSlice{ reflect.MakeSlice(s.Type(), l, l) }
	n.VSet(0, v)
	n.Overwrite(1, s)
	s.setValue(n.Value)
}

func (s *VSlice) PrependSlice(o *VSlice) {
	l := s.Len() + o.Len()
	n := VSlice{ reflect.MakeSlice(s.Type(), l, l) }
	n.Overwrite(0, o)
	n.Overwrite(o.Len(), s)
	s.setValue(n.Value)
}

func (s *VSlice) Repeat(count int) *VSlice {
	length := s.Len() * count
	capacity := s.Cap()
	if capacity < length {
		capacity = length
	}
	destination := MakeSlice(s, length, capacity)
	for start, end := 0, s.Len(); count > 0; count-- {
		reflect.Copy(destination.Slice(start, end), s.Value)
		start = end
		end += s.Len()
	}
	return destination
}

func (s *VSlice) Flatten() {
	if s != nil {
		n := reflect.MakeSlice(s.Type(), 0, 0)
		for i := 0; i < s.Len(); i++ {
			if v, ok := s.At(i).(Flattenable); ok {
				v.Flatten()
				n = reflect.AppendSlice(n, reflect.ValueOf(v))
			} else {
				n = reflect.Append(n, reflect.ValueOf(v))
			}
		}
		s.Value = n
	}
}

func (s VSlice) equal(o VSlice) (r bool) {
	if s.Len() == o.Len() {
		r = true
		for i := s.Len() - 1; i > -1; i-- {
			switch v := s.At(i).(type) {
			case Equatable:		r = v.Equal(o.At(i))
			default:			r = v == o.At(i)
			}
			if !r {
				return
			}
		}
	}
	return
}

func (s VSlice) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case *VSlice:		r = o != nil && s.equal(*o)
	case VSlice:		r = s.equal(o)
	default:				if v := reflect.ValueOf(o); v.Type() == s.Type() {
								r = s.equal(VSlice{ v })
							} else {
								r = s.Len() > 0 && s.At(0) == o
							}							
	}
	return
}

func (s VSlice) Car() (h interface{}) {
	if s.Len() > 0 {
		h = s.At(0)
	}
	return
}

func (s VSlice) Cdr() (t VSlice) {
	if s.Len() > 1 {
		t.Value = s.Slice(1, s.Len() - 1)
	}
	return
}

func (s *VSlice) Rplaca(v interface{}) {
	switch {
	case s == nil:			s.setValue(VWrap(List(v)).Value)
	case s.Len() == 0:		s.Append(v)
	default:				s.Set(0, v)
	}
}

func (s *VSlice) Rplacd(v interface{}) {
	if s == nil {
		s.setValue(VWrap(List(v)).Value)
	} else {
		ReplaceSlice := func(v VSlice) {
			if l := v.Len(); l < s.Cap() {
				reflect.Copy(s.Slice(1, s.Len() - 1), v.Value)
				s.SetLen(l + 1)
			} else {
				l++
				n := reflect.MakeSlice(s.Type(), l, l)
				n.Index(0).Set(s.VAt(0))
				reflect.Copy(n.Slice(1, l - 1), v.Value)
				s.Value = n
			}
		}

		switch v := v.(type) {
		case *VSlice:			ReplaceSlice(*v)
		case VSlice:			ReplaceSlice(v)
		case *[]reflect.Value:	ReplaceSlice(*VWrap(*v))
		case []reflect.Value:	ReplaceSlice(*VWrap(v))
		case nil:				s.SetLen(1)
		default:				s.Set(1, v)
								s.SetLen(2)
		}
	}
}