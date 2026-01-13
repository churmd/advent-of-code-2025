package dial

import (
	"fmt"
	"regexp"
	"strconv"
)

func FollowRotationsCountZeros(start int, commands []string) (int, error) {
	count := 0
	currentStart := start
	for _, c := range commands {
		zeros, err := RotateByTickCountZeros(currentStart, c)
		if err != nil {
			return 0, fmt.Errorf("error with command %s: %w", c, err)
		}
		count = count + zeros

		n, err := Rotate(currentStart, c)
		if err != nil {
			return 0, fmt.Errorf("error with command %s: %w", c, err)
		}
		currentStart = n

	}

	return count, nil
}

func FollowRotations(start int, commands []string) ([]int, error) {
	var result []int
	currentStart := start
	for _, c := range commands {
		n, err := Rotate(currentStart, c)
		if err != nil {
			return nil, fmt.Errorf("error with command %s: %w", c, err)
		}
		currentStart = n
		result = append(result, n)
	}

	return result, nil
}

type command struct {
	isLeft bool
	amount int
}

func parseCommand(s string) (command, error) {
	r := regexp.MustCompile(`^(L|R)(\d+$)`)
	groups := r.FindStringSubmatch(s)
	if len(groups) != 3 {
		return command{}, fmt.Errorf("command is not formatted correctly: %s", s)
	}

	n, err := strconv.Atoi(groups[2])
	if err != nil {
		return command{}, fmt.Errorf("could not get number: %w", err)
	}

	c := command{
		isLeft: groups[1] == "L",
		amount: n,
	}
	return c, nil
}

func RotateByTickCountZeros(start int, command string) (int, error) {
	c, err := parseCommand(command)
	if err != nil {
		return 0, err
	}

	currentStart := start
	count := 0
	for {
		if c.amount <= 0 {
			return count, nil
		}
		currentStart := rotate(currentStart, c)
		c.amount--
		if currentStart == 0 {
			count++
		}
	}

}

func Rotate(start int, command string) (int, error) {
	c, err := parseCommand(command)
	if err != nil {
		return 0, err
	}
	return rotate(start, c), nil
}

func rotate(start int, c command) int {
	if c.isLeft {
		return rotateLeft(start, c.amount)
	}

	return rotateRight(start, c.amount)
}

func rotateLeft(start int, amount int) int {
	return mod((start - amount), 100)
}

func rotateRight(start int, amount int) int {
	return mod((start + amount), 100)
}

func mod(a, b int) int {
	return (a%b + b) % b
}
