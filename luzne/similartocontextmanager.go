//https://abhinavg.net/posts/context-manager-antipattern/

package main

import (
"fmt"
"os" )

func main(){
myFile:="variadics.go"
myFile2:="variadics2.go"
z, closing := fileOpen(myFile)
defer closing()

z2, closing := fileOpen(myFile2)
defer closing()

fmt.Println(z)
fmt.Println(z2)
}

func fileOpen(name string) (MyFILE *os.File, cleanup func() ){
fileopened, err := os.OpenFile(name, os.O_RDONLY, 0644)
if err!=nil {
fmt.Println("problem with this file")
}
return fileopened, func() {fmt.Println("closing") 
                           fileopened.Close()}

}

