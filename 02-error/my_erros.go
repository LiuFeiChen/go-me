package main


type NotFoundError struct {
	Msg string
	Err error
}


func (n *NotFoundError) Error() string{
	return n.Msg
}

func (n *NotFoundError) Unwrap() error {
	return n.Err
}

