package passwd

type Interface interface {
	Generate(password string) (string, error)
	Compare(hashed, password string) error
}
