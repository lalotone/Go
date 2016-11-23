package main

import (
"fmt"
"net/http"
"strconv"
//"time"
)
var i int = 0
var convStr string
var z string

/*func main(){
    for i = 0; i < 1500; i++ {
        z = strconv.Itoa(i)
        url := "http://www.urbanosdezaragoza.es/frm_esquemaparadatime.php?poste=" + z
        resp, err := http.Get(url)
        if err != nil {
            fmt.Println("Houston, we've got problems")
        }else{
            if resp.StatusCode == 200{
                fmt.Println("OK: El poste "+z+" existe")
            }else{
                fmt.Println("WARN: El poste "+z+" NO existe")
            }
        }
    }
}*/

//FAILED ATTEMPT WITH GO ****** THREADS
func ping(fromPost, toPost int, message chan string){
    for i = fromPost; i < toPost; i++ {
        z := strconv.Itoa(i)
        url := "http://www.urbanosdezaragoza.es/frm_esquemaparadatime.php?poste=" + z
        resp, err := http.Get(url)
        if err != nil {
            fmt.Println("Houston, we've got problems")
        }else{
            if resp.StatusCode == 200{
                dataOk := "OK: El poste "+z+" existe"
                message <- dataOk
            }else{
                dataBad := "WARN: El poste "+z+" NO existe"
                message <- dataBad
            }
        }
    }
}
func printer(message chan string){
    for{
        msgOne := <- message
        fmt.Println(msgOne)
    }
}
func main(){
    message := make(chan string)
    go ping(0, 200, message)
    go printer(message)
    fmt.Println("Done")
}
