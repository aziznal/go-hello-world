package main

import (
	"fmt"
	"runtime"
)

func main() {
	// fmt.Println("Hello World!")
	//
	// k := "Hello World, Again!";
	//
	// fmt.Println(k);

	// loop();

	// switchCase()

	arrays()
}

func loop() {
	for i := 0; i <= 10; i++ {
		fmt.Print(fmt.Sprint(i) + ", ")
	}

	var foo float32 = 0

	for foo < 10 {
		foo += 0.1
	}
}

func switchCase() {
	os := runtime.GOOS

	switch os {
	case "linux":
		fmt.Println("You're awesome")
	default:
		fmt.Println(os)
	}
}

func arrays() {
	var messages []string = []string{"Hello", "darkness", "my", "old", "friend"}

    fmt.Printf("type of messages is %T\n", messages)

	fmt.Println(messages)

    fmt.Println("length of messages: " +fmt.Sprint(len(messages)))

    fmt.Println("capacity of messages: " +fmt.Sprint(cap(messages)))
}
