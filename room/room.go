package room //包名

import "fmt" //导入标准库

// 方法名首字母必须大写
func Test(name string, age int) string {
	fmt.Printf("%s is %d years old.\n", name, age)
	s := fmt.Sprintf("the name:%s,the age:%d. \n", name, age) //格式化输出
	return s
}

// 返回多个值
func Test_Return_More(name string, age int) (string, error) {
	s := fmt.Sprintf("the name:%s,the age:%d. \n", name, age)
	return s, nil
}
