package main

import (
	"fmt"
	"slices"
	"time"
)

// 定义结构体
type Student struct {
	ID        int
	Name      string
	Age       int
	CreatedAt time.Time
}

// 初始化结构体的几种方式
func createStudents() []Student {
	// 方式1：顺序初始化
	student1 := Student{1, "张三", 20, time.Now()}

	// 方式2：键值对初始化
	student2 := Student{
		ID:   2,
		Name: "李四",
		Age:  21,
	}

	// 方式3：new关键字
	student3 := new(Student)
	student3.ID = 3
	student3.Name = "王五"
	student3.Age = 22

	return []Student{student1, student2, *student3}
}

// 增：添加学生
func addStudent(students []Student, student Student) []Student {
	return append(students, student)
}

// 删：删除指定ID的学生
func deleteStudent(students []Student, id int) []Student {
	for i, s := range students {
		if s.ID == id {
			// return append(students[:i], students[i+1:]...)
			return slices.Delete(students, i, i+1)
		}
	}
	return students
}

// 改：更新学生信息
func updateStudent(students []Student, student Student) {
	for i, s := range students {
		if s.ID == student.ID {
			students[i] = student
			return
		}
	}
}

// 查：查找指定ID的学生
func findStudent(students []Student, id int) *Student {
	for _, s := range students {
		if s.ID == id {
			return &s
		}
	}
	return nil
}

func main() {
	students := createStudents()

	// 测试CRUD操作
	fmt.Printf("初始学生列表：%+v\n", students)

	// 测试添加
	newStudent := Student{4, "赵六", 23, time.Now()}
	students = addStudent(students, newStudent)

	// 测试查找
	if student := findStudent(students, 2); student != nil {
		fmt.Printf("找到学生：%+v\n", *student)
	}

	// 测试更新
	updateStudent(students, Student{2, "李四改", 25, time.Now()})

	// 测试删除
	students = deleteStudent(students, 1)

	fmt.Printf("最终学生列表：%+v\n", students)
}
