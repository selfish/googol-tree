package tree

import (
	. "github.com/selfish/googol-tree/googol"
)

type Node func(*Blackboard) bool
type Modifier func(*Blackboard) bool
type Condition func(*Blackboard) bool

func Root(blackboard *Blackboard, node Node) bool {
	return node(blackboard)
}


func Sequence(nodes ...Node) func(*Blackboard) bool {
	return func(blackboard *Blackboard) bool {
		for _, node := range nodes {
			if res := node(blackboard); res == false {
				return false
			}
		}
		return true
	}
}

func While(cond Condition) func(*Blackboard) bool {
	return func(blackboard *Blackboard) bool {
		for cond(blackboard) == true {
		}
		return true
	}
}

func Task(fn Modifier) func(*Blackboard) bool {
	return func(blackboard *Blackboard) bool {
		return fn(blackboard)
	}
}
