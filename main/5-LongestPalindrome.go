package main
/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
import "fmt"

func MapKey(l, r int) string {
	return fmt.Sprintf("%d:%d", l, r)
}

func LongestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	retStrMaxLength := 1
	retStr := s[0:1]
	flag := make(map[string]bool)
	flag[MapKey(0,0)] = true
	for r := 1; r < len(s); r++ {
		for l := 0; l < r; l++ {
			if s[l] == s[r] && (r - l <= 2 || flag[(MapKey(l + 1, r - 1))]) {
				flag[(MapKey(l, r))] = true
				if r -l + 1 > retStrMaxLength {
					retStrMaxLength = r - l + 1
					retStr = s[l:r + 1]
				}
			}
		}
	}
	return retStr
}

func main() {
	s := "abcdeffedcba"
	fmt.Println(LongestPalindrome(s))
}
