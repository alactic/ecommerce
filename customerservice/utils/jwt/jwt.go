package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(details map[string]string) string {
	// var secretkey = os.Get("mysecretkey")
	mySigningKey := []byte("elvisSecreyKey")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo":       "bar",
		"firstname": details["firstname"],
		"lastname":  details["lastname"],
		"email":     details["email"],
		"nbf":       time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Errorf("error generated while creating token : %s", err.Error())
	}
	// fmt.Println(tokenString, err)
	return tokenString
}

func DecodeJWT(tokenBearer string) {
	// Token from another example.  This token is expired
	var tokenString = tokenBearer

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("elvisSecreyKey"), nil
	})

	if token.Valid {
		// byte, _ := json.Marshal(token.Claims)
		// fmt.Println("You look nice today :: ", string(byte))
		// fmt.Println("You look nice today :: ", token.Claims)
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println(claims["foo"], claims["firstname"])
		} else {
			fmt.Println(err)
		}
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
		} else {
			fmt.Println("Couldn't handle this token:", err)
		}
	} else {
		fmt.Println("Couldn't handle this token:", err)
	}
}
