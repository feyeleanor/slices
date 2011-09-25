package slices

import (
	"fmt"
	"github.com/feyeleanor/raw"
	"rand"
	"reflect"
)

func RWrap(i interface{}) (s RSlice) {
	switch v := i.(type) {
	case *RSlice:		s = *v
	case RSlice:		s = v
	default:			if v := reflect.ValueOf(i); v.Kind() == reflect.Slice {
							if !v.CanAddr() {
								x := reflect.New(v.Type()).Elem()
								x.Set(v)
								v = x
							}
							s = RSlice{ v }
						} else {
							panic(v.Kind())
						}
	}
	return
}

func RList(n... interface{}) RSlice {
	return RSlice{ reflect.ValueOf(n) }
}

type RSlice struct {
	reflect.Value
}

func (s *RSlice) setValue(v reflect.Value) {
	s.Value = raw.MakeAddressable(s.Value)
	s.Value.Set(v)
}

func (s *RSlice) MakeAddressable()					{ s.Value = raw.MakeAddressable(s.Value) }
func (s *RSlice) SetValue(i interface{})			{ s.setValue(reflect.ValueOf(i)) }
func (s *RSlice) At(i int) interface{}				{ return s.Index(i).Interface() }
func (s *RSlice) Set(i int, value interface{})		{ s.Index(i).Set(reflect.ValueOf(value)) }
func (s *RSlice) VSet(i int, value reflect.Value)	{ s.Index(i).Set(value) }
func (s *RSlice) Clear(i int)						{ s.Index(i).Set(reflect.Zero(s.Type().Elem())) }

func (s RSlice) Swap(i, j int) {
	temp := s.Index(i).Interface()
	s.Index(i).Set(s.Index(j))
	s.Index(j).Set(reflect.ValueOf(temp))
}

func (s RSlice) RestrictTo(i, j int) {
	s.setValue(s.Slice(i, j))
}

func (s *RSlice) Cut(i, j int) {
	l := s.Len()
	if i < 0 {
		i = 0
	}
	if j > l {
		j = l
	}
	if j > i {
		if m := l - (j - i); m > 0 && l > m {
			reflect.Copy(s.Slice(i, m), s.Slice(j, l))
			for k := m; k < l; k++ {
				s.Clear(k)
			}
			s.MakeAddressable()
			s.SetLen(m)
		}
	}
}

func (s *RSlice) Trim(i, j int) {
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
			s.Clear(k)
		}
		s.MakeAddressable()
		s.SetLen(j - i)
	}
}

func (s *RSlice) Delete(i int) {
	n := s.Len()
	if i > -1 && i < n {
		l := n - 1
		reflect.Copy(s.Slice(i, l), s.Slice(i + 1, n))
		s.Clear(l)
		s.MakeAddressable()
		s.SetLen(l)
	}
}

func (s *RSlice) DeleteIf(f interface{}) {
	p := 0
	switch f := f.(type) {
	case reflect.Value:				switch f.Kind() {
									case reflect.Func:		ft := f.Type()
															if ft.NumIn() > 0 && ft.NumOut() > 0 && ft.Out(0).Kind() == reflect.Bool {
																for i := 0; i < s.Len(); i++ {
																	v := s.Index(i)
																	if i != p {
																		s.VSet(p, v)
																	}
																	if !f.Call([]reflect.Value{reflect.ValueOf(v.Interface())})[0].Bool() {
																		p++
																	}
																}
															} else {
																panic(f)
															}

									default:				for i := 0; i < s.Len(); i++ {
																v := s.Index(i)
																if i != p {
																	s.VSet(p, v)
																}
																if v.Interface() != f.Interface() {
																	p++
																}
															}
									}

	case func(reflect.Value) bool:	for i := 0; i < s.Len(); i++ {
										v := s.Index(i)
										if i != p {
											s.VSet(p, v)
										}
										if !f(v) {
											p++
										}
									}

	case func(interface{}) bool:	for i := 0; i < s.Len(); i++ {
										v := s.At(i)
										if i != p {
											s.Set(p, v)
										}
										if !f(v) {
											p++
										}
									}

	default:						s.DeleteIf(reflect.ValueOf(f))
									return
	}
	s.MakeAddressable()
	s.SetLen(p)
}

