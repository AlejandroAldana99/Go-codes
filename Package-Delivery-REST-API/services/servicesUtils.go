package services

import (
	"time"

	"github.com/AlejandroAldana99/mvp_api/constants"
	"github.com/AlejandroAldana99/mvp_api/models"
)

func validCoordinates(lat float64, lng float64) bool {
	return (lat > -90.0 && lat < 90.0) && (lng > -180.0 && lng < 180.0)
}

func completeSize(data models.PackageData) (models.PackageData, bool) {
	special := false
	if data.Weight < constants.LimitWeightS {
		data.Size = constants.MeasureS
	} else if data.Weight < constants.LimitWeightM {
		data.Size = constants.MeasureM
	} else if data.Weight < constants.LimitWeightL {
		data.Size = constants.MeasureL
	} else {
		data.Size = constants.MeasureSpecial
		data.Service = constants.NameSpecialService
		special = true
	}

	return data, special
}

func compareTime(start time.Time, orderTime time.Time) bool {
	end := start.Add(time.Minute * 2)
	return orderTime.After(start) && orderTime.Before(end)
}

func compareStatus(orderStatus string) bool {
	return orderStatus == constants.OnWayStatus || orderStatus == constants.DeliveredStatus
}
