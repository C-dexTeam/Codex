package hasherService

import "github.com/matthewhartstonge/argon2"

func HashPassword(password string) (hashedPassword string, err error) {
	argon := argon2.DefaultConfig()

	hash, err := argon.HashEncoded([]byte(password))
	if err != nil {
		return "", err
	}

	return string(hash), nil

}

func CompareHashAndPassword(hashedPassword string, password string) (ok bool, err error) {
	return argon2.VerifyEncoded([]byte(password), []byte(hashedPassword))
}
