package main // import github.com/selfish/googol-tree

import (
	"fmt"
	. "github.com/selfish/googol-tree/googol"
	"github.com/selfish/googol-tree/tree"
	. "github.com/selfish/googol-tree/util"
	"math"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	// _ = naive(0, 1000, true)

	_ = multiNaive(10000, 99, 100, true)
}

func naive(sampleSize int, cardCount int, optimizeSample bool) (DidWin bool) {

	if optimizeSample == true {
		sampleSize = int(math.Ceil(float64(cardCount) * (1 / math.E)))
	}

	blackboard := Blackboard{
		// Runtime configuration:
		SampleSize: sampleSize,
		// Runtime variables:
		ChosenValue: 0,
		MaxInSample: 0,
		Index:       0,
		// Game constants:
		NumberSet: RandomIntSlice(cardCount, 10000),
	}

	tree.Root(
		&blackboard,
		tree.Sequence(
			tree.While(TakeSample),
			tree.While(SeekHigher),
			tree.Task(CalculateResult),
		),
	)

	return blackboard.DidWin
}

func multiNaive(gameCount int, sampleSize int, cardCount int, optimizeSample bool) (winRate float64) {

	// Create array of blackboards:
	games := make([]bool, gameCount, gameCount)

	// Play games:
	for i := 0; i < gameCount; i++ {
		fmt.Printf("%d \t| ", i+1)
		games[i] = naive(sampleSize, cardCount, optimizeSample)
	}

	// Print Summary:
	wins := CountWins(games)
	winRate = float64(wins) / float64(gameCount) * 100

	fmt.Printf("--SUM--\t| Victory: %d \t| Losses: %d \t| Win rate: %.2f%% \t| e-rate: %.2f%%\n", wins, gameCount-wins, winRate, 100/math.E)
	return winRate
}

//func optimizer(gameCount int, sampleSize int, cardCount int, optimizeSample bool) (optimalSample float64) {
//
//
//	obb := OptiBlackboard{
//
//	}
//
// otree.Root(
//		&obb,
//		otree.Sequence(
//			otree.For(
//				0, 100, 10,
//				otree.Repeat(
//					100,
//					otree.Task(),
//				)
//			),
//			otree.For(
//				0, 100, 10,
//				otree.Task(),
//			),
//		),
// )
//
//	games := make([]Blackboard, gameCount, gameCount)
//	rand.Seed(time.Now().UnixNano())
//
//	for i := 0; i < gameCount; i++ {
//		games[i] = naive()
//	}
//
//	wins := CountWins(games)
//
//	fmt.Printf("Victory: %d \t| Losses: %d \t| Win rate: %.2f%%\n", wins, gameCount-wins, float64(wins)/float64(gameCount)*100)
//}
