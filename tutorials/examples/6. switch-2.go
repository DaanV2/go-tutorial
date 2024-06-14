switch time.Now().Weekday() {
case time.Saturday, time.Sunday:
    fmt.Println("It's the weekend")
default:
    fmt.Println("It's a weekday")
}

switch time.Now().Weekday() {
case time.Sunday:
    fallthrough
case time.Saturday:
    fmt.Println("It's the weekend")
default:
    fmt.Println("It's a weekday")
}