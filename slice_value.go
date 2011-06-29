package slices

import "fmt"
import "reflect"

func MakeSlice(t interface{}, length, capacity int) (s *SliceValue) {
	switch t := t.(type) {
	case reflect.Type:	s = &SliceValue{ reflect.MakeSlice(t, length, capacity) }
	case Typed:			s = MakeSlice(t.Type(), length, capacity)
	default:			s = MakeSlice(reflect.ValueOf(t), length, capacity)
	}
	return
}

func VSlice(i interface{}) (s *SliceValue) {
	switch v := reflect.ValueOf(i); v.Kind() {
	case reflect.Slice:						if !v.CanAddr() {
												x := reflect.New(v.Type()).Elem()
												x.Set(v)
												v = x
											}
											s = &SliceValue{ v }

	case reflect.Ptr, reflect.Interface:	s = VSlice(v.Elem())
	default:								panic(v)
	}
	return
}

type SliceValue struct {
	reflect.Value
}

func (s *SliceValue) At(i int) interface{} {
	return s.Index(i).Interface()
}

func (s *SliceValue) Set(i int, value interface{}) {
	s.Index(i).Set(reflect.ValueOf(value))
}

func (s *SliceValue) Clear(i int) {
	s.Index(i).Set(reflect.Zero(s.Type().Elem()))
}

func (s *SliceValue) Each(f func(interface{})) {
	for i := 0; i < s.Len(); i++ {
		f(s.At(i))
	}
}

func (s *SliceValue) String() (t string) {
	s.Each(func( v interface{}) {
		if len(t) > 0 {
			t += " "
		}
		t += fmt.Sprintf("%v", v)
	})
	return fmt.Sprintf("(%v)", t)
}

func (s *SliceValue) BlockCopy(destination, source, count int) {
	reflect.Copy(s.Slice(destination, destination + count), s.Slice(source, source + count))
}

func (s *SliceValue) BlockClear(start, count int) {
	panic(s)
}

func (s *SliceValue) Overwrite(offset int, source interface{}) {
	switch source := source.(type) {
	case *SliceValue:		s.Overwrite(offset, *source)
	case SliceValue:		if offset == 0 {
								reflect.Copy(s.Value, source.Value)
							} else {
								reflect.Copy(s.Slice(offset, s.Len()), source.Value)
							}
	default:				switch v := reflect.ValueOf(source); v.Kind() {
							case reflect.Slice:		s.Overwrite(offset, VSlice(source))
							default:				s.Set(offset, v.Interface())
							}
	}
}

func (s *SliceValue) Reallocate(length, capacity int) {
	switch {
	case length > capacity:		s.Reallocate(capacity, capacity)

	case capacity != s.Cap():	x := reflect.MakeSlice(s.Type(), length, capacity)
								reflect.Copy(x, s.Value)
								s.Value = x

	default:					s.Value = s.Slice(0, length)
	}
}

//	SPLIT INTO Append() AND AppendSliceValue()

func (s *SliceValue) Append(i interface{}) {
	switch v := i.(type) {
	case *SliceValue:		s.Append(*v)
	case SliceValue:		s.setValue(reflect.AppendSlice(s.Value, v.Value))
	default:				switch v := reflect.ValueOf(i); v.Kind() {
							case reflect.Slice:		s.setValue(reflect.AppendSlice(s.Value, v))
							default:				s.setValue(reflect.Append(s.Value, v))
							}
	}
}

//	SPLIT INTO Prepend() AND PrependSliceValue()

func (s *SliceValue) Prepend(i interface{}) {
	l := s.Len()
	switch v := i.(type) {
	case *SliceValue:		s.Prepend(*v)
	case SliceValue:		l += v.Len()
							n := MakeSlice(s, l, l)
							n.Overwrite(0, v)
							n.Overwrite(v.Len(), s)
							s.setValue(n.Value)
	default:				switch v := reflect.ValueOf(i); v.Kind() {
							case reflect.Slice:		l += v.Len()
													n := MakeSlice(s, l, l)
													n.Overwrite(0, VSlice(i))
													n.Overwrite(v.Len(), s)
													s.setValue(n.Value)

							default:				l++
													n := MakeSlice(s, l, l)
													n.Overwrite(0, i)
													n.Overwrite(1, s)
													s.setValue(n.Value)
						}
	}
}

func (s *SliceValue) Repeat(count int) *SliceValue {
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

func (s *SliceValue) Flatten() {
	
}

func (s SliceValue) equal(o SliceValue) (r bool) {
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

func (s SliceValue) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case *SliceValue:		r = o != nil && s.equal(*o)
	case SliceValue:		r = s.equal(o)
	default:				if v := reflect.ValueOf(o); v.Type() == s.Type() {
								r = s.equal(SliceValue{ v })
							} else {
								r = s.Len() > 0 && s.At(0) == o
							}							
	}
	return
}

func (s *SliceValue) setValue(v reflect.Value) {
	if !s.CanAddr() {
		x := reflect.New(s.Type()).Elem()
		x.Set(s.Value)
		s.Value = x
	}
	s.Value = v
}

func (s *SliceValue) SetValue(i interface{}) {
	s.setValue(reflect.ValueOf(i))
}


func (s *SliceValue) Feed(c chan<- interface{}, f func(x interface{}) interface{}) {
	go func() {
		for i, l := 0, s.Len(); i < l; i++ {
			c <- f(s.Index(i).Interface())
		}
		close(c)
	}()
}

func (s *SliceValue) Pipe(f func(x interface{}) interface{}) (c chan interface{}) {
	c = make(chan interface{})
	s.Feed(c, f)
	return
}