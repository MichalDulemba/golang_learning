package main

import (
"fmt"
 "time"
"image"
_ "image/jpeg"
"os"
)


func main(){
    fmt.Println("testtest")
    var g float32
    g = 349.0
    fmt.Println("sa ",g)
    begin := time.Now()
    for i:=0; i<100000000; i++{
        fmt.Println("pleplepel",i)
        filecontent, err := os.Open("./test.jpg")
        if err!=nil {
        fmt.Println("problem1",err)}

        img, _, err := image.DecodeConfig(filecontent)
        if err!=nil {
                fmt.Println(os.Stderr, "problem2", err)}
        fmt.Println(img.Width, img.Height)
        //print(filecontent, error)
//         if i % 2 == 0 {
//             fmt.Println("test failed")
//             }
        filecontent.Close()
    }
    finish := time.Since(begin)
    fmt.Println("this took ", finish)
    //go anotherfunc()
    //go secondfunc()
    //time.Sleep(40*time.Second)
}

// func anotherfunc() {
//     fmt.Println("nowy test")
//     for i:=0; i<1000; i++{
//     fmt.Println("first func",i)
//     time.Sleep(2* time.Second)
//     }
//     }
//
// func secondfunc() {
//     for i:=0; i<500; i++{
//         fmt.Println("second func",i)
//         time.Sleep(1 * time.Second)
//         }
//
//     }

