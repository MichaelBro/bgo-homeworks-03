package main

import (
	"bgo-homeworks-03/pkg/card"
	"fmt"
)

func main() {
	master := &card.Card{
		Id:       1,
		Issuer:   "MasterCard",
		Balance:  73_000_00,
		Currency: "RUB",
		Number:   "5500 6600 5500 1234",
		Icon:     "/images/icon.svg",
		Transactions: []card.Transaction{
			{
				Id:        1,
				Amount:    1_203_91,
				Timestamp: 1608768400,
				MCC:       "5812",
				Status:    "processing",
			},
			{
				Id:        2,
				Amount:    2_000_00,
				Timestamp: 1608768400,
				MCC:       card.MccRefill,
				Status:    "success",
			},
			{
				Id:        3,
				Amount:    735_55,
				Timestamp: 1608854400,
				MCC:       "5411",
				Status:    "processing",
			},
			{
				Id:        4,
				Amount:    203_91,
				Timestamp: 1608853818,
				MCC:       "5812",
				Status:    "processing",
			},
		},
	}

	anyTransaction := card.Transaction{
		Id:        5,
		Amount:    200_000_00,
		Timestamp: 1608853818,
		MCC:       "5815",
		Status:    "reject",
	}

	card.AddTransaction(master, anyTransaction)

	mccArrayToSum := []string{"5812", "5411"}
	sum := card.SumByMCC(master.Transactions, mccArrayToSum)

	category := card.TranslateMCC(master.Transactions[1].MCC)

	fmt.Println("MasterCard: ", master)
	fmt.Println("Sum by MCC: ", sum)
	fmt.Println("Category by MCC: ", category)

	category = card.TranslateMCC("6666")
	fmt.Println("Category by MCC: ", category)

	category = card.TranslateMCC(master.Transactions[2].MCC)
	fmt.Println("Category by MCC: ", category)

}
