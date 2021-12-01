package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const YEAR = "2021"

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing input (eg 'go run scaffolding.go 05 binaryboarding')")
	}

	os.MkdirAll(YEAR, os.ModePerm)

	day := "day" + os.Args[1]

	basePath := filepath.Join(YEAR, day)
	os.MkdirAll(basePath, os.ModePerm)

	data := []byte("")
	dataPath := filepath.Join(basePath, "data")
	err := ioutil.WriteFile(dataPath, data, 0644)
	if err != nil {
		log.Fatal(err)
	}

	main := []byte("package main\n\nfunc main() {\n\t\n}")
	mainPath := filepath.Join(basePath, "main.go")
	err = ioutil.WriteFile(mainPath, main, 0644)
	if err != nil {
		log.Fatal(err)
	}

	utils := []byte("package main")
	utilsPath := filepath.Join(basePath, "utils.go")
	err = ioutil.WriteFile(utilsPath, utils, 0644)
	if err != nil {
		log.Fatal(err)
	}

	pckg := os.Args[2]
	pckgPath := filepath.Join(basePath, pckg)
	os.MkdirAll(pckgPath, os.ModePerm)

	part1 := []byte(`package ` + pckg + "\n\n//Part1 ...\nfunc Part1(input string) int {\n\treturn 0\n}")
	part1Path := filepath.Join(pckgPath, "part1.go")
	err = ioutil.WriteFile(part1Path, part1, 0644)
	if err != nil {
		log.Fatal(err)
	}

	part1Test := []byte(`package ` + pckg + `

import (
	"reflect"
	"testing"
)
var testCasesPart1 = []struct {
	input  string
	output int
}{
	{
		` + "``" + `,
		0,
	},
}

func TestPart1(t *testing.T) {
	for i, tc := range testCasesPart1 {
		expected := tc.output
		actual := Part1(tc.input)
		if !reflect.DeepEqual(expected, actual) {
			t.Fatalf("TestCase %d failed.\n\tExpected: %v\n\tGot: %v", i, expected, actual)
		}
	}
}
`)
	part1TestPath := filepath.Join(pckgPath, "part1_test.go")
	err = ioutil.WriteFile(part1TestPath, part1Test, 0644)
	if err != nil {
		log.Fatal(err)
	}

	part2 := []byte("package " + pckg + "\n\n//Part2 ...\nfunc Part2(input string) int {\n\treturn 0\n}")
	part2Path := filepath.Join(pckgPath, "part2.go")
	err = ioutil.WriteFile(part2Path, part2, 0644)
	if err != nil {
		log.Fatal(err)
	}

	part2Test := []byte(`package ` + pckg + `

import (
	"reflect"
	"testing"
)
var testCasesPart2 = []struct {
	input  string
	output int
}{
	{
		` + "``" + `,
		0,
	},
}

func TestPart2(t *testing.T) {
	for i, tc := range testCasesPart2 {
		expected := tc.output
		actual := Part2(tc.input)
		if !reflect.DeepEqual(expected, actual) {
			t.Fatalf("TestCase %d failed.\n\tExpected: %v\n\tGot: %v", i, expected, actual)
		}
	}
}
`)
	part2TestPath := filepath.Join(pckgPath, "part2_test.go")
	err = ioutil.WriteFile(part2TestPath, part2Test, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
