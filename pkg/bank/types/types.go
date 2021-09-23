package types

//Money представляет собой денежную сумму 
type Money int64

//Currency представляет код валюты
type Currency string


//Код валют
const (
	TJS Currency="TJS"
	RUB Currency="RUB"
	USD Currency="USD"
)


//PAN представляет номер карты
type PAN string

//Card представляет информацию о платёжной карте 
type Card struct{
	ID int
	PAN PAN
	Balance Money
	MinBalance Money
	Currency Currency
	Color string
	Name string
	Active bool
}

//Payment представляет информацию о платеже.
type Payment struct{
	ID int
	Amount Money
}