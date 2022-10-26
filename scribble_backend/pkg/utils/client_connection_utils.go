package utils

import "errors"

func Remove(s []string, ID string) ([]string, error) {
	idx := -1
	for index, val := range s {
		if val == ID {
			idx = index
			break
		}
	}

	if idx == -1 {
		return s, errors.New("ID doesn't exits")
	}

	s[idx] = s[len(s)-1]
	return s[:len(s)-1], nil
}
