package googol

import (
	"fmt"
	. "github.com/selfish/googol-tree/util"
)

type MultiBlackboard struct {
	Blackboards []Blackboard
}

type OptiBlackboard struct {
	Blackboards []Blackboard
}

type Blackboard struct {
	DidWin      bool
	SampleSize  int
	ChosenValue int
	MaxInSample int
	Index       int
	NumberSet   []int
}

func TakeSample(bb *Blackboard) bool {
	// Stop iteration if sample size reached
	if bb.Index > (bb.SampleSize - 1) {
		return false
	}

	// Update max if found higher:
	if bb.NumberSet[bb.Index] > bb.MaxInSample {
		bb.MaxInSample = bb.NumberSet[bb.Index]
	}

	// Update Iteration Status
	bb.Index = bb.Index + 1
	return true
}

func SeekHigher(bb *Blackboard) bool {
	// Stop iteration if reched end
	if bb.Index >= len(bb.NumberSet) {
		bb.ChosenValue = bb.NumberSet[len(bb.NumberSet)-1]
		return false
	}

	// Update max if found higher:
	if bb.NumberSet[bb.Index] > bb.MaxInSample {
		bb.ChosenValue = bb.NumberSet[bb.Index]
		return false
	}

	// Update Iteration Status:
	bb.Index = bb.Index + 1
	return true
}

func CalculateResult(bb *Blackboard) bool {
	bb.DidWin = bb.ChosenValue == MaxIntSlice(bb.NumberSet)
	fmt.Printf("Victory: %t \t| Chosen: %d \t| Max: %d\n", bb.DidWin, bb.ChosenValue, MaxIntSlice(bb.NumberSet))
	return bb.DidWin
}
