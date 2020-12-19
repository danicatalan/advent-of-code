package shuttlesearch

import (
	"log"
	"strconv"
	"strings"
)

//Part2 ...
func Part2(input string) int {
	buses := make(map[int]int)
	for i, b := range strings.Split(input, ",") {
		if b == "x" {
			continue
		}
		bi, err := strconv.Atoi(b)
		if err != nil {
			log.Fatal(err)
		}
		buses[i] = bi
	}

	timestamp, step := 0, 1
	for k, v := range buses {
		for (timestamp+k)%v != 0 {
			timestamp += step
		}
		step *= v
	}
	return timestamp
}
