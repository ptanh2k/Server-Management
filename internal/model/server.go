package model

import "time"

type Server struct {
	ServerID    uint16    `json:"server_id,omitempty"`
	ServerName  string    `json:"server_name"`
	Status      bool      `json:"status"`
	CreatedTime time.Time `json:"created_time"`
	LastUpdated time.Time `json:"last_updated"`
	Domain      string    `json:"domain"`
	CreatedBy   uint16    `json:"created_by"`
}
