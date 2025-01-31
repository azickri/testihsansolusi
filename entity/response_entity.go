package entity

type ResponseRegister struct {
	Number string `json:"no_rekening"`
}

type ResponseBalance struct {
	Balance int64 `json:"saldo"`
}

type ResponseError struct {
	Remark string `json:"remark"`
}
