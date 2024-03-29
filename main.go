package main

import (
	"fmt"
	"strings"
)

// задачи с leetcode
func main() {
	// TwoSum([]int{2, 7, 11, 15}, 9)
	// fmt.Println(PalindromNumber(1221))
	// fmt.Println(RomanToInt("MCMXCIV"))
	// fmt.Println(GetCommonPrefix([]string{"flaz", "flaow", "flaowesrs", "flaag"}))
	// fmt.Println(RemoveDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	// fmt.Println(RemoveElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
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

// RemoveDuplicates - remove the duplicates in-place such that each unique element appears only once
func RemoveDuplicates(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return l
	}

	j := 0
	for i := 1; i < l; i++ {
		if nums[j] != nums[i] {
			j++
			nums[j] = nums[i]
		}
	}

	return j + 1
}

// RemoveElement - remove in place element in array
func RemoveElement(nums []int, val int) int {
	v := 0
	for i := range nums {
		if nums[i] != val {
			v++
		}
	}

	for i := 0; i < len(nums); i++ {
		if i == v {
			break
		}

		if nums[i] == val {
			j := 1
			for nums[i] == val {
				nums[i] = nums[i+j]
				nums[i+j] = val
				j++
			}
		}
	}

	return v
}

// StrStr - return index of substring start
func StrStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	if haystack == "" || len(haystack) < len(needle) {
		return -1
	}

	res := -1
	for i := range haystack {
		if haystack[i] == needle[0] {
			c := 0
			for j := i; j < i+len(needle); j++ {
				if j == len(haystack) || haystack[j] != needle[c] {
					break
				}

				c++
				if c == len(needle) {
					return i
				}
			}
		}
	}

	return res
}

// StrStr2 - return index of substring start
func StrStr2(haystack string, needle string) int {
	if haystack == needle || len(needle) < 1 {
		return 0
	}

	split := strings.Split(haystack, needle)

	fmt.Println(split)

	if len(split[0]) < len(haystack) {
		return len(split[0])
	}

	return -1
}

// LengthOfLongestSubstring - return len of longest substring, thanks cap
func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	if len(s) == 1 {
		return 1
	}

	var substrings [][]byte

	l := len(s)
	max := 0
	for i := 0; i < l; i++ {
		dict := make(map[byte]struct{})

		var substr []byte
		for j := i; j <= l; j++ {
			if j == l {
				substrings = append(substrings, substr)
				break
			}

			if _, ok := dict[s[j]]; ok {
				substrings = append(substrings, substr)
				break
			}

			substr = append(substr, s[j])

			dict[s[j]] = struct{}{}
		}

		if len(substrings[i]) > max {
			max = len(substrings[i])
		}
	}

	return max
}

// SearchInsert - search int value in int array, if we cant find in array value then we will place where it should be. Return index.
func SearchInsert(nums []int, target int) int {
	if nums[0] > target {
		return 0
	}

	for i := range nums {
		if nums[i] == target {
			return i
		}

		if i == len(nums)-1 {
			return len(nums)
		}

		if nums[i] < target && target < nums[i+1] {
			return i + 1
		}
	}

	return 0
}

// MaxSubArray - find max sum of subarray elements. Use Kodane's algth.
func MaxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}

	cur := nums[0]
	far := nums[0]
	for i := 1; i < len(nums); i++ {
		cur = max(nums[i], cur+nums[i])
		far = max(far, cur)
	}

	return far
}

// SummArrayView - ...
func SummArrayView(digits []int) []int {
	// run away from cycle
	if digits[len(digits)-1] != 9 {
		digits[len(digits)-1] += 1
		return digits
	}

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] = digits[i] + 1
			break
		} else {
			digits[i] = 0
			if i == 0 {
				digits = append([]int{1}, digits...)
			}
		}
	}

	fmt.Println(digits)
	return digits
}

// СlimbStairs - returns the number of options for passing a ladder of n steps
func СlimbStairs(n int) int {
	if n == 2 || n == 1 {
		return n
	}

	res := 3
	prev := 2
	for i := 3; i < n; i++ {
		fmt.Println(res)
		tmp := res
		res = res + prev
		prev = tmp
	}

	return res
}

// MySqrt - return sqrt of value
func MySqrt(x int) int {
	l, r := 0, x+1
	for l < r {
		m := l + (r-l)/2
		if m*m > x {
			r = m
		} else {
			l = m + 1
		}
	}

	return l - 1
}
