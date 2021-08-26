package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

/**
 * 自定义结构体 可以额外配置参数
 */
type SomeStandardClaims struct {
	UserName string
	jwt.StandardClaims   // jwt提供的结构体
}

type Conf struct {
	Secret []byte
	ExpireAt int64
}

type JWT struct {
	Secret []byte
	ExpireAt int64
}

func New(conf *Conf) *JWT {
	return &JWT{Secret: conf.Secret, ExpireAt: conf.ExpireAt}
}

// GenToken 生成JWT
func (j *JWT) GenerateToken(username string) (string, int64, error) {
	c := SomeStandardClaims{
		UserName: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(j.ExpireAt)).Unix(), // 过期时间
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	value, err := token.SignedString(j.Secret)
	if err != nil {
		return "", 0, err
	}
	return value, c.ExpiresAt, nil
}

// ParseToken 解析JWT
func (j *JWT) ParseToken(tokenString string) (*SomeStandardClaims, error) {
	// 解析token
	someStandardClaims := new(SomeStandardClaims)
	token, err := jwt.ParseWithClaims(tokenString, someStandardClaims, func(token *jwt.Token) (i interface{}, err error) {
		return j.Secret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*SomeStandardClaims)
	if ok == false || token.Valid == false { // 校验token
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
