package pipe

import (
	"reflect"
)

func InterfaceOf(ps ...interface{}) []interface{} {
	return ps
}

func assert(b bool, text string) {
	if b {
		panic(text)
	}
}

func Default(b bool, d interface{}) interface{} {
	if !b{
		return
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
