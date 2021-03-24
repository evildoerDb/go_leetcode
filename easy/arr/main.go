package main

import "fmt"

// 计算一维数组的动态和
//`给你一个数组 nums 。数组「动态和」的计算公式为：runningSum[i] = sum(nums[0]…nums[i]) 。
//
//请返回 nums 的动态和。
//
//
//
//示例 1：
//
//输入：nums = [1,2,3,4]
//输出：[1,3,6,10]
//解释：动态和计算过程为 [1, 1+2, 1+2+3, 1+2+3+4] 。
//`
func runningSum(nums []int) []int {

	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return nums
}

func main() {
	a := []int{1, 2, 3, 4}
	res := runningSum(a)

	for _, value := range res {
		fmt.Println(value)
	}
}
