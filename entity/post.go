package entity

import "time"

// Post represents a user post object
type Post struct {
	ID       string
	Caption  string
	ImageURL string
	PostedTimestamp time.Duration
}
