package models

import "time"

type Tag struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	OwnerID   *int      `json:"ownerId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
