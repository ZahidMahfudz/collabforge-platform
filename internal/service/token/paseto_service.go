package token

import (
	"encoding/hex"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/zahidmahfudz/collabforge-platform/config"
)

var Logger = config.Logger

type PasetoService struct {
	parser paseto.Parser
	key paseto.V4SymmetricKey
}

func NewPasetoService() *PasetoService {
	secretHex := config.GetEnv("PASETO_SECRET_KEY")

	//decode hex string ke byte
	secretKey, err := hex.DecodeString(secretHex)
	if err != nil {
		Logger.Fatalf("Failed to decode PASETO secret key: %v", err)
	}

	//buat symmetric key dari byte
	key, err := paseto.V4SymmetricKeyFromBytes(secretKey)
	if err != nil {
		Logger.Fatalf("Failed to create PASETO symmetric key: %v", err)
	}

	return &PasetoService{
		parser: paseto.NewParser(),
		key: key,
	}
}

func (p *PasetoService) GenerateAccessToken(UserId string, Email string, duration time.Duration) (string, error) {
	now := time.Now()
	expired := now.Add(duration)

	token := paseto.NewToken()

	token.SetIssuedAt(now)
	token.SetExpiration(expired)

	token.SetString("user_id", UserId)
	token.SetString("email", Email)

	signed := token.V4Encrypt(p.key, nil)

	return signed, nil
}

func (p *PasetoService) VerifyToken(tokenStr string) (*paseto.Token, error) {
	token, err := p.parser.ParseV4Local(p.key, tokenStr, nil)
	if err != nil {
		return nil, err
	}
	return token, nil
}