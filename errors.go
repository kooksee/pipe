package pipe

import (
	"github.com/juju/errors"
)

func Assert(b bool, text string, args ...interface{}) {
	if b {
		panic(errors.Errorf(text, args).Error())
	}
}

func AssertErr(err error, format string, args ...interface{}) {
	if err == nil {
		return
	}

	e := errors.Annotatef(err, format, args...)
	e.(*errors.Err).SetLocation(1)
	panic(e.Error())
}
