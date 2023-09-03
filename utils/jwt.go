package utils

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

const (
	Expire      = 1000 * 60 * 60 * 24              // token过期时间
	AppSecret   = "ukc8BDbRigUDaY6pZFfWus2jZWLPHO" // 秘钥
	TokenHeader = "token"
)

// 生成token字符串的方法
func GetJwtToken(id int64) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Duration(Expire)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims
	tokenString, err := token.SignedString([]byte(AppSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 判断token是否存在与有效
func CheckToken(jwtToken string) bool {
	if jwtToken == "" {
		return false
	}
	_, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(AppSecret), nil
	})
	if err != nil {
		return false
	}
	return true
}

// 根据token字符串获取会员id
func GetMemberIdByJwtToken(request *http.Request) string {
	jwtToken := request.Header.Get(TokenHeader)
	if jwtToken == "" {
		return ""
	}
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(AppSecret), nil
	})
	if err != nil || !token.Valid {
		return ""
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok {
		return claims["id"].(string)
	}
	return ""
}
