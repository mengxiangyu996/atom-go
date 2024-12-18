package token

import (
	"atom-go/config"
	"atom-go/internal/dto"
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// 后台授权声明
type AdminClaims struct {
	UserId   int    `json:"userId"`
	UserType string `json:"userType"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Status   string `json:"status"`
	jwt.RegisteredClaims
}

// 获取授权声明
func GetAdminClaims(user dto.UserTokenResponse) *AdminClaims {

	return &AdminClaims{
		UserId:   user.UserId,
		UserType: user.UserType,
		Username: user.Username,
		Nickname: user.Nickname,
		Status:   user.Status,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(config.Data.Token.ExpireTime))), // 过期时间，默认30分钟
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                                                // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                                                // 生效时间
			Issuer:    "atom-go",                                                                                     // 签发人
			Subject:   "admin",                                                                                       // 主题
		},
	}
}

// 生成token
func (a *AdminClaims) GenerateToken() string {

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, a).SignedString([]byte(config.Data.Token.Secret))
	if err != nil {
		return ""
	}

	return token
}

// 解析token
func ParseAdminToken(ctx *gin.Context) (*AdminClaims, error) {

	authorization := ctx.GetHeader(config.Data.Token.Header)
	if authorization == "" {
		return nil, errors.New("请先登录")
	}

	tokenSplit := strings.Split(authorization, " ")
	if len(tokenSplit) != 2 || tokenSplit[0] != "Bearer" {
		return nil, errors.New("authorization format error")
	}

	token, err := jwt.ParseWithClaims(tokenSplit[1], &AdminClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Data.Token.Secret), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("token格式错误")
			}
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token已过期")
			}
			if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token未生效")
			}
			return nil, errors.New("token校验失败")
		}
		return nil, err
	}

	if claims, ok := token.Claims.(*AdminClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("token校验失败")
}
