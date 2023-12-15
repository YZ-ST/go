package main

import "fmt"

func main() {
    // 创建一个示例map
    myMap := map[string]int{
        "apple":  1,
        "banana": 2,
        "cherry": 3,
    }

    // 遍历map并打印键和值
    for key, value := range myMap {
        fmt.Println("map's key:", key)
        fmt.Println("map's value:", value)
    }

    fmt.Println("----------")

    // _ 是一个特殊的变量名，任何赋予它的值都会被丢弃
    for _, v := range myMap{
        fmt.Println("map's val:", v)
    }
}
