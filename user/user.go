package user

type Kisi struct {
	Adi    string
	Soyadi string
	Yasi   int
	Baba   *Kisi
	Anne   *Kisi
}

func NewKisi(adi, soyadi string, yasi int) *Kisi {
	return &Kisi{adi, soyadi, yasi, nil, nil}
}

func (k *Kisi) SetBaba(baba *Kisi) {
	k.Baba = baba
}

func (k *Kisi) SetAnne(anne *Kisi) {
	k.Anne = anne
}
