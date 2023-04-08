package contextmanager

type CtxKey uint

const (
	CtxUserID CtxKey = iota
	CtxIsLoggedIn
)
