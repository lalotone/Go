package main

import(
    "net"
    "fmt"
    "bufio"
    "os"
)
//var foo string = "foo"
var inputData string

func check(e error){
    if(e != nil){
        panic(e)
    }
}
func main(){
    conn, err := net.Dial("tcp", "127.0.0.1:777")
    check(err)
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan(){
	inputData = scanner.Text()
	fmt.Println(inputData)
	if len(inputData) > 3{
            fmt.Fprintln(conn, inputData)
        }
    }
}
