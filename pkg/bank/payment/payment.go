package payment

import "bank/pkg/bank/types"

//Max возвращает максимальний платёж
func Max(payments []types.Payment) types.Payment {
	k := 0
	maxPayment := payments[0]
	for i := 0; i < len(payments); i++ {
		if payments[i].Amount > maxPayment.Amount {
			maxPayment = payments[i]
			k = i
		}
	}
	return payments[k]
}

//PaymentSource реализует слай источников оплаты
func PaymentSources(cards []types.Card) []types.PaymentSource {
	var slicePaymentSource []types.PaymentSource
	for _, card := range cards {
		if card.Balance <= 0 || !card.Active {
			continue
		} else {
			spSource := types.PaymentSource{
				Type:    card.Name,
				Number:  string(card.PAN),
				Balance: card.Balance,
			}
			slicePaymentSource = append(slicePaymentSource, spSource)
		}
	}
	return slicePaymentSource
}
