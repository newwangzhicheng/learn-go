package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
)

// Student student struct
type Student struct {
	name string
	age  int
}

// Person
type Person interface {
	hello() string
}

func (stu *Student) hello(Person string) string {
	return person
}

func main() {
	str1 := "golang"
	str2 := "go语言"
	fmt.Println(reflect.TypeOf(str2[2]).Kind()) // uint8
	fmt.Println(str1[2], string(str1[2]))       // 108 l
	fmt.Printf("%d %c\n", str2[2], str2[2])     // 232 è
	fmt.Println("len(str2): ", len(str2))       // len(str2):  8

	slice1 := make([]float32, 3, 5)       // 声明一个长度为3，容量为5的切片
	fmt.Println(len(slice1), cap(slice1)) // 3 5
	slice2 := append(slice1, 1, 23, 4)    // 可扩展性
	fmt.Println(len(slice2), cap(slice2)) // 6 12
	sub1 := slice2[3:]                    // 截取子切片
	combin := append(sub1, slice1...)     // 合并切片
	fmt.Println(sub1, combin)             // [1 23 4] [1 23 4 0 0 0]

	map1 := make(map[string]int) // 仅声明
	map2 := map[string]string{   // 声明并初始化
		"a": "1",
		"b": "2",
	}
	map1["j"] = 2     // 赋值
	map2["a"] = "cap" // 修改

	str := "golang"
	p := &str // p指向str
	*p = "javascript"
	fmt.Println(str) //改变指针的值，也就改变了str

	_, err := os.Open("this.txt")
	if err != nil {
		fmt.Println(err)
	}

	err3 := testError()
	fmt.Println(err3)

	stu := &Student{
		name: "wzc",
	}
	stu2 := new(Student)
	fmt.Println(stu)
	fmt.Println(stu2)

	testDefer()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("msg: ", r)
			fmt.Println("after defer")
		}
	}()
	testPanic()
}

func addNminus(num1, num2 int) (add, minus int) {
	add = num1 + num2
	minus = num1 - num2
	return
}

func testError() (err error) {
	_, err2 := os.Open("this.txt")
	if err2 != nil {
		return errors.New("读取文件失败")
	}
	return nil
}

/*
返回
first
defer2
defer1
*/
func testDefer() {
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
	fmt.Println("first")
}

func testPanic() {
	fmt.Println("befor panic")
	panic(300)
	fmt.Println("after panic") // unreachable code
}
