package DP

/*  Given a non-negative number represented as an array of digits,
plus one to the number.
*/
func PlusOne(number []int) []int {
	carry := (number[0] + 1) % 10
	number[0] = number[0] + 1
	var result []int
	lengthOfNumber := len(number)
	if carry == 1 {
		for i := 0; carry > 0 && i < lengthOfNumber; i++ {
			carry = (number[i] + carry) % 10
			result = append(result, (number[i]+carry)/10)
		}
	}
	if carry == 1 {
		result = append([]int{1}, result...)
	}
	return result
}
