package main

import (
	"sort"
)

/**
零钱兑换  --- 贪心算法
*/

//func main() {
//	coins := []int{1, 2, 5}
//	amount := 11
//	numCoins := coinChange(coins, amount)
//	fmt.Printf("最少需要的硬币数：%d\n", numCoins)
//}

func coinChange(coins []int, amount int) int {
	// 面额按从大到小排序
	// 使用sort.IntSlice类型将整数数组转换为切片
	// sort.Reverse函数将该切片作为参数传递，然后将整个切片反向排序
	sort.Sort(sort.Reverse(sort.IntSlice(coins)))

	// 记录硬币数量
	count := 0
	// 从最大面额的硬币开始尝试
	index := 0

	for amount > 0 && index < len(coins) {
		if coins[index] <= amount {
			// 尝试使用当前面额的硬币数量
			numCoins := amount / coins[index]
			count += numCoins
			amount -= numCoins * coins[index]
		}
		// 尝试下一个面额的硬币
		index++
	}

	if amount == 0 {
		return count
	}
	// 返回-1表示无法兑换
	return -1
}
