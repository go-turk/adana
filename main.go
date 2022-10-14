package main

import (
	"fmt"
	"github.com/go-turk/adana/user"
)

func main() {
	polatAlemdar := user.NewKisi("Polat", "Alemdar", 55)
	// Polatın babası
	omerCandan := user.NewKisi("Ömer", "Candan", 76)
	// Polatın annesi
	nazifeCandan := user.NewKisi("Nazife", "Candan", 72)
	// Polatın Baba Dedesi
	ahmetCandan := user.NewKisi("Ahmet", "Candan", 98)
	// Polatın Babaannesi
	ayseCandan := user.NewKisi("Ayşe", "Candan", 96)
	// Polatın Anne Dedesi
	necipAksakalli := user.NewKisi("Necip", "Aksakallı", 98)
	// Polatın Anneannesi
	necibeAksakalli := user.NewKisi("Necibe", "Aksakallı", 96)
	// tanımalamalrı yapalım
	polatAlemdar.SetBaba(omerCandan)
	polatAlemdar.SetAnne(nazifeCandan)
	omerCandan.SetBaba(ahmetCandan)
	omerCandan.SetAnne(ayseCandan)
	nazifeCandan.SetBaba(necipAksakalli)
	nazifeCandan.SetAnne(necibeAksakalli)

	// istersek şöyle de tanımlayabilirdir
	polatAlemdar.Baba.Baba = ahmetCandan
	polatAlemdar.Baba.Anne = ayseCandan
	polatAlemdar.Anne.Baba = necipAksakalli
	polatAlemdar.Anne.Anne = necibeAksakalli

	fmt.Println(polatAlemdar)
	fmt.Println(polatAlemdar.Baba)
	fmt.Println(polatAlemdar.Baba.Baba)
	fmt.Println(polatAlemdar.Baba.Anne)
	fmt.Println(polatAlemdar.Anne)
	fmt.Println(polatAlemdar.Anne.Baba)
	fmt.Println(polatAlemdar.Anne.Anne)
}
