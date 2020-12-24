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
		Number:   "4400 6600 5500 1234",
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
				MCC:       "6540",
				Status:    "success",
			},
			{
				Id:        3,
				Amount:    735_55,
				Timestamp: 1608854400,
				MCC:       "5411",
				Status:    "processing",
			},
		},
	}
	fmt.Println(master)
}
