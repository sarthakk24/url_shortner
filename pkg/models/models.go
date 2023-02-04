package models

type User struct {
	UserID   string
	Name     string
	Email    string
	Password string
}

type Url struct {
	UrlId       string
	BaseUrl     string
	RedirectUrl string
}