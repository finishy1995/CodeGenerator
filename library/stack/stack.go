package stack

import "container/list"

type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	return &Stack{
		list: list.New(),
	}
}

func (stack *Stack) Push(value interface{}) {
	stack.list.PushBack(value)
}

func (stack *Stack) Pop() interface{} {
	v := stack.list.Back()
	if v != nil {
		stack.list.Remove(v)
		return v.Value
	}
	return nil
}

func (stack *Stack) LazyPop() interface{} {
	v := stack.list.Back()
	if v != nil {
		return v.Value
	}
	return nil
}

func (stack *Stack) Len() int {
	return stack.list.Len()
}
