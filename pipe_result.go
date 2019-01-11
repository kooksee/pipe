package pipe

import (
	"encoding/json"
	"reflect"
)

type _data struct {
	_values []reflect.Value
}

func (t *_data) String() string {
	if len(t._values) < 1 || !t._values[0].IsValid() {
		return ""
	}
	return t._values[0].String()
}

func (t *_data) Json() string {
	var _res []interface{}
	for _, _p := range t._values {
		if !_p.IsValid() {
			_res = append(_res, nil)
		} else {
			_res = append(_res, _p.Interface())
		}
	}

	dt, err := json.Marshal(_res)
	assert(err != nil, "data json error(%s)", err)

	return string(dt)
}
