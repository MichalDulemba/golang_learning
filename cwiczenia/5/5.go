package main

import (
	"fmt"
	"reflect"
)

type lody struct {
	nazwalodow string
	smak       string
}
type osobaZlodami struct {
	name       string
	secondname string
	icecream   []lody
}

func main() {
	lodymalinowe := lody{
		nazwalodow: "mali",
		smak:       "malinowe",
	}
	lodytruskawkowe := lody{
		nazwalodow: "rewela",
		smak:       "truskawka",
	}
	lodywloskie := lody{
		nazwalodow: "kreciolki",
		smak:       "smietana",
	}
	osoba := osobaZlodami{
		name:       "Nadia",
		secondname: "Dulemba",
		icecream:   []lody{lodymalinowe, lodytruskawkowe},
	}
	osoba2 := osobaZlodami{
		name:       "Igor",
		secondname: "Dulemba_",
		icecream:   []lody{lodywloskie},
	}
	fmt.Println(osoba)
	fmt.Println(osoba2)
	fmt.Println(osoba.icecream)
	fmt.Println(reflect.TypeOf(osoba))

	for _, value := range osoba.icecream {
		fmt.Println(value)
	}

	lista := make(map[string]osobaZlodami)
	lista[osoba.secondname] = osoba
	lista[osoba2.secondname] = osoba2

	fmt.Println(lista)

	drugalista := map[string]osobaZlodami{
		osoba.secondname:  osoba,
		osoba2.secondname: osoba2,
	}
	fmt.Println(drugalista)
	fmt.Println("\n\n\n\n")

	type vehicle struct {
		doors int
		color string
	}

	type truck struct {
		vehicle
		fourWheel bool
		luxury    bool
	}

	type sedan struct {
		vehicle
		fourWheel bool
		luxury    bool
	}

	camry := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "blue",
		},
		// color:     "blue",
		fourWheel: true,
		luxury:    false,
	}
	fmt.Println(camry)

	largeford := truck{
		vehicle: vehicle{
			doors: 3,
			color: "red",
		},
		fourWheel: false,
		luxury:    true,
	}
	fmt.Println(largeford.color)
	fmt.Println("\n\n\n\n")

	d := struct {
		name       string
		secondname string
	}{name: "sded", secondname: "sdfsdfsdf"}
	fmt.Println(d)
}
