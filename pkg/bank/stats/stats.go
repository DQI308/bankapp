package stats

import "bank/pkg/bank/types"

//Avg расчитывает средную суммуплатежа.
func Avg(payments []types.Payment)types.Money{
	sum:=types.Money(0)
	cnt:=types.Money(0)
	for _,payment:=range payments{
		sum+=payment.Amount
		cnt++
	}
	result:=types.Money(sum/cnt)
	return result
}
