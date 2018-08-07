/*
function to check if given two strings are k anagrams
of each other.

K anagrams are those strings whcich can be formed from
one string to another, if k characters are replaced.
if given two string `grammar` and `anagram` they are
k anagrams of each other if k=3. This means `grammar`
can change to `anagram` if less than three characters
are replaced, in this case it is `n`,`a`by `m`,`r`.
*/
package DP

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func KAnagrams(s1, s2 string, k int) bool {
	var c1 = make(map[int]int)
	var c2 = make(map[int]int)
	for _, val := range s1 {
		c1[int(val-'a')]++
	}

	for _, val := range s2 {
		c2[int(val-'a')]++
	}
	count := 0

	for key, val := range c1 {
		count = count + diff(val, c2[key])
	}
	if count > k {
		return false
	}
	return true
}
