package answer_three

// RunlongestCommonPrefix 函数用于查找字符串数组中的最长公共前缀
// 参数 strSlice 是一个字符串切片，包含需要比较的字符串
// 返回值是这些字符串的最长公共前缀字符串
func RunlongestCommonPrefix(strSlice []string) string {
	// 如果输入的字符串切片为空，直接返回空字符串
	if len(strSlice) == 0 {
		return ""
	}

	// 遍历第一个字符串的每个字符
	for s := 0; s < len(strSlice[0]); s++ {
		// 获取当前字符作为标准比较字符
		standardByte := strSlice[0][s]
		// 遍历其他字符串进行比较
		for i := 1; i < len(strSlice); i++ {
			// 如果当前位置超出其他字符串的长度，或者字符不匹配，则返回已匹配的前缀
			if s >= len(strSlice[i]) || strSlice[i][s] != standardByte {
				return strSlice[0][:s]
			}

		}
	}
	// 如果所有字符都匹配，则返回第一个字符串本身
	return strSlice[0]

}
