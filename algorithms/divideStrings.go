package algorithms

import "fmt"

func DivideStr(strs []string) [][]string {
	res := make([][]string, 0)
	for i := range strs {
		chars := []rune(strs[i])
		char := []string{}
		for i := 0; i < len(chars); i++ {
			char = append(char, string(chars[i]))
		}
		res = append(res, char)
	}
	fmt.Println(res)
	return res
}
