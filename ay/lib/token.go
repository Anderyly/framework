package lib

import "github.com/gin-gonic/gin"

var _ Token = (*token)(nil)

type Token interface {
	GenerateToken(string, int64) (string, error)
}

const (
	TokenAccessKey  = "token_access_"
	TokenRefreshKey = "token_refresh_"
	TokenUserIdKey  = "token_im_user_id_"
	TokenDnoKey     = "token_dno_"
)

type token struct {
	Ctx *gin.Context
}

func NewToken(ctx *gin.Context) Token {
	return &token{
		Ctx: ctx,
	}
}

type MakeToken struct {
	AccessToken  string
	RefreshToken string
	Dno          string
	Expiry       int64
	UserId       int
}

func (t *token) GenerateToken(str string, expire int64) (res string, err error) {
	if res, err = NewStr().AuthCode(str, true, "", expire); err != nil {
		return
	}

	return
}
