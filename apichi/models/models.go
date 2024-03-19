package models

import (
    "time"
)
// User represents a user in the system
type Tasks struct {
    ID        string    `json:"id"`
    Title     string    `json:"title"`
    CreatedAt time.Time `json:"created_at"`
}
