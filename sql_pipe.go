package pipe

import (
	"fmt"
	"reflect"
	"strings"
)

func Sql(format string, a ...interface{}) *_func {
	return nil
}

type Model interface {
	Name() string
}

func SqlD(d interface{}) string {
	_t := reflect.TypeOf(d)
	switch _t.Kind() {
	case reflect.Int:
		return fmt.Sprintf("%d", d)
	case reflect.String:
		return fmt.Sprintf("'%s'", d)
	case reflect.Invalid:
		return "null"
	default:
		return ""
	}
}

func Insert(model Model) string {
	rv := reflect.ValueOf(model)
	assert(!rv.IsValid(), "params is nil")
	assert(rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Struct, "params type is not match")

	sql := "insert into " + model.Name()
	var _fields []string
	var _vs []string
	_t := rv.Elem().Type()
	for i := 0; i < rv.Elem().NumField(); i++ {
		fieldName := If(_t.Field(i).Tag.Get("json") == "", _t.Field(i).Tag.Get("db"), _t.Field(i).Tag.Get("json")).(string)
		_fields = append(_fields, fieldName)
		_vs = append(_vs, SqlD(rv.Elem().Field(i).Interface()))
	}
	sql += "(" + strings.Join(_fields, ",") + ")"
	sql += " value (" + strings.Join(_vs, ",") + ");"

	return sql
}
