package main

import "fmt"

/**
回溯算法  -- 全排列
*/

func main() {
	nums := []int{1, 2, 3}
	permutations := permute(nums)
	for _, permutation := range permutations {
		fmt.Println(permutation)
	}
}

// 生成全排列的入口函数
func permute(nums []int) [][]int {
	result := [][]int{} // 存储结果的切片
	backtrack(nums, []int{}, &result)
	return result
}

// 回溯函数，用于生成全排列
func backtrack(nums []int, current []int, result *[][]int) {
	// 当前排列长度等于数组长度时，将其加入结果切片
	if len(current) == len(nums) {
		temp := make([]int, len(current))
		copy(temp, current) // 复制当前排列到临时切片
		*result = append(*result, temp)
		return
	}

	// 尝试每个未使用的元素
	for _, num := range nums {
		if contains(current, num) {
			continue // 如果当前数字已经在排列中，跳过
		}
		current = append(current, num)
		backtrack(nums, current, result)   // 递归调用
		current = current[:len(current)-1] // 回溯，移除最后一个元素
	}
}

// 判断切片中是否包含某个元素
func contains(nums []int, num int) bool {
	for _, n := range nums {
		if n == num {
			return true
		}
	}
	return false
}
