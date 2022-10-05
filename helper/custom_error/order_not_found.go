package custom_error

type OrderNotFound struct {
}

func (e *OrderNotFound) Error() string {
	return "order not found"
}