func (s RSlice) Each(f interface{}) {
	switch f := f.(type) {
	case func(reflect.Value):				for i := 0; i < s.Len(); i++ { f(s.Index(i)) }
	case func(int, reflect.Value):			for i := 0; i < s.Len(); i++ { f(i, s.Index(i)) }
	case func(interface{}, reflect.Value):	for i := 0; i < s.Len(); i++ { f(i, s.Index(i)) }
	case func(interface{}):					for i := 0; i < s.Len(); i++ { f(s.At(i)) }
	case func(int, interface{}):			for i := 0; i < s.Len(); i++ { f(i, s.At(i)) }
	case func(interface{}, interface{}):	for i := 0; i < s.Len(); i++ { f(i, s.At(i)) }
	}
}

func (s RSlice) String() (t string) {
	s.Each(func( v interface{}) {
		if len(t) > 0 {
			t += " "
		}
		t += fmt.Sprintf("%v", v)
	})
	return fmt.Sprintf("(%v)", t)
}

func (s RSlice) BlockCopy(destination, source, count int) {
	end := source + count
	if end > s.Len() {
		end = s.Len()
	}
	reflect.Copy(s.Slice(destination, s.Len()), s.Slice(source, end))
}

func (s RSlice) BlockClear(start, count int) {
	end := start + count
	if end > s.Len() {
		end = s.Len()
	}
	for i := start; i < end; i++ {
		s.Clear(i)
	} 
}

func (s RSlice) Overwrite(offset int, source interface{}) {
	switch source := source.(type) {
	case *RSlice:		s.Overwrite(offset, *source)
	case RSlice:		if offset == 0 {
							reflect.Copy(s.Value, source.Value)
						} else {
							reflect.Copy(s.Slice(offset, s.Len()), source.Value)
						}
	default:			switch v := reflect.ValueOf(source); v.Kind() {
						case reflect.Slice:		s.Overwrite(offset, RWrap(source))
						default:				s.Set(offset, v.Interface())
						}
	}
}

func (s *RSlice) Reallocate(length, capacity int) {
	switch {
	case length > capacity:		s.Reallocate(capacity, capacity)

	case capacity != s.Cap():	x := reflect.MakeSlice(s.Type(), length, capacity)
								reflect.Copy(x, s.Value)
								s.setValue(x)

	default:					s.setValue(s.Slice(0, length))
	}
}

func (s *RSlice) Extend(n int) {
	c := s.Cap()
	l := s.Len() + n
	if l > c {
		c = l
	}
	s.Reallocate(l, c)
}

func (s *RSlice) Expand(i, n int) {
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
		reflect.Copy(x.Slice(i + n, l), s.Slice(i, s.Len()))
		s.setValue(x)
	} else {
		for j := l - 1; j >= i; j-- {
			s.Index(j).Set(s.Index(j - n))
		}
		s.SetLen(l)
	}
}

func (s RSlice) Depth() (c int) {
	for i := s.Len() - 1; i > -1; i-- {
		if v, ok := s.At(i).(Nested); ok {
			if r := v.Depth() + 1; r > c {
				c = r
			}
		}
	}
	return
}

func (s RSlice) Reverse() {
	end := s.Len() - 1
	for i := 0; i < end; i++ {
		s.Swap(i, end)
		end--
	}
}

func (s *RSlice) Append(v interface{}) {
	switch v := v.(type) {
	case reflect.Value:		s.setValue(reflect.Append(s.Value, v))
	case RSlice:			s.setValue(reflect.AppendSlice(s.Value, v.Value))
	default:				switch v := reflect.ValueOf(v); v.Kind() {
							case reflect.Slice:			s.setValue(reflect.AppendSlice(s.Value, v))
							default:					s.setValue(reflect.Append(s.Value, v))
							}
	}
}

