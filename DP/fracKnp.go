package DP

type Knap struct {
	weight int
	value  int
}
type KnapArr []Knap

func swap(a *Knap, b *Knap) {
	var temp *Knap
	temp = a
	a = b
	b = temp
	return
}
func SortByRatio(KnArray KnapArr, length int) {
	// cycle through Knap array
	//	find the value of value/weight
	//	sort the entries based on the ratio

	for j := 0; j < length; j++ {
		for i := 0; i < length; i++ {
			if KnArray[i].value/KnArray[i].weight > KnArray[i+1].value/KnArray[i].weight {
				swap(&KnArray[i], &KnArray[i+1])
			}
		}

	}
	return
}

func OptimumPurcahse(KnArray KnapArr, cumulative int, length int) int {
	var sum int = 0
	SortByRatio(KnArray, length)
	for i := 0; i < length; i++ {
		sum = sum + KnArray[i].value
		if cumulative <= sum {
			break
		}
	}
	//	keep adding value to result, until weight is less than cumulative
	// return result
	return sum
}
