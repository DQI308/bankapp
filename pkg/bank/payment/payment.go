package payment

import "bank/pkg/bank/types"

//Max возвращает максимальний платёж
func Max(payments []types.Payment) types.Payment {
	k := 0
	maxPayment := payments[0]
	for i:=0;i<len(payments);i++{
     if payments[i].Amount>maxPayment.Amount{
		 maxPayment=payments[i]
		 k=i
	 }
	}
	return payments[k]
}
