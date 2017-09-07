package entities

type User struct  {
	OpenId string `json:"open_id"`
	NickName string `json:"nick_name"`
	PhoneNumber string `json:"phone_number"`
	Gender string `json:"gender"`
	City string `json:"city`
	Province string `json:"province"`
	Country string `json:"Country"`
	AvatarUrl string `json:"avatar_url"`
	Source string `json:"source"`
	CreatedAt string `json:"created_at"`
	DeletedAt string `json:"deleted_at"`
}