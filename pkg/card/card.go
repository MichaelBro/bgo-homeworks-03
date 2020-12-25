package card

type Transaction struct {
	Id        int64
	Amount    int64
	Timestamp int64
	MCC       string
	Status    string
}

type Card struct {
	Id           int64
	Issuer       string
	Balance      int64
	Currency     string
	Number       string
	Icon         string
	Transactions []Transaction
}

const MccRefill = "6000"

func AddTransaction(card *Card, transaction Transaction) {
	card.Transactions = append(card.Transactions, transaction)
}

func SumByMCC(transactions []Transaction, mcc []string) int64 {
	var sum int64
	for _, transaction := range transactions {
		if IsMCCInTransaction(transaction.MCC, mcc) {
			sum += transaction.Amount
		}
	}
	return sum
}

func IsMCCInTransaction(mcc string, array []string) (status bool) {
	status = false
	for _, currentMcc := range array {
		if mcc == currentMcc {
			status = true
		}
	}
	return
}
