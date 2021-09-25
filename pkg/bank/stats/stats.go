package stats

import "github.com/DQI308/bank/pkg/bank/types"

//Avg рассчитывает средную сумму платежа
func Avg(payments []types.Payment)types.Money{
	sum:=types.Money(0)
	base:=types.Money(0)
	for _,payment:=range payments{
		sum+=payment.Amount
		base+=1
	} 
	result:=sum/base
	return result
}