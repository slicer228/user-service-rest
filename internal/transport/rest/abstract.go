package rest

type Server interface {
	MustStart()
	GracefulShutdown()
}
