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

	// arrays()

	maps()
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

	fmt.Println("length of messages: " + fmt.Sprint(len(messages)))

	fmt.Println("capacity of messages: " + fmt.Sprint(cap(messages)))
}

type Color struct {
	R, G, B int
}

func maps() {
	var colors = map[string]Color{
        "green": {
            G: 255,
        },
    }

    colors["red"] = Color{R: 255}

	fmt.Println(colors)


    colors["blue"] = Color{B: 255}

    fmt.Println("blue: " + fmt.Sprint(colors["blue"]))

    _, ok := colors["red"]

    fmt.Println("Is red a key in colors? " + fmt.Sprint(ok))
}
