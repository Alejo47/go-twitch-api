package users

import "time"

type Pagination struct {
	Cursor string `json:"cursor"`
}

type FollowRequest struct {
	After  string `json:"after"`
	First  string `json:"first"`
	FromID string `json:"from_id"`
	ToID   string `json:"to_id"`
}

type Num1 struct {
	Active  bool   `json:"active"`
	ID      string `json:"id"`
	Version string `json:"version"`
	Name    string `json:"name"`
}
type Num2 struct {
	Active  bool   `json:"active"`
	ID      string `json:"id"`
	Version string `json:"version"`
	Name    string `json:"name"`
}
type Num3 struct {
	Active  bool   `json:"active"`
	ID      string `json:"id"`
	Version string `json:"version"`
	Name    string `json:"name"`
}
type Panel struct {
	Num1 Num1 `json:"1"`
	Num2 Num2 `json:"2"`
	Num3 Num3 `json:"3"`
}
type Overlay struct {
	Num1 Num1 `json:"1"`
}
type Num1 struct {
	Active  bool   `json:"active"`
	ID      string `json:"id"`
	Version string `json:"version"`
	Name    string `json:"name"`
	X       int    `json:"x"`
	Y       int    `json:"y"`
}
type Num2 struct {
	Active bool `json:"active"`
}
type Component struct {
	Num1 Num1 `json:"1"`
	Num2 Num2 `json:"2"`
}
type UserExtensionsData struct {
	Panel     Panel     `json:"panel"`
	Overlay   Overlay   `json:"overlay"`
	Component Component `json:"component"`
}

type ActiveUserExtensionsData struct {
	ID          string   `json:"id"`
	Version     string   `json:"version"`
	Name        string   `json:"name"`
	CanActivate bool     `json:"can_activate"`
	Type        []string `json:"type"`
}

type UserData struct {
	ID              string `json:"id"`
	Login           string `json:"login"`
	DisplayName     string `json:"display_name"`
	Type            string `json:"type"`
	BroadcasterType string `json:"broadcaster_type"`
	Description     string `json:"description"`
	ProfileImageURL string `json:"profile_image_url"`
	OfflineImageURL string `json:"offline_image_url"`
	ViewCount       int    `json:"view_count"`
	Email           string `json:"email"`
}

type FollowData struct {
	FromID     string    `json:"from_id"`
	FromName   string    `json:"from_name"`
	ToID       string    `json:"to_id"`
	ToName     string    `json:"to_name"`
	FollowedAt time.Time `json:"followed_at"`
}

type UserExtensionsResponse struct {
	UserExtensionsData []Data `json:"data"`
}

type ActiveUserExtensionsResponse struct {
	ActiveUserExtensionsData Data `json:"data"`
}

type UserResponse struct {
	UserData []Data `json:"data"`
}

type FollowsResponse struct {
	Total      int        `json:"total"`
	FollowData []Data     `json:"data"`
	Pagination Pagination `json:"pagination"`
}
