package recipe

import "time"

// Recipe represents the recipe data model
type Recipe struct {
	ID                 int       `json:"id,omitempty"`
	Title              string    `json:"title,omitempty"`
	MakingTime         string    `json:"preparation_time,omitempty"`
	Serves             string    `json:"serves,omitempty"`
	Ingredients        string    `json:"ingredients,omitempty"`
	Cost               string    `json:"cost,omitempty"`
	FormattedCreatedAt string    `json:"created_at,omitempty"`
	FormattedUpdatedAt string    `json:"updated_at,omitempty"`
	CreatedAt          time.Time `json:"-"`
	UpdatedAt          time.Time `json:"-"`
}
