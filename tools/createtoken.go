package tools

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

const Key = "My Secret"

func CreateToken(user, password string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":     user,
		"password": password,
		"time":     time.Now().Unix(),
		"exp":      time.Now().Add(10 * time.Minute * time.Duration(1)).Unix(),
	})
	//  把token已约定的加密方式和加密秘钥加密，当然也可以使用不对称加密
	tokenString, _ := token.SignedString([]byte(Key))
	s, err := token.SigningString()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(s)

	return tokenString
	//  登录时候，把tokenString返回给客户端，然后需要登录的页面就在header上面附此字符串
	//  eg: header["Authorization"] = "bearer "+tokenString
}

func ValidToken(tokenstring string) bool {
	hmacSampleSecret := []byte(Key)
	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSampleSecret, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["user"], claims["password"])
		return true
	} else {
		fmt.Println(err)
		return false
	}

}
