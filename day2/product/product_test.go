package product_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/churmd/advent-of-code-2025/day2/product"
)

func TestIsInvalid(t *testing.T) {
	invalidtestCases := []int{11, 22, 99, 1010, 1188511885}
	for _, tc := range invalidtestCases {
		t.Run(fmt.Sprintf("invalid %d", tc), func(t *testing.T) {
			if !product.IsInvalid(tc) {
				t.Errorf("expected %d to be invalid", tc)
			}
		})
	}

	validtestCases := []int{101, 123344, 543545}
	for _, tc := range validtestCases {
		t.Run(fmt.Sprintf("valid %d", tc), func(t *testing.T) {
			if product.IsInvalid(tc) {
				t.Errorf("expected %d to be valid", tc)
			}
		})
	}
}

func TestIsInvalid2(t *testing.T) {
	invalidtestCases := []int{11, 22, 99, 1010, 1188511885, 446446, 38593859, 565656, 824824824, 2121212121}
	for _, tc := range invalidtestCases {
		t.Run(fmt.Sprintf("invalid %d", tc), func(t *testing.T) {
			if !product.IsInvalid2(tc) {
				t.Errorf("expected %d to be invalid", tc)
			}
		})
	}

	validtestCases := []int{101, 123344, 543545}
	for _, tc := range validtestCases {
		t.Run(fmt.Sprintf("valid %d", tc), func(t *testing.T) {
			if product.IsInvalid2(tc) {
				t.Errorf("expected %d to be valid", tc)
			}
		})
	}
}

func TestFindAllInvalidIDs(t *testing.T) {
	testcases := []struct {
		start    int
		end      int
		expected []int
	}{
		{
			start:    11,
			end:      22,
			expected: []int{11, 22},
		},
		{
			start:    95,
			end:      115,
			expected: []int{99},
		},

		{
			start:    998,
			end:      1012,
			expected: []int{1010},
		},
		{
			start:    1188511880,
			end:      1188511890,
			expected: []int{1188511885},
		},
		{
			start:    222220,
			end:      222224,
			expected: []int{222222},
		},
		{
			start:    446443,
			end:      446449,
			expected: []int{446446},
		},
		{
			start:    38593856,
			end:      38593862,
			expected: []int{38593859},
		},
		{
			start: 1698522,
			end:   1698528,
		},
	}
	for _, tc := range testcases {
		t.Run(fmt.Sprintf("start %d end %d", tc.start, tc.end), func(t *testing.T) {
			p := product.ProductIDRange{
				Start: tc.start,
				End:   tc.end,
			}
			result := product.FindAllInvalidIDs(p, product.IsInvalid)
			if !slices.Equal(result, tc.expected) {
				t.Errorf("expected:\n%v\ngot:\n%v\n", tc.expected, result)
			}
		})
	}
}
