package model

type Stock struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Company string `json:"company"`
	Price   int    `json:"price"`
}
