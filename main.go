package main

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
import (
	"answer-question-demo/answer_one"
)

// 主函数，程序的入口点
func main() {
	// 调用runPalindromic函数，并传入一个整数切片
	// 该切片是一个回文数列，从5递减到1再递增回5
	answer_one.RunPalindromic(0)
}
