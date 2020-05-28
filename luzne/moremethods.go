package main

import (
  "fmt"
)

type silnik_samochodowy struct{
  cylindry int
  moc int
  olej string
  Stan string
}

func (s *silnik_samochodowy) odpal(sposob string) (succ string, error error){
  fmt.Println("start engine with", sposob)
  s.Stan = "ON"
  return "successfull", nil
}

func (s *silnik_samochodowy) zgas(sposob string) (succ string, error error){
  fmt.Println("zgas", sposob)
  s.Stan = "OFF"
  return "successfull", nil
}

func (s *silnik_samochodowy) check() (state string, error error){
  fmt.Println("silnik w stanie", s.Stan)
  return "pleple", nil
}

type maszyna interface{
  odpal(sposob string) ( string,  error)
  check() ( string,  error)
}
func wyslijDoNaprawy(m maszyna) (zgloszenie_id int, error error){
  // *m.Stan="aaa"
  return 42, nil
}

func main(){
    nowakia := silnik_samochodowy{
      cylindry:5,
      moc:300,
      olej:"elf",
      Stan: "wylaczony",
    }

    fmt.Println(nowakia)
    nowakia.odpal("kluczykiem")
    fmt.Println(nowakia)
    nowakia.zgas("zdalnie")
    fmt.Println(nowakia)

    nowakia.check()

    // out, err := wyslijDoNaprawy(&nowakia)
    // fmt.Println(out, err)
    // fmt.Println(nowakia)
}
