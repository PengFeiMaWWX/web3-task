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
}
