package auth

type Store interface {
	Add(hash, email string)
	Validate(hash string) bool
}
