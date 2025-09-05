package main

import "fmt"

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	sliceA := make([]int, 0)
	sliceA = append(sliceA, 1)
	fmt.Println(sliceA)

	go func() {

		for {
			<-ch1
			fmt.Println("1")
			ch2 <- 2

		}

	}()

	go func() {
		for {
			<-ch2
			fmt.Println("2")
			ch1 <- 1
		}

	}()

	//ch1 <- 1
	ch2 <- 2

}
