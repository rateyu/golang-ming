// MyProject/models/person.go
package models

import "fmt"

// Person 结构体代表一个人
type Person struct {
	Name string
	Age  int
}

// Greet 方法使 Person 实例能够打招呼
func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}
