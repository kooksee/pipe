package pipe

import (
	"fmt"
	"reflect"
)

type _func struct {
	params []reflect.Value
}

func (t *_func) Pipe(fn interface{}) *_func {
	t.assert(fn)

	return &_func{params: reflect.ValueOf(fn).Call(t.params)}
}

func (t *_func) assert(fn interface{}) {
	assert(fn == nil, "the func is nil")

	_v := reflect.ValueOf(fn)
	assert(_v.Kind() != reflect.Func, fmt.Sprintf("the params(%s) is not func", _v.Type().Name()))
}

func (t *_func) P(tags ...string) {
	for _, p := range t.params {
		fmt.Println(p.Kind().String(), p.Type().String(), p.Interface())
	}

	_p := "\n"
	if len(tags) > 0 {
		_p = tags[0] + _p
	}
	fmt.Println(_p)
}

func (t *_func) Map(fn func(i int, v interface{}) interface{}) *_func {
	t.assert(fn)

	_f := &_func{}
	for i, p := range t.params {
		_f.params = append(_f.params, reflect.ValueOf(fn(i, p.Interface())))
	}
	return _f
}

func (t *_func) Reduce(fn interface{}) *_func {
	t.assert(fn)

	_fn := reflect.ValueOf(fn)
	if len(t.params) < 2 {
		return &_func{params: t.params}
	}

	rs := _fn.Call([]reflect.Value{t.params[0], t.params[1]})
	assert(len(rs) != 1, "must return one value")

	_res := rs[0]
	for i := 2; i < len(t.params); i++ {
		rs = _fn.Call([]reflect.Value{_res, t.params[i]})
		assert(len(rs) != 1, "must return one value")
		_res = rs[0]
	}
	return &_func{params: []reflect.Value{_res}}
}

func (t *_func) Any(fn func(v interface{}) bool) bool {
	t.assert(fn)

	for _, p := range t.params {
		if fn(p.Interface()) {
			return true
		}
	}
	return false
}

func (t *_func) Every(fn func(v interface{}) bool) bool {
	t.assert(fn)

	for _, p := range t.params {
		if !fn(p.Interface()) {
			return false
		}
	}
	return true
}

func (t *_func) MustNotError() {
	t.Each(func(_ int, v interface{}) {
		if IsError(v) {
			panic(v.(error).Error())
		}
	})
}

func (t *_func) FilterError() *_func {
	return t.Filter(func(_ int, v interface{}) bool {
		return !IsError(v)
	})
}

func (t *_func) Filter(fn func(i int, v interface{}) bool) *_func {
	t.assert(fn)

	var vs []reflect.Value
	for i, p := range t.params {
		if fn(i, p.Interface()) {
			vs = append(vs, p)
		}
	}
	return &_func{params: vs}
}

func (t *_func) ToSlice() *_func {
	var _ps []reflect.Value
	_ds := t.params[0]
	for i := 0; i < _ds.Len(); i++ {
		_ps = append(_ps, _ds.Index(i))
	}
	t.params = _ps
	return t
}

func (t *_func) Each(fn interface{}) {
	t.assert(fn)
	_fn := reflect.ValueOf(fn)
	_t := _fn.Type()
	assert(_t.NumIn() > 2, "the func input num is more than 2(%d)", _t.NumIn())
	assert(_t.NumOut() != 0, "the func output num is not equal(%d)", _t.NumOut())

	for i, p := range t.params {
		if p.Kind() == reflect.Invalid {
			if _t.NumIn() == 1 {
				p = reflect.New(_t.In(0)).Elem()
			}

			if _t.NumIn() == 2 {
				p = reflect.New(_t.In(1)).Elem()
			}
		}

		var _pi []reflect.Value
		if _t.NumIn() == 1 {
			_pi = []reflect.Value{p}
		}

		if _t.NumIn() == 2 {
			_pi = []reflect.Value{reflect.ValueOf(i), p}
		}

		_fn.Call(_pi)
	}
}

func DataRange(s, e, t int) *_func {
	var _ps []reflect.Value
	for i := s; i < e; i += t {
		_ps = append(_ps, reflect.ValueOf(i))
	}
	return &_func{
		params: _ps,
	}
}

func DataArray(ps interface{}) *_func {
	_d := reflect.ValueOf(ps)
	var _ps []reflect.Value
	for i := 0; i < _d.Len(); i++ {
		_ps = append(_ps, _d.Index(i))
	}
	return &_func{
		params: _ps,
	}
}

func Data(ps ...interface{}) *_func {
	var vs []reflect.Value
	for _, v := range ps {
		vs = append(vs, reflect.ValueOf(v))
	}
	return &_func{
		params: vs,
	}
}
