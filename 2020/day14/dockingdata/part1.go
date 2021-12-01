package dockingdata

import (
	"log"
	"strconv"
	"strings"
)

//Part1 ...
func Part1(input string) int {
	mem := make(map[int]int64)
	mask := ""
	for _, instruction := range strings.Split(input, "\n") {
		kv := strings.Split(instruction, " = ")
		k, v := kv[0], kv[1]
		switch k {
		case "mask":
			mask = v
		default:
			addr, result := getMemoryAddress(k), applyMask(mask, v)
			mem[addr] = result
		}
	}
	sum := 0
	for _, v := range mem {
		sum += int(v)
	}
	return sum
}

func applyMask(mask, value string) int64 {
	result := ""
	base := "000000000000000000000000000000000000"
	value = toBinary(value)
	value = base[:len(base)-len(value)] + value

	result = value
	for i, c := range mask {
		if c != 'X' {
			result = result[:i] + string(c) + result[i+1:]
		}
	}
	return toDecimal(result)
}

func getMemoryAddress(v string) int {
	addr := v[4 : len(v)-1]
	n, err := strconv.Atoi(addr)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func toBinary(v string) string {
	n, err := strconv.Atoi(v)
	if err != nil {
		log.Fatal(err)
	}
	return strconv.FormatInt(int64(n), 2)
}

func toDecimal(v string) int64 {
	n, err := strconv.ParseInt(v, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return n
}
