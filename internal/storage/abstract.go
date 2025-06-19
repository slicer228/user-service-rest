package storage

type Connection interface {
	NewSession() (Session, error)
}

type Session interface {
}

type UserInteractor interface {
}
