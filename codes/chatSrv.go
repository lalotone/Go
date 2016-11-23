package main

import(
    "net"
    "bufio"
    "fmt"
)
func check(e error){
    if e!=nil{
        panic(e)
    }
}
func main(){
    link, err := net.Listen("tcp", ":777")
    check(err)
    conn, _ := link.Accept()
    check(err)
    for{
        data, _ := bufio.NewReader(conn).ReadString('\n')
        //check(err)
        fmt.Print(string(data))
    }
}
