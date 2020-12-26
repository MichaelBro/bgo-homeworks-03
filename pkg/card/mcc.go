package card

func TranslateMCC(code string) string {
	mcc := map[string]string{
		MccRefill: "Пополнение",
		"5411":    "Супермаркеты",
		"5533":    "Автоуслуги",
		"5812":    "Рестораны",
		"5912":    "Аптеки",
	}

	category, status := mcc[code]

	if status {
		return category
	}

	return "Category undefined."
}