func (s *RSlice) Prepend(v interface{}) {
	switch v := v.(type) {
	case reflect.Value:		l := s.Len() + 1
							n := RSlice{ reflect.MakeSlice(s.Type(), 0, l) }
							switch v.Kind() {
							case reflect.Slice:		n.setValue(reflect.AppendSlice(n.Value, v))
							default:				n.setValue(reflect.Append(n.Value, v))
							}
							n.setValue(reflect.AppendSlice(n.Value, s.Value))
							s.setValue(n.Value)

	case RSlice:			l := s.Len() + v.Len()
							n := RSlice{ reflect.MakeSlice(s.Type(), 0, l) }
							n.setValue(reflect.AppendSlice(n.Value, v.Value))
							n.setValue(reflect.AppendSlice(n.Value, s.Value))
							s.setValue(n.Value)

	default:				s.Prepend(reflect.ValueOf(v))
	}
}

func (s *RSlice) Repeat(count int) *RSlice {
	length := s.Len() * count
	capacity := s.Cap()
	if capacity < length {
		capacity = length
	}
	destination := RSlice{ reflect.MakeSlice(s.Type(), length, capacity) }
	for start, end := 0, s.Len(); count > 0; count-- {
		reflect.Copy(destination.Slice(start, end), s.Value)
		start = end
		end += s.Len()
	}
	return &destination
}

func (s *RSlice) Flatten() {
	if CanFlatten(s) {
		sl := s.Len()
		n := reflect.MakeSlice(s.Type(), 0, sl)
		for i := 0; i < sl; i++ {
			v := s.At(i)
			if v, ok := v.(Flattenable); ok {
				v.Flatten()
			}
			switch v := v.(type) {
			case RSlice:				n = reflect.AppendSlice(n, v.Value)
			case []reflect.Value:		n = reflect.AppendSlice(n, reflect.ValueOf(VSlice(v)))
			case VSlice:				n = reflect.AppendSlice(n, reflect.ValueOf(v))
			case reflect.Value:			if v.Kind() == reflect.Slice {
											n = reflect.AppendSlice(n, v)
										} else {
											n = reflect.Append(n, v)
										}
			default:					if v := reflect.ValueOf(v); v.Kind() == reflect.Slice {
											n = reflect.AppendSlice(n, v)
										} else {
											n = reflect.Append(n, v)
										}
			}
		}
		s.Value = n
	}
}

func (s RSlice) equal(o RSlice) (r bool) {
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

func (s RSlice) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case RSlice:			r = s.equal(o)
	default:				if v := reflect.ValueOf(o); v.Type() == s.Type() {
								r = s.equal(RSlice{ v })
							} else {
								r = s.Len() > 0 && s.At(0) == o
							}							
	}
	return
}

func (s RSlice) Car() (h interface{}) {
	if s.Len() > 0 {
		h = s.At(0)
	}
	return
}

func (s RSlice) Cdr() (t RSlice) {
	if s.Len() > 1 {
		t.Value = s.Slice(1, s.Len())
	} else {
		t.Value = reflect.MakeSlice(s.Type(), 0, 0)
	}
	return
}

func (s *RSlice) Rplaca(v interface{}) {
	switch {
	case s == nil:			s.setValue(RWrap(Slice{v}).Value)
	case s.Len() == 0:		s.Append(v)
	default:				s.Set(0, v)
	}
}

func (s *RSlice) Rplacd(v interface{}) {
	if s == nil {
		s.setValue(RWrap(Slice{v}).Value)
	} else {
		s.MakeAddressable()
		ReplaceSlice := func(v RSlice) {
			if l := v.Len(); l < s.Cap() {
				reflect.Copy(s.Slice(1, s.Len()), v.Value)
				s.SetLen(l + 1)
			} else {
				l++
				n := reflect.MakeSlice(s.Type(), l, l)
				n.Index(0).Set(s.Index(0))
				reflect.Copy(n.Slice(1, l), v.Value)
				s.Value = n
			}
		}

		switch x := v.(type) {
		case reflect.Value:		s.Set(1, x)
								s.SetLen(2)
		case RSlice:			ReplaceSlice(x)
		case []reflect.Value:	ReplaceSlice(RWrap(x))
		case nil:				s.SetLen(1)
		default:				s.Set(1, v)
								s.SetLen(2)
		}
	}
}

