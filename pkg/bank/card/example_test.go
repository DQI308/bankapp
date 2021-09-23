package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleWithdraw_positiv() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: true}, 10_000_00)
	fmt.Println(result.Balance)
	//Output:
	//1000000
}

func ExampleWithdraw_noMoney() {
	result := Withdraw(types.Card{Balance: 0, Active: true}, 10_000_00)
	fmt.Println(result.Balance)
	//Output:
	//0

}

func ExampleWithdraw_inactive() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: false}, 10_000_00)
	fmt.Println(result.Balance)
	//Output:
	//2000000
}

func ExampleWithdraw_limit() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: true}, 10_000_000)
	fmt.Println(result.Balance)
	//Output:
	//2000000
}

func ExampleDeposit_positiv() {
	card := types.Card{
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: types.TJS,
		Color:    "White",
		Name:     "DQI",
		Active:   true,
	}
	Deposit(&card, 10_000)
	fmt.Println(card.Balance)
	//Output:
	//10000

}

func ExampleDeposit_negativbalance() {
	card := types.Card{
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  types.Money(-1000000),
		Currency: types.TJS,
		Color:    "White",
		Name:     "DQI",
		Active:   true,
	}
	Deposit(&card, 20_000_00)
	fmt.Println(card.Balance)
	//Output:
	//1000000

}
func ExampleDeposit_inactive() {
	card := types.Card{
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: types.TJS,
		Color:    "White",
		Name:     "DQI",
		Active:   false,
	}
	Deposit(&card, 10_000)
	fmt.Println(card.Balance)
	//Output:
	//0
}

func ExampleDeposit_limit() {
	card := types.Card{
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: types.TJS,
		Color:    "White",
		Name:     "DQI",
		Active:   true,
	}
	Deposit(&card, 10_000_000)
	fmt.Println(card.Balance)
	//Output:
	//0
}

func ExampleAddBonus_positiv() {
	card := types.Card{
		PAN:        "5058 xxxx xxxx 0001",
		Balance:    1_00,
		MinBalance: 10_000,
		Currency:   types.TJS,
		Color:      "White",
		Name:       "DQI",
		Active:     true,
	}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	//Output:
	//124
}

func ExampleAddBonus_inactive() {
	card := types.Card{
		PAN:        "5058 xxxx xxxx 0001",
		Balance:    1_00,
		MinBalance: 10_000,
		Currency:   types.TJS,
		Color:      "White",
		Name:       "DQI",
		Active:     false,
	}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	//Output:
	//100
}

func ExampleAddBonus_minbalanse() {
	card := types.Card{
		PAN:        "5058 xxxx xxxx 0001",
		Balance:    types.Money(-1),
		MinBalance: 10_000,
		Currency:   types.TJS,
		Color:      "White",
		Name:       "DQI",
		Active:     true,
	}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	//Output:
	//-1
}

func ExampleAddBonus_limit() {
	card := types.Card{
		PAN:        "5058 xxxx xxxx 0001",
		Balance:    1_00,
		MinBalance: 10_000_000,
		Currency:   types.TJS,
		Color:      "White",
		Name:       "DQI",
		Active:     true,
	}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	//Output:
	//24757
}
