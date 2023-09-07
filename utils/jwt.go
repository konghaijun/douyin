package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type MyClaims struct {
	//除了满足下面的Claims，还需要以下用户信息
	Uid int64
	//jwt中标准的Claims
	jwt.StandardClaims
}

// 使用指定的 secret 签名声明一个 key ，便于后续获得完整的编码后的字符串token
var key = []byte("jiguangweb")

// GetToken 生成token的方法
func GetToken(uid int64) (string, error) {

	c := &MyClaims{
		Uid: uid, //自定义字段
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),                          //发布时间
			ExpiresAt: time.Now().Add(time.Hour * 10 * 24).Unix(), //过期时间
			Issuer:    "tiktok123456",                             //签发人
		},
	}
	//使用指定的签名方法创建签名对象
	//这里使用HS256加密算法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseToken 解析token
func ParseToken(tokenString string) (*MyClaims, bool) {
	token, _ := jwt.ParseWithClaims(tokenString, &MyClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return key, nil
		})
	if token != nil {
		//校验token
		if claims, ok := token.Claims.(*MyClaims); ok {
			if token.Valid {
				return claims, true
			} else {
				return claims, false
			}
		}
	}
	return nil, false
}

// JWTMiddleWare 鉴权中间件，鉴权并设置token信息到Context
func JWTMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Query("token")
		if tokenStr == "" {
			tokenStr = c.PostForm("token")
		}
		//用户不存在
		if tokenStr == "" {
			c.JSON(http.StatusOK, "用户不存在")
			c.Abort() //阻止执行
			return
		}
		//验证token
		tokenStruck, ok := ParseToken(tokenStr)
		if !ok {
			c.JSON(http.StatusOK, "token不正确")
			c.Abort() //阻止执行
			return
		}
		//token超时
		if time.Now().Unix() > tokenStruck.ExpiresAt {
			c.JSON(http.StatusOK, "token过期")
			c.Abort() //阻止执行
			return
		}
		c.Set("user_id", tokenStruck.Uid)
		c.Next()
	}
}
