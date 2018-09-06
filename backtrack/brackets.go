package geeksforgeeks

/* this function parses a string and takes out
all the non paired brackets
*/
func isBracket(in rune) bool {
	return in == ')' || in == '('
}

func isValidBrackets(in string) bool {
	var count int
	for _, val := range in {
		if val == '(' {
			count++
		}
		if val == ')' {
			count--
		}
		if count < 0 {
			return false
		}
	}
	return count == 0
}

func EleminateBrackets(in string) []string {
	// returns a set of possibiliies for
	// a valid string
	var visited map[string]bool
	var queue []string
	var res []string
	queue = append(queue, in) // append the initial string
	for len(queue) > 0 {

		temp := queue[len(queue)-1]
		visited[temp] = true
		queue = queue[:len(queue)-1]

		if isValidBrackets(temp) {
			res = append(res, temp)
		}
	}

}
