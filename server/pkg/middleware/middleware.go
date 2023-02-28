package middleware

type Middleware struct {
	SignKey []byte
}

func New(key []byte) *Middleware {
	return &Middleware{
		SignKey: key,
	}
}
