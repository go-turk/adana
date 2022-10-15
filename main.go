package main

import "fmt"

type Kisi struct {
	Ad    string
	Soyad string
	Yas   int
	Anne  *Kisi
}

func NewKisi(ad string, soyad string, yas int) *Kisi {
	return &Kisi{
		Ad:    ad,
		Soyad: soyad,
		Yas:   yas,
	}
}

func main() {
	polatAlemdar := NewKisi("Polat", "Alemdar", 30)
	nazifeHanim := NewKisi("Nazife", "Candan", 76)
	polatAlemdar.Anne = nazifeHanim

	fmt.Println(polatAlemdar.Ad)
	fmt.Println(polatAlemdar.Anne.Ad)
	fmt.Println(polatAlemdar.Anne.Anne)

}
