package main

import (
	// "fmt"
	erfenfa "DSA/erfenfa"
	"fmt"
	// "DSA/stack"
	// "DSA/stack/dandiaozhan"
	// "DSA/string/kmp"
)

func main() {
	//result := stack.DecodeString("3[a2[c]]")
	//resutl1 := dandiaozhan.DailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})
	//for i := 0; i < len(resutl1); i++ {
	//	fmt.Println(resutl1[i])
	//}
	//result3 := zhan.LargestRectangleArea([]int{2, 1, 5, 6, 2, 3})
	//for i := 0; i < len(result3); i++ {
	//	fmt.Println(result3[i])
	////}
	//value := dandiaozhan.Trap([]int{0, 1, 0, 2, 1})
	//fmt.Println("解码结果:", result)
	//fmt.Println(value)
	//str3 := kmp.StrStr3("sttttttadbutsad", "sad")
	//fmt.Println(str3)
	//fmt.Println(kmp.RotateString("abcde630ccc", "630cccabcde"))
	//fmt.Println(erfenfa.MySqrt(15))
	output := erfenfa.Search([]int{1, 3, 6, 10}, 11)
	fmt.Print(output)
}

//package main
//
//import (
//	"fmt"
//	"io/ioutil"
//	"net/http"
//)
//
//func main() {
//	resp, err := http.Get("https://www.baidu.com/") //发送HTTP GET请求
//	if err != nil {
//		fmt.Printf("get failed, err:%v\n", err)
//		return
//	}
//	defer resp.Body.Close()                //延迟关闭服务 关键字确保在函数返回前关闭响应体。defer会延迟执行其后的语句，直到包含它的函数即将返回时才执行
//	body, err := ioutil.ReadAll(resp.Body) //读取响应体的全部内容到内存中
//	if err != nil {
//		fmt.Printf("read from resp.Body failed, err:%v\n", err)
//		return
//	}
//	fmt.Print(string(body))
//}
