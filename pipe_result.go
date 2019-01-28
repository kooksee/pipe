package pipe

import (
	"encoding/json"
	"reflect"
)

type _data struct {
	_values []reflect.Value
}

func (t *_data) String() string {
	if len(t._values) < 1 || !t._values[0].IsValid() || t._values[0].Kind() != reflect.String {
		return ""
	}
	return t._values[0].String()
}

func (t *_data) Value(v interface{}) error {
	return json.Unmarshal([]byte(t.Json()), v)
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