func (s RSlice) SetIntersection(o RSlice) (r RSlice) {
	cache := make(map[interface{}]bool)
	s.Each(func(v interface{}) {
		if ok := cache[v]; !ok {
			cache[v] = true
		}
	})

	results := []interface{}{}
	o.Each(func(v interface{}) {
		if _, ok := cache[v]; ok {
			cache[v] = false, false
			results = append(results, v)
		}
	})
	return RList(results...)
}

func (s RSlice) SetUnion(o RSlice) (r RSlice) {
	cache := make(map[interface{}]bool)
	s.Each(func(v interface{}) {
		if ok := cache[v]; !ok {
			cache[v] = true
		}
	})
	o.Each(func(v interface{}) {
		if ok := cache[v]; !ok {
			cache[v] = true
		}
	})

	results := []interface{}{}
	for k, _ := range cache {
		results = append(results, k)
	}
	return RList(results...)
}

func (s RSlice) SetDifference(o RSlice) (r RSlice) {
	left := make(map[interface{}]bool)
	right := make(map[interface{}]bool)
	s.Each(func(v interface{}) {
		if ok := left[v]; !ok {
			left[v] = true
		}
	})
	o.Each(func(v interface{}) {
		if ok := right[v]; !ok {
			right[v] = true
		}
	})

	results := []interface{}{}
	for k, _ := range left {
		if ok := right[k]; ok {
			right[k] = false, false
		} else {
			results = append(results, k)
		}
	}
	for k, _ := range right {
		if ok := left[k]; !ok {
			results = append(results, k)
		}
	}
	return RList(results...)
}

func (s RSlice) Find(v interface{}) (i int, found bool) {
	for j := 0; j < s.Len(); j++ {
		if s.At(j) == v {
			i = j
			found = true
			break
		}
	}
	return
}

func (s RSlice) FindN(v interface{}, n int) (i ISlice) {
	i = make(ISlice, 0, 0)
	for j := 0; j < s.Len(); j++ {
		if s.At(j) == v {
			i = append(i, j)
			if len(i) == n {
				break
			}
		}
	}
	return
}

func (s *RSlice) KeepIf(f interface{}) {
	p := 0
	l := s.Len()
	switch f := f.(type) {
	case reflect.Value:				for i := 0; i < l; i++ {
										v := s.Index(i)
										if i != p {
											s.VSet(p, v)
										}
										if v.Interface() == f.Interface() {
											p++
										}
									}

	case func(reflect.Value) bool:	for i := 0; i < l; i++ {
										v := s.Index(i)
										if i != p {
											s.VSet(p, v)
										}
										if f(v) {
											p++
										}
									}

	case func(interface{}) bool:	for i := 0; i < l; i++ {
										v := s.Index(i)
										if i != p {
											s.VSet(p, v)
										}
										if f(v.Interface()) {
											p++
										}
									}

	default:						for i := 0; i < l; i++ {
										v := s.Index(i)
										if i != p {
											s.VSet(p, v)
										}
										if v.Interface() == f {
											p++
										}
									}
	}
	s.MakeAddressable()
	s.SetLen(p)
}

func (s RSlice) ReverseEach(f interface{}) {
	switch f := f.(type) {
	case func(reflect.Value):					for i := s.Len() - 1; i > -1; i-- { f(s.Index(i)) }
	case func(int, reflect.Value):				for i := s.Len() - 1; i > -1; i-- { f(i, s.Index(i)) }
	case func(interface{}, reflect.Value):		for i := s.Len() - 1; i > -1; i-- { f(i, s.Index(i)) }
	case func(interface{}):						for i := s.Len() - 1; i > -1; i-- { f(s.At(i)) }
	case func(int, interface{}):				for i := s.Len() - 1; i > -1; i-- { f(i, s.At(i)) }
	case func(interface{}, interface{}):		for i := s.Len() - 1; i > -1; i-- { f(i, s.At(i)) }
	}
}

