package pipe

import (
	"errors"
	"reflect"
)

type _try struct {
	err    error
	params []reflect.Value
}

func (t *_try) Then(fn interface{}) *_try {
	assertFn(fn)

	if t.err != nil {
		return t
	}

	_fn := reflect.ValueOf(fn)
	_t := _fn.Type()

	//assert(len(t.params) != _t.NumIn(), "the params num is not match(%d,%d)", len(t.params), _t.NumIn())
	//assert(_t.NumOut() != 0, "the output params num is not 0(%d)", _t.NumOut())

	var _res []reflect.Value
	for i, p := range t.params {
		if !p.IsValid() {
			p = reflect.New(_t.In(i)).Elem()
		}
		_res = append(_res, p)
	}

	return &_try{params: _fn.Call(_res)}
}

func (t *_try) Catch(fn func(err error)) {

	if t.err == nil {
		return
	}

	fn(t.err)
}

func Try(fn interface{}) *_try {
	assertFn(fn)

	t := &_try{}
	defer func() {
		defer func() {
			if r := recover(); r != nil {
				switch r.(type) {
				case error:
					t.err = r.(error)
				case string:
					t.err = errors.New(r.(string))
				}
			}
		}()
		t.params = reflect.ValueOf(fn).Call([]reflect.Value{})
	}()
	return t
}
