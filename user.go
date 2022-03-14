package auth_jwt

type User struct {
	GUID string `json:"guid" binding:"required"`
}

type Refresh struct {
	GUID string `json:"guid" binding:"required"`
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type TokenPair struct {
	AccessToken string `json:"access_token" binding:"required"`
	RefreshToken string `json:"refresh_token" binding:"required"`
}
