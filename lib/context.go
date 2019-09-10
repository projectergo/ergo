package ergo

import (
	"errors"
	"reflect"
)

type Context struct {
	processes  map[int]*Process
	registered map[string]*Process
	functions  map[string][]*fun
	evalStack  *stack
}

func NewContext(functions ...*fun) *Context {
	context := &Context{
		functions: make(map[string][]*fun),
		evalStack: new(stack),
		processes: make(map[int]*Process),
	}
	for _, function := range functions {
		context.functions[function.name] = append(context.functions[function.name], function)
	}
	return context
}

func Define(
	name string,
	function interface{},
) *fun {
	ergoFunction, err := newFun(name, function)
	if err != nil {
		panic(err)
	}
	return ergoFunction
}

func Export(function *fun) *fun {
	function.exported = true
	return function
}

func Eval(name string, ctx *Context, args ...interface{}) (interface{}, error) {
	candidates := ctx.functions[name]
	for _, candidate := range candidates {
		if candidate.match(ctx, args...) {
			if ctx.evalStack.isEmpty() && !candidate.exported {
				return nil, errors.New("ergo function is not exported")
			}
			ctx.evalStack.push(candidate)

			var valueArgs = convertArgsToReflectValues(ctx, candidate, args)

			outputVals := candidate.call(valueArgs)

			_ = ctx.evalStack.pop()
			// End: Handle stack pop

			switch len(outputVals) {
			case 0:
				return nil, nil
			case 1:
				return outputVals[0].Interface(), nil
			default:
				vals := make([]interface{}, 0)
				for _, val := range outputVals {
					vals = append(vals, val.Interface())
				}
				break
			}
			return nil, nil
		}
	}
	return nil, errors.New("no function matched")
}

func (ctx *Context) IsRegistered(name string) bool {
	return ctx.registered[name] != nil
}

func (ctx *Context) addProcess(process *Process) {
	ctx.processes[process.pid] = process
}

func convertArgsToReflectValues(ctx *Context, candidate *fun, args []interface{}) []reflect.Value {
	var valueArgs []reflect.Value
	valueArgs = append(valueArgs, reflect.ValueOf(ctx))
	for i, arg := range args {
		var atomType = getTypeAtom()
		var parameterType = candidate.reflection.In(i + 1)
		if parameterType.Implements(atomType) {
			atom := getAtom(arg, parameterType)
			valueArgs = append(valueArgs, atom)
		} else {
			valueArgs = append(valueArgs, reflect.ValueOf(arg))
		}
	}
	return valueArgs
}
