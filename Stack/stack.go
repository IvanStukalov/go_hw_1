package Stack

type (
	Stack struct {
		top *node
		len int
	}
	node struct {
		value byte
		prev  *node
	}
)

func New() *Stack {
	return &Stack{nil, 0}
}

func (stack *Stack) Len() int {
	return stack.len
}

func (stack *Stack) Peek() byte {
	return stack.top.value
}

func (stack *Stack) Pop() byte {
	ptr := stack.top
	stack.top = ptr.prev
	stack.len--
	return ptr.value
}

func (stack *Stack) Push(value byte) {
	newNode := &node{value, stack.top}
	stack.top = newNode
	stack.len++
}
