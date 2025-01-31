package entity

type BodyRegister struct {
	Name  string `validate:"required" json:"nama"`
	NIK   string `validate:"required" json:"nik"`
	Phone string `validate:"required" json:"no_hp"`
}

type BodyDepositOrWithdraw struct {
	Number  string `json:"no_rekening"`
	Nominal int64  `json:"nominal"`
}
