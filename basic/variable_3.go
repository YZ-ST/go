// 程式起始點(定義文件屬於甚麼package)
package main

// 匯入套件
import (
	"fmt"
)


// main() go 的進入點(會顯找pagkage main 的 main() 作為進入點)
func main() {
	// %d:decimal %s:string %f:float %T:type
	var a int = 10
	var b string = "Hello"
	var c string = fmt.Sprintf("%d_%s", a, b)

	// fmt.Println: 不支援複雜格式輸出，會自動加換行，加入空格
	fmt.Println(c)

	// fmt.Printf: 支援複雜格式輸出，不會自動加換行
	fmt.Printf(c)

	// type of a
	fmt.Printf("\n%T\n", a)
	fmt.Println("%T", a)
}