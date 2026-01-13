package dial_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/churmd/advent-of-code-2025/day1/dial"
)

func TestFollowRotations(t *testing.T) {
	t.Run("lists the result of each command", func(t *testing.T) {
		commands := []string{
			"L68",
			"L30",
			"R48",
			"L5",
			"R60",
			"L55",
			"L1",
			"L99",
			"R14",
			"L82",
		}
		actual, err := dial.FollowRotations(50, commands)
		if err != nil {
			t.Errorf("expected no error got %s", err)
		}
		expected := []int{82, 52, 0, 95, 55, 0, 99, 0, 14, 32}
		if !slices.Equal(actual, expected) {
			t.Errorf("Expected:\n%v\nGot:\n%v\n", expected, actual)
		}
	})
}

func TestRotate(t *testing.T) {
	testcases := []struct {
		start    int
		command  string
		expected int
	}{
		{
			start:    11,
			command:  "R8",
			expected: 19,
		},
		{
			start:    50,
			command:  "L68",
			expected: 82,
		},
		{
			start:    82,
			command:  "L30",
			expected: 52,
		},
		{
			start:    52,
			command:  "R48",
			expected: 0,
		},
	}
	for _, tc := range testcases {
		name := fmt.Sprintf("start %d command %s expected %d", tc.start, tc.command, tc.expected)
		t.Run(name, func(t *testing.T) {
			actual, err := dial.Rotate(tc.start, tc.command)
			if err != nil {
				t.Errorf("expected no error got %s", err)
			}
			if actual != tc.expected {
				t.Errorf("expected %d got %d", tc.expected, actual)
			}
		})
	}

	invalidCommandCases := []string{
		"r76",
		"l786",
		"R",
		"L",
		"543543",
		"fdgdfg",
		"LR567",
		"R65L",
	}
	for _, tc := range invalidCommandCases {
		t.Run("invalid command "+tc, func(t *testing.T) {
			_, err := dial.Rotate(0, tc)
			if err == nil {
				t.Errorf("expected error but got nil")
			}
		})
	}
}

func TestRotateByTickCountZeros(t *testing.T) {
	actual, err := dial.RotateByTickCountZeros(50, "R1000")
	if err != nil {
		t.Errorf("expected no error got %s", err)
	}
	if actual != 10 {
		t.Errorf("expected 10 got %d", actual)
	}
}
