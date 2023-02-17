package main

import (
	"fmt"
	"math"
)

func DefineTwoStrings(str1 string, str2 string) bool {
	opt := 0

	if len(str1) == len(str2) {
		for i := 0; i < len(str1); i++ {
			if string(str1[i]) != string(str2[i]) {
				opt++
			}
		}
		return opt == 1
	}

	mpCharStr1, mpCharStr2 := make(map[string]int), make(map[string]int)

	for _, char := range str1 {
		mpCharStr1[string(char)]++
	}

	for _, char := range str2 {
		mpCharStr2[string(char)]++
	}

	for k := range mpCharStr1 {
		if _, ok := mpCharStr2[k]; !ok {
			mpCharStr2[k]++
			opt++
		}
	}

	for k := range mpCharStr2 {
		if _, ok := mpCharStr1[k]; !ok {
			delete(mpCharStr2, k)
			opt++
		}

		if mpCharStr2[k] != mpCharStr1[k] {
			diff := int(math.Abs(float64(mpCharStr1[k]) - float64(mpCharStr2[k])))
			if mpCharStr2[k] > mpCharStr1[k] {
				mpCharStr2[k] -= diff
			}
			if mpCharStr2[k] < mpCharStr1[k] {
				mpCharStr2[k] += diff
			}
			opt += diff
		}
	}

	return opt == 1
}

func main() {
	fmt.Println("GIVEN INPUT 1:", "telkom")
	fmt.Println("GIVEN INPUT 2:", "telecom")
	fmt.Println("RESULT:", DefineTwoStrings("telkom", "telecom"))

	fmt.Println("")

	fmt.Println("GIVEN INPUT 1:", "telkom")
	fmt.Println("GIVEN INPUT 2:", "tlkom")
	fmt.Println("RESULT:", DefineTwoStrings("telkom", "tlkom"))

	fmt.Println("")

	fmt.Println("GIVEN INPUT 1:", "telekom")
	fmt.Println("GIVEN INPUT 2:", "telecom")
	fmt.Println("RESULT:", DefineTwoStrings("talekom", "telecom"))
}
