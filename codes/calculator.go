package main

import(
"fmt" 
)

var x,y,option,result int

func options(){
    fmt.Println("What do you want to do?")
    fmt.Println("1.-Sum")
    fmt.Println("2.-Subtract")
}
func sum(x,y int)(int){
    sum := x + y
    return sum
}
func subtract(a,b int)(int){
    subtract := a - b
    return subtract
}
func main(){
    fmt.Println("CALCULATOR")
    fmt.Println("Enter first number")
    fmt.Scan(&x)
    fmt.Println("Enter second number")
    fmt.Scan(&y)
    options()
    fmt.Scan(&option)
    switch option{
        case 1:
            result := sum(x,y)
            fmt.Println(result)
        break
        case 2:
            fmt.Println(subtract(x,y))
        break
    }
}
