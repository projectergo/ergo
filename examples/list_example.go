package examples

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

func RunListExamples() {
	list := []int{5, 3, 1, 2}
	println(list)
	res, err := Eval("reverse_list", listCtx, list)
	if err != nil {
		log.Fatal(err)
	}
	println(res)
}
