package main

import "fmt"

type irtibatBilgi struct {
	eposta     string
	postaKodu int
}

type kisi struct {
	ad      string
	soyad   string
	irtibatBilgi
}

func main() {
	kisiler := kisi{
		ad:    "Mehmet",
		soyad: "Kavak",
		irtibatBilgi: irtibatBilgi{
			eposta:     "mehmetsukru.kavak07@gmail.com",
			postaKodu: 34100,
		},
	}
	kisiler.ekranaYaz()
	kisilerPointer := &kisiler
	kisilerPointer.adDegistir("Mehmet Şükrü")
	fmt.Println("")
	kisiler.ekranaYaz()

}

func (k kisi) ekranaYaz(){
	fmt.Printf("%+v", k)
}

func (pointerKisi *kisi) adDegistir(yeniAd string){
	(*pointerKisi).ad = yeniAd
}
