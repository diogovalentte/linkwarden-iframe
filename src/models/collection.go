package models

import "time"

// Collection represents a Linkwarden collection
type Collection struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	Color       string    `json:"color"`
	ParentID    *int      `json:"parentId,omitempty`
	IsPublic    bool      `json:"isPublic"`
	OwnerID     *int      `json:"ownerId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
