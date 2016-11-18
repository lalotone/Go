package main

import (
"fmt"
"net/http"
"strconv"
"time"
)
var i int = 0
var convStr string
//var c chan string = make(chan string)

//WORKING FIRST STUFF
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
//END WORKING FIRST STUFF

//Return 2 houndred posts
/*func returnTH(c chan string){
  for i = 0; i < 200; i++ {	
    convStr = strconv.Itoa(i)
    url := "http://www.urbanosdezaragoza.es/frm_esquemaparadatime.php?poste=" + convStr
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Houston, we've got problems")
    }else{
        if resp.StatusCode == 200{
            c <- "OK: The bus post "+convStr+" exists"
            printChannel(c)
        }else{
            c <- "WARN: The bus post "+convStr+" does not exist"
        }   
    }
  }
}
func printChannel(c chan string){
	r := <- c
	fmt.Println(r)
}*/

func main(){
    for i = 0; i < 1500; i++ {
		convStr = strconv.Itoa(i)
        go func() {
            convStr = strconv.Itoa(i)
			url := "http://www.urbanosdezaragoza.es/frm_esquemaparadatime.php?poste=" + convStr
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println("Houston, we've got problems")
			}else{
			if resp.StatusCode == 200{
				c <- "OK: The bus post "+convStr+" exists"
				printChannel(c)
			}else{
				c <- "WARN: The bus post "+convStr+" does not exist"
			}   
                  }
                   }
							  }
}
