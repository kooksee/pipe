package pipe

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

func ToInt(p string) int {
	r, err := strconv.Atoi(p)
	assert(err != nil, "can not convert %s to int,error(%s)", p, err)
	return r
}

func InterfaceOf(ps ...interface{}) []interface{} {
	return ps
}

func assert(b bool, text string, args ...interface{}) {
	if b {
		panic(fmt.Sprintf(text, args...))
	}
}

func If(b bool, trueVal, falseVal interface{}) interface{} {
	_t1 := reflect.ValueOf(trueVal)
	_t2 := reflect.ValueOf(falseVal)

	assert(_t1.Kind() == reflect.Func && _t1.Type().NumOut() != 1, "the output must be one")
	assert(_t2.Kind() == reflect.Func && _t2.Type().NumOut() != 1, "the output must be one")

	var _res reflect.Value
	if b {
		_res = _t1
	} else {
		_res = _t2
	}

	if _res.Kind() == reflect.Func {
		_res = _res.Call([]reflect.Value{})[0]
	}

	if !_res.IsValid() {
		return nil
	}

	return _res.Interface()
}

func IsError(p interface{}) bool {
	if p == nil {
		return false
	}

	_, ok := p.(error)
	return ok
}

func IsPtr(p interface{}) bool {
	return reflect.TypeOf(p).Kind() == reflect.Ptr
}

func Type(p interface{}) reflect.Kind {
	return reflect.TypeOf(p).Kind()
}

func Fn(f interface{}, params ...interface{}) func() interface{} {
	return func() interface{} {
		t := reflect.TypeOf(f)
		assert(t.Kind() != reflect.Func, "err -> Wrap: please input func")

		var vs []reflect.Value
		for i, p := range params {
			if p == nil {
				vs = append(vs, reflect.New(t.In(i)).Elem())
			} else {
				vs = append(vs, reflect.ValueOf(p))
			}
		}

		out := reflect.ValueOf(f).Call(vs)
		if !out[0].IsValid() {
			return nil
		}

		return out[0]
	}
}

func assertFn(fn interface{}) {
	assert(fn == nil, "the func is nil")

	_v := reflect.ValueOf(fn)
	assert(_v.Kind() != reflect.Func, "the params(%s) is not func type", _v.Type())
}

func P(d ...interface{}) {
	for _, i := range d {
		dt, err := json.MarshalIndent(i, "", "\t")
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(reflect.ValueOf(i).String(), "->", string(dt))
	}
}
