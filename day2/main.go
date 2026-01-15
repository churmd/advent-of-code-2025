package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/churmd/advent-of-code-2025/day2/product"
)

func main() {
	ranges := ParseInput(input)

	answer1 := product.SumInvalidIDs(ranges, product.IsInvalid)
	fmt.Println("What do you get if you add up all of the invalid IDs?")
	fmt.Println(answer1)

	answer2 := product.SumInvalidIDs(ranges, product.IsInvalid2)
	fmt.Println("What do you get if you add up all of the invalid IDs using these new rules?")
	fmt.Println(answer2)
}

func ParseInput(s string) []product.ProductIDRange {
	var result []product.ProductIDRange
	ranges := strings.Split(input, ",")
	for _, r := range ranges {
		productIDs := strings.Split(r, "-")
		start, err := strconv.Atoi(productIDs[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(productIDs[1])
		if err != nil {
			panic(err)
		}
		pir := product.ProductIDRange{
			Start: start,
			End:   end,
		}
		result = append(result, pir)
	}

	return result
}

const input = `26803-38596,161-351,37-56,9945663-10044587,350019-413817,5252508299-5252534634,38145069-38162596,1747127-1881019,609816-640411,207466-230638,18904-25781,131637-190261,438347308-438525264,5124157617-5124298820,68670991-68710448,8282798062-8282867198,2942-5251,659813-676399,57-99,5857600742-5857691960,9898925025-9899040061,745821-835116,2056-2782,686588904-686792094,5487438-5622255,325224-347154,352-630,244657-315699,459409-500499,639-918,78943-106934,3260856-3442902,3-20,887467-1022885,975-1863,5897-13354,43667065-43786338`
