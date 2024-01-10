package main

import (
	"fmt"
	"time"
)

type Medicine struct {
	Name         string
	Manufactured time.Time
}

func main() {
	// Припустимо, що сьогодні 10 січня 2024 року
	today := time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC)

	medicines := []Medicine{
		{Name: "MedicineA", Manufactured: time.Date(2023, time.July, 1, 0, 0, 0, 0, time.UTC)},
		{Name: "MedicineB", Manufactured: time.Date(2023, time.August, 1, 0, 0, 0, 0, time.UTC)},
		{Name: "MedicineC", Manufactured: time.Date(2023, time.October, 1, 0, 0, 0, 0, time.UTC)},
	}

	forSale, forUse, forDisposal := categorizeMedicines(medicines, today)

	fmt.Println("For Sale:", forSale)
	fmt.Println("For Use:", forUse)
	fmt.Println("For Disposal:", forDisposal)
}

func categorizeMedicines(medicines []Medicine, today time.Time) ([]Medicine, []Medicine, []Medicine) {
	var forSale, forUse, forDisposal []Medicine

	sixMonthsAgo := today.AddDate(0, -6, 0)
	for _, medicine := range medicines {
		if medicine.Manufactured.After(sixMonthsAgo) {
			forSale = append(forSale, medicine)
		} else if medicine.Manufactured.Before(sixMonthsAgo) && medicine.Manufactured.After(sixMonthsAgo.AddDate(0, -6, 0)) {
			forUse = append(forUse, medicine)
		} else {
			forDisposal = append(forDisposal, medicine)
		}
	}

	return forSale, forUse, forDisposal
}
