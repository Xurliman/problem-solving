package _980_find_unique_binary

import "strings"

//	input: "111","011","001",
//
// output: "000"
func findDifferentBinaryString(nums []string) string {
	var res strings.Builder
	res.Grow(len(nums))
	for i := 0; i < len(nums); i++ {
		if string(nums[i][i]) == "0" {
			res.WriteString("1")
		} else {
			res.WriteString("0")
		}
	}

	return res.String()
}
