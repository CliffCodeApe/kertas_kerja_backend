package token

import (
	"crypto/rsa"
	"log"
	"os"

	"kertas_kerja/config"

	"github.com/golang-jwt/jwt/v5"
)

var jwtConfig *jwtStruct

type jwtStruct struct {
	jwtLifeTime        uint
	jwtRefreshLifeTime uint
	privateKey         *rsa.PrivateKey
	publicKey          *rsa.PublicKey
}

func Load() {
	cfg := config.Get()

	privateKeyBytes, err := os.ReadFile("private.pem")

	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)

	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	publicKeyBytes, err := os.ReadFile("public.pem")

	if err != nil {
		log.Fatalf("Failed to load public key: %v", err)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)

	if err != nil {
		log.Fatalf("Failed to parse public key: %v", err)
	}

	jwtConfig = &jwtStruct{
		jwtLifeTime:        cfg.AccessTokenLifeTime,
		jwtRefreshLifeTime: cfg.RefreshTokenLifeTime,
		publicKey:          publicKey,
		privateKey:         privateKey,
	}
}
