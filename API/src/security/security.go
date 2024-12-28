package security

import "golang.org/x/crypto/bcrypt"

// Hash recieves a password and hash this to save in the database
func Hash(password string) ([]byte, error){
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
// VerifyPassword verify a passward against hash and returns an error if they are not equivalent
func VerifyPassword(password, passwordHash string) error{
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
}