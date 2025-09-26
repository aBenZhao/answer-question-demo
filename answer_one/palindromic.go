package answer_one

import (
	"fmt"
	"strconv"
)

// runPalindromic 是一个用于处理回文数组的函数
// 参数:
//
//	ints: 整数切片，可能包含回文数序列
func RunPalindromic(palindromic int) {
	if palindromic < 0 {
		fmt.Println("不是回文数,为负数:", palindromic)
		return
	}
	// 转换成byte数组
	palindromicStr := strconv.Itoa(palindromic)
	// 转换成切片
	palindromicSlice := []byte(palindromicStr)

	// 遍历切片
	for index := 0; index < len(palindromicSlice); index++ {
		if palindromicSlice[index] != palindromicSlice[len(palindromicSlice)-1-index] {
			fmt.Println("不是回文数:", string(palindromicSlice))
			break
		}

		// 奇数回文会在这里结束
		if index == len(palindromicSlice)-1-index {
			fmt.Println("是奇数回文数:", string(palindromicSlice))
			break
		}

		// 偶数回文会在这里结束
		if index+1 == len(palindromicSlice)-1-index {
			fmt.Println("是偶数回文数:", string(palindromicSlice))
			break
		}

	}

}
