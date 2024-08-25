package internal

type Request interface {
	Do(cep string) (*Message, error)
}
