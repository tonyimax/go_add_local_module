package greetings //包名

import (
	"errors"
	"fmt" //导入标准库
	"math/rand"
)

// 函数
// name参数 类型string
// 返回类型 string
func Hello(name string) string {
	msg := fmt.Sprintf("hello: %v", name) //返回格式化字符串
	return msg                            //返回字符串
}

// 随机取输出格式串
func randomFormat() string {
	//输出格式数组
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))] //随机取一个格式输出
}

// 返回随机格式输出字符串
func HelloByRand(name string) string {
	s := fmt.Sprintf(randomFormat(), name)
	return s
}

// 根据输入字符串随机格式化输出，并处理错误
func OutputName(s string) (string, error) {
	if s == "" {
		return s, errors.New("name参数为空，请输入") //错误处理
	}
	msg := fmt.Sprintf(randomFormat(), s)
	return msg, nil
}

// 根据输入字符串随机格式化输出，并处理错误，处理数组
// 传入数组
// 返回处理后的数组
func OutputNames(strs []string) (map[string]string, error) {
	msgs := make(map[string]string)
	for _, s := range strs {
		msg, err := OutputName(s)
		if err != nil {
			return nil, err
		}
		msgs[s] = msg
	}
	return msgs, nil
}

// 处理整形数组
// 传入整形数组
// 返回处理后的数组
func ProcNumArray(nums []int32) (map[int32]int32, int, error) {
	arr := make(map[int32]int32)
	for i, num := range nums {
		fmt.Printf("%d->:%d\n", i, num)
		arr[int32(i)] = num * num
	}
	return arr, len(arr), nil
}
