package popolneniye

import "popolneniye/bank"

func Deposit(card *bank.Card, amount bank.Money) {

	if !card.Active {
		return
	}
	if card.Balance < 0 {
		return
	}
	if amount > 50_000 {
		return
	}

	card.Balance = card.Balance + amount

}
