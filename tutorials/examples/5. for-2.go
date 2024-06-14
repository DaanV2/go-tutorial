package main

import "fmt"

func main() {
    for {
        fmt.Println("loop")
        break
    }

    for n := range 6 {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }

    i := []int{1, 2, 3}
    for k, v := range i {
        fmt.Println(k, v)
    }
}