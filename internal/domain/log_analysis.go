package domain

import "time"

type LogAnalysis struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	UserID          uint      `json:"user_id"`
	Filename        string    `json:"filename"`
	TotalRequests   int       `json:"total_requests"`
	UniqueIPs       int       `json:"unique_ips"`
	ErrorCount      int       `json:"error_count"`
	AverageResponse float64   `json:"average_response"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
