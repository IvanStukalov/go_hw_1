package Stack

type (
	Stack struct {
		top *node
		len int
	}
	node struct {
		value interface{}
		prev  *node
	}
)

func New() *Stack {
	return &Stack{nil, 0}
}

func (stack *Stack) Len() int {
	return stack.len
}

func (stack *Stack) Peek() interface{} {
	return stack.top.value
}

func (stack *Stack) Pop() interface{} {
	ptr := stack.top
	stack.top = ptr.prev
	stack.len--
	return ptr.value
}

func (stack *Stack) Push(value interface{}) {
	newNode := &node{value, stack.top}
	stack.top = newNode
	stack.len++
}
