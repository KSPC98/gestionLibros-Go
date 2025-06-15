package security

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

// Encapsulación: campo privado
type Usuario struct {
	email    string
	password string
}

func (u *Usuario) Autenticar(email, password string) error {
	if u.email != email || u.password != encryptPassword(password) {
		return errors.New("credenciales inválidas")
	}
	return nil
}

// Método privado (minúscula)
func encryptPassword(pwd string) string {
	hash := sha256.Sum256([]byte(pwd))
	return hex.EncodeToString(hash[:])
}
