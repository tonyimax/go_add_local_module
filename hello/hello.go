package main

import (
	"fmt"
	"log"
	"rsc.io/quote"
)
import "com.test/greetings" //自定义模块1  go mod edit -replace com.test/greetings=../greetings   && go mod tidy
import "com.test/room"      //自定义模块2  go mod edit -replace com.test/room=../room  && go mod tidy

func main() {
	//调用标准库及quote模块方法
	fmt.Println(quote.Go())
	fmt.Println(quote.Hello())
	//调用本地模块方法
	s := greetings.Hello("nance")
	fmt.Println("return string: %v", s)
	s1 := room.Test("reverse", 25)
	fmt.Println("===>:%v", s1)
	a, b := room.Test_Return_More("lucy", 55) //接收函数返回的多个值
	fmt.Println("===>:%s,%d", a, b)
	//随机格式输出
	fmt.Println(greetings.HelloByRand("YEAR"))

	//数据定义及初始化
	maps := []string{"高德", "谷歌", "百度", "腾讯"}
	//常规遍历，每次返回索引
	for i := 0; i < len(maps); i++ {
		fmt.Printf("===>地图名:%s<===\n", maps[i]) //格式化输出
		fmt.Println("***>", maps[i], "<***")    //同时输出多个对象
	}
	//范围遍历，每次返回索引与值
	for i, v := range maps {
		fmt.Println(i, v)
	}
	//调用自定义模块中的函数处理数组
	newMaps, err := greetings.OutputNames(maps)
	if err == nil {
		for i, v := range newMaps {
			fmt.Println("===> ", i, v)
		}
	} else {
		fmt.Println(err)
	}

	nums := []int32{1, 3, 2, 6, 7, 9, 8, 4, 0}
	fmt.Println("===>处理前:", nums)
	n, c, e := greetings.ProcNumArray(nums)
	fmt.Println("===>处理后:", n, c, e)

	//日志模块使用
	log.SetPrefix("go日志前缀:hello.og: ---> ")
	log.SetFlags(3) //0:不显示日期时间 1:显示日期 2:显示时间 3.显示日期与时间
	log.Fatal(e, c, n)
}
