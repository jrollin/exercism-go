package isogram

import (
	s "strings"
)

func IsIsogram(word string) bool {

	if len(word) == 0 {
		return true
	}

	noSpaceAndLower := s.Replace(s.ToLower(word), " ", "", -1)
	noSpaceAndLowerAndHyphen := s.Replace(noSpaceAndLower, "-", "", -1)
	for _, w := range noSpaceAndLowerAndHyphen {
		if s.Count(noSpaceAndLowerAndHyphen, string(w)) > 1 {
			return false
		}
	}
	return true
}
