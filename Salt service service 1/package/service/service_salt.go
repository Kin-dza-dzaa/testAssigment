package service

import (
	"math/rand"
	"strings"
)

const (
	SALT_LEN     = 32
	LETTER_BYTES = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
)

type SaltService struct {
}

func (s *SaltService) GenerateSalt() string {
	sb := strings.Builder{}
	sb.Grow(SALT_LEN)

	for i := 0; i < SALT_LEN; i++ {
		sb.WriteByte(LETTER_BYTES[rand.Intn(len(LETTER_BYTES))])
	}

	return sb.String()
}
