package examples

import (
	. "ergo/lib"
	"log"
)

var (
	basicCtx = NewContext(
		Export(
			Define("myFunc", func(context *Context, text string) {
				println(text)
			})),
		Export(
			Define("myFunc", func(context *Context, i int) int {
				return i + 1
			}),
		),
		Define("factorial", func(context *Context, n AtomZero) int {
			return 1
		}),
		Export(
			Define("factorial", func(context *Context, n int) int {
				if res, err := Eval("factorial", context, n-1); err != nil {
					log.Fatal(err)
				} else {
					return n * res.(int)
				}
				return 0
			}),
		),
	)
)

func RunBasicExamples() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Spawn a concurrent process
	process := basicCtx.Spawn("factorial", 10)
	go func(proc *Process) {
		res := <-process.Result()
		if res.Error != nil {
			log.Fatal(res.Error)
		}
		println(res.Result.(int))
	}(process)

	// Prints Hello world
	_, err := Eval("myFunc", basicCtx, "Hello World")
	if err != nil {
		log.Fatal(err)
	}

	// Returns 4 + 1
	res, err := Eval("myFunc", basicCtx, 4)
	if err != nil {
		log.Fatal(err)
	}
	println(res.(int))

	// Calculates 5! recursively
	res, err = Eval("factorial", basicCtx, 5)
	if err != nil {
		log.Fatal(err)
	}
	println(res.(int))
}
