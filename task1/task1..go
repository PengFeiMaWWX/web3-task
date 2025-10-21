package main

import (
	"fmt"
)

/*
*
1、只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。

	找出那个只出现了一次的元素。可以使用 for 循环遍历数组，
	结合 if 条件判断和 map 数据结构来解决，
	例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
*/
//使用map记录每个数字出现的次数，找出只出现一次的数字
func singleNumber(a []int) int {

	// 创建map用于记录每个数字出现的次数
	numCount := make(map[int]int)
	//第一次遍历：统计每个数字出现的次数
	for _, num := range a {
		/**
		1. 读取当前值：temp := numCount[num]
		2. 增加值：temp = temp + 1
		3. 写回map：numCount[num] = temp
		*/
		numCount[num]++
	}

	//第二次遍历， 统计出
	for num, count := range numCount {
		if count == 1 {
			return num
		}
	}
	return -1
}

/*
*
2、回文数

考察：数字操作、条件判断
题目：判断一个整数是否是回文数
*/
func isPalindrome(x int) bool {

	// 特殊情况处理
	if x < 0 {
		return false // 负数不是回文
	}

	if x%10 == 0 && x != 0 {
		return false // 以0结尾的数不是回文数
	}

	revertedNumber := 0
	// 当原始数字大于反转后的数字时继续循环
	for x > revertedNumber {

		revertedNumber = revertedNumber*10 + x%10
		// 移除x的最后一位
		x /= 10
		fmt.Printf("x = %d, revertedNumber = %d\n", x, revertedNumber)

	}
	// 整数除法会自动截断小数部分，只保留整数部分
	return x == revertedNumber || x == revertedNumber/10

}

func isValidString(s string) bool {

	//使用切片模拟栈
	stack := []rune{}

	//定义括号映射关系
	pairs := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	//遍历字符串中的每个元素
	for _, c := range s {
		// 如果是开括号，压入对应的闭括号
		if closing, isOpen := pairs[c]; isOpen {
			stack = append(stack, closing)
		} else {
			// 遇到闭括号时检查
			if len(stack) == 0 || stack[len(stack)-1] != c {
				return false
			}
			// 匹配成功， 弹出栈顶（中从切片（slice）中移除最后一个元素）
			/**切片表达式
			:表示从索引 0 开始
			len(stack)-1表示结束索引（不包含该索引）
			具体操作
			len(stack)：获取切片当前长度（元素个数）
			len(stack)-1：计算新切片的结束索引（比原长度小1）
			stack[:len(stack)-1]：创建一个新切片，包含原切片从索引 0 到 len(stack)-2的所有元素
			stack = ...：将新切片赋值回原变量，完成"弹出"操作
			*/
			stack = stack[:len(stack)-1]

		}
	}
	// 检查栈是否为空
	return len(stack) == 0
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	//以第一个字符串为基准
	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		// 逐个比较，直到找到不匹配的位置
		j := 0
		for j < len(prefix) && j < len(strs[i]) && prefix[j] == strs[i][j] {
			j++
		}
		// 更新前缀为匹配部分
		prefix = prefix[:j]

		// 如果前缀为空，提前结束
		if prefix == "" {
			return ""
		}
	}
	return prefix

}

func main() {

	// 题目一
	var a = []int{2, 2, 3, 3, 4}
	fmt.Println(singleNumber(a))

	// 题目二
	testNumbers := []int{12321, 1221, 123, 10, -121}
	for _, num := range testNumbers {
		fmt.Printf("\n测试数字: %d\n", num)
		result := isPalindrome(num)
		fmt.Printf("是否是回文: %t\n", result)
	}

	// 题目三
	//testCases := []struct {
	//	input    string
	//	expected bool
	//}{
	//	{"()", true},
	//	{"()[]{}", true},
	//	{"(]", false},
	//	{"([)]", false},
	//	{"{[]}", true},
	//	{"", true}, // 空字符串有效
	//}
	//
	//for _, tc := range testCases {
	//	result := isValidString(tc.input)
	//	fmt.Printf("输入: %-8s 预期: %-5t 实际: %-5t %s\n",
	//		tc.input, tc.expected, result,
	//		map[bool]string{true: "✓", false: "✗"}[result == tc.expected])
	//}

	// 题目四
	testCases := []struct {
		input  []string
		expect string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"a"}, "a"},
		{[]string{"", "abc"}, ""},
		{[]string{"same", "same", "same"}, "same"},
		{[]string{"prefix", "preface", "preview"}, "pre"},
		{[]string{}, ""},
	}

	for _, tc := range testCases {
		result := longestCommonPrefix(tc.input)
		fmt.Printf("输入: %v\n预期: %q\n实际: %q\n匹配: %t\n\n",
			tc.input, tc.expect, result, result == tc.expect)
	}

}
