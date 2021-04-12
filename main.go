package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main()  {
	in := bufio.NewReader(os.Stdin)
	str1, _ := in.ReadString(',')
	str2, _ := in.ReadString('\n')
	fmt.Print(findDiff(str1, str2))

}

func splitter(str string) []rune{
	strSlice := strings.Fields(strings.ToLower(strings.Trim(str,",")))
	var resSlice []rune
	for _, el := range strSlice{
		for _, v := range el {
			resSlice = append(resSlice,v)
		}
	}
	return resSlice
}

func toMap (runeSlice []rune) map[rune]int{
	counter := make(map[rune]int)
	for _, el := range runeSlice{
		counter[el]++
	}
	return counter
}

func findDiff(a,b string) bool{
	var m1 map[rune]int
	var m2 map[rune]int
	var res bool
	m1 = toMap(splitter(a))
	m2 = toMap(splitter(b))
	for k, _ := range m1 {
			if m1[k] != m2[k] {
				res = false
				break
			}else {res = true}
			}

		return res
}