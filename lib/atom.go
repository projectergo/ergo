package ergo

import (
	"reflect"
)

type Atom interface {
	Match(val interface{}) bool
}

func getTypeAtom() reflect.Type {
	return reflect.TypeOf((*Atom)(nil)).Elem()
}

type AtomZero int

func (AtomZero) Match(val interface{}) bool {
	return reflect.DeepEqual(val, 0)
}

type AtomOne int

func (AtomOne) Match(val interface{}) bool {
	return reflect.DeepEqual(val, 1)
}

type AtomEmptyArray []interface{}

func (AtomEmptyArray) Match(val interface{}) bool {
	reflection := reflect.ValueOf(val)
	return (reflection.Kind() == reflect.Slice || reflection.Kind() == reflect.Array) &&
		reflection.Len() == 0
}

func getAtom(val interface{}, t reflect.Type) reflect.Value {
	if t.ConvertibleTo(reflect.TypeOf([]interface{}{})) {
		return reflect.ValueOf(make([]interface{}, 0))
	}
	if reflect.TypeOf(val).ConvertibleTo(t) {
		return reflect.ValueOf(val).Convert(t)
	}
	return reflect.Value{}
}
