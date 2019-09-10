package ergo

import (
	"errors"
	"reflect"
)

type fun struct {
	name       string
	function   interface{}
	reflection reflect.Type
	exported   bool
}

func newFun(name string, function interface{}) (*fun, error) {
	reflection := reflect.TypeOf(function)
	if reflection.Kind() != reflect.Func {
		return nil, errors.New("second Argument should be a function")
	}
	if reflection.NumIn() == 0 || reflection.In(0).String() != "*ergo.Context" {
		return nil, errors.New("function's first argument should be context")
	}
	return &fun{
		name:       name,
		function:   function,
		reflection: reflect.TypeOf(function),
		exported:   false,
	}, nil
}

func (fun *fun) match(context *Context, args ...interface{}) bool {
	var atomType = getTypeAtom()
	for i, arg := range args {
		var t = fun.reflection.In(i + 1)
		if t.Implements(atomType) {
			atom := getAtom(arg, t)
			match := callMatchMethod(t, atom, arg)
			if !match {
				return false
			}
		} else {
			if !reflect.DeepEqual(reflect.TypeOf(arg), fun.reflection.In(i+1)) {
				return false
			}
		}
	}
	return true
}

func (fun *fun) call(args []reflect.Value) []reflect.Value {
	fn := reflect.ValueOf(fun.function)
	return fn.Call(args)

}

func callMatchMethod(t reflect.Type, atom reflect.Value, val interface{}) bool {
	method, _ := t.MethodByName("Match")
	return method.Func.
		Call([]reflect.Value{
			atom,
			reflect.ValueOf(val),
		})[0].Bool()
}
