package main

import (
	"fmt"
	"strings"
)

// 定义结构体
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// 值接收者方法：获取全名
func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

// 指针接收者方法：更新年龄
func (p *Person) UpdateAge(newAge int) {
	p.Age = newAge
}

// 值接收者方法：检查是否成年
func (p Person) IsAdult() bool {
	return p.Age >= 18
}

// 指针接收者方法：更新名字为大写
func (p *Person) UppercaseName() {
	p.FirstName = strings.ToUpper(p.FirstName)
	p.LastName = strings.ToUpper(p.LastName)
}

// 创建新的Person实例
func NewPerson(firstName, lastName string, age int) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func main() {
	// 创建Person实例
	person := NewPerson("张", "三", 20)

	// 使用值接收者方法
	fmt.Printf("全名: %s\n", person.FullName())
	fmt.Printf("是否成年: %v\n", person.IsAdult())

	// 使用指针接收者方法
	person.UpdateAge(21)
	fmt.Printf("更新后的年龄: %d\n", person.Age)

	person.UppercaseName()
	fmt.Printf("转换为大写后的姓名: %s\n", person.FullName())

	// 方法集示例
	// 值类型可以调用值接收者方法
	p1 := Person{"李", "四", 25}
	fmt.Printf("p1全名: %s\n", p1.FullName())

	// 指针类型可以调用所有方法
	p2 := &Person{"王", "五", 30}
	p2.UpdateAge(31)
	fmt.Printf("p2更新后的年龄: %d\n", p2.Age)
}
