package main

import (
	"fmt"
)

// задачи с leetcode
func main() {
	// TwoSum([]int{2, 7, 11, 15}, 9)
	// fmt.Println(PalindromNumber(1221))
	// fmt.Println(RomanToInt("MCMXCIV"))
	// fmt.Println(GetCommonPrefix([]string{"flaz", "flaow", "flaowesrs", "flaag"}))
}

// TwoSum - nums содержит массив целозначныйх чисел, target число которое надо получить из суммы двух чисел в nums.
// Нужно вернуть массив из индексов двух чисел которые в сумме давали бы target
func TwoSum(nums []int, target int) {

	// вариант решения в лоб
	/*
		for i := range nums {
			for j := range nums {
				if nums[i]+nums[j] == target && i != j {
					result = append(result, i)
					result = append(result, j)

					fmt.Println(result)
					return
				}
			}
		}
	*/

	m := make(map[int]int)
	for i, v := range nums {
		if j, ok := m[target-v]; ok && j != i {
			fmt.Println([]int{j, i})
			return
		}

		m[v] = i
	}

	fmt.Println("Жопа чё, такой комбинации нет братишка")
}

// PalindromNumber - нужно выяснить является ли x палиндромом
func PalindromNumber(x int) bool {
	var (
		sum     = 0
		compare = x
	)

	for x > 0 {
		sum = (sum * 10) + x%10
		x = x / 10

	}

	return compare == sum
}

// RomanToInt - convert string of Roman number to int
func RomanToInt(s string) int {
	var result int

	letters := map[uint8]int{73: 1, 76: 50, 86: 5, 88: 10, 67: 100, 68: 500, 77: 1000}

	pv := 0

	for i := len(s) - 1; i >= 0; i-- {
		cv := letters[s[i]]
		if pv > cv {
			result = result - cv
		} else {
			result = result + cv
		}

		pv = cv
	}

	return result
}

// GetCommonPrefix - return common prefix
func GetCommonPrefix(s []string) string {
	var result string
	max := len(s[0])

LOOP:
	for i := 0; i < max; i++ {
		for j := 0; j < len(s)-1; j++ {
			if s[j][i] != s[j+1][i] {
				continue LOOP
			}
		}

		if s[len(s)-1][i] == s[len(s)-2][i] {
			result += string(s[len(s)-2][i])
		}
	}

	return result
}

// IsValidBrackets - validate brackets subsequence
func IsValidBrackets(s string) bool {
	ls := len(s)
	if ls == 1 || ls%2 != 0 {
		return false
	}

	p := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	var stack []byte
	for i := range s {
		l := len(stack)
		if l > 0 && p[stack[l-1]] == s[i] {
			stack = stack[:l-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}