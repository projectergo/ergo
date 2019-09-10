# Ergo
STILL VERY MUCH WIP. PLEASE SEND YOUR FEEDBACK.
## Idea
Ergo is a library/framework that tries to encapsulate some of Erlang's
idioms and paradigms in Golang (hence the name, which also means "work" in Greek),
thus providing golang with a different approach to concurrency than the CSP (Communicating Sequential Processes)
model. This should in theory take advantage Erlang's design while leveraging
Go's deployability and performance. 
The idea with Ergo is eventually to be able to transfer any possible Erlang
to Go, and to mix and mash the two programming paradigms with ease to easily 
develop scalable and fault tolerant distributed systems.

### What can Ergo provide?
Ergo can provide:
 - Function overloading
 - Pure functional programming in Go
 - Hot Code Reloading (WIP)
 - Pattern Matching
 - Actor concurrency model (WIP)
 - Spawning distributed processes (WIP)
 
Currently you can:
 - Define and export functions in an Ergo context (something like a module)
 Functions that are not exported can only be called by functions defined in that
 same context.
 - Use one of the predefined atoms (AtomZero, AtomOne, AtomEmptyArray) for
 pattern matching
 - Create your own atoms for pattern matching.
 - Spawn an Ergo process.
 - Register an Ergo process and give it a name.
 
 
### Plans for Ergo.
Ergo will eventually supply:
 - Sending to/Receiving from other Ergo processes
 - An Ergo Node for deploying Ergo processes
 - Provide an std context that will contain most of the functions found in
 Ergo's stdlib.
 - And many others!
 
 
## Example Usage

```go
package main

import (
	. "ergo/lib"
	"log"
)

var (
	basicCtx = NewContext(
		Export(
			Define("factorial", func(context *Context, n AtomZero) int {
                            return 1
                        }),
		),
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

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
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
```

```go
package main

import (
	. "ergo/lib"
	"log"
)

var (
	listCtx = NewContext(
		Export(
			Define("reverse_list", func(ctx *Context, array AtomEmptyArray) []int {
				return []int{}
			}),
		),
		Export(
			Define("reverse_list", func(ctx *Context, array []int) []int {
				res, err := Eval("reverse_list", ctx, array[1:])
				if err != nil {
					log.Fatal(err)
				}
				arr := res.([]int)
				arr = append(arr, array[0])
				return arr
			}),
		),
	)
)

func main() {
	list := []int{5, 3, 1, 2}
	println(list)
	// Reverses list
	res, err := Eval("reverse_list", listCtx, list)
	if err != nil {
		log.Fatal(err)
	}
	println(res)
}

```