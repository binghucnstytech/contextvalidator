package util

import "strings"

// this function take a string representation of nested containers
// and validate if a given structure was correctly structured
func IsContatinersDefinitionValid(src string) bool {
	if src == "" || len(src)%2 == 1 {
		return false
	}

	//expect that src includes 3 types of brackets only, that is, [], {} and ()
	separators := []string{"[]", "{}", "()"}

	for {
		originalLength := len(src)
		for _, sep := range separators {
			//remove available pairs
			src = strings.Join(strings.Split(src, sep), "")

			if len(src) == 0 {
				return true
			}
		}

		//length is not zero but no more pair
		if len(src) == originalLength {
			return false
		}
	}
}
