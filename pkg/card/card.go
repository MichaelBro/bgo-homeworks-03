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
