package pipe

import (
	"fmt"
	"reflect"
	"strconv"
)

func ToInt(p string) int {
	r, err := strconv.Atoi(p)
	assert(err != nil, fmt.Sprintf("can not convert %s to int,error(%s)", p, err.Error()))
	return r
}

func InterfaceOf(ps ...interface{}) []interface{} {
	return ps
}

func assert(b bool, text string) {
	if b {
		panic(text)
	}
}

func If(b bool, trueVal, falseVal interface{}) interface{} {
	if b {
		return trueVal
	}
	return falseVal
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
