package card

import "bank/pkg/bank/types"

//Футкция выпуска карт
func IssueCard(currency types.Currency, color string, name string) types.Card {
	card := types.Card{
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: currency,
		Color:    color,
		Name:     name,
		Active:   true,
	}
	return card
}

//Функция снятия средств из карты
func Withdraw(card types.Card, amount types.Money) types.Card {
	if card.Active == false {
		return card
	}
	if card.Balance < amount {
		return card
	}
	if amount < 0 {
		return card
	}
	if amount > 20_000_00 {
		return card
	}
	card.Balance -= amount
	return card
}

//Функция Зачисления средств
func Deposit(card *types.Card, amount types.Money) {
	if !(*card).Active {
		return
	}
	if amount > 5000000 {
		return
	}
	if amount < 0 {
		return
	}
	(*card).Balance += amount
	return
}

//Функция зачисления бонуса
func AddBonus(card *types.Card, percent int, dayslnMonth int, dayslnYear int) {
	bonus := (int((*card).MinBalance) * percent * dayslnMonth) / dayslnYear / 100
	if !(*card).Active || (*card).Balance <= 0 || bonus > 5_000_00 {
		return
	}
	(*card).Balance += types.Money(bonus)
}


//Функция суммирования денег на всех картах
func Total(cards []types.Card)types.Money{
	sum:=types.Money(0)
	var card types.Card
	for _,card=range cards{
		if card.Balance>=0 && card.Active{
			sum+=card.Balance
		}
	}
	return sum
}