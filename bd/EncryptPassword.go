package bd

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(pass string) (string, error) {
	/*
		lo minimo es de 6
		power de 6 para un usuario normal
		power de 8 para un superusuario
		power de mas de 12 tardar muchisimo
	*/
	power := 6
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), power)
	return string(bytes), err
}
