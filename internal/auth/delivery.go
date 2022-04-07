package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/yaml.v3"
	"os"
)

var configFile = "./configs/config.yaml"

func GetToken(username string) (string, error) {
	tokenConfig, err := parseToken()
	if err != nil {
		return "", fmt.Errorf("failed parse token")
	}

	signingKey := []byte(tokenConfig)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": username,
	})
	tokenString, err := token.SignedString(signingKey)
	return tokenString, err
}

func VerifyToken(tokenString string) (jwt.Claims, error) {
	tokenConfig, err := parseToken()
	if err != nil {
		return nil, fmt.Errorf("failed parse token")
	}

	signingKey := []byte(tokenConfig)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}

func parseToken() (string, error) {
	configContent, err := os.ReadFile(configFile)
	if err != nil {
		return "", fmt.Errorf("error parse delivery config: %s", err)
	}

	var config struct {
		Token struct {
			Key string
		}
	}

	err = yaml.Unmarshal(configContent, &config)
	if err != nil {
		return "", fmt.Errorf("error unmarshal delivery config: %s", err)
	}

	return config.Token.Key, nil
}
