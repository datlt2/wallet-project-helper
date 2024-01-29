package jwtHelper

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

// Create token depend on private key
func CreateToken(userID string) (string, error) {
	log.Infof("HP-Create token for userID %v", userID)
	// Create claim
	claims := jwt.MapClaims{
		"user_id": userID,
		"start":   time.Now().Format(time.RFC3339),
		"expired": time.Now().Add(time.Hour).Format(time.RFC3339), // Token hết hạn sau 1 giờ
	}

	// Generate private key
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedToken, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// Verify token depend on public key
func VerifyToken(tokenString string) (*jwt.Token, error) {
	log.Infof("HP-Verify token %v", tokenString)

	// Remove unnecessary string
	sep := strings.Split(tokenString, " ")
	if len(sep) > 1 {
		tokenString = sep[len(sep)-1]
	} else {
		tokenString = sep[0]
	}

	// Generate public key
	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))

	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

// Verify token info
func VerifyTokenInfo(userID string, tokenInfo jwt.MapClaims) (bool, error) {
	if userID != tokenInfo["user_id"] {
		return false, errors.New("Token is not correct")
	}
	startTime, _ := time.Parse(time.RFC3339, tokenInfo["start"].(string))
	expiredTime, _ := time.Parse(time.RFC3339, tokenInfo["expired"].(string))

	if time.Now().After(startTime) && time.Now().Before(expiredTime) {
		return true, nil
	}
	return false, errors.New("Token is expired")
}

const (
	privateKey = `
-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQCZ4ipNuSKYdvjSCF4ehCdEEejebfMVXHQUhsX74DkLE/KxQyHU
rCYZgycvhYLV1+JQWh3WLGvSBIc2Oy1bka3fpAWrgSw16NdrTRysBvLAOBrkZKxF
s85OAlLot+g/uUYJJgPbwtdQN/AhLbO+0hHxqwbSGjNtduvBL59rtRFMXwIDAQAB
AoGABf4Al7Y12qoHVmZtA9MxlDz+TGsLpDFNc98rpns8vWmxcaWjk5yAf03UIM2a
IqmdbnOT5dsk97Clcz8qrP4nPxa2e3qt9pI8AG4rW4W+ktJLqJF4/hUITjqIn97s
Cs62Nr52mo2s/Ulo2KuUesYa32zMTrDRhYijNOQ0w3+P7EECQQDJFkQIeY9Iqado
KFQrsgomkyoG6d7YxXaarsa9sJoUSUeC8f/hadd9KgYA1yCs/bvqoruUhJtlH7s+
GVSBlFvtAkEAw+f2M4mY9pdtzugix8hHPDcwn2iQyL0UHFIJJSkHRjBfQwjt6LIw
qhvED3R09SogHej0TXxtihlP1UUd/WB3+wJBALm2NJLXIXtsI83QIRxRy4ogs9m6
uDLe+1CURbv+k+5UVxUaRsV5qKhR3UV9aRIbLCfgrvjVF2bnTVhmsGMUD2kCQBed
cpQr1gCDqHz0hpzGi6+6h2Xv2OQZxr7TNL8B/xp64kDWZrdSI4Od7ThftWmINm7k
ke51PU8UVgdDWaYdZtkCQErNRIT9RP4uedykhaOlyDG3R3QTLXPpAYl5T/zdqv/4
3Ee6Lf4x03EFmw6/T0IVHrmWyPq8TltmkK3wFX1UJ34=
-----END RSA PRIVATE KEY-----
`

	publicKey = `
-----BEGIN RSA PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCZ4ipNuSKYdvjSCF4ehCdEEeje
bfMVXHQUhsX74DkLE/KxQyHUrCYZgycvhYLV1+JQWh3WLGvSBIc2Oy1bka3fpAWr
gSw16NdrTRysBvLAOBrkZKxFs85OAlLot+g/uUYJJgPbwtdQN/AhLbO+0hHxqwbS
GjNtduvBL59rtRFMXwIDAQAB
-----END RSA PUBLIC KEY-----
`
)
