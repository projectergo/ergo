package ergo

var (
	pidCount = 0
)

type Process struct {
	pid    int
	ctx    *Context
	result chan struct {
		Result interface{}
		Error  error
	}
}

func newProcess(ctx *Context) *Process {
	pidCount++
	var process = &Process{pid: pidCount, ctx: ctx, result: make(chan struct {
		Result interface{}
		Error  error
	})}
	ctx.addProcess(process)
	return process
}

func (process *Process) Ctx() *Context {
	return process.ctx
}

func (process *Process) Pid() int {
	return process.pid
}

func (process *Process) Result() chan struct {
	Result interface{}
	Error  error
} {
	return process.result
}

func (ctx *Context) Spawn(name string, args ...interface{}) *Process {
	process := newProcess(ctx)
	go func(ctx *Context, process *Process, name string, args ...interface{}) {
		res, err := Eval(name, ctx, args...)
		if err != nil {
			process.Result() <- struct {
				Result interface{}
				Error  error
			}{Result: nil, Error: err}
		}
		process.Result() <- struct {
			Result interface{}
			Error  error
		}{Result: res, Error: nil}
	}(ctx, process, name, args...)
	return process
}

func (ctx *Context) Register(name string, process *Process) *Process {
	if process.IsRegistered() || ctx.IsRegistered(name) {
		return nil
	}
	ctx.registered[name] = process
	return process
}

func (process *Process) IsRegistered() bool {
	for _, proc := range process.ctx.registered {
		if proc == process {
			return true
		}
	}
	return false
}
