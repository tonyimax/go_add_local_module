package greetings

import (
	"regexp"
	"testing"
)

func TestHello(t *testing.T) {
	name := "HelloWorld"
	want := regexp.MustCompile(`\b` + name + `\b`) //正则，表达式无法编译会引发异常
	//msg, err := OutputName("HelloWorld")
	msg, err := OutputName("") //模拟出错，传入空字符串
	//正则匹配
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("HelloWorld") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// 测试函数
func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`) //正则，表达式无法编译会引发异常
	msg, err := OutputName("Gladys")
	//正则匹配
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// 测试函数，传入空值
func TestHelloEmpty(t *testing.T) {
	msg, err := OutputName("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
