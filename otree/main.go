package otree

import (
	. "github.com/selfish/googol-tree/googol"
)

type Node func(*OptiBlackboard) bool
type Modifier func(*OptiBlackboard) bool
type Condition func(*OptiBlackboard) bool

func Root(blackboard *OptiBlackboard, node Node) bool {
	return node(blackboard)
}

func Sequence(nodes ...Node) func(*OptiBlackboard) bool {
	return func(blackboard *OptiBlackboard) bool {
		for _, node := range nodes {
			if res := node(blackboard); res == false {
				return false
			}
		}
		return true
	}
}

func For(start int, end int, step int, node Node) func(*OptiBlackboard) bool {
	return func(blackboard *OptiBlackboard) bool {
		for i := start; i <= end; i += step {
			node(blackboard)
		}
		return true
	}
}

func Repeat(count int, node Node) func(*OptiBlackboard) bool {
	return func(blackboard *OptiBlackboard) bool {
		for i := 0; i < count; i++ {
			node(blackboard)
		}
		return true
	}
}

func While(cond Condition) func(*OptiBlackboard) bool {
	return func(blackboard *OptiBlackboard) bool {
		for cond(blackboard) == true {
		}
		return true
	}
}

func Task(fn Modifier) func(*OptiBlackboard) bool {
	return func(blackboard *OptiBlackboard) bool {
		return fn(blackboard)
	}
}

func Null() func(*OptiBlackboard) bool {
	return func(_ *OptiBlackboard) bool {
		return true
	}
}
