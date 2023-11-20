package common

import (
	"SpiderShop-Restfull-API/module/user/entity"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(user *entity.UserCreate, JWT_SECRET_KEY string) (string, error) {
	// create claims
	claims := &Claims{
		UserID:   user.UserID,
		Account:  user.Account,
		Password: user.Password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Hạn sử dụng JWT trong 1 ngày.
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// create token sign
	tokenString, err := token.SignedString([]byte(JWT_SECRET_KEY))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func PraseToken(tokenString string, JWT_SECRET_KEY string) (*Claims, error) {
	// Parse token từ chuỗi tokenString.
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JWT_SECRET_KEY, nil
	})
	if err != nil {
		return nil, err
	}
	// Kiểm tra lỗi và xác minh token.
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		// Xử lý dữ liệu từ claims (email và password).
		return claims, nil
	}
	// Xác thực không thành công.
	return nil, errors.New(TOKEN_NOT_APPROPRIATE)

}
