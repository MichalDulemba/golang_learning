package main
import (
"fmt")

func main(){




z:=owoc{"jablko", "duzy"}
noweDanie:=danieOwocowe{
jakiOwoc:z,
nazwaDania:"salatka"}
fmt.Println(noweDanie)
fmt.Println(z)
noweDanie.wyrzuc()
//z.pokaz()
another(z)
another(noweDanie)
}

type owoc struct{
     rodzaj string
     rozmiar string
}

type danieOwocowe struct{
     jakiOwoc owoc
     nazwaDania string
}

type jedzenie interface{
wyrzuc()
}

func another(j jedzenie){
  fmt.Println("another interface func", j)
}
func (x danieOwocowe) wyrzuc() {
     fmt.Println("wyrzuc danie", x.nazwaDania)
}
func (x owoc) wyrzuc(){
     fmt.Println("pokaz", x.rodzaj)
}
