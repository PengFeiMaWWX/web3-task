package main

import (
	"fmt"
	"reflect"
	"sort"
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

/*
*
3、字符串
有效的括号

考察：字符串处理、栈的使用

题目：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
*/
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

/*
*

	4、最长公共前缀

	  考察：字符串处理、循环嵌套

	  题目：查找字符串数组中的最长公共前缀
*/
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

/*
*
基本值类型
5、加一

难度：简单

考察：数组操作、进位处理

题目：给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
*/
func plusOne(digits []int) []int {

	// 从最后一位开始，向前遍历
	for i := len(digits) - 1; i >= 0; i-- {

		// 如果当前为小于9，加1后直接返回
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		// 否则当前位置0，进位1传递给前一位
		digits[i] = 0
	}
	/**
	[]int{1}：创建一个只包含数字 1 的新切片
	digits...：使用 ...语法将原 digits 切片的所有元素展开
	append()：将展开的 digits 元素追加到新切片 [1]后面
	*/
	return append([]int{1}, digits...)
}

/*
*
6、引用类型：切片

	删除有序数组中的重复项：给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。

不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，一个快指针 j 用于遍历数组，当 nums[i] 与 nums[j] 不相等时，
将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位。
*/
func removeDuplicates(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	// 使用双指针法
	i := 0 // 慢指针 ，指向最后一个不重复的元素的位置
	for j := 0; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}

	}
	return i + 1

}

/*
*
7、合并区间：以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，
将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。
*/
func merge(intervals [][]int) [][]int {

	//1、按照区间起始位置进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	//2、初始化结果切片， 放入第一个区间
	merged := [][]int{intervals[0]}

	//3、遍历剩余区间
	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		lastMerged := merged[len(merged)-1]

		// 检查当前区间是否与最后一个合并区间重叠
		if current[0] <= lastMerged[1] {
			// 有重叠，合并区间（取结束位置的较大值）
			if current[1] > lastMerged[1] {
				lastMerged[1] = current[1]
			}
		} else {
			// 无重叠，直接添加到结果中
			merged = append(merged, current)
		}
	}
	return merged

}

// 辅助函数：取最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
*
8、两数之和

考察：数组遍历、map使用

题目：给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
*/
func twoSum(nums []int, target int) []int {
	// 创建一个map，键为元素值，值为对应的索引
	numMap := make(map[int]int)

	// 遍历数组
	for i, num := range nums {
		// 计算补数：target - 当前元素值
		complement := target - num
		// 检查补数是否在map中
		if index, found := numMap[complement]; found {
			// 如果找到，返回补数的索引和当前索引
			return []int{index, i}
		}
		// 将当前元素和索引存入map
		numMap[num] = i
	}
	// 未找到解，返回nil（根据题目假设，通常不会执行到这里）
	return nil
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
	testCases3 := []struct {
		input    string
		expected bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"", true}, // 空字符串有效
	}

	for _, tc := range testCases3 {
		result := isValidString(tc.input)
		fmt.Printf("输入: %-8s 预期: %-5t 实际: %-5t %s\n",
			tc.input, tc.expected, result,
			map[bool]string{true: "✓", false: "✗"}[result == tc.expected])
	}

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

	// 题目无
	// 测试用例
	testCases5 := []struct {
		input  []int
		expect []int
	}{
		// 普通情况：无进位
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		// 普通进位：中间位进位
		{[]int{1, 2, 9}, []int{1, 3, 0}},
		// 连续进位：多位连续进位
		{[]int{1, 9, 9}, []int{2, 0, 0}},
		// 全9情况：位数增加
		{[]int{9, 9, 9}, []int{1, 0, 0, 0}},
		// 单个数字
		{[]int{0}, []int{1}},
		// 边界情况：空数组（题目要求非空数组，但做防御性测试）
		{[]int{}, []int{1}},
	}

	// 执行测试
	for i, tc := range testCases5 {
		result := plusOne(tc.input)
		if reflect.DeepEqual(result, tc.expect) {
			fmt.Printf("测试用例 %d 通过: 输入 %v → 输出 %v\n", i+1, tc.input, result)
		} else {
			fmt.Printf("测试用例 %d 失败: 输入 %v → 期望 %v, 实际 %v\n", i+1, tc.input, tc.expect, result)
		}
	}

	// 测试用例
	testCases6 := []struct {
		input    []int
		expected int
		result   []int
	}{
		{[]int{1, 1, 2}, 2, []int{1, 2}},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5, []int{0, 1, 2, 3, 4}},
		{[]int{}, 0, []int{}},
		{[]int{1}, 1, []int{1}},
		{[]int{1, 2, 3}, 3, []int{1, 2, 3}},
		{[]int{1, 1, 1, 1}, 1, []int{1}},
	}

	for _, tc := range testCases6 {
		// 复制输入数组以避免修改原始测试数据
		inputCopy := make([]int, len(tc.input))
		copy(inputCopy, tc.input)

		// 执行函数
		newLength := removeDuplicates(inputCopy)

		// 验证结果
		if newLength != tc.expected {
			fmt.Printf("测试失败: 输入 %v, 期望长度 %d, 实际长度 %d\n", tc.input, tc.expected, newLength)
			continue
		}

		// 验证数组前newLength个元素
		valid := true
		for i := 0; i < newLength; i++ {
			if inputCopy[i] != tc.result[i] {
				valid = false
				break
			}
		}

		if valid {
			fmt.Printf("测试通过: 输入 %v -> 输出 %v (长度 %d)\n", tc.input, inputCopy[:newLength], newLength)
		} else {
			fmt.Printf("测试失败: 输入 %v, 期望 %v, 实际 %v\n", tc.input, tc.result, inputCopy[:newLength])
		}
	}

	// 测试用例
	testCases7 := [][][]int{
		{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, // 标准情况
		{{1, 4}, {4, 5}},                    // 端点相接
		{{1, 4}, {2, 3}},                    // 完全包含
		{{1, 2}},                            // 单个区间
	}

	for i, intervals := range testCases7 {
		result := merge(intervals)
		fmt.Printf("测试用例7 %d: 输入 %v -> 输出 %v\n", i+1, intervals, result)
	}

	// 测试用例8
	tests := []struct {
		nums   []int
		target int
	}{
		{[]int{2, 7, 11, 15}, 9},
		{[]int{3, 2, 4}, 6},
		{[]int{3, 3}, 6},
	}

	for _, test := range tests {
		result := twoSum(test.nums, test.target)
		fmt.Printf("nums = %v, target = %d -> %v\n", test.nums, test.target, result)
	}
}
