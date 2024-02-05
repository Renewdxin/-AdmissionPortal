package verify

import (
	"fmt"
	"math/rand"
	"time"
)

type VerificationCodeAdapter struct{}

func NewVerificationCodeAdapter() *VerificationCodeAdapter {
	return &VerificationCodeAdapter{}
}

func (s VerificationCodeAdapter) GenerateCode() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	code := rand.Intn(1000000)
	return fmt.Sprintf("%06d", code)
}
