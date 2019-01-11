package pipe_test

import (
	"errors"
	"fmt"
	"github.com/kooksee/pipe"
	"testing"
)

type t1 struct {
	A string
	b int
}

func TestP(t *testing.T) {
	pipe.Data([]int{1, 2, 3}, []int{1, 2, 3}).P()
	pipe.Data(t1{A: "dd", b: 1}, &t1{A: "sss", b: 2}).P()
}

func TestFilter(t *testing.T) {
	pipe.Data(t1{A: "dd", b: 1}, &t1{A: "sss", b: 2}).Filter(func(i int, v interface{}) bool {
		return !pipe.IsPtr(v)
	}).P("test filter")
}

func TestMap(t *testing.T) {
	pipe.Data(t1{A: "dd", b: 1}, t1{A: "sss", b: 2}).Map(func(i int, v interface{}) interface{} {
		_t := v.(t1)
		_t.b = 100000
		return _t
	}).P("test map")

	pipe.Data(t1{A: "dd", b: 1}, t1{A: "sss", b: 2}).Map(func(i int, v interface{}) interface{} {
		_t := v.(t1)
		_t.b = 100000
		return _t
	}).Each(func(i int, a ...interface{}) {
		fmt.Println(a)
	})

	pipe.Data(t1{A: "dd", b: 1}, t1{A: "sss", b: 2}).Map(func(i int, v interface{}) interface{} {
		_t := v.(t1)
		_t.b = 100000
		return _t
	}).Pipe(func(a ...interface{}) {
		fmt.Println(a)
	})

	pipe.Data(nil, &t1{}).Map(func(i int, v *t1) *t1 {
		if v == nil {
			return nil
		}

		fmt.Println(v.b)

		v.b = 100000

		return v
	}).Map(func(v *t1) *t1 {
		if v == nil {
			return nil
		}

		fmt.Println("map2", v.b)
		v.b = 222000000
		return v
	}).Each(func(v interface{}) {
		fmt.Println(v)
	})
}

func TestArray(t *testing.T) {
	var ddd []int
	ddd = append(ddd, 1, 2, 34)
	pipe.DataArray(ddd).Each(func(i, n int) {
		fmt.Println(i, n)
	})

	pipe.DataRange(1, 100, 3).P()
	pipe.DataRange(1, 100, 3).Pipe(func(a ...int) []int {
		fmt.Println(a)
		return a
	}).ToSlice().Each(func(i, n int) {
		fmt.Println(i, n)
	})
}

func TestReduce(t *testing.T) {

	pipe.Data(t1{A: "dd", b: 1}, t1{A: "sss", b: 2}).Map(func(i int, v interface{}) interface{} {
		_t := v.(t1)
		_t.b = 100000
		return _t
	}).Reduce(func(s t1, v t1) int {
		return s.b + v.b
	}).Pipe(func(d int) int {
		fmt.Println("pppp", d)
		return d
	}).P("test reduce")
}

func TestEach(t *testing.T) {
	pipe.Data(1, 2, 3, 4).Each(func(a ...interface{}) {
		fmt.Println(a)
	})

	pipe.Data(1, 2, 3, 4).Each(func(i int, a interface{}) {
		fmt.Println(i, a)
	})

	pipe.Data(1, 2, 3, 4).Each(func(a interface{}) {
		fmt.Println(a)
	})
}

func TestPipe(t *testing.T) {
	pipe.Data(1, "dd").Pipe(func(a int, b string) (int, string) {
		fmt.Println("callback success ok", a, b)
		return a, b
	}).Pipe(func(a int, b string) {
		fmt.Println("callback ", a, b)
	})

	pipe.Data(1, 2, 3, 4).Pipe(func(a ...int) {
		fmt.Println(a)
	})
}

func TestIsError(t *testing.T) {
	fmt.Println(pipe.IsError(errors.New("")))
	fmt.Println(pipe.IsError(nil))

}

func TestError(t *testing.T) {
	//pipe.Data(1, 2, 3, errors.New("sss")).MustNotError()
	pipe.Data(1, 2, 3, nil).MustNotError()
}
