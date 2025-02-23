package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
)

type Jwt struct {
	signingKey []byte
}

func NewJwt(signingKey string) *Jwt {
	return &Jwt{signingKey: []byte(signingKey)}
}

func (j *Jwt) CreateToken(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.signingKey)
}

// ParseToken 解析并验证 JWT Token
func (j *Jwt) ParseToken(tokenString string) (map[string]any, error) {
	// 解析 Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("无效的签名方法")
		}
		// 返回密钥
		return j.signingKey, nil
	})
	if err != nil {
		return nil, errors.New("无效的 Token")
	}

	// 验证 Token 是否有效
	if !token.Valid {
		return nil, errors.New("无效的 Token")
	}

	// 提取 Claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("无效的 Claims")
	}

	// 检查 Token 是否过期
	exp, ok := claims["exp"].(float64)
	if !ok {
		return nil, errors.New("无效的过期时间")
	}
	if time.Now().Unix() > int64(exp) {
		return nil, errors.New("Token 已过期")
	}

	// 返回 Claims
	return claims, nil
}
