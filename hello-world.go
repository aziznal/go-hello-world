package main

import "fmt"

func main() {
    fmt.Println("Hello World!")

    k := "Hello World, Again!";

    fmt.Println(k);

    loop();
}

func loop() {
    for i := 0; i <= 10; i++ {
        fmt.Print(fmt.Sprint(i) + ", ");
    }

    var foo float32 = 0

    for ; foo < 10; {
        foo += 0.1
    }
}

