package answer_four

import (
	"fmt"
	"strconv"
	"strings"
)

func IncrementIntegerArr(numSlice []int64) {
	if len(numSlice) == 0 || (len(numSlice) == 1 && numSlice[0] == 0) {
		fmt.Println(numSlice)
		return
	}

	digitalize := numSliceToInt64(numSlice)
	digitalize++

	result := int64ToDigits(digitalize)
	fmt.Println(result)

}

func numSliceToInt64(numSlice []int64) int64 {
	digitalizeStrSlice := make([]string, 0, len(numSlice))
	for _, value := range numSlice {
		digitalizeStrSlice = append(digitalizeStrSlice, strconv.Itoa(int(value)))
	}
	digitalizeStr := strings.Join(digitalizeStrSlice, "")
	parseInt, err := strconv.ParseInt(digitalizeStr, 10, 64)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	digitalize := int64(parseInt)
	return digitalize
}

func int64ToDigits(digitalize int64) []int64 {
	var digitSlice []int64
	for digitalize > 0 {
		digit := digitalize % 10
		digitSlice = append(digitSlice, digit)
		digitalize /= 10
	}

	return reversedCopy(digitSlice)
}

func reversedCopy[T any](s []T) []T {
	n := len(s)
	// 创建同长度的新切片
	reversed := make([]T, n)
	// 逆序复制
	for i := 0; i < n; i++ {
		reversed[i] = s[n-1-i]
	}
	return reversed
}
