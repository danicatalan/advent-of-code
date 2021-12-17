package sevensegmentsearch

import (
	"sort"
	"strconv"
	"strings"
)

//Part2 determines all of the wire/segment connections and decodes the four-digit output values.
func Part2(input string) int {
	entries := strings.Split(input, "\n")
	count := 0
	for _, entry := range entries {
		count += decode(entry)
	}
	return count
}

func decode(input string) int {
	signalPatterns := strings.Split(strings.Split(input, DELIMITER)[0], " ")
	for i, signal := range signalPatterns {
		signalPatterns[i] = sortChar(signal)
	}
	sort.Slice(signalPatterns, func(i, j int) bool {
		return len(signalPatterns[i]) < len(signalPatterns[j])
	})

	dictionary := make(map[string]string)
	dictionary["1"] = signalPatterns[0]
	dictionary["7"] = signalPatterns[1]
	dictionary["4"] = signalPatterns[2]
	dictionary["8"] = signalPatterns[9]
	for _, signal := range signalPatterns[6:9] {
		if !containsAll(signal, dictionary["1"]) {
			dictionary["6"] = signal
		}
		if containsAll(signal, dictionary["4"]) {
			dictionary["9"] = signal
		}
		if containsAll(signal, dictionary["1"]) && !containsAll(signal, dictionary["4"]) {
			dictionary["0"] = signal
		}
	}
	for _, signal := range signalPatterns[3:6] {
		if containsAll(signal, dictionary["1"]) {
			dictionary["3"] = signal
		}
		if containsAll(dictionary["6"], signal) {
			dictionary["5"] = signal
		}
		if !containsAll(signal, dictionary["1"]) && !containsAll(dictionary["6"], signal) {
			dictionary["2"] = signal
		}
	}

	reversedDictionary := make(map[string]string)
	for number, signal := range dictionary {
		reversedDictionary[signal] = number
	}

	outputValue := strings.Split(strings.Split(input, DELIMITER)[1], " ")
	temp := ""
	for _, val := range outputValue {
		temp += reversedDictionary[sortChar(val)]
	}
	output, _ := strconv.Atoi(temp)
	return output
}

func containsAll(s string, substr string) bool {
	for _, char := range substr {
		if !strings.Contains(s, string(char)) {
			return false
		}
	}
	return true
}

func sortChar(input string) string {
	characters := strings.Split(input, "")
	sort.Strings(characters)
	return strings.Join(characters, "")
}
