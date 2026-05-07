package bcrypt

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports"
	"golang.org/x/crypto/bcrypt"
)

type BcryptHasher struct {
	cost int
}

func NewBcryptHasher(cost int) ports.PasswordHasher {
	if cost == 0 {
		cost = bcrypt.DefaultCost
	}
	return &BcryptHasher{cost: cost}
}

func (h *BcryptHasher) Hash(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (h *BcryptHasher) Compare(hased, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hased), []byte(password))
	return err == nil
}
