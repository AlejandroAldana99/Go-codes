package models

import "time"

type Status uint

const (
	Pass Status = iota
	Fail
	Warn
)

func (s Status) String() string {
	statusNames := [...]string{
		"pass",
		"fail",
		"warn",
	}
	return statusNames[s]
}

type HealthStatus struct {
	Status      string                  `json:"status"`
	Description string                  `json:"description"`
	Version     string                  `json:"version"`
	Details     []HealthComponentDetail `json:"details"`
}

type HealthComponentDetail struct {
	Name        string    `json:"componentName"`
	Type        string    `json:"componentType"`
	ID          string    `json:"componentId"`
	MetricValue float32   `json:"metricValue"`
	MetricUnit  string    `json:"metricUnit"`
	Time        time.Time `json:"time"`
	Status      string    `json:"status"`
}
