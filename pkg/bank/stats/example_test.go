package stats

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:1,
			Amount:10_000,
		},{
			ID:2,
			Amount:10_000,
		},
	}
	result:=Avg(payments)
	fmt.Println(result)
	//Output:
	//20000

}