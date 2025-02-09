package similarity

import "github.com/emirpasic/gods/sets/hashset"

func JaccardSimilarity(x []string, y []string) float64 {

	sX := hashset.New()
	for _, elem := range x {
		sX.Add(elem)
	}

	sY := hashset.New()
	for _, elem := range y {
		sY.Add(elem)
	}

	return float64(sX.Intersection(sY).Size()) / float64(sX.Union(sY).Size())
}
