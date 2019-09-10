package ergo

type stack struct {
	head *node
}

type node struct {
	val  *fun
	next *node
}

func newNode(fn *fun) *node {
	return &node{val: fn, next: nil}
}

func (stack *stack) push(fn *fun) {
	node := newNode(fn)
	if stack.head == nil {
		stack.head = node
		return
	}
	node.next = stack.head
	stack.head = node
}

func (stack *stack) pop() *fun {
	if stack.head == nil {
		return nil
	}
	fn := stack.head
	stack.head = stack.head.next
	return fn.val
}

func (stack *stack) isEmpty() bool {
	return stack.head == nil
}
