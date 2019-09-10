package main

import "ergo/examples"

func main() {
	end := make(chan struct{})
	//examples.RunBasicExamples()
	examples.RunListExamples()
	_ = <-end
}
