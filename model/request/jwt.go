package request

import (
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

// Custom claims structure
type CustomClaims struct {
	UUID        uuid.UUID
	ID          uint
	Username    string
	NickName    string
	AuthorityId string
	BufferTime  int64
	jwt.StandardClaims
}


type MyClaims struct {
	CreateTime uint `json:"iat"`
	Host string `json:"aud"`
	Jti jti `json:"jti"`
	jwt.StandardClaims
}

type jti struct {
	Uid uint `json:"id"`
	Type string `json:"type"`
}