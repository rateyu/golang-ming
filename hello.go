package main

import (
	"fmt"
	"golang-ming/models"
)

type Person struct {
	Name string
	Age  int
}

func main() {

	// 直接初始化
	s := []int{1, 2, 3}
	fmt.Println(s[0])

	s = []int{1, 2, 3, 4, 5}
	// 删除索引为2的元素（元素3）
	s = append(s[:1], s[3:]...)
	for i, mys := range s {
		fmt.Println(i,mys)
	}


	s = s[:0] // 清空切片

	// 使用 make 创建切片，指定长度和容量
	s1 := make([]int, 3, 5) // 长度为3，容量为5
	s1 = append(s1, 6,7,8,9)
	fmt.Println(len(s1))
	for i, v := range s1 {
		fmt.Println(i,v)
	}

	// 从数组创建切片
	arr := [5]int{1, 2, 3, 4, 5}
	s2 := arr[1:4] // 创建一个包含元素 2, 3, 4 的切片
	fmt.Println(s2[0])

	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	copy(dst, src) // 将 src 的内容复制到 dst


	// 创建 Person 类型的实例
	person := models.Person{Name: "Alice", Age: 30}
	// 调用 Greet 方法
	person.Greet()

	// 创建 Person 类型的切片
	people := []Person{}

	// 添加 Person 实例到切片中
	people = append(people, Person{Name: "Alice", Age: 30})
	people = append(people, Person{Name: "Bob", Age: 25})

	// 遍历切片
	for _, person := range people {
		fmt.Printf("%s is %d years old.\n", person.Name, person.Age)
	}

	// 修改切片中的元素
	if len(people) > 1 {
		people[1].Age = 26 // 修改 Bob 的年龄
	}

	// 再次遍历切片以显示修改后的数据
	for _, person := range people {
		fmt.Printf("%s is now %d years old.\n", person.Name, person.Age)
	}



	// 创建一个映射，键类型为 string，值类型为 Person
	peoplemap := make(map[string]Person)

	// 向映射中添加元素
	peoplemap["Alice"] = Person{Name: "Alice", Age: 30}
	peoplemap["Bob"] = Person{Name: "Bob", Age: 25}

	// 访问映射中的元素
	alice := peoplemap["Alice"]
	fmt.Println("Alice's age:", alice.Age)

	// 修改映射中的元素
	if person, ok := peoplemap["Bob"]; ok {
		person.Age = 26
		peoplemap["Bob"] = person // 必须重新赋值，因为结构体是值类型
	}

	// 遍历映射
	for key, value := range peoplemap {
		fmt.Printf("%s is %d years old.\n", key, value.Age)
	}

	// 删除映射中的元素
	delete(peoplemap, "Bob")
	fmt.Println("After deleting Bob:", peoplemap)


}
