package answer_two

import "fmt"

const smallStartParenthesesStr string = "("
const smallEndParenthesesStr string = ")"
const midStartParenthesesStr string = "{"
const midEndParenthesesStr string = "}"
const bigStartParenthesesStr string = "["
const bigEndParenthesesStr string = "]"

func RunValidParentheses(str string) {
	// 检查字符串中的括号是否成对匹配
	pair := checkParenthesesPair(str)
	if !pair {
		fmt.Println("括号成对匹配错误")
		return
	}
	// 检查字符串中的括号是否正确嵌套
	nested := checkParenthesesNested(str)
	if nested {
		fmt.Println("括号嵌套正确")
	} else {
		fmt.Println("括号嵌套错误")
	}

}

func checkParenthesesNested(str string) bool {
	// 将字符串转换为字节切片以便遍历
	strSlice := []byte(str)
	for i := 0; i < len(strSlice); i++ {
		if smallEndParenthesesStr == string(strSlice[i]) {
			if i == 0 {
				return false
			}
			if smallStartParenthesesStr == string(strSlice[i-1]) {
				if 2 < len(strSlice) {
					strSlice = append(strSlice[:i-1], strSlice[i+1:]...)
					i = -1
					continue
				}
				return true
			}
			return false
		}

		if midEndParenthesesStr == string(strSlice[i]) {
			if i == 0 {
				return false
			}
			if midStartParenthesesStr == string(strSlice[i-1]) {
				if 2 < len(strSlice) {
					strSlice = append(strSlice[:i-1], strSlice[i+1:]...)
					i = -1
					continue
				}
				return true
			}
			return false
		}

		if bigEndParenthesesStr == string(strSlice[i]) {
			if i == 0 {
				return false
			}
			if bigStartParenthesesStr == string(strSlice[i-1]) {
				if 2 < len(strSlice) {
					strSlice = append(strSlice[:i-1], strSlice[i+1:]...)
					i = -1
					continue
				}
				return true
			}
			return false
		}
	}
	return true

}

// 检查字符串中的括号是否成对匹配
func checkParenthesesPair(str string) bool {

	// 初始化计数器，用于记录不同类型的括号数量
	smallStartParentheses := 0 // 小括号左括号数量
	smallEndParentheses := 0   // 小括号右括号数量
	midStartParentheses := 0   // 中括号左括号数量
	midEndParentheses := 0     // 中括号右括号数量
	bigStartParentheses := 0   // 大括号左括号数量
	bigEndParentheses := 0     // 大括号右括号数量

	// 将字符串转换为字节切片以便遍历
	strSlice := []byte(str)
	// 遍历字符串中的每个字符
	for i := 0; i < len(strSlice); i++ {
		value := string(strSlice[i])
		// 使用switch语句判断字符类型并更新对应计数器
		switch value {
		case smallStartParenthesesStr: // 遇到小括号左括号
			smallStartParentheses++
		case smallEndParenthesesStr: // 遇到小括号右括号
			smallEndParentheses++
		case midStartParenthesesStr: // 遇到中括号左括号
			midStartParentheses++
		case midEndParenthesesStr: // 遇到中括号右括号
			midEndParentheses++
		case bigStartParenthesesStr: // 遇到大括号左括号
			bigStartParentheses++
		case bigEndParenthesesStr: // 遇到大括号右括号
			bigEndParentheses++
		default: // 遇到非括号字符
			fmt.Println("存在其他字符")
			return false
		}
	}

	// 检查各类括号的数量是否匹配
	if smallStartParentheses != smallEndParentheses || midStartParentheses != midEndParentheses || bigStartParentheses != bigEndParentheses {
		return false
	}

	return true
}
