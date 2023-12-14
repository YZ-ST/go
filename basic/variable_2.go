// 程式起始點(定義文件屬於甚麼package)
package main

// 匯入套件
import (
	"fmt"
)

// 宣告變數
var variable_test string = "variable_test"
var variable_test_2 string = "variable_test_2"
var variable_test_3 string = "3"

var number_test1 int = 1
var number_test2 int = 2

// main() go 的進入點(會顯找pagkage main 的 main() 作為進入點)
func main() {
	fmt.Println(variable_test + "\n" + variable_test_2) // 字串相加
	// fmt.Println(number_test1 + variable_test_3) // 字串與數字不能相加
}