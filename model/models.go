package models

type Stock struct {
	SockID  int64  `json:"sockid"`
	Name    string `json:"name"`
	Price   int64  `json:"price"`
	Company string `json:"company"`
}