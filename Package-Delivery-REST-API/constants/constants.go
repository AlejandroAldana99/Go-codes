package constants

const (
	MeasureS            = "S"
	MeasureM            = "M"
	MeasureL            = "L"
	MeasureSpecial      = "especial"
	NameOrdinaryService = "estandar"
	NameSpecialService  = "especial"
	LimitWeightS        = 5.0
	LimitWeightM        = 15.0
	LimitWeightL        = 25.0
	CreateStatus        = "creado"
	PickupStatus        = "recolectado"
	StationStatus       = "en_estacion"
	OnWayStatus         = "en_ruta"
	DeliveredStatus     = "entregada"
	CancelStatus        = "cancelada"
	GeneralStatus       = "Done"
	GenericName         = "unknown"
	GenericHeader       = "Authorization"
	AdminRole           = "Admin"
	Key                 = "secret"
)

var StatusList = map[string]bool{
	"creado":      true,
	"recolectado": true,
	"en_estacion": true,
	"en_ruta":     true,
	"entregada":   true,
	"cancelada":   true,
}
