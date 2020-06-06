package longestPath

import (
	"regexp"
	"strings"
)

func longestPath(fileSystem string) int {
	pat := regexp.MustCompile("(\t*)([ \\w]+)([ \\.\\w+]*)?")
	// pat := regexp.MustCompile("(\t*)([^\f]*)")

	lines := pat.FindAllStringSubmatchIndex(fileSystem, -1)
	// fmt.Println(lines)
	stack := make([]int, 1)
	maxLen := 0
	prevLevel := 0
	for _, indexes := range lines {
		tabs_num := indexes[3] - indexes[2]
		chars_num := indexes[5] - indexes[4] + indexes[7] - indexes[6]
		hasDot := (indexes[7] - indexes[6]) > 0
		if tabs_num == 0 {
			stack = stack[0:1]
			stack[0] = chars_num
		} else if tabs_num > prevLevel {
			stack = append(stack, stack[prevLevel]+chars_num+1)
		} else if tabs_num == prevLevel {
			stack[prevLevel] = stack[prevLevel-1] + chars_num + 1
		} else { // tabs < level
			stack = stack[:tabs_num+1]
			stack[tabs_num] = stack[tabs_num-1] + chars_num + 1
		}
		if hasDot && stack[tabs_num] > maxLen {
			maxLen = stack[tabs_num]
		}
		prevLevel = tabs_num
		// fmt.Println(stack)
	}
	return maxLen
}

func longestPath_v2(fileSystem string) int {
	pat := regexp.MustCompile("(\t*)([^\f]*)")
	lines := pat.FindAllStringSubmatchIndex(fileSystem, -1)
	// fmt.Println(lines)
	stack := make([]int, 1)
	maxLen := 0
	prevLevel := 0
	for _, indexes := range lines {
		tabs_num := indexes[3] - indexes[2]
		chars_num := indexes[5] - indexes[4]
		hasDot := strings.Contains(fileSystem[indexes[4]:indexes[5]], ".")

		if tabs_num == 0 {
			stack = stack[0:1]
			stack[0] = chars_num
		} else if tabs_num > prevLevel {
			stack = append(stack, stack[prevLevel]+chars_num+1)
		} else if tabs_num == prevLevel {
			stack[prevLevel] = stack[prevLevel-1] + chars_num + 1
		} else { // tabs < level
			stack = stack[:tabs_num+1]
			stack[tabs_num] = stack[tabs_num-1] + chars_num + 1
		}
		if hasDot && stack[tabs_num] > maxLen {
			maxLen = stack[tabs_num]
		}
		prevLevel = tabs_num
		// fmt.Println(stack)
	}
	return maxLen
}

func longestPath_v3(fileSystem string) int {
	pat := regexp.MustCompile("(\t*)([^\f]*)")
	lines := pat.FindAllStringSubmatchIndex(fileSystem, -1)
	// fmt.Println(lines)
	stack := make([]uint8, 1)
	maxLen := uint8(0)
	prevLevel := uint8(0)
	for _, indexes := range lines {
		tabs_num := uint8(indexes[3] - indexes[2])
		chars_num := uint8(indexes[5] - indexes[4])
		hasDot := strings.Contains(fileSystem[indexes[4]:indexes[5]], ".")

		if tabs_num == 0 {
			stack = stack[0:1]
			stack[0] = chars_num
		} else if tabs_num > prevLevel {
			stack = append(stack, stack[prevLevel]+chars_num+1)
		} else if tabs_num == prevLevel {
			stack[prevLevel] = stack[prevLevel-1] + chars_num + 1
		} else { // tabs < level
			stack = stack[:tabs_num+1]
			stack[tabs_num] = stack[tabs_num-1] + chars_num + 1
		}
		if hasDot && stack[tabs_num] > maxLen {
			maxLen = stack[tabs_num]
		}
		prevLevel = tabs_num
		// fmt.Println(stack)
	}
	return int(maxLen)
}
