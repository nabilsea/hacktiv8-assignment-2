package custom_error

type ItemAlreadyExists struct {
}

func (e *ItemAlreadyExists) Error() string {
	return "item already exists"
}
