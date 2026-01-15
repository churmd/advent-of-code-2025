package product

import (
	"strconv"
	"strings"
)

type ProductIDRange struct {
	Start int
	End   int
}

func SumInvalidIDs(proudctRanges []ProductIDRange, isInvalid func(n int) bool) int {
	total := 0
	for _, p := range proudctRanges {
		for _, i := range FindAllInvalidIDs(p, isInvalid) {
			total += i
		}
	}
	return total
}

func FindAllInvalidIDs(p ProductIDRange, isInvalid func(n int) bool) []int {
	var result []int
	for i := p.Start; i <= p.End; i++ {
		if isInvalid(i) {
			result = append(result, i)
		}
	}
	return result
}

func IsInvalid(n int) bool {
	s := strconv.Itoa(n)
	if len(s)%2 != 0 {
		return false
	}

	midpoint := len(s) / 2
	left := s[0:midpoint]
	right := s[midpoint:]
	return left == right
}

func IsInvalid2(n int) bool {
	s := strconv.Itoa(n)
	length := len(s)
	for i := 1; i <= len(s)/2; i++ {
		quotient := length / i
		remainder := length % i
		if remainder == 0 {
			prefix := s[0:i]
			repeat := strings.Repeat(prefix, quotient)
			if repeat == s {
				return true
			}
		}
	}
	return false
}
