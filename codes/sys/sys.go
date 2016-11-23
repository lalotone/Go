package tools 

import(
    "fmt"
    //"os"
    "io/ioutil"
    "strconv"
    "strings"
    "bufio"
    "os"
)

var temp int
var strData string
var ramData []string

func check(e error){
    if e != nil{
        panic(e)
    }
}
func getTempStatus(){
    tempRawData, err := ioutil.ReadFile("/sys/class/thermal/thermal_zone0/temp")
    strData = string(tempRawData)
    strDat := strings.Replace(strData, "\n", "", -1)
    temp, err = strconv.Atoi(strDat)
    check(err)
    fmt.Println("System temp: ",temp / 1000, "ºC")
}
func getRamStatus(){
    ramRawData, err := os.Open("/proc/meminfo") 
    check(err)
    defer ramRawData.Close()
    scanner := bufio.NewScanner(ramRawData)
    for scanner.Scan(){
        ramData = append(ramData, scanner.Text())
    }
    ramData = ramData[0:3]
    //z := len(lines)
    //fmt.Println(z)
    for i := 0; i < len(ramData); i++{
        fmt.Println(ramData[i])
    }
}

//func main(){
//    getTempStatus()
//    getRamStatus()
//}



