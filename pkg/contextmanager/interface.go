package contextmanager

type Interface interface {
	GetUserID() uint
	IsLoggedIn() bool
}
