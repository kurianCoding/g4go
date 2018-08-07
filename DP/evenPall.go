package DP

/*given a pallindromic number, this function will count the maximum occuring
digit. It will map the number to that digit, if its the only one. If there
are more than one, then it will map it to the least of those*/

func mapPall(num int) int {
	//this function finds the digit to which pallindrome is to be matched
	//find the occurences of each digit in the pallindrome
	countDig := make(map[int]int)
	temp := 0
	for i := 0; num > 0; i++ {
		temp = num % 10
		num = num / 10
		countDig[temp] = countDig[temp] + 1

	}
	temp = 0
	for _, val := range countDig {
		if temp < val {
			temp = val
		}

	}
	for _, val := range countDig {
		if val != temp {
			delete(countDig, val)
		}
	}
	temp = ^int(0) //max int
	for key, _ := range countDig {
		if temp < key {
			temp = key
		}

	}

	return temp
}
