package main

import (
	"fmt"
)

var intData = 100

func userLogin(name []string) [][]string {
	var varName *string
	fmt.Printf("ptr 的值为 : %f\n", varName)

	var ok string
	fmt.Println(ok)

	mapData := make(map[string]string)
	fmt.Println(mapData)

	type User struct {
		Name string
		Age  int
	}

	type UU struct {
		data User
		size int
	}

	userStruct := User{"hangman", 18}

	fmt.Println(userStruct.Age)

	UUData := UU{userStruct, 11}
	fmt.Println(UUData.data.Name)

	return [][]string{name, name}
}

func modifyData(input *int) {
	*input -= 1
}

func factorial(n int) int {
	// 基准条件
	if n == 1 {
		return 1
	}
	// 递归条件
	return n * factorial(n-1)
}

type Person interface {
	GetName(name string) string
	GetAge() int
}

func main() {
	fmt.Println(userLogin([]string{"hangman", ""}))
	modifyData(&intData)
	fmt.Println(intData)
	fmt.Println(factorial(5))
}
