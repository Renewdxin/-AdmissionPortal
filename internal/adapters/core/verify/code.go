package verify

import (
	"fmt"
	"math/rand"
	"time"
)

type VerificationCodeService struct{}

func NewVerificationCodeService() *VerificationCodeService {
	return &VerificationCodeService{}
}

func (s *VerificationCodeService) GenerateCode() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	code := rand.Intn(1000000)
	return fmt.Sprintf("%06d", code)
}
