/*
a set of words form square set if they read same in rows
as they do in columns. for example ["ball","area","lead","lady"].
form a word square because

b a l l
a r e a
l e a d
l a d y
They read the same in rows and columns.
*/
package DP

func IsSquare(list []string) bool {
	var transList []string
	length := len(list)
	for i := 0; i < length; i++ {

		var tempWord []byte

		for _, val := range list {

			tempWord = append(tempWord, val[i])
		}
		transList = append(transList, string(tempWord))
	}

	for key, val := range list {
		if val != transList[key] {
			return false
		}
	}
	return true
}
