package models

type PackageData struct {
	Weight  float64 `json:"weight"`
	Size    string  `json:"size"`
	Service string  `json:"service,omitempty" default:"Standard"`
}
