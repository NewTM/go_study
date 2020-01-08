package main

import (
	"fmt"
	"reflect"
)

type Test interface {
	hell()
}

type Student struct {
	Name string `学生姓名`
	Age  int    `a:"aaaa"b:"bbbb"`
}

func (a *Student) hell() {
	fmt.Println(a.Name)
}

func main() {
	s := Student{}
	rt := reflect.TypeOf(s)
	fieldName, ok := rt.FieldByName("Name")
	if ok {
		fmt.Println(fieldName.Tag)
	}

	fieldAge, ok2 := rt.FieldByName("Age")
	if ok2 {
		a := fieldAge.Tag.Get("a")
		b := fieldAge.Tag.Get("b")
		fmt.Println(a, b) //aaaa bbbb
	}

	fmt.Println(rt.Name())          //Student
	fmt.Println(rt.NumField())      //2
	fmt.Println(rt.PkgPath())       //main
	fmt.Println(rt.String())        //main.Student
	fmt.Println(rt.Kind().String()) //struct  返回该接口的具体分类
	for i := 0; i < rt.NumField(); i++ {
		fmt.Println(`这个类型的字段包含`, rt.Field(i).Name, i)
	}
}
