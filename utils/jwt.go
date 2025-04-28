package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("your-secret-key") // 这里替换为你的密钥

// 定义 JWT 载荷结构
type Claims struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Roles []int  `json:"roles"` // 角色列表
	jwt.StandardClaims
}

// 创建 JWT
func GenerateJWT(id int, name string, phone string, rules []int) (string, error) {
	claims := Claims{
		ID:    id,
		Name:  name,
		Phone: phone,
		Roles: rules,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(), // 设置过期时间
			Issuer:    "your-app",
		},
	}

	// 使用密钥签名生成 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// 解析 JWT 并返回解析后的字符串
func GetJWT(headerToken string) (*Claims, error) {
	var claims Claims

	// 解析 JWT
	token, err := jwt.ParseWithClaims(headerToken, &claims, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法
		return jwtSecret, nil
	})

	// 错误处理
	if err != nil {
		return &claims, err
	}

	// 如果 token 无效
	if !token.Valid {
		return &claims, jwt.NewValidationError("invalid token", jwt.ValidationErrorMalformed)
	}

	return &claims, nil
}
