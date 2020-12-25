package dockingdata

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

//Part2 ...
func Part2(input string) int {
	mem := make(map[int]int64)
	mask := ""

	for _, instruction := range strings.Split(input, "\n") {
		kv := strings.Split(instruction, " = ")
		k, v := kv[0], kv[1]
		switch k {
		case "mask":
			mask = v
		default:
			addr := getMemoryAddress(k)
			value, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			base := "000000000000000000000000000000000000"
			addrBin := toBinary(fmt.Sprint(addr))
			addrBin = base[:len(base)-len(addrBin)] + addrBin
			result := applyMask2(mask, addrBin)

			addrs := getAddr(result)

			for _, a := range strings.Split(addrs, ",") {
				addr := int(toDecimal(a))
				mem[addr] = int64(value)
			}
		}
	}
	sum := 0
	for _, v := range mem {
		sum += int(v)
	}
	return sum
}

func applyMask2(mask, value string) string {
	result := ""
	result = value
	for i, c := range mask {
		if c != '0' {
			result = result[:i] + string(c) + result[i+1:]
		}
	}
	return result
}

func getAddr(s string) string {
	if strings.Count(s, "X") == 0 {
		return s
	}
	i := strings.Index(s, "X")
	s0, s1 := s[:i]+"0"+s[i+1:], s[:i]+"1"+s[i+1:]
	return getAddr(s0) + "," + getAddr(s1)
}
