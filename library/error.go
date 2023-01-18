package library

type Error struct {
	Message string
}

func (e Error) Validate() Error {
	return Error{
		Message: e.Message,
	}
}
