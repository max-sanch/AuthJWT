package auth_jwt

type User struct {
	GUID string `json:"guid" binding:"required"`
}
