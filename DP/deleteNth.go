/* this function checks if the given array has more than the
   number of repetitions for any given element. If it does
   have, then it will remove those repetitions and return
   the cleared array*/
package DP

func LimitRep(arr []int, limit int) []int {
	var checkRep = make(map[int]int)
	var result []int
	for i := 0; i < len(arr); i++ {
		if checkRep[arr[i]] < limit {
			checkRep[arr[i]]++
			result = append(result, arr[i])

		}

	}
	return result
}
