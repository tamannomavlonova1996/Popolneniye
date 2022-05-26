package main

import (
	"fmt"
	"popolneniye/bank"
	"popolneniye/popolneniye"
)

func main() {

	toma := bank.Card{
		ID:       1,
		PAN:      "20250297",
		Balance:  50,
		Currency: "TJS",
		Color:    "white",
		Name:     "Tamanno",
		Active:   true,
	}

	popolneniye.Deposit(&toma, 40)
	fmt.Println(toma)
	//fmt.Println(toma.Balance)

}
