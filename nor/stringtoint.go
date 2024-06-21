package nor

import (
	"errors"
	"strconv"
)

// conver string list to int list
func StringToInt(str []string) ([]int, error) {
	var res []int
	for _, n := range str {
		if !check(n) && n != "" {
			return nil, errors.New("Error!")
		}
		if n != "" {
			num, _ := strconv.Atoi(n)
			res = append(res, num)
		}

	}
	return res, nil
}

// check if any is of the data contain string or any illigal data
func check(str string) bool {
	for _, r := range str {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}