package errors

type PanicResponse interface {
	Code() int
	Object() any
}