func (s RSlice) ReplaceIf(f interface{}, r interface{}) {
	var replacement		reflect.Value
	var ok 				bool

	if replacement, ok = r.(reflect.Value); !ok {
		replacement = reflect.ValueOf(r)
	}
	l := s.Len()
	switch f := f.(type) {
	case reflect.Value:				fi := f.Interface()
									for i := 0; i < l; i++ {
										if s.At(i) == fi {
											s.VSet(i, replacement)
										}
									}

	case func(reflect.Value) bool:	for i := 0; i < l; i++ {
										if f(s.Index(i)) {
											s.VSet(i, replacement)
										}
									}

	case func(interface{}) bool:	for i := 0; i < l; i++ {
										if f(s.At(i)) {
											s.VSet(i, replacement)
										}
									}

	default:						for i := 0; i < l; i++ {
										if s.At(i) == f {
											s.VSet(i, replacement)
										}
									}
	}
}

func (s *RSlice) Replace(o interface{}) {
	switch o := o.(type) {
	case reflect.Value:		*s = RSlice{o}
	case RSlice:			*s = o
	default:				*s = RWrap(o)
	}
}

func (s RSlice) Select(f interface{}) interface{} {
	l := s.Len()
	r := reflect.MakeSlice(s.Type(), 0, l / 4)
	switch f := f.(type) {
	case reflect.Value:				fi := f.Interface()
									for i := 0; i < l; i++ {
										v := s.Index(i)
										if v.Interface() == fi {
											r = reflect.Append(r, v)
										}
									}

	case func(reflect.Value) bool:	for i := 0; i < l; i++ {
										v := s.Index(i)
										if f(v) {
											r = reflect.Append(r, v)
										}
									}

	case func(interface{}) bool:	for i := 0; i < l; i++ {
										v := s.Index(i)
										if f(v.Interface()) {
											r = reflect.Append(r, v)
										}
									}

	default:						for i := 0; i < l; i++ {
										v := s.Index(i)
										if v.Interface() == f {
											r = reflect.Append(r, v)
										}
									}
	}
	return r.Interface()
}

func (s *RSlice) Uniq() {
	var v	reflect.Value
	var vi	interface{}

	l := s.Len()
	if l > 0 {
		p := 0
		m := make(map[interface{}] bool)
		for i := 0; i < l; i++ {
			v = s.Index(i)
			vi = v.Interface()
			if ok := m[vi]; !ok {
				m[vi] = true
				s.VSet(p, v)
				p++
			}
		}
		s.MakeAddressable()
		s.SetLen(p)
	}
}

func (s RSlice) Shuffle() {
	l := s.Len() - 1
	for i := 0; i < s.Len(); i++ {
		s.Swap(i, i + rand.Intn(l - i))
	}
}

func (s RSlice) ValuesAt(n ...int) interface{} {
	r := reflect.MakeSlice(s.Type(), 0, len(n))
	for _, v := range n {
		r = reflect.Append(r, s.Index(v))
	}
	return r.Interface()
}

func (s *RSlice) Insert(i int, v interface{}) {
	switch v := v.(type) {
	case reflect.Value:		l := s.Len() + 1
							n := reflect.MakeSlice(s.Type(), l, l)
							reflect.Copy(n, s.Slice(0, i))
							n.Index(i).Set(v)
							reflect.Copy(n.Slice(i + 1, l), s.Slice(i, l - 1))
							s.Value = n

	case RSlice:			l := s.Len() + v.Len()
							n := reflect.MakeSlice(s.Type(), l, l)
							reflect.Copy(n, s.Slice(0, i))
							reflect.Copy(n.Slice(i, l), v.Value)
							reflect.Copy(n.Slice(i + v.Len(), l), s.Slice(i, s.Len()))
							s.Value = n

	case []interface{}:		s.Insert(i, RWrap(v))
	default:				s.Insert(i, reflect.ValueOf(v))
	}
}