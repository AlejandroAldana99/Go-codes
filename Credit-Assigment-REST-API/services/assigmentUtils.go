package services

import (
	"github.com/AlejandroAldana99/Credit-Assigment-REST-API/models"
)

func calculateLoans(invest float64) []models.CombinationsData {
	combinatios := []models.CombinationsData{}
	combinatios = append(combinatios, assign(invest))

	return combinatios
}

func assign(invest float64) models.CombinationsData {
	data := models.CombinationsData{}
	if invest < 300 {
		return data
	}

	credit_type_300_max := int(invest / 300)
	credit_type_500_max := int(invest / 500)
	credit_type_700_max := int(invest / 700)
	var combination int32

	for i := 0; i <= credit_type_300_max; i++ {
		for j := 0; j <= credit_type_500_max; j++ {
			for k := 0; k <= credit_type_700_max; k++ {
				combination = int32((300 * i) + (500 * j) + (700 * k))

				if combination == int32(invest) {
					data.CreditType300 = i
					data.CreditType500 = j
					data.CreditType700 = k
					return data
				}

				if combination > int32(invest) {
					break
				}
			}
		}
	}

	return data
}
